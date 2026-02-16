// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dogmaAttributeCategories, err := UnmarshalDogmaAttributeCategories(bytes)
//    bytes, err = dogmaAttributeCategories.Marshal()

package dogmaAttributeCategories

import "encoding/json"

func UnmarshalDogmaAttributeCategories(data []byte) (DogmaAttributeCategories, error) {
	var r DogmaAttributeCategories
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DogmaAttributeCategories) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DogmaAttributeCategories struct {
	Key         int64  `json:"_key"`
	Description string `json:"description"`
	Name        string `json:"name"`
}
