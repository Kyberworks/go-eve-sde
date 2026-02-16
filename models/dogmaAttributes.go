// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dogmaAttributes, err := UnmarshalDogmaAttributes(bytes)
//    bytes, err = dogmaAttributes.Marshal()

package main

import "encoding/json"

func UnmarshalDogmaAttributes(data []byte) (DogmaAttributes, error) {
	var r DogmaAttributes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DogmaAttributes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DogmaAttributes struct {
	Key                 int64   `json:"_key"`
	AttributeCategoryID int64   `json:"attributeCategoryID"`
	DataType            int64   `json:"dataType"`
	DefaultValue        float64 `json:"defaultValue"`
	Description         string  `json:"description"`
	DisplayWhenZero     bool    `json:"displayWhenZero"`
	HighIsGood          bool    `json:"highIsGood"`
	Name                string  `json:"name"`
	Published           bool    `json:"published"`
	Stackable           bool    `json:"stackable"`
}
