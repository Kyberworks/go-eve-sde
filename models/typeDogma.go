// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    typeDogma, err := UnmarshalTypeDogma(bytes)
//    bytes, err = typeDogma.Marshal()

package main

import "encoding/json"

func UnmarshalTypeDogma(data []byte) (TypeDogma, error) {
	var r TypeDogma
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TypeDogma) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TypeDogma struct {
	Key             int64            `json:"_key"`
	DogmaAttributes []DogmaAttribute `json:"dogmaAttributes"`
}

type DogmaAttribute struct {
	AttributeID int64   `json:"attributeID"`
	Value       float64 `json:"value"`
}
