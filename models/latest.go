// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    latest, err := UnmarshalLatest(bytes)
//    bytes, err = latest.Marshal()

package main

import "time"

import "encoding/json"

func UnmarshalLatest(data []byte) (Latest, error) {
	var r Latest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Latest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Latest struct {
	Key         string    `json:"_key"`
	BuildNumber int64     `json:"buildNumber"`
	ReleaseDate time.Time `json:"releaseDate"`
}
