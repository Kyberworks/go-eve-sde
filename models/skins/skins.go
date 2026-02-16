// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    skins, err := UnmarshalSkins(bytes)
//    bytes, err = skins.Marshal()

package skins

import "encoding/json"

func UnmarshalSkins(data []byte) (Skins, error) {
	var r Skins
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Skins) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Skins struct {
	Key                int64   `json:"_key"`
	AllowCCPDevs       bool    `json:"allowCCPDevs"`
	InternalName       string  `json:"internalName"`
	SkinMaterialID     int64   `json:"skinMaterialID"`
	Types              []int64 `json:"types"`
	VisibleSerenity    bool    `json:"visibleSerenity"`
	VisibleTranquility bool    `json:"visibleTranquility"`
}
