// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    npcCorporations, err := UnmarshalNpcCorporations(bytes)
//    bytes, err = npcCorporations.Marshal()

package main

import "encoding/json"

func UnmarshalNpcCorporations(data []byte) (NpcCorporations, error) {
	var r NpcCorporations
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NpcCorporations) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type NpcCorporations struct {
	Key                        int64       `json:"_key"`
	CeoID                      int64       `json:"ceoID"`
	Deleted                    bool        `json:"deleted"`
	Description                Description `json:"description"`
	Extent                     string      `json:"extent"`
	HasPlayerPersonnelManager  bool        `json:"hasPlayerPersonnelManager"`
	InitialPrice               int64       `json:"initialPrice"`
	MemberLimit                int64       `json:"memberLimit"`
	MinSecurity                float64     `json:"minSecurity"`
	MinimumJoinStanding        int64       `json:"minimumJoinStanding"`
	Name                       Description `json:"name"`
	SendCharTerminationMessage bool        `json:"sendCharTerminationMessage"`
	Shares                     int64       `json:"shares"`
	Size                       string      `json:"size"`
	StationID                  int64       `json:"stationID"`
	TaxRate                    float64     `json:"taxRate"`
	TickerName                 string      `json:"tickerName"`
	UniqueName                 bool        `json:"uniqueName"`
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
