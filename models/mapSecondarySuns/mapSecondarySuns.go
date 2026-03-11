// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    mapSecondarySuns, err := UnmarshalMapSecondarySuns(bytes)
//    bytes, err = mapSecondarySuns.Marshal()

package mapSecondarySuns

import "encoding/json"

func UnmarshalMapSecondarySuns(data []byte) (MapSecondarySuns, error) {
	var r MapSecondarySuns
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MapSecondarySuns) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MapSecondarySuns struct {
	Key                int64    `json:"_key"`
	EffectBeaconTypeID int64    `json:"effectBeaconTypeID"`
	Position           Position `json:"position"`
	SolarSystemID      int64    `json:"solarSystemID"`
	TypeID             int64    `json:"typeID"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
