// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    controlTowerResources, err := UnmarshalControlTowerResources(bytes)
//    bytes, err = controlTowerResources.Marshal()

package main

import "encoding/json"

func UnmarshalControlTowerResources(data []byte) (ControlTowerResources, error) {
	var r ControlTowerResources
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ControlTowerResources) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ControlTowerResources struct {
	Key       int64      `json:"_key"`
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Purpose        int64 `json:"purpose"`
	Quantity       int64 `json:"quantity"`
	ResourceTypeID int64 `json:"resourceTypeID"`
}
