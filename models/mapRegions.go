// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    mapRegions, err := UnmarshalMapRegions(bytes)
//    bytes, err = mapRegions.Marshal()

package main

import "encoding/json"

func UnmarshalMapRegions(data []byte) (MapRegions, error) {
	var r MapRegions
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MapRegions) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MapRegions struct {
	Key              int64       `json:"_key"`
	ConstellationIDs []int64     `json:"constellationIDs"`
	Description      Description `json:"description"`
	FactionID        int64       `json:"factionID"`
	Name             Description `json:"name"`
	NebulaID         int64       `json:"nebulaID"`
	Position         Position    `json:"position"`
	WormholeClassID  int64       `json:"wormholeClassID"`
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

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
