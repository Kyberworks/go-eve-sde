// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dbuffCollections, err := UnmarshalDbuffCollections(bytes)
//    bytes, err = dbuffCollections.Marshal()

package main

import "encoding/json"

func UnmarshalDbuffCollections(data []byte) (DbuffCollections, error) {
	var r DbuffCollections
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DbuffCollections) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DbuffCollections struct {
	Key                            int64                           `json:"_key"`
	AggregateMode                  string                          `json:"aggregateMode"`
	DeveloperDescription           string                          `json:"developerDescription"`
	ItemModifiers                  []Modifier                      `json:"itemModifiers"`
	LocationGroupModifiers         []LocationGroupModifier         `json:"locationGroupModifiers"`
	LocationModifiers              []Modifier                      `json:"locationModifiers"`
	LocationRequiredSkillModifiers []LocationRequiredSkillModifier `json:"locationRequiredSkillModifiers"`
	OperationName                  string                          `json:"operationName"`
	ShowOutputValueInUI            string                          `json:"showOutputValueInUI"`
}

type Modifier struct {
	DogmaAttributeID int64 `json:"dogmaAttributeID"`
}

type LocationGroupModifier struct {
	DogmaAttributeID int64 `json:"dogmaAttributeID"`
	GroupID          int64 `json:"groupID"`
}

type LocationRequiredSkillModifier struct {
	DogmaAttributeID int64 `json:"dogmaAttributeID"`
	SkillID          int64 `json:"skillID"`
}
