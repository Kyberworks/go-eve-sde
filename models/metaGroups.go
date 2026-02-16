// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    metaGroups, err := UnmarshalMetaGroups(bytes)
//    bytes, err = metaGroups.Marshal()

package main

import "encoding/json"

func UnmarshalMetaGroups(data []byte) (MetaGroups, error) {
	var r MetaGroups
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MetaGroups) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MetaGroups struct {
	Key   int64 `json:"_key"`
	Color Color `json:"color"`
	Name  Name  `json:"name"`
}

type Color struct {
	B float64 `json:"b"`
	G float64 `json:"g"`
	R float64 `json:"r"`
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
