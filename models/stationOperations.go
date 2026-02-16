// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    stationOperations, err := UnmarshalStationOperations(bytes)
//    bytes, err = stationOperations.Marshal()

package main

import "encoding/json"

func UnmarshalStationOperations(data []byte) (StationOperations, error) {
	var r StationOperations
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StationOperations) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StationOperations struct {
	Key                 int64         `json:"_key"`
	ActivityID          int64         `json:"activityID"`
	Border              float64       `json:"border"`
	Corridor            float64       `json:"corridor"`
	Description         Description   `json:"description"`
	Fringe              float64       `json:"fringe"`
	Hub                 float64       `json:"hub"`
	ManufacturingFactor float64       `json:"manufacturingFactor"`
	OperationName       Description   `json:"operationName"`
	Ratio               float64       `json:"ratio"`
	ResearchFactor      float64       `json:"researchFactor"`
	Services            []int64       `json:"services"`
	StationTypes        []StationType `json:"stationTypes"`
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

type StationType struct {
	Key   int64 `json:"_key"`
	Value int64 `json:"_value"`
}
