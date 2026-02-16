// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    certificates, err := UnmarshalCertificates(bytes)
//    bytes, err = certificates.Marshal()

package main

import "encoding/json"

func UnmarshalCertificates(data []byte) (Certificates, error) {
	var r Certificates
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Certificates) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Certificates struct {
	Key            int64       `json:"_key"`
	Description    Description `json:"description"`
	GroupID        int64       `json:"groupID"`
	Name           Description `json:"name"`
	RecommendedFor []int64     `json:"recommendedFor"`
	SkillTypes     []SkillType `json:"skillTypes"`
}

type Description struct {
	De string `json:"de"`
	En string `json:"en"`
	Es string `json:"es"`
	Fr string `json:"fr"`
	Ja string `json:"ja"`
	Ko string `json:"ko"`
	Ru string `json:"ru"`
	Zh string `json:"zh"`
}

type SkillType struct {
	Key      int64 `json:"_key"`
	Advanced int64 `json:"advanced"`
	Basic    int64 `json:"basic"`
	Elite    int64 `json:"elite"`
	Improved int64 `json:"improved"`
	Standard int64 `json:"standard"`
}
