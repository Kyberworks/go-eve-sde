// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    races, err := UnmarshalRaces(bytes)
//    bytes, err = races.Marshal()

package races

import "encoding/json"

func UnmarshalRaces(data []byte) (Races, error) {
	var r Races
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Races) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Races struct {
	Key         int64       `json:"_key"`
	Description Description `json:"description"`
	IconID      int64       `json:"iconID"`
	Name        Description `json:"name"`
	ShipTypeID  int64       `json:"shipTypeID"`
	Skills      []Skill     `json:"skills"`
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

type Skill struct {
	Key   int64 `json:"_key"`
	Value int64 `json:"_value"`
}
