// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    freelanceJobSchemas, err := UnmarshalFreelanceJobSchemas(bytes)
//    bytes, err = freelanceJobSchemas.Marshal()

package main

import "encoding/json"

func UnmarshalFreelanceJobSchemas(data []byte) (FreelanceJobSchemas, error) {
	var r FreelanceJobSchemas
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FreelanceJobSchemas) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FreelanceJobSchemas struct {
	Key   int64   `json:"_key"`
	Value []Value `json:"_value"`
}

type Value struct {
	Key                            string                          `json:"_key"`
	ContentTags                    []string                        `json:"contentTags"`
	Description                    Description                     `json:"description"`
	IconID                         string                          `json:"iconID"`
	MaxContributionsPerParticipant MaxContributionsPerParticipant  `json:"maxContributionsPerParticipant"`
	Parameters                     []Parameter                     `json:"parameters"`
	ProgressDescription            Description                     `json:"progressDescription"`
	RewardDescription              Description                     `json:"rewardDescription"`
	TargetDescription              Description                     `json:"targetDescription"`
	Title                          Description                     `json:"title"`
	ContributionMultiplier         *ContributionMultiplier         `json:"contributionMultiplier,omitempty"`
	MaxProgressPerContribution     *MaxContributionsPerParticipant `json:"maxProgressPerContribution,omitempty"`
}

type ContributionMultiplier struct {
	DefaultValue     int64       `json:"defaultValue"`
	Description      Description `json:"description"`
	IconID           IconID      `json:"iconID"`
	MaxValue         int64       `json:"maxValue"`
	MinValue         float64     `json:"minValue"`
	Title            Description `json:"title"`
	UnsetDescription Description `json:"unsetDescription"`
}

type Description struct {
	De string `json:"de"`
	En string `json:"en"`
	Es string `json:"es"`
	Fr string `json:"fr"`
	Ja string `json:"ja"`
	Ko string `json:"ko"`
	Ru string `json:"ru"`
	Zh string `json:"zh"`
}

type MaxContributionsPerParticipant struct {
	Description        Description `json:"description"`
	IconID             IconID      `json:"iconID"`
	Title              Description `json:"title"`
	UnsetDescription   Description `json:"unsetDescription"`
	AcceptedValueTypes []string    `json:"acceptedValueTypes,omitempty"`
}

type Parameter struct {
	Key          string        `json:"_key"`
	Matcher      *Matcher      `json:"matcher,omitempty"`
	ItemDelivery *ItemDelivery `json:"itemDelivery,omitempty"`
	Boolean      *Boolean      `json:"boolean,omitempty"`
}

type Boolean struct {
	ChoiceLabel Description `json:"choiceLabel"`
	Default     bool        `json:"default"`
	Description Description `json:"description"`
	IconID      string      `json:"iconID"`
	OptionFalse Option      `json:"optionFalse"`
	OptionTrue  Option      `json:"optionTrue"`
	Title       Description `json:"title"`
}

type Option struct {
	Description Description `json:"description"`
	Title       Description `json:"title"`
}

type ItemDelivery struct {
	DeliveryLocation Matcher                        `json:"deliveryLocation"`
	Description      Description                    `json:"description"`
	IconID           IconID                         `json:"iconID"`
	InventoryType    MaxContributionsPerParticipant `json:"inventoryType"`
	Title            Description                    `json:"title"`
}

type Matcher struct {
	AcceptedValueTypes []string    `json:"acceptedValueTypes"`
	Description        Description `json:"description"`
	IconID             string      `json:"iconID"`
	MaxEntries         int64       `json:"maxEntries"`
	Title              Description `json:"title"`
	UnsetDescription   Description `json:"unsetDescription"`
	Optional           *bool       `json:"optional,omitempty"`
	Type               *string     `json:"type,omitempty"`
}

type IconID string

const (
	Inventory IconID = "inventory"
	Isk       IconID = "isk"
	Personal  IconID = "personal"
)
