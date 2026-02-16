// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    marketGroups, err := UnmarshalMarketGroups(bytes)
//    bytes, err = marketGroups.Marshal()

package main

import "encoding/json"

func UnmarshalMarketGroups(data []byte) (MarketGroups, error) {
	var r MarketGroups
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MarketGroups) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MarketGroups struct {
	Key         int64       `json:"_key"`
	Description Description `json:"description"`
	HasTypes    bool        `json:"hasTypes"`
	IconID      int64       `json:"iconID"`
	Name        Description `json:"name"`
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
