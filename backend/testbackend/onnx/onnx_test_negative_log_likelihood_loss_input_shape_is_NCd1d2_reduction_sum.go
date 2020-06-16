package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("NegativeLogLikelihoodLoss", "TestNegativeLogLikelihoodLossInputShapeIsNCd1d2ReductionSum", NewTestNegativeLogLikelihoodLossInputShapeIsNCd1d2ReductionSum)
}

/*
&ir.ModelProto{
    IrVersion:   6,
    OpsetImport: {
        &ir.OperatorSetIdProto{Domain:"", Version:12},
    },
    ProducerName:    "backend-test",
    ProducerVersion: "",
    Domain:          "",
    ModelVersion:    0,
    DocString:       "",
    Graph:           &ir.GraphProto{
        Node: {
            &ir.NodeProto{
                Input:     {"input", "target"},
                Output:    {"loss"},
                Name:      "",
                OpType:    "NegativeLogLikelihoodLoss",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:          "reduction",
                        RefAttrName:   "",
                        DocString:     "",
                        Type:          3,
                        F:             0,
                        I:             0,
                        S:             {0x73, 0x75, 0x6d},
                        T:             (*ir.TensorProto)(nil),
                        G:             (*ir.GraphProto)(nil),
                        SparseTensor:  (*ir.SparseTensorProto)(nil),
                        Floats:        nil,
                        Ints:          nil,
                        Strings:       nil,
                        Tensors:       nil,
                        Graphs:        nil,
                        SparseTensors: nil,
                    },
                },
                DocString: "",
            },
        },
        Name:              "test_negative_log_likelihood_loss_input_shape_is_NCd1d2_reduction_sum",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "input",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:5},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:6},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:6},
                                        Denotation: "",
                                    },
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
            &ir.ValueInfoProto{
                Name: "target",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 6,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:6},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:6},
                                        Denotation: "",
                                    },
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
        },
        Output: {
            &ir.ValueInfoProto{
                Name: "loss",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{},
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
        },
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestNegativeLogLikelihoodLossInputShapeIsNCd1d2ReductionSum version: 6.
func NewTestNegativeLogLikelihoodLossInputShapeIsNCd1d2ReductionSum() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "NegativeLogLikelihoodLoss",
		Title:  "TestNegativeLogLikelihoodLossInputShapeIsNCd1d2ReductionSum",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xdd, 0x1, 0xa, 0x45, 0xa, 0x5, 0x69, 0x6e, 0x70, 0x75, 0x74, 0xa, 0x6, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x4, 0x6c, 0x6f, 0x73, 0x73, 0x22, 0x19, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x6b, 0x65, 0x6c, 0x69, 0x68, 0x6f, 0x6f, 0x64, 0x4c, 0x6f, 0x73, 0x73, 0x2a, 0x13, 0xa, 0x9, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3, 0x73, 0x75, 0x6d, 0xa0, 0x1, 0x3, 0x12, 0x45, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x6c, 0x69, 0x68, 0x6f, 0x6f, 0x64, 0x5f, 0x6c, 0x6f, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x5f, 0x69, 0x73, 0x5f, 0x4e, 0x43, 0x64, 0x31, 0x64, 0x32, 0x5f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x6d, 0x5a, 0x1f, 0xa, 0x5, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x6, 0xa, 0x2, 0x8, 0x6, 0x5a, 0x1c, 0xa, 0x6, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x12, 0xa, 0x10, 0x8, 0x6, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x6, 0xa, 0x2, 0x8, 0x6, 0x62, 0xe, 0xa, 0x4, 0x6c, 0x6f, 0x73, 0x73, 0x12, 0x6, 0xa, 0x4, 0x8, 0x1, 0x12, 0x0, 0x42, 0x2, 0x10, 0xc},

		/*

		   &ir.NodeProto{
		     Input:     []string{"input", "target"},
		     Output:    []string{"loss"},
		     Name:      "",
		     OpType:    "NegativeLogLikelihoodLoss",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000217600)(name:"reduction" type:STRING s:"sum" )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 5, 6, 6),
				tensor.WithBacking([]float32{0.5488135, 0.71518934, 0.60276335, 0.5448832, 0.4236548, 0.6458941, 0.4375872, 0.891773, 0.96366274, 0.3834415, 0.79172504, 0.5288949, 0.56804454, 0.92559665, 0.071036056, 0.0871293, 0.020218397, 0.83261985, 0.77815676, 0.87001216, 0.9786183, 0.7991586, 0.46147937, 0.7805292, 0.11827443, 0.639921, 0.14335328, 0.9446689, 0.5218483, 0.41466194, 0.2645556, 0.7742337, 0.45615032, 0.56843394, 0.0187898, 0.6176355, 0.6120957, 0.616934, 0.94374806, 0.6818203, 0.3595079, 0.43703195, 0.6976312, 0.06022547, 0.6667667, 0.67063785, 0.21038257, 0.12892629, 0.31542835, 0.36371076, 0.57019675, 0.43860152, 0.9883738, 0.10204481, 0.20887676, 0.16130951, 0.6531083, 0.2532916, 0.46631077, 0.2444256, 0.15896958, 0.11037514, 0.6563296, 0.13818295, 0.19658236, 0.36872518, 0.82099324, 0.09710128, 0.8379449, 0.09609841, 0.97645944, 0.4686512, 0.9767611, 0.6048455, 0.7392636, 0.039187793, 0.28280696, 0.12019656, 0.2961402, 0.11872772, 0.31798318, 0.41426298, 0.064147495, 0.6924721, 0.56660146, 0.2653895, 0.5232481, 0.09394051, 0.5759465, 0.9292962, 0.31856894, 0.6674104, 0.13179787, 0.7163272, 0.2894061, 0.18319136, 0.5865129, 0.020107547, 0.82894003, 0.004695476, 0.6778165, 0.27000797, 0.735194, 0.96218854, 0.24875315, 0.57615733, 0.5920419, 0.5722519, 0.22308163, 0.952749, 0.44712538, 0.84640867, 0.6994793, 0.29743695, 0.81379783, 0.39650574, 0.8811032, 0.5812729, 0.8817354, 0.6925316, 0.7252543, 0.50132436, 0.95608366, 0.6439902, 0.42385504, 0.6063932, 0.019193199, 0.30157483, 0.66017354, 0.2900776, 0.6180154, 0.4287687, 0.13547407, 0.29828233, 0.5699649, 0.59087276, 0.57432526, 0.6532008, 0.65210325, 0.43141845, 0.8965466, 0.36756188, 0.43586493, 0.89192337, 0.806194, 0.7038886, 0.10022689, 0.9194826, 0.7142413, 0.998847, 0.1494483, 0.86812603, 0.16249293, 0.6155596, 0.123819984, 0.8480082, 0.807319, 0.56910074, 0.4071833, 0.069166996, 0.69742876, 0.45354268, 0.7220556, 0.8663823, 0.9755215, 0.8558034, 0.011714084, 0.35997805, 0.72999054, 0.17162968, 0.5210366, 0.05433799, 0.19999653, 0.018521795, 0.7936977, 0.22392468, 0.34535167, 0.9280813, 0.7044144, 0.03183893, 0.16469416, 0.6214784, 0.5772286, 0.23789282, 0.934214, 0.6139659, 0.5356328, 0.58991, 0.730122, 0.311945, 0.39822108, 0.20984375, 0.186193, 0.9443724, 0.73955077, 0.49045882, 0.22741462, 0.25435647, 0.05802916, 0.43441662, 0.3117959, 0.6963435, 0.37775183, 0.17960368, 0.024678728, 0.06724963, 0.67939276, 0.45369685, 0.5365792, 0.8966713, 0.9903389, 0.21689698, 0.6630782, 0.26332238, 0.020651, 0.7583786, 0.32001716, 0.3834639, 0.5883171, 0.8310484, 0.6289818, 0.8726507, 0.27354205, 0.7980468, 0.18563594, 0.95279163, 0.68748826, 0.21550767, 0.9473706, 0.7308558, 0.25394166, 0.21331197, 0.5182007, 0.025662718, 0.20747007, 0.42468548, 0.37416998, 0.46357542, 0.27762872, 0.58678436, 0.8638556, 0.11753186, 0.5173791, 0.13206811, 0.7168597, 0.3960597, 0.5654213, 0.18327984, 0.14484777, 0.48805627, 0.35561273, 0.94043195, 0.76532525, 0.7486636, 0.9037197, 0.08342244, 0.55219245, 0.58447605, 0.96193635, 0.29214752, 0.24082878, 0.10029394, 0.01642963, 0.9295293, 0.66991657, 0.7851529, 0.28173012, 0.58641016, 0.06395527, 0.4856276, 0.97749513, 0.87650526, 0.33815897, 0.96157014, 0.23170163, 0.9493188, 0.9413777, 0.79920256, 0.6304479, 0.87428796, 0.29302028, 0.84894353, 0.6178767, 0.013236858, 0.3472335, 0.14814086, 0.9818294, 0.4783703, 0.49739137, 0.63947254, 0.3685846, 0.13690028, 0.82211775, 0.18984792, 0.511319, 0.22431703, 0.09784448, 0.8621915, 0.97291946, 0.9608347, 0.9065555, 0.7740473, 0.33314514, 0.08110139, 0.40724117, 0.23223414, 0.13248764, 0.053427182, 0.72559434, 0.011427458, 0.77058077, 0.14694664, 0.07952208, 0.08960304, 0.6720478, 0.24536721, 0.42053947, 0.5573688, 0.8605512, 0.7270443, 0.2703279, 0.1314828, 0.05537432, 0.30159864, 0.26211816, 0.45614058, 0.68328136, 0.6956254, 0.28351885, 0.37992695, 0.18115096, 0.7885455, 0.056848075, 0.6969972, 0.7786954, 0.7774076, 0.25942257, 0.37381315, 0.58759964, 0.2728219, 0.3708528, 0.19705428, 0.45985588, 0.0446123, 0.79979587, 0.07695644, 0.5188351, 0.3068101, 0.57754296, 0.9594333, 0.6455702, 0.035362437, 0.43040243, 0.51001686, 0.5361775, 0.6813925, 0.2775961, 0.12886056, 0.39267567, 0.9564057, 0.1871309, 0.90398395, 0.54380596, 0.4569114, 0.8820414, 0.45860395, 0.72416764, 0.39902532, 0.9040444, 0.69002503, 0.69962204, 0.3277204, 0.75677866, 0.6360611, 0.24002028, 0.16053882, 0.7963915, 0.9591666, 0.45813882, 0.59098417, 0.85772264, 0.45722345, 0.9518745, 0.5757512, 0.8207671, 0.9088437, 0.8155238, 0.15941447, 0.62889844, 0.39843425, 0.06271295, 0.42403224, 0.25868407, 0.8490383, 0.033304628, 0.9589827, 0.35536885, 0.3567069, 0.016328502, 0.18523233, 0.4012595, 0.9292914, 0.09961493, 0.94530153, 0.86948854, 0.4541624, 0.3267009, 0.23274413, 0.6144647, 0.03307459, 0.015606064, 0.42879573, 0.06807408, 0.251941, 0.22116092, 0.2531912, 0.13105524, 0.012036223, 0.1154843, 0.61848027, 0.9742562, 0.990345, 0.4090541, 0.16295442, 0.63876176, 0.49030533, 0.9894098, 0.065304205, 0.7832344, 0.2883985, 0.24141861, 0.66250455, 0.24606319, 0.6658591, 0.51730853, 0.42408898, 0.5546878, 0.28705153, 0.7065747, 0.41485688, 0.36054555, 0.8286569, 0.92496693, 0.04600731, 0.23262699, 0.34851936, 0.8149665, 0.98549145, 0.9689717, 0.90494835, 0.29655626, 0.99201125, 0.24942005, 0.10590615, 0.9509526, 0.23342025, 0.68976825, 0.05835636, 0.7307091, 0.8817202, 0.2724369, 0.3790569, 0.3742962, 0.74878824, 0.23780724, 0.1718531, 0.44929165, 0.3044684, 0.8391891, 0.23774183, 0.50238943, 0.9425836, 0.6339977, 0.8672894, 0.9402097, 0.75076485, 0.69957507, 0.96796554, 0.9944008, 0.45182168, 0.07086978, 0.29279402, 0.1523547, 0.41748637, 0.13128933, 0.6041178, 0.38280806, 0.89538586, 0.96779466, 0.5468849, 0.27482358, 0.59223044, 0.8967612, 0.40673333, 0.55207825, 0.27165276, 0.45544416, 0.40171355, 0.24841346, 0.5058664, 0.31038082, 0.37303486, 0.5249705, 0.75059503, 0.33350748, 0.92415875, 0.8623186, 0.048690297, 0.25364253, 0.44613552, 0.10462789, 0.348476, 0.7400975, 0.68051445, 0.6223844, 0.7105284, 0.20492369, 0.3416981, 0.6762425, 0.8792348, 0.54367805, 0.28269964, 0.030235259, 0.7103368, 0.007884104, 0.37267908, 0.5305372, 0.92211145, 0.08949455, 0.40594232, 0.0243132, 0.34261099, 0.62223107, 0.27906793, 0.20974995, 0.11570323, 0.5771403, 0.69527, 0.67195714, 0.948861}),
			),

			tensor.New(
				tensor.WithShape(3, 6, 6),
				tensor.WithBacking([]int32{1, 4, 0, 0, 2, 3, 1, 0, 0, 0, 0, 2, 4, 0, 2, 3, 4, 3, 4, 4, 4, 2, 1, 3, 4, 4, 0, 0, 2, 4, 3, 3, 0, 0, 1, 4, 4, 0, 0, 0, 1, 1, 1, 4, 1, 4, 0, 4, 3, 2, 0, 3, 2, 3, 2, 2, 1, 3, 3, 0, 4, 3, 3, 0, 0, 1, 4, 4, 1, 0, 0, 0, 3, 4, 4, 2, 2, 3, 3, 4, 2, 1, 0, 4, 1, 3, 0, 3, 2, 1, 0, 0, 3, 4, 1, 0, 0, 2, 1, 1, 0, 1, 1, 0, 0, 3, 3, 3}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]float32{-58.83849}),
			),
		},
	}
}
