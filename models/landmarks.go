// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    landmarks, err := UnmarshalLandmarks(bytes)
//    bytes, err = landmarks.Marshal()

package main

import "encoding/json"

func UnmarshalLandmarks(data []byte) (Landmarks, error) {
	var r Landmarks
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Landmarks) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Landmarks struct {
	Key         int64       `json:"_key"`
	Description Description `json:"description"`
	Name        Description `json:"name"`
	Position    Position    `json:"position"`
}

type Description struct {
	De string `json:"de"`
	En string `json:"en"`
	Es string `json:"es"`
	Fr string `json:"fr"`
	Ja string `json:"ja"`
	Ko string `json:"ko"`
	Ru string `json:"ru"`
	Zh string `json:"zh"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
