// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    typeBonus, err := UnmarshalTypeBonus(bytes)
//    bytes, err = typeBonus.Marshal()

package main

import "encoding/json"

func UnmarshalTypeBonus(data []byte) (TypeBonus, error) {
	var r TypeBonus
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TypeBonus) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TypeBonus struct {
	Key         int64       `json:"_key"`
	RoleBonuses []RoleBonus `json:"roleBonuses"`
	Types       []Type      `json:"types"`
}

type RoleBonus struct {
	Bonus      float64   `json:"bonus"`
	BonusText  BonusText `json:"bonusText"`
	Importance int64     `json:"importance"`
	UnitID     int64     `json:"unitID"`
}

type BonusText struct {
	De string `json:"de"`
	En string `json:"en"`
	Es string `json:"es"`
	Fr string `json:"fr"`
	Ja string `json:"ja"`
	Ko string `json:"ko"`
	Ru string `json:"ru"`
	Zh string `json:"zh"`
}

type Type struct {
	Key   int64       `json:"_key"`
	Value []RoleBonus `json:"_value"`
}
