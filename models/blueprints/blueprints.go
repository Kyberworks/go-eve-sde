// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    blueprints, err := UnmarshalBlueprints(bytes)
//    bytes, err = blueprints.Marshal()

package blueprints

import "encoding/json"

func UnmarshalBlueprints(data []byte) (Blueprints, error) {
	var r Blueprints
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Blueprints) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Blueprints struct {
	Key                int64      `json:"_key"`
	Activities         Activities `json:"activities"`
	BlueprintTypeID    int64      `json:"blueprintTypeID"`
	MaxProductionLimit int64      `json:"maxProductionLimit"`
}

type Activities struct {
	Copying          Copying       `json:"copying"`
	Manufacturing    Manufacturing `json:"manufacturing"`
	ResearchMaterial Copying       `json:"research_material"`
	ResearchTime     Copying       `json:"research_time"`
}

type Copying struct {
	Time int64 `json:"time"`
}

type Manufacturing struct {
	Materials []Material `json:"materials"`
	Products  []Material `json:"products"`
	Time      int64      `json:"time"`
}

type Material struct {
	Quantity int64 `json:"quantity"`
	TypeID   int64 `json:"typeID"`
}
