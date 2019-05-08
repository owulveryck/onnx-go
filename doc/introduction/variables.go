package main

import (
	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
)

var (
	backend    *gorgonnx.Graph
	model      *onnx.Model
	mnistTable = []string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}
	emotionTable = []string{
		"neutral",
		"happiness",
		"surprise",
		"sadness",
		"anger",
		"disgust",
		"fear",
		"contempt",
	}
	models = map[string]modelDemo{
		"mnist": modelDemo{
			height: 28,
			width:  28,
			table:  mnistTable,
		},
		"emotion": modelDemo{
			height: 64,
			width:  64,
			table:  emotionTable,
		},
	}
)

type modelDemo struct {
	height int
	width  int
	table  []string
}
