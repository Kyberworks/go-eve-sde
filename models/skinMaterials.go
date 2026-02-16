// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    skinMaterials, err := UnmarshalSkinMaterials(bytes)
//    bytes, err = skinMaterials.Marshal()

package main

import "encoding/json"

func UnmarshalSkinMaterials(data []byte) (SkinMaterials, error) {
	var r SkinMaterials
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkinMaterials) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SkinMaterials struct {
	Key           int64       `json:"_key"`
	DisplayName   DisplayName `json:"displayName"`
	MaterialSetID int64       `json:"materialSetID"`
}

type DisplayName struct {
	De string `json:"de"`
	En string `json:"en"`
	Es string `json:"es"`
	Fr string `json:"fr"`
	Ja string `json:"ja"`
	Ko string `json:"ko"`
	Ru string `json:"ru"`
	Zh string `json:"zh"`
}
