// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    agentsInSpace, err := UnmarshalAgentsInSpace(bytes)
//    bytes, err = agentsInSpace.Marshal()

package agentsInSpace

import "encoding/json"

func UnmarshalAgentsInSpace(data []byte) (AgentsInSpace, error) {
	var r AgentsInSpace
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AgentsInSpace) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AgentsInSpace struct {
	Key           int64 `json:"_key"`
	DungeonID     int64 `json:"dungeonID"`
	SolarSystemID int64 `json:"solarSystemID"`
	SpawnPointID  int64 `json:"spawnPointID"`
	TypeID        int64 `json:"typeID"`
}
