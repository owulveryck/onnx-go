package gorgonnx

import (
	"errors"

	"github.com/owulveryck/onnx-go"
	"gorgonia.org/gorgonia"
)

type stableSoftmax struct {
	axis int
}

func init() {
	register("Softmax", newStableSoftmax)
}

func newStableSoftmax() operator {
	return &stableSoftmax{}
}

func (s *stableSoftmax) apply(g *Graph, ns ...*Node) error {
	n := ns[0]
	children := getOrderedChildren(g.g, n)
	err := checkCondition(children, 1)
	if err != nil {
		return err
	}
	a := children[0].gorgoniaNode
	coerced, err := coerceInto2D(a, s.axis)
	if err != nil {
		return err
	}
	stableCoerced, err := stabilizeNode(coerced)
	if err != nil {
		return err
	}
	var exp, sum *gorgonia.Node
	if exp, err = gorgonia.Exp(stableCoerced); err == nil {
		axis := 1
		if exp.IsScalar() {
			axis = 0
		}
		sum, err = gorgonia.Sum(exp, axis)
		if err == nil {
			if sum.IsScalar() {
				tmp, err := gorgonia.HadamardDiv(exp, sum)
				if err != nil {
					return err
				}
				n.gorgoniaNode, err = gorgonia.Reshape(tmp, children[0].gorgoniaNode.Shape())
				return err
			}
			a, b, err := gorgonia.Broadcast(exp, sum, gorgonia.NewBroadcastPattern(nil, []byte{1}))
			if err != nil {
				return err
			}
			tmp, err := gorgonia.HadamardDiv(a, b)
			if err != nil {
				return err
			}
			n.gorgoniaNode, err = gorgonia.Reshape(tmp, children[0].gorgoniaNode.Shape())
			return err
		}
		return err
	}
	return err
}

// coerceInto2D does a reshape of the node following this rule:
// For an arbitrary n-dimensional tensor input ∈ [a₀, a₁, ..., aₖ₋₁, a_k, ..., aₙ₋₁]
// and k is the axis provided, then input will be coerced into a 2-dimensional tensor
// with dimensions [a₀ * ... * aₖ₋₁, a_k * ... * aₙ₋₁].
// For the default case where axis=1,
// this means the input tensor will be coerced into a 2D tensor of dimensions [a₀, a₁ * ... * aₙ₋₁],
// where a₀ is often the batch size.
// In this situation, we must have a₀ = N and a₁ * ... * aₙ₋₁ = D.
// Each of these dimensions must be matched correctly, or else the operator will throw errors.
func coerceInto2D(n *gorgonia.Node, k int) (*gorgonia.Node, error) {
	if len(n.Shape()) == 2 {
		return n, nil
	}
	if k > len(n.Shape()) {
		return nil, errors.New("softmax cannot be applied on an axis > len(shape()) of the input")
	}
	row := 1
	col := 1
	for i, shape := range n.Shape() {
		if i < k {
			row = row * shape
		} else {
			col = col * shape
		}
	}
	return gorgonia.Reshape(n, []int{row, col})
}

// stabilize a node consits of subtracting the max value from every row
// this function returns an error is the input is not a tensor or a matrix
func stabilizeNode(n *gorgonia.Node) (*gorgonia.Node, error) {
	if len(n.Shape()) != 2 {
		return nil, errors.New("can only stabilize a vector or a matrix")
	}

	var m1 *gorgonia.Node
	var err error

	if n.Shape()[0] == 1 {
		m1, err = gorgonia.Max(n)
	} else {
		m1, err = gorgonia.Max(n, 1)
	}

	if err != nil {
		return nil, err
	}

	a1, b1, err := gorgonia.Broadcast(n, m1, gorgonia.NewBroadcastPattern(nil, []byte{1}))
	if err != nil {
		return nil, err
	}
	return gorgonia.Sub(a1, b1)
}

func (s *stableSoftmax) init(o onnx.Operation) error {
	axis, ok := o.Attributes["axis"]
	if !ok {
		s.axis = 1
		return nil
	}
	err := errors.New("axis in not an int")
	if axis, ok := axis.(int64); ok {
		s.axis = int(axis)
		err = nil
	}
	if s.axis < 0 {
		return &onnx.ErrNotImplemented{
			Operator: "Softmax",
			Message:  "Negative axis not supported",
		}
	}
	return err
}
