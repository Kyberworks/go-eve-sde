// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    mapAsteroidBelts, err := UnmarshalMapAsteroidBelts(bytes)
//    bytes, err = mapAsteroidBelts.Marshal()

package mapAsteroidBelts

import "encoding/json"

func UnmarshalMapAsteroidBelts(data []byte) (MapAsteroidBelts, error) {
	var r MapAsteroidBelts
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MapAsteroidBelts) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MapAsteroidBelts struct {
	Key            int64      `json:"_key"`
	CelestialIndex int64      `json:"celestialIndex"`
	OrbitID        int64      `json:"orbitID"`
	OrbitIndex     int64      `json:"orbitIndex"`
	Position       Position   `json:"position"`
	Radius         float64    `json:"radius"`
	SolarSystemID  int64      `json:"solarSystemID"`
	Statistics     Statistics `json:"statistics"`
	TypeID         int64      `json:"typeID"`
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
	RotationRate   float64 `json:"rotationRate"`
	SpectralClass  string  `json:"spectralClass"`
	SurfaceGravity float64 `json:"surfaceGravity"`
	Temperature    float64 `json:"temperature"`
}
