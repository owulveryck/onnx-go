// This file contains specific handler to run the demos

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/x/gorgonnx"
)

func init() {
	http.HandleFunc("/image", imagePostHandler)
	http.HandleFunc("/model", modelPostHandler)
}

type img struct {
	Image string
	Model string `json:"model"`
}

func imagePostHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var payload img
	err := dec.Decode(&payload)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	var res results
	if md, ok := models[payload.Model]; ok {
		img, err := processPicture(payload.Image, md.height, md.width)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}
		res, err = process(img, md.height, md.width, md.table)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	fmt.Printf("%v", res)
	fmt.Fprintf(w, "%v", res)
}
func modelPostHandler(w http.ResponseWriter, r *http.Request) {
	// reset the backend and the model
	// Create a backend receiver
	backend = gorgonnx.NewGraph()
	// Create a model and set the execution backend
	model = onnx.NewModel(backend)
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
	err = model.UnmarshalBinary(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
