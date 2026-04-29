// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    archetypes, err := UnmarshalArchetypes(bytes)
//    bytes, err = archetypes.Marshal()

package archetypes

import "encoding/json"

func UnmarshalArchetypes(data []byte) (Archetypes, error) {
	var r Archetypes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Archetypes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Archetypes struct {
	Key         int64       `json:"_key"`
	Description Description `json:"description"`
	Title       Description `json:"title"`
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
