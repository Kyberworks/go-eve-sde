// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    graphics, err := UnmarshalGraphics(bytes)
//    bytes, err = graphics.Marshal()

package main

import "encoding/json"

func UnmarshalGraphics(data []byte) (Graphics, error) {
	var r Graphics
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Graphics) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Graphics struct {
	Key         int64  `json:"_key"`
	GraphicFile string `json:"graphicFile"`
}
