// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    agentTypes, err := UnmarshalAgentTypes(bytes)
//    bytes, err = agentTypes.Marshal()

package main

import "encoding/json"

func UnmarshalAgentTypes(data []byte) (AgentTypes, error) {
	var r AgentTypes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AgentTypes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AgentTypes struct {
	Key  int64  `json:"_key"`
	Name string `json:"name"`
}
