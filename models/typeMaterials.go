// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    typeMaterials, err := UnmarshalTypeMaterials(bytes)
//    bytes, err = typeMaterials.Marshal()

package main

import "encoding/json"

func UnmarshalTypeMaterials(data []byte) (TypeMaterials, error) {
	var r TypeMaterials
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TypeMaterials) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TypeMaterials struct {
	Key       int64      `json:"_key"`
	Materials []Material `json:"materials"`
}

type Material struct {
	MaterialTypeID int64 `json:"materialTypeID"`
	Quantity       int64 `json:"quantity"`
}
