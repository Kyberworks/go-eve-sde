// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    mapPlanets, err := UnmarshalMapPlanets(bytes)
//    bytes, err = mapPlanets.Marshal()

package main

import "encoding/json"

func UnmarshalMapPlanets(data []byte) (MapPlanets, error) {
	var r MapPlanets
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MapPlanets) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MapPlanets struct {
	Key             int64      `json:"_key"`
	AsteroidBeltIDs []int64    `json:"asteroidBeltIDs"`
	Attributes      Attributes `json:"attributes"`
	CelestialIndex  int64      `json:"celestialIndex"`
	MoonIDs         []int64    `json:"moonIDs"`
	OrbitID         int64      `json:"orbitID"`
	Position        Position   `json:"position"`
	Radius          int64      `json:"radius"`
	SolarSystemID   int64      `json:"solarSystemID"`
	Statistics      Statistics `json:"statistics"`
	TypeID          int64      `json:"typeID"`
}

type Attributes struct {
	HeightMap1   int64 `json:"heightMap1"`
	HeightMap2   int64 `json:"heightMap2"`
	Population   bool  `json:"population"`
	ShaderPreset int64 `json:"shaderPreset"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Statistics struct {
	Density        float64 `json:"density"`
	Eccentricity   float64 `json:"eccentricity"`
	EscapeVelocity float64 `json:"escapeVelocity"`
	Locked         bool    `json:"locked"`
	MassDust       float64 `json:"massDust"`
	MassGas        float64 `json:"massGas"`
	OrbitPeriod    float64 `json:"orbitPeriod"`
	OrbitRadius    float64 `json:"orbitRadius"`
	Pressure       float64 `json:"pressure"`
	RotationRate   float64 `json:"rotationRate"`
	SpectralClass  string  `json:"spectralClass"`
	SurfaceGravity float64 `json:"surfaceGravity"`
	Temperature    float64 `json:"temperature"`
}
