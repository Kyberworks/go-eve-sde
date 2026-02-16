// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    mapConstellations, err := UnmarshalMapConstellations(bytes)
//    bytes, err = mapConstellations.Marshal()

package main

import "encoding/json"

func UnmarshalMapConstellations(data []byte) (MapConstellations, error) {
	var r MapConstellations
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MapConstellations) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MapConstellations struct {
	Key             int64    `json:"_key"`
	FactionID       int64    `json:"factionID"`
	Name            Name     `json:"name"`
	Position        Position `json:"position"`
	RegionID        int64    `json:"regionID"`
	SolarSystemIDs  []int64  `json:"solarSystemIDs"`
	WormholeClassID int64    `json:"wormholeClassID"`
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

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
