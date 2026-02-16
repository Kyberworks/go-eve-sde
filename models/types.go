// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    types, err := UnmarshalTypes(bytes)
//    bytes, err = types.Marshal()

package main

import "encoding/json"

func UnmarshalTypes(data []byte) (Types, error) {
	var r Types
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Types) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Types struct {
	Key         int64   `json:"_key"`
	GroupID     int64   `json:"groupID"`
	Mass        float64 `json:"mass"`
	Name        Name    `json:"name"`
	PortionSize int64   `json:"portionSize"`
	Published   bool    `json:"published"`
}

type Name struct {
	De string `json:"de"`
	En string `json:"en"`
	Es string `json:"es"`
	Fr string `json:"fr"`
	Ja string `json:"ja"`
	Ko string `json:"ko"`
	Ru string `json:"ru"`
	Zh string `json:"zh"`
}
