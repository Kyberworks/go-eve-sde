// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    groups, err := UnmarshalGroups(bytes)
//    bytes, err = groups.Marshal()

package main

import "encoding/json"

func UnmarshalGroups(data []byte) (Groups, error) {
	var r Groups
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Groups) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Groups struct {
	Key                  int64 `json:"_key"`
	Anchorable           bool  `json:"anchorable"`
	Anchored             bool  `json:"anchored"`
	CategoryID           int64 `json:"categoryID"`
	FittableNonSingleton bool  `json:"fittableNonSingleton"`
	Name                 Name  `json:"name"`
	Published            bool  `json:"published"`
	UseBasePrice         bool  `json:"useBasePrice"`
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
