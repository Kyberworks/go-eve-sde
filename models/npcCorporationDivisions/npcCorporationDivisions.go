// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    npcCorporationDivisions, err := UnmarshalNpcCorporationDivisions(bytes)
//    bytes, err = npcCorporationDivisions.Marshal()

package npcCorporationDivisions

import "encoding/json"

func UnmarshalNpcCorporationDivisions(data []byte) (NpcCorporationDivisions, error) {
	var r NpcCorporationDivisions
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NpcCorporationDivisions) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type NpcCorporationDivisions struct {
	Key            int64  `json:"_key"`
	DisplayName    string `json:"displayName"`
	InternalName   string `json:"internalName"`
	LeaderTypeName Name   `json:"leaderTypeName"`
	Name           Name   `json:"name"`
}

type Name struct {
	De string `json:"de"`
	En string `json:"en"`
	Es string `json:"es"`
	Fr string `json:"fr"`
	Ja string `json:"ja"`
	Ko string `json:"ko"`
	Ru string `json:"ru"`
	Zh string `json:"zh"`
}
