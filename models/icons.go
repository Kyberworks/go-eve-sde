// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    icons, err := UnmarshalIcons(bytes)
//    bytes, err = icons.Marshal()

package main

import "encoding/json"

func UnmarshalIcons(data []byte) (Icons, error) {
	var r Icons
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Icons) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Icons struct {
	Key      int64  `json:"_key"`
	IconFile string `json:"iconFile"`
}
