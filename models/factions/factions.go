// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    factions, err := UnmarshalFactions(bytes)
//    bytes, err = factions.Marshal()

package factions

import "encoding/json"

func UnmarshalFactions(data []byte) (Factions, error) {
	var r Factions
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Factions) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Factions struct {
	Key                  int64       `json:"_key"`
	CorporationID        int64       `json:"corporationID"`
	Description          Description `json:"description"`
	FlatLogo             string      `json:"flatLogo"`
	FlatLogoWithName     string      `json:"flatLogoWithName"`
	IconID               int64       `json:"iconID"`
	MemberRaces          []int64     `json:"memberRaces"`
	MilitiaCorporationID int64       `json:"militiaCorporationID"`
	Name                 Description `json:"name"`
	ShortDescription     Description `json:"shortDescription"`
	SizeFactor           float64     `json:"sizeFactor"`
	SolarSystemID        int64       `json:"solarSystemID"`
	UniqueName           bool        `json:"uniqueName"`
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
