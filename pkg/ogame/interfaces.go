package ogame

import (
	"time"
)

type Celestial interface {
	GetCoordinate() Coordinate
	GetDiameter() int64
	GetFields() Fields
	GetID() CelestialID
	GetImg() string
	GetName() string
	GetType() CelestialType
}

// BaseOgameObj base interface for all ogame objects (buildings, technologies, ships, defenses)
type BaseOgameObj interface {
	ConstructionTime(nbr, universeSpeed int64, acc BuildAccelerators, hasTechnocrat, isDiscoverer bool) time.Duration
	GetID() ID
	GetName() string
	GetPrice(int64) Resources
	GetRequirements() map[ID]int64
	IsAvailable(CelestialType, LazyResourcesBuildings, LazyFacilities, LazyResearches, int64, CharacterClass) bool
}

// DefenderObj base interface for all defensive units (ships, defenses)
type DefenderObj interface {
	BaseOgameObj
	DefenderConstructionTime(nbr, universeSpeed int64, acc DefenseAccelerators) time.Duration
	GetRapidfireAgainst() map[ID]int64
	GetRapidfireFrom() map[ID]int64
	GetShieldPower(Researches) int64
	GetStructuralIntegrity(Researches) int64
	GetWeaponPower(Researches) int64
}

// Ship interface implemented by all ships units
type Ship interface {
	DefenderObj
	GetCargoCapacity(techs Researches, probeRaids, isCollector, isPioneers bool) int64
	GetFuelConsumption(techs Researches, fleetDeutSaveFactor float64, isGeneral bool) int64
	GetSpeed(techs Researches, isCollector, isGeneral bool) int64
}

// Defense interface implemented by all defenses units
type Defense interface {
	DefenderObj
}

// Levelable base interface for all levelable ogame objects (buildings, technologies)
type Levelable interface {
	BaseOgameObj
	GetLevel(LazyResourcesBuildings, LazyFacilities, LazyResearches) int64
}

// Technology interface that all technologies implement
type Technology interface {
	Levelable
	TechnologyConstructionTime(nbr, universeSpeed int64, acc TechAccelerators, hasTechnocrat, isDiscoverer bool) time.Duration
}

// Building interface that all buildings implement
type Building interface {
	Levelable
	BuildingConstructionTime(nbr, universeSpeed int64, acc BuildingAccelerators) time.Duration
	DeconstructionPrice(lvl int64, techs Researches) Resources
}

type BuildAccelerators interface {
	TechAccelerators
	BuildingAccelerators
	DefenseAccelerators
}

type TechAccelerators interface {
	GetResearchLab() int64
}

type DefenseAccelerators interface {
	GetNaniteFactory() int64
	GetShipyard() int64
}

type BuildingAccelerators interface {
	GetNaniteFactory() int64
	GetRoboticsFactory() int64
}
