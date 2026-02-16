// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    categories, err := UnmarshalCategories(bytes)
//    bytes, err = categories.Marshal()

package categories

import "encoding/json"

func UnmarshalCategories(data []byte) (Categories, error) {
	var r Categories
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Categories) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Categories struct {
	Key       int64 `json:"_key"`
	Name      Name  `json:"name"`
	Published bool  `json:"published"`
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
