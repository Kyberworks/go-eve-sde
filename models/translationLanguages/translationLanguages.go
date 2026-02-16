// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    translationLanguages, err := UnmarshalTranslationLanguages(bytes)
//    bytes, err = translationLanguages.Marshal()

package translationLanguages

import "encoding/json"

func UnmarshalTranslationLanguages(data []byte) (TranslationLanguages, error) {
	var r TranslationLanguages
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TranslationLanguages) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TranslationLanguages struct {
	Key  string `json:"_key"`
	Name string `json:"name"`
}
