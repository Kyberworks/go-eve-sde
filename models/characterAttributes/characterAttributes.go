// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    characterAttributes, err := UnmarshalCharacterAttributes(bytes)
//    bytes, err = characterAttributes.Marshal()

package characterAttributes

import "encoding/json"

func UnmarshalCharacterAttributes(data []byte) (CharacterAttributes, error) {
	var r CharacterAttributes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CharacterAttributes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CharacterAttributes struct {
	Key              int64  `json:"_key"`
	Description      string `json:"description"`
	IconID           int64  `json:"iconID"`
	Name             Name   `json:"name"`
	Notes            string `json:"notes"`
	ShortDescription string `json:"shortDescription"`
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
