// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    mercenaryTacticalOperations, err := UnmarshalMercenaryTacticalOperations(bytes)
//    bytes, err = mercenaryTacticalOperations.Marshal()

package mercenaryTacticalOperations

import "encoding/json"

func UnmarshalMercenaryTacticalOperations(data []byte) (MercenaryTacticalOperations, error) {
	var r MercenaryTacticalOperations
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MercenaryTacticalOperations) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MercenaryTacticalOperations struct {
	Key               int64       `json:"_key"`
	AnarchyImpact     int64       `json:"anarchy_impact"`
	Description       Description `json:"description"`
	DevelopmentImpact int64       `json:"development_impact"`
	InfomorphBonus    int64       `json:"infomorph_bonus"`
	Name              Description `json:"name"`
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
