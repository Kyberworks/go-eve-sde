// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    corporationActivities, err := UnmarshalCorporationActivities(bytes)
//    bytes, err = corporationActivities.Marshal()

package main

import "encoding/json"

func UnmarshalCorporationActivities(data []byte) (CorporationActivities, error) {
	var r CorporationActivities
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CorporationActivities) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CorporationActivities struct {
	Key  int64 `json:"_key"`
	Name Name  `json:"name"`
}

type Name struct {
	De string `json:"de"`
	En string `json:"en"`
	Es string `json:"es"`
	Fr string `json:"fr"`
	Ja string `json:"ja"`
	Ko string `json:"ko"`
	Ru string `json:"ru"`
	Zh string `json:"zh"`
}
