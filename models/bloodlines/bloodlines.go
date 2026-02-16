// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    bloodlines, err := UnmarshalBloodlines(bytes)
//    bytes, err = bloodlines.Marshal()

package bloodlines

import "encoding/json"

func UnmarshalBloodlines(data []byte) (Bloodlines, error) {
	var r Bloodlines
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Bloodlines) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Bloodlines struct {
	Key           int64       `json:"_key"`
	Charisma      int64       `json:"charisma"`
	CorporationID int64       `json:"corporationID"`
	Description   Description `json:"description"`
	IconID        int64       `json:"iconID"`
	Intelligence  int64       `json:"intelligence"`
	Memory        int64       `json:"memory"`
	Name          Description `json:"name"`
	Perception    int64       `json:"perception"`
	RaceID        int64       `json:"raceID"`
	Willpower     int64       `json:"willpower"`
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
