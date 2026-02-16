// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    ancestries, err := UnmarshalAncestries(bytes)
//    bytes, err = ancestries.Marshal()

package ancestries

import "encoding/json"

func UnmarshalAncestries(data []byte) (Ancestries, error) {
	var r Ancestries
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Ancestries) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Ancestries struct {
	Key              int64       `json:"_key"`
	BloodlineID      int64       `json:"bloodlineID"`
	Charisma         int64       `json:"charisma"`
	Description      Description `json:"description"`
	IconID           int64       `json:"iconID"`
	Intelligence     int64       `json:"intelligence"`
	Memory           int64       `json:"memory"`
	Name             Description `json:"name"`
	Perception       int64       `json:"perception"`
	ShortDescription string      `json:"shortDescription"`
	Willpower        int64       `json:"willpower"`
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
