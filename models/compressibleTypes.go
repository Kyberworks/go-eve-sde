// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    compressibleTypes, err := UnmarshalCompressibleTypes(bytes)
//    bytes, err = compressibleTypes.Marshal()

package main

import "encoding/json"

func UnmarshalCompressibleTypes(data []byte) (CompressibleTypes, error) {
	var r CompressibleTypes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CompressibleTypes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CompressibleTypes struct {
	Key              int64 `json:"_key"`
	CompressedTypeID int64 `json:"compressedTypeID"`
}
