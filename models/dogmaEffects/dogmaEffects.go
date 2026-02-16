// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dogmaEffects, err := UnmarshalDogmaEffects(bytes)
//    bytes, err = dogmaEffects.Marshal()

package dogmaEffects

import "encoding/json"

func UnmarshalDogmaEffects(data []byte) (DogmaEffects, error) {
	var r DogmaEffects
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DogmaEffects) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DogmaEffects struct {
	Key                  int64  `json:"_key"`
	DisallowAutoRepeat   bool   `json:"disallowAutoRepeat"`
	DischargeAttributeID int64  `json:"dischargeAttributeID"`
	DurationAttributeID  int64  `json:"durationAttributeID"`
	EffectCategoryID     int64  `json:"effectCategoryID"`
	ElectronicChance     bool   `json:"electronicChance"`
	GUID                 string `json:"guid"`
	IsAssistance         bool   `json:"isAssistance"`
	IsOffensive          bool   `json:"isOffensive"`
	IsWarpSafe           bool   `json:"isWarpSafe"`
	Name                 string `json:"name"`
	PropulsionChance     bool   `json:"propulsionChance"`
	Published            bool   `json:"published"`
	RangeChance          bool   `json:"rangeChance"`
}
