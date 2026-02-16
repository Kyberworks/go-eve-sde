// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    sde, err := UnmarshalSde(bytes)
//    bytes, err = sde.Marshal()

package main

import "time"

import "encoding/json"

func UnmarshalSde(data []byte) (Sde, error) {
	var r Sde
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Sde) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Sde struct {
	Key         string    `json:"_key"`
	BuildNumber int64     `json:"buildNumber"`
	ReleaseDate time.Time `json:"releaseDate"`
}
