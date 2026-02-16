// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dynamicItemAttributes, err := UnmarshalDynamicItemAttributes(bytes)
//    bytes, err = dynamicItemAttributes.Marshal()

package dynamicItemAttributes

import "encoding/json"

func UnmarshalDynamicItemAttributes(data []byte) (DynamicItemAttributes, error) {
	var r DynamicItemAttributes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DynamicItemAttributes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DynamicItemAttributes struct {
	Key                int64                `json:"_key"`
	AttributeIDs       []AttributeID        `json:"attributeIDs"`
	InputOutputMapping []InputOutputMapping `json:"inputOutputMapping"`
}

type AttributeID struct {
	Key int64   `json:"_key"`
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

type InputOutputMapping struct {
	ApplicableTypes []int64 `json:"applicableTypes"`
	ResultingType   int64   `json:"resultingType"`
}
