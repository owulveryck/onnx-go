package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Resize", "TestResizeDownsampleLinear", NewTestResizeDownsampleLinear)
}

// NewTestResizeDownsampleLinear version: 4.
func NewTestResizeDownsampleLinear() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Resize",
		Title:  "TestResizeDownsampleLinear",
		ModelB: []byte{0x8, 0x4, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x9a, 0x1, 0xa, 0x29, 0xa, 0x1, 0x58, 0xa, 0x6, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x1, 0x59, 0x22, 0x6, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x2a, 0x11, 0xa, 0x4, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x6, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0xa0, 0x1, 0x3, 0x12, 0x1d, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x5a, 0x1b, 0xa, 0x1, 0x58, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x4, 0x5a, 0x14, 0xa, 0x6, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x4, 0x62, 0x1b, 0xa, 0x1, 0x59, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x2, 0x42, 0x2, 0x10, 0xa},

		/*

		   &pb.NodeProto{
		     Input:     []string{"X", "scales"},
		     Output:    []string{"Y"},
		     Name:      "",
		     OpType:    "Resize",
		     Attributes: ([]*pb.AttributeProto) (len=1 cap=1) {
		    (*pb.AttributeProto)(0xc000119400)(name:"mode" type:STRING s:"linear" )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 2, 4),
				tensor.WithBacking([]float32{1, 2, 3, 4, 5, 6, 7, 8}),
			),

			tensor.New(
				tensor.WithShape(4),
				tensor.WithBacking([]float32{1, 1, 0.6, 0.6}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 1, 2),
				tensor.WithBacking([]float32{1, 2.6666665}),
			),
		},
	}
}
