// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    cloneGrades, err := UnmarshalCloneGrades(bytes)
//    bytes, err = cloneGrades.Marshal()

package cloneGrades

import "encoding/json"

func UnmarshalCloneGrades(data []byte) (CloneGrades, error) {
	var r CloneGrades
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CloneGrades) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CloneGrades struct {
	Key    int64   `json:"_key"`
	Name   string  `json:"name"`
	Skills []Skill `json:"skills"`
}

type Skill struct {
	Level  int64 `json:"level"`
	TypeID int64 `json:"typeID"`
}
