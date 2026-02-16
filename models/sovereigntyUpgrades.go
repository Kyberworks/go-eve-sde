// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    sovereigntyUpgrades, err := UnmarshalSovereigntyUpgrades(bytes)
//    bytes, err = sovereigntyUpgrades.Marshal()

package main

import "encoding/json"

func UnmarshalSovereigntyUpgrades(data []byte) (SovereigntyUpgrades, error) {
	var r SovereigntyUpgrades
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SovereigntyUpgrades) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SovereigntyUpgrades struct {
	Key                    int64  `json:"_key"`
	Fuel                   Fuel   `json:"fuel"`
	MutuallyExclusiveGroup string `json:"mutually_exclusive_group"`
	PowerAllocation        int64  `json:"power_allocation"`
	WorkforceAllocation    int64  `json:"workforce_allocation"`
}

type Fuel struct {
	HourlyUpkeep int64 `json:"hourly_upkeep"`
	StartupCost  int64 `json:"startup_cost"`
	TypeID       int64 `json:"type_id"`
}
