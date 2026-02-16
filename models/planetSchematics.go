// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    planetSchematics, err := UnmarshalPlanetSchematics(bytes)
//    bytes, err = planetSchematics.Marshal()

package main

import "encoding/json"

func UnmarshalPlanetSchematics(data []byte) (PlanetSchematics, error) {
	var r PlanetSchematics
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlanetSchematics) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PlanetSchematics struct {
	Key       int64   `json:"_key"`
	CycleTime int64   `json:"cycleTime"`
	Name      Name    `json:"name"`
	Pins      []int64 `json:"pins"`
	Types     []Type  `json:"types"`
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

type Type struct {
	Key      int64 `json:"_key"`
	IsInput  bool  `json:"isInput"`
	Quantity int64 `json:"quantity"`
}
