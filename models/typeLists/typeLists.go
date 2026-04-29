// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    typeLists, err := UnmarshalTypeLists(bytes)
//    bytes, err = typeLists.Marshal()

package typeLists

import "encoding/json"

func UnmarshalTypeLists(data []byte) (TypeLists, error) {
	var r TypeLists
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TypeLists) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TypeLists struct {
	Key             int64   `json:"_key"`
	IncludedTypeIDs []int64 `json:"includedTypeIDs"`
	Name            string  `json:"name"`
}
