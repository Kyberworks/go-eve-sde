// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    mapStargates, err := UnmarshalMapStargates(bytes)
//    bytes, err = mapStargates.Marshal()

package mapStargates

import "encoding/json"

func UnmarshalMapStargates(data []byte) (MapStargates, error) {
	var r MapStargates
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MapStargates) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MapStargates struct {
	Key           int64       `json:"_key"`
	Destination   Destination `json:"destination"`
	Position      Position    `json:"position"`
	SolarSystemID int64       `json:"solarSystemID"`
	TypeID        int64       `json:"typeID"`
}

type Destination struct {
	SolarSystemID int64 `json:"solarSystemID"`
	StargateID    int64 `json:"stargateID"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
