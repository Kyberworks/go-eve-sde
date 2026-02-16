// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    stationServices, err := UnmarshalStationServices(bytes)
//    bytes, err = stationServices.Marshal()

package main

import "encoding/json"

func UnmarshalStationServices(data []byte) (StationServices, error) {
	var r StationServices
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *StationServices) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type StationServices struct {
	Key         int64       `json:"_key"`
	ServiceName ServiceName `json:"serviceName"`
}

type ServiceName struct {
	De string `json:"de"`
	En string `json:"en"`
	Es string `json:"es"`
	Fr string `json:"fr"`
	Ja string `json:"ja"`
	Ko string `json:"ko"`
	Ru string `json:"ru"`
	Zh string `json:"zh"`
}
