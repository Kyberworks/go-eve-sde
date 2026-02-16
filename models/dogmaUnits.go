// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dogmaUnits, err := UnmarshalDogmaUnits(bytes)
//    bytes, err = dogmaUnits.Marshal()

package main

import "encoding/json"

func UnmarshalDogmaUnits(data []byte) (DogmaUnits, error) {
	var r DogmaUnits
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DogmaUnits) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DogmaUnits struct {
	Key         int64       `json:"_key"`
	Description Description `json:"description"`
	DisplayName Description `json:"displayName"`
	Name        string      `json:"name"`
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
