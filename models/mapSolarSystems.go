// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    mapSolarSystems, err := UnmarshalMapSolarSystems(bytes)
//    bytes, err = mapSolarSystems.Marshal()

package main

import "encoding/json"

func UnmarshalMapSolarSystems(data []byte) (MapSolarSystems, error) {
	var r MapSolarSystems
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MapSolarSystems) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MapSolarSystems struct {
	Key             int64      `json:"_key"`
	Border          bool       `json:"border"`
	ConstellationID int64      `json:"constellationID"`
	Hub             bool       `json:"hub"`
	International   bool       `json:"international"`
	Luminosity      float64    `json:"luminosity"`
	Name            Name       `json:"name"`
	PlanetIDs       []int64    `json:"planetIDs"`
	Position        Position   `json:"position"`
	Position2D      Position2D `json:"position2D"`
	Radius          float64    `json:"radius"`
	RegionID        int64      `json:"regionID"`
	Regional        bool       `json:"regional"`
	SecurityClass   string     `json:"securityClass"`
	SecurityStatus  float64    `json:"securityStatus"`
	StarID          int64      `json:"starID"`
	StargateIDs     []int64    `json:"stargateIDs"`
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

type Position2D struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
