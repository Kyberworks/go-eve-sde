// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    planetResources, err := UnmarshalPlanetResources(bytes)
//    bytes, err = planetResources.Marshal()

package planetResources

import "encoding/json"

func UnmarshalPlanetResources(data []byte) (PlanetResources, error) {
	var r PlanetResources
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *PlanetResources) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type PlanetResources struct {
	Key   int64 `json:"_key"`
	Power int64 `json:"power"`
}
