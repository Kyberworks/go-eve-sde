// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    mapStars, err := UnmarshalMapStars(bytes)
//    bytes, err = mapStars.Marshal()

package main

import "encoding/json"

func UnmarshalMapStars(data []byte) (MapStars, error) {
	var r MapStars
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MapStars) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MapStars struct {
	Key           int64      `json:"_key"`
	Radius        int64      `json:"radius"`
	SolarSystemID int64      `json:"solarSystemID"`
	Statistics    Statistics `json:"statistics"`
	TypeID        int64      `json:"typeID"`
}

type Statistics struct {
	Age           float64 `json:"age"`
	Life          float64 `json:"life"`
	Luminosity    float64 `json:"luminosity"`
	SpectralClass string  `json:"spectralClass"`
	Temperature   float64 `json:"temperature"`
}
