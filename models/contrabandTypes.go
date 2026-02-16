// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    contrabandTypes, err := UnmarshalContrabandTypes(bytes)
//    bytes, err = contrabandTypes.Marshal()

package main

import "encoding/json"

func UnmarshalContrabandTypes(data []byte) (ContrabandTypes, error) {
	var r ContrabandTypes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContrabandTypes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ContrabandTypes struct {
	Key      int64     `json:"_key"`
	Factions []Faction `json:"factions"`
}

type Faction struct {
	Key              int64   `json:"_key"`
	AttackMinSEC     float64 `json:"attackMinSec"`
	ConfiscateMinSEC float64 `json:"confiscateMinSec"`
	FineByValue      float64 `json:"fineByValue"`
	StandingLoss     float64 `json:"standingLoss"`
}
