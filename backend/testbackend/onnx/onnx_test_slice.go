package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Slice", "TestSlice", NewTestSlice)
}

/*
&ir.ModelProto{
    IrVersion:   4,
    OpsetImport: {
        &ir.OperatorSetIdProto{Domain:"", Version:10},
    },
    ProducerName:    "backend-test",
    ProducerVersion: "",
    Domain:          "",
    ModelVersion:    0,
    DocString:       "",
    Graph:           &ir.GraphProto{
        Node: {
            &ir.NodeProto{
                Input:     {"x", "starts", "ends", "axes", "steps"},
                Output:    {"y"},
                Name:      "",
                OpType:    "Slice",
                Domain:    "",
                Attribute: nil,
                DocString: "",
            },
        },
        Name:              "test_slice",
        Initializer:       nil,
        SparseInitializer: nil,
        DocString:         "",
        Input:             {
            &ir.ValueInfoProto{
                Name: "x",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:20},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:10},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:5},
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
                Name: "starts",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 7,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
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
                Name: "ends",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 7,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
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
                Name: "axes",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 7,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
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
                Name: "steps",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 7,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
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
                Name: "y",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:10},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:5},
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
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestSlice version: 4.
func NewTestSlice() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Slice",
		Title:  "TestSlice",
		ModelB: []byte{0x8, 0x4, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xbb, 0x1, 0xa, 0x28, 0xa, 0x1, 0x78, 0xa, 0x6, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0xa, 0x4, 0x65, 0x6e, 0x64, 0x73, 0xa, 0x4, 0x61, 0x78, 0x65, 0x73, 0xa, 0x5, 0x73, 0x74, 0x65, 0x70, 0x73, 0x12, 0x1, 0x79, 0x22, 0x5, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0xa, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x14, 0xa, 0x2, 0x8, 0xa, 0xa, 0x2, 0x8, 0x5, 0x5a, 0x14, 0xa, 0x6, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x5a, 0x12, 0xa, 0x4, 0x65, 0x6e, 0x64, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x5a, 0x12, 0xa, 0x4, 0x61, 0x78, 0x65, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x5a, 0x13, 0xa, 0x5, 0x73, 0x74, 0x65, 0x70, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x62, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0xa, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0xa},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x", "starts", "ends", "axes", "steps"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Slice",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(20, 10, 5),
				tensor.WithBacking([]float32{1.7640524, 0.4001572, 0.978738, 2.2408931, 1.867558, -0.9772779, 0.95008844, -0.1513572, -0.10321885, 0.41059852, 0.14404356, 1.4542735, 0.7610377, 0.121675014, 0.44386324, 0.33367434, 1.4940791, -0.20515826, 0.3130677, -0.85409576, -2.5529897, 0.6536186, 0.8644362, -0.742165, 2.2697546, -1.4543657, 0.045758516, -0.18718386, 1.5327792, 1.4693588, 0.15494743, 0.37816253, -0.88778573, -1.9807965, -0.34791216, 0.15634897, 1.2302907, 1.2023798, -0.3873268, -0.30230275, -1.048553, -1.420018, -1.7062702, 1.9507754, -0.5096522, -0.4380743, -1.2527953, 0.7774904, -1.6138978, -0.21274029, -0.89546657, 0.3869025, -0.51080513, -1.1806322, -0.028182229, 0.42833188, 0.06651722, 0.3024719, -0.6343221, -0.36274117, -0.67246044, -0.35955316, -0.8131463, -1.7262826, 0.17742614, -0.40178093, -1.6301984, 0.46278226, -0.9072984, 0.051945396, 0.7290906, 0.12898292, 1.1394007, -1.2348258, 0.40234163, -0.6848101, -0.87079716, -0.5788497, -0.31155252, 0.05616534, -1.1651498, 0.9008265, 0.46566245, -1.5362437, 1.4882522, 1.8958892, 1.1787796, -0.17992483, -1.0707526, 1.0544517, -0.40317693, 1.222445, 0.20827498, 0.97663903, 0.3563664, 0.7065732, 0.01050002, 1.7858706, 0.12691209, 0.40198937, 1.8831507, -1.347759, -1.270485, 0.9693967, -1.1731234, 1.9436212, -0.41361898, -0.7474548, 1.922942, 1.4805148, 1.867559, 0.90604466, -0.86122566, 1.9100649, -0.26800337, 0.8024564, 0.947252, -0.15501009, 0.61407936, 0.9222067, 0.37642553, -1.0994008, 0.2982382, 1.3263859, -0.69456786, -0.14963454, -0.43515354, 1.8492638, 0.67229474, 0.40746182, -0.76991606, 0.5392492, -0.6743327, 0.031830557, -0.6358461, 0.67643327, 0.57659084, -0.20829876, 0.3960067, -1.0930616, -1.4912575, 0.4393917, 0.1666735, 0.63503146, 2.3831449, 0.94447947, -0.91282225, 1.1170163, -1.3159074, -0.4615846, -0.0682416, 1.7133427, -0.74475485, -0.82643855, -0.09845252, -0.6634783, 1.1266359, -1.0799315, -1.1474687, -0.43782005, -0.49803245, 1.929532, 0.9494208, 0.08755124, -1.2254355, 0.844363, -1.0002153, -1.5447711, 1.1880298, 0.3169426, 0.9208588, 0.31872764, 0.8568306, -0.6510256, -1.0342429, 0.6815945, -0.80340964, -0.6895498, -0.4555325, 0.017479159, -0.35399392, -1.3749512, -0.6436184, -2.2234032, 0.62523144, -1.6020577, -1.1043833, 0.05216508, -0.739563, 1.5430146, -1.2928569, 0.26705086, -0.039282817, -1.1680934, 0.5232767, -0.17154633, 0.77179056, 0.82350415, 2.163236, 1.336528, -0.36918184, -0.23937918, 1.0996596, 0.6552637, 0.64013153, -1.616956, -0.024326125, -0.7380309, 0.2799246, -0.09815039, 0.9101789, 0.3172182, 0.78632796, -0.4664191, -0.94444627, -0.4100497, -0.017020414, 0.37915173, 2.259309, -0.042257152, -0.955945, -0.34598178, -0.463596, 0.48148146, -1.540797, 0.06326199, 0.15650654, 0.23218104, -0.5973161, -0.23792173, -1.424061, -0.49331987, -0.54286146, 0.41605005, -1.1561824, 0.7811981, 1.4944845, -2.069985, 0.42625874, 0.676908, -0.63743705, -0.3972718, -0.13288058, -0.29779088, -0.30901298, -1.6760038, 1.1523316, 1.0796186, -0.81336427, -1.4664243, 0.5210649, -0.57578796, 0.14195317, -0.31932843, 0.69153875, 0.6947491, -0.7255974, -1.383364, -1.5829384, 0.6103794, -1.1888592, -0.5068163, -0.596314, -0.052567296, -1.9362798, 0.1887786, 0.52389103, 0.08842209, -0.31088617, 0.097400166, 0.39904633, -2.7725928, 1.9559124, 0.39009333, -0.6524086, -0.39095336, 0.49374178, -0.11610394, -2.0306845, 2.064493, -0.11054066, 1.0201727, -0.69204986, 1.5363771, 0.2863437, 0.60884386, -1.0452534, 1.2111453, 0.68981814, 1.3018463, -0.6280876, -0.48102713, 2.3039167, -1.0600158, -0.1359497, 1.1368914, 0.09772497, 0.5829537, -0.39944902, 0.37005588, -1.3065269, 1.6581306, -0.11816405, -0.6801782, 0.6663831, -0.4607198, -1.3342584, -1.3467175, 0.69377315, -0.15957344, -0.13370156, 1.0777438, -1.1268258, -0.7306777, -0.3848798, 0.09435159, -0.042171452, -0.2868872, -0.0616264, -0.10730527, -0.7196044, -0.812993, 0.27451634, -0.8909151, -1.1573553, -0.31229225, -0.15766701, 2.2567234, -0.7047003, 0.9432607, 0.7471883, -1.1889449, 0.77325296, -1.1838807, -2.6591723, 0.60631955, -1.7558906, 0.45093447, -0.6840109, 1.6595508, 1.0685093, -0.4533858, -0.6878376, -1.2140774, -0.44092262, -0.28035548, -0.36469355, 0.15670386, 0.5785215, 0.34965447, -0.76414394, -1.4377915, 1.3645319, -0.6894492, -0.6522936, -0.52118933, -1.8430696, -0.477974, -0.4796558, 0.6203583, 0.6984571, 0.003770889, 0.93184835, 0.339965, -0.015682112, 0.16092817, -0.19065349, -0.3948495, -0.26773354, -1.1280113, 0.2804417, -0.9931236, 0.8416313, -0.24945858, 0.04949498, 0.4938368, 0.6433145, -1.5706234, -0.20690368, 0.8801789, -1.6981058, 0.38728046, -2.2555642, -1.0225068, 0.038630553, -1.6567152, -0.98551077, -1.471835, 1.648135, 0.16422775, 0.5672903, -0.2226751, -0.35343176, -1.6164742, -0.29183736, -0.7614922, 0.8579239, 1.1411018, 1.4665787, 0.85255194, -0.5986539, -1.1158969, 0.7666632, 0.3562928, -1.7685385, 0.3554818, 0.8145198, 0.058925588, -0.18505368, -0.8076485, -1.4465348, 0.800298, -0.30911446, -0.23346666, 1.7327212, 0.6845011, 0.370825, 0.1420618, 1.5199949, 1.7195894, 0.9295051, 0.5822246, -2.094603, 0.12372191, -0.13010696, 0.09395323, 0.9430461, -2.7396772, -0.56931204, 0.26990435, -0.46684554, -1.4169061, 0.8689635, 0.27687192, -0.97110456, 0.3148172, 0.8215857, 0.005292646, 0.8005648, 0.078260176, -0.39522898, -1.1594205, -0.085930765, 0.19429293, 0.87583274, -0.11510747, 0.4574156, -0.964612, -0.78262913, -0.1103893, -1.0546285, 0.8202478, 0.46313033, 0.27909577, 0.3389041, 2.0210435, -0.4688642, -2.2014413, 0.1993002, -0.050603542, -0.51751906, -0.97882986, -0.43918952, 0.18133843, -0.5028167, 2.4124537, -0.96050435, -0.79311734, -2.28862, 0.25148442, -2.0164065, -0.53945464, -0.27567053, -0.70972794, 1.7388726, 0.99439436, 1.3191369, -0.8824188, 1.128594, 0.49600095, 0.77140594, 1.0294389, -0.90876323, -0.42431763, 0.86259604, -2.6556191, 1.5133281, 0.55313206, -0.045703962, 0.22050765, -1.0299352, -0.34994337, 1.1002843, 1.298022, 2.696224, -0.07392467, -0.65855294, -0.51423395, -1.0180418, -0.07785475, 0.38273242, -0.03424228, 1.0963469, -0.2342158, -0.34745064, -0.5812685, -1.6326345, -1.5677677, -1.179158, 1.3014281, 0.8952603, 1.3749641, -1.3322116, -1.9686247, -0.6600563, 0.17581895, 0.49869028, 1.0479722, 0.28427967, 1.7426687, -0.22260568, -0.9130792, -1.6812183, -0.8889713, 0.24211796, -0.8887203, 0.9367425, 1.4123276, -2.369587, 0.8640523, -2.239604, 0.40149906, 1.2248706, 0.064856105, -1.2796892, -0.5854312, -0.26164544, -0.18224478, -0.20289683, -0.10988278, 0.21348006, -1.2085737, -0.24201983, 1.5182612, -0.38464543, -0.4438361, 1.0781974, -2.5591846, 1.1813786, -0.63190377, 0.16392857, 0.09632136, 0.9424681, -0.26759475, -0.6780258, 1.2978458, -2.364174, 0.020334182, -1.3479254, -0.7615734, 2.0112567, -0.044595428, 0.1950697, -1.7815628, -0.7290447, 0.1965574, 0.3547577, 0.61688656, 0.008627899, 0.5270042, 0.4537819, -1.8297404, 0.037005723, 0.76790243, 0.5898798, -0.36385882, -0.8056265, -1.1183119, -0.13105401, 1.1330799, -1.9518042, -0.6598917, -1.1398025, 0.7849575, -0.5543096, -0.47063765, -0.21694957, 0.44539326, -0.392389, -3.046143, 0.5433119, 0.43904296, -0.21954103, -1.0840366, 0.35178012, 0.37923554, -0.47003287, -0.21673147, -0.9301565, -0.17858909, -1.5504293, 0.41731882, -0.9443685, 0.23810315, -1.405963, -0.5900577, -0.110489406, -1.6606998, 0.115147874, -0.37914756, -1.7423562, -1.3032428, 0.60512006, 0.895556, -0.13190864, 0.40476182, 0.22384356, 0.32962298, 1.285984, -1.5069984, 0.67646074, -0.38200897, -0.22425893, -0.30224973, -0.3751471, -1.2261962, 0.1833392, 1.670943, -0.05613302, -0.0013850428, -0.687299, -0.11747455, 0.46616644, -0.37024245, -0.45380405, 0.40326455, -0.91800475, 0.25249663, 0.8203218, 1.3599485, -0.09038201, 1.3675972, 1.0344099, -0.99621266, -1.2179385, -0.30496365, 1.0289356, -0.07228701, -0.6006576, 1.5522432, 0.28690448, -2.3205943, 0.31716064, 0.52004063, 0.22560866, 0.4497121, -0.067275606, -1.3183959, -0.370704, -0.94561577, -0.9327409, -1.2630683, 0.45248908, 0.097896144, -0.44816536, -0.64933795, -0.023423105, 1.0791948, -2.0042157, 0.37687653, -0.545712, -1.8845859, -1.945703, -0.9127835, 0.21950956, 0.39306292, -0.9389816, 1.017021, 1.4229835, 0.39608657, -0.59140265, 1.1244192, 0.7553957, 0.86740744, -0.6564637, -2.8345544, 2.116791, -1.6108783, -0.035768073, 2.3807454, 0.33057675, 0.94924647, -1.5023966, -1.7776669, -0.5327028, 1.0907497, -0.34624946, -0.7946363, 0.19796729, 1.0819352, -1.4449402, -1.210543, -0.7886692, 1.0946383, 0.23482153, 2.1321535, 0.9364457, -0.035095178, 1.2650778, 0.21149701, -0.70492136, 0.67997485, -0.6963267, -0.2903971, 1.3277828, -0.10128149, -0.8031414, -0.46433768, 1.0217906, -0.55254066, -0.38687086, -0.51029277, 0.1839255, -0.38548976, -1.6018361, -0.8871809, -0.932789, 1.2433194, 0.81267405, 0.58725935, -0.50535834, -0.81579155, -0.5075176, -1.0518801, 2.4972005, -2.2453218, 0.56400853, -1.2845523, -0.10434349, -0.98800194, -1.177629, -1.1401963, 1.7549862, -0.13298842, -0.7657022, 0.55578697, 0.010349315, 0.72003376, -1.8242567, 0.30360392, 0.7726948, -1.6615983, 0.44819528, 1.6961815, -0.014857704, 0.82140595, 0.67057043, -0.7075057, 0.039766736, -1.5669947, -0.45130304, 0.26568797, 0.7231005, 0.024612125, 0.71998376, -1.1029062, -0.10169727, 0.019279385, 1.8495913, -0.21416666, -0.49901664, 0.021351224, -0.91911346, 0.19275385, -0.3650552, -1.7913276, -0.058586553, -0.3175431, -1.6324233, -0.06713416, 1.4893559, 0.5213038, 0.6119272, -1.3414967, 0.47689837, 0.14844958, 0.5290452, 0.4226286, -1.3597807, -0.041400813, -0.75787085, -0.050084095, -0.8974009, 1.3124703, -0.8589724, -0.8989422, 0.07458641, -1.0770991, -0.4246633, -0.8299646, 1.411172, 0.78580385, -0.057469517, -0.39121705, 0.9409176, 0.4052041, 0.49805242, -0.026192237, -1.68823, -0.112465985, -0.5324899, 0.6450553, 1.0118425, -0.65795106, 0.46838522, 1.735879, -0.66771275, 1.6819217, -0.85258585, 0.022959756, -0.011145612, 0.0114989, -0.837678, -0.5911831, -0.66772026, 0.3269626, 0.33003512, 2.2259443, 1.370989, -0.50984323, 0.3248696, 0.997118, 0.030601824, -0.069641575, 0.05157494, 0.8672766, -0.84832054, -0.32566947, 0.47043315, 0.31144708, 0.23958276, -0.36980116, 0.9725358, 2.1338682, 0.4064155, -0.1931767, 0.7557403, -0.53913265, -0.74969035, 0.032808747, -2.5827966, -1.1539503, -0.34796184, -1.3533889, -1.0326431, -0.43674833, -1.6429653, -0.40607178, -0.53527015, 0.025405208, 1.154184, 0.17250441, 0.021062022, 0.099454455, 0.22739278, -1.0167387, -0.11477532, 0.30875126, -1.37076, 0.8656529, 1.0813761, -0.63137597, -0.24133779, -0.87819034, 0.69938046, -1.0612223, -0.222477, -0.8589199, 0.05095428, -1.7942293, 1.3264617, -0.9646064, 0.059894685, -0.21252304, -0.7621145, -0.88778013, 0.93639857, -0.5256406, 0.2711702, -0.80149686, -0.64718145, 0.47224715, 0.9304085, -0.17531641, -1.4219198, 1.997956, -0.8565493, -1.5415874, 2.5944245, -0.4040323, -1.4617327, -0.6834398, 0.3675449, 0.19031155, -0.8517292, 1.8227236, -0.5215797, -1.1846865, 0.9606934, 1.3290628, -0.8174931, -1.4013473, 1.0304383, -2.0473237, -1.2266216, 0.96744615, -0.055352546, -0.26393735, 0.3528166, -0.15277442, -1.2986867, 1.2760754, 1.325014, 0.20533256, 0.045134015, 2.339625, -0.27643284, -0.25957698, 0.36448124, 1.471322, 1.5927707, -0.25857264, 0.30833125, -1.3780835, -0.3119761, -0.84029037, -1.0068318, 1.6815767, -0.79228663, -0.5316059, 0.36584878, 1.2978252, 0.48111513, 2.759355, -0.074667975, 0.25871643, 0.27560067, 1.4350494, 0.5072389, -0.1162297, -0.9474886, 0.24444346, 1.4013448, -0.4103818, 0.5289436, 0.24614778, 0.86351967, -0.8047537, 2.346647, -1.2791611, -0.36555108, 0.9380925, 0.29673317, 0.82998616, -0.49610233, -0.074804984, 0.012231983, 1.5692596, 0.69042903, 0.7966721, -0.6579261, 0.9688826, 0.22558166, 1.3891454, 2.0140603, -0.30676576, -0.40630314, -0.86404496, -0.14357951, -0.38202545, 0.3595044, -0.14456682, -0.36159927, 1.0645851, -0.9378802, 0.43310794, -0.40594172, 0.7243685, 1.3852615, -0.30309826, 0.44103292, 0.17879286, -0.7994224, 0.2407875, 0.2891205, 0.41287082, -0.1983989, 0.0941923, -1.1476109, -0.35811406}),
			),

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]int64{0, 0}),
			),

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]int64{3, 10}),
			),

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]int64{0, 1}),
			),

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]int64{1, 1}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 10, 5),
				tensor.WithBacking([]float32{1.7640524, 0.4001572, 0.978738, 2.2408931, 1.867558, -0.9772779, 0.95008844, -0.1513572, -0.10321885, 0.41059852, 0.14404356, 1.4542735, 0.7610377, 0.121675014, 0.44386324, 0.33367434, 1.4940791, -0.20515826, 0.3130677, -0.85409576, -2.5529897, 0.6536186, 0.8644362, -0.742165, 2.2697546, -1.4543657, 0.045758516, -0.18718386, 1.5327792, 1.4693588, 0.15494743, 0.37816253, -0.88778573, -1.9807965, -0.34791216, 0.15634897, 1.2302907, 1.2023798, -0.3873268, -0.30230275, -1.048553, -1.420018, -1.7062702, 1.9507754, -0.5096522, -0.4380743, -1.2527953, 0.7774904, -1.6138978, -0.21274029, -0.89546657, 0.3869025, -0.51080513, -1.1806322, -0.028182229, 0.42833188, 0.06651722, 0.3024719, -0.6343221, -0.36274117, -0.67246044, -0.35955316, -0.8131463, -1.7262826, 0.17742614, -0.40178093, -1.6301984, 0.46278226, -0.9072984, 0.051945396, 0.7290906, 0.12898292, 1.1394007, -1.2348258, 0.40234163, -0.6848101, -0.87079716, -0.5788497, -0.31155252, 0.05616534, -1.1651498, 0.9008265, 0.46566245, -1.5362437, 1.4882522, 1.8958892, 1.1787796, -0.17992483, -1.0707526, 1.0544517, -0.40317693, 1.222445, 0.20827498, 0.97663903, 0.3563664, 0.7065732, 0.01050002, 1.7858706, 0.12691209, 0.40198937, 1.8831507, -1.347759, -1.270485, 0.9693967, -1.1731234, 1.9436212, -0.41361898, -0.7474548, 1.922942, 1.4805148, 1.867559, 0.90604466, -0.86122566, 1.9100649, -0.26800337, 0.8024564, 0.947252, -0.15501009, 0.61407936, 0.9222067, 0.37642553, -1.0994008, 0.2982382, 1.3263859, -0.69456786, -0.14963454, -0.43515354, 1.8492638, 0.67229474, 0.40746182, -0.76991606, 0.5392492, -0.6743327, 0.031830557, -0.6358461, 0.67643327, 0.57659084, -0.20829876, 0.3960067, -1.0930616, -1.4912575, 0.4393917, 0.1666735, 0.63503146, 2.3831449, 0.94447947, -0.91282225, 1.1170163, -1.3159074, -0.4615846}),
			),
		},
	}
}
