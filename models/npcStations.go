// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    npcStations, err := UnmarshalNpcStations(bytes)
//    bytes, err = npcStations.Marshal()

package main

import "encoding/json"

func UnmarshalNpcStations(data []byte) (NpcStations, error) {
	var r NpcStations
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NpcStations) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type NpcStations struct {
	Key                      int64    `json:"_key"`
	CelestialIndex           int64    `json:"celestialIndex"`
	OperationID              int64    `json:"operationID"`
	OrbitID                  int64    `json:"orbitID"`
	OrbitIndex               int64    `json:"orbitIndex"`
	OwnerID                  int64    `json:"ownerID"`
	Position                 Position `json:"position"`
	ReprocessingEfficiency   float64  `json:"reprocessingEfficiency"`
	ReprocessingHangarFlag   int64    `json:"reprocessingHangarFlag"`
	ReprocessingStationsTake float64  `json:"reprocessingStationsTake"`
	SolarSystemID            int64    `json:"solarSystemID"`
	TypeID                   int64    `json:"typeID"`
	UseOperationName         bool     `json:"useOperationName"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
