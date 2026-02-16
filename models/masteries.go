// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    masteries, err := UnmarshalMasteries(bytes)
//    bytes, err = masteries.Marshal()

package main

import "encoding/json"

func UnmarshalMasteries(data []byte) (Masteries, error) {
	var r Masteries
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Masteries) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Masteries struct {
	Key   int64   `json:"_key"`
	Value []Value `json:"_value"`
}

type Value struct {
	Key   int64   `json:"_key"`
	Value []int64 `json:"_value"`
}
