// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dungeons, err := UnmarshalDungeons(bytes)
//    bytes, err = dungeons.Marshal()

package dungeons

import "encoding/json"

func UnmarshalDungeons(data []byte) (Dungeons, error) {
	var r Dungeons
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Dungeons) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Dungeons struct {
	Key              int64       `json:"_key"`
	AllowedShipsList []int64     `json:"allowedShipsList"`
	ArchetypeID      int64       `json:"archetypeID"`
	Description      Description `json:"description"`
	FactionID        int64       `json:"factionID"`
	Name             Description `json:"name"`
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
