// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    npcCharacters, err := UnmarshalNpcCharacters(bytes)
//    bytes, err = npcCharacters.Marshal()

package npcCharacters

import "time"

import "encoding/json"

func UnmarshalNpcCharacters(data []byte) (NpcCharacters, error) {
	var r NpcCharacters
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NpcCharacters) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type NpcCharacters struct {
	Key           int64     `json:"_key"`
	BloodlineID   int64     `json:"bloodlineID"`
	Ceo           bool      `json:"ceo"`
	CorporationID int64     `json:"corporationID"`
	Gender        bool      `json:"gender"`
	LocationID    int64     `json:"locationID"`
	Name          Name      `json:"name"`
	RaceID        int64     `json:"raceID"`
	StartDate     time.Time `json:"startDate"`
	UniqueName    bool      `json:"uniqueName"`
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
