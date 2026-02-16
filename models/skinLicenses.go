// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    skinLicenses, err := UnmarshalSkinLicenses(bytes)
//    bytes, err = skinLicenses.Marshal()

package main

import "encoding/json"

func UnmarshalSkinLicenses(data []byte) (SkinLicenses, error) {
	var r SkinLicenses
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SkinLicenses) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SkinLicenses struct {
	Key           int64 `json:"_key"`
	Duration      int64 `json:"duration"`
	LicenseTypeID int64 `json:"licenseTypeID"`
	SkinID        int64 `json:"skinID"`
}
