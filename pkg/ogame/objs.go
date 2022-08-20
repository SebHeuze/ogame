package ogame

// All ogame objects
var (
	AllianceDepot                = register[*allianceDepot](newAllianceDepot) // Buildings
	CrystalMine                  = register[*crystalMine](newCrystalMine)
	CrystalStorage               = register[*crystalStorage](newCrystalStorage)
	DeuteriumSynthesizer         = register[*deuteriumSynthesizer](newDeuteriumSynthesizer)
	DeuteriumTank                = register[*deuteriumTank](newDeuteriumTank)
	FusionReactor                = register[*fusionReactor](newFusionReactor)
	MetalMine                    = register[*metalMine](newMetalMine)
	MetalStorage                 = register[*metalStorage](newMetalStorage)
	MissileSilo                  = register[*missileSilo](newMissileSilo)
	NaniteFactory                = register[*naniteFactory](newNaniteFactory)
	ResearchLab                  = register[*researchLab](newResearchLab)
	RoboticsFactory              = register[*roboticsFactory](newRoboticsFactory)
	SeabedDeuteriumDen           = register[*seabedDeuteriumDen](newSeabedDeuteriumDen)
	ShieldedMetalDen             = register[*shieldedMetalDen](newShieldedMetalDen)
	Shipyard                     = register[*shipyard](newShipyard)
	SolarPlant                   = register[*solarPlant](newSolarPlant)
	SpaceDock                    = register[*spaceDock](newSpaceDock)
	LunarBase                    = register[*lunarBase](newLunarBase)
	SensorPhalanx                = register[*sensorPhalanx](newSensorPhalanx)
	JumpGate                     = register[*jumpGate](newJumpGate)
	Terraformer                  = register[*terraformer](newTerraformer)
	UndergroundCrystalDen        = register[*undergroundCrystalDen](newUndergroundCrystalDen)
	SolarSatellite               = register[*solarSatellite](newSolarSatellite)
	AntiBallisticMissiles        = register[*antiBallisticMissiles](newAntiBallisticMissiles) // Defense
	GaussCannon                  = register[*gaussCannon](newGaussCannon)
	HeavyLaser                   = register[*heavyLaser](newHeavyLaser)
	InterplanetaryMissiles       = register[*interplanetaryMissiles](newInterplanetaryMissiles)
	IonCannon                    = register[*ionCannon](newIonCannon)
	LargeShieldDome              = register[*largeShieldDome](newLargeShieldDome)
	LightLaser                   = register[*lightLaser](newLightLaser)
	PlasmaTurret                 = register[*plasmaTurret](newPlasmaTurret)
	RocketLauncher               = register[*rocketLauncher](newRocketLauncher)
	SmallShieldDome              = register[*smallShieldDome](newSmallShieldDome)
	Battlecruiser                = register[*battlecruiser](newBattlecruiser) // Ships
	Battleship                   = register[*battleship](newBattleship)
	Bomber                       = register[*bomber](newBomber)
	ColonyShip                   = register[*colonyShip](newColonyShip)
	Cruiser                      = register[*cruiser](newCruiser)
	Deathstar                    = register[*deathstar](newDeathstar)
	Destroyer                    = register[*destroyer](newDestroyer)
	EspionageProbe               = register[*espionageProbe](newEspionageProbe)
	HeavyFighter                 = register[*heavyFighter](newHeavyFighter)
	LargeCargo                   = register[*largeCargo](newLargeCargo)
	LightFighter                 = register[*lightFighter](newLightFighter)
	Recycler                     = register[*recycler](newRecycler)
	SmallCargo                   = register[*smallCargo](newSmallCargo)
	Crawler                      = register[*crawler](newCrawler)
	Reaper                       = register[*reaper](newReaper)
	Pathfinder                   = register[*pathfinder](newPathfinder)
	ArmourTechnology             = register[*armourTechnology](newArmourTechnology) // Technologies
	Astrophysics                 = register[*astrophysics](newAstrophysics)
	CombustionDrive              = register[*combustionDrive](newCombustionDrive)
	ComputerTechnology           = register[*computerTechnology](newComputerTechnology)
	EnergyTechnology             = register[*energyTechnology](newEnergyTechnology)
	EspionageTechnology          = register[*espionageTechnology](newEspionageTechnology)
	GravitonTechnology           = register[*gravitonTechnology](newGravitonTechnology)
	HyperspaceDrive              = register[*hyperspaceDrive](newHyperspaceDrive)
	HyperspaceTechnology         = register[*hyperspaceTechnology](newHyperspaceTechnology)
	ImpulseDrive                 = register[*impulseDrive](newImpulseDrive)
	IntergalacticResearchNetwork = register[*intergalacticResearchNetwork](newIntergalacticResearchNetwork)
	IonTechnology                = register[*ionTechnology](newIonTechnology)
	LaserTechnology              = register[*laserTechnology](newLaserTechnology)
	PlasmaTechnology             = register[*plasmaTechnology](newPlasmaTechnology)
	ShieldingTechnology          = register[*shieldingTechnology](newShieldingTechnology)
	WeaponsTechnology            = register[*weaponsTechnology](newWeaponsTechnology)
)

type ObjsStruct struct{ m map[ID]BaseOgameObj }

func (o ObjsStruct) ByID(id ID) BaseOgameObj {
	return o.m[id]
}

var Objs = ObjsStruct{m: make(map[ID]BaseOgameObj)}

func register[T BaseOgameObj](constructorFn func() T) T {
	inst := constructorFn()
	Objs.m[inst.GetID()] = inst
	return inst
}

// Defenses array of all defenses objects
var Defenses = []Defense{
	RocketLauncher,
	LightLaser,
	HeavyLaser,
	GaussCannon,
	IonCannon,
	PlasmaTurret,
	SmallShieldDome,
	LargeShieldDome,
	AntiBallisticMissiles,
	InterplanetaryMissiles,
}

// Ships array of all ships objects
var Ships = []Ship{
	LightFighter,
	HeavyFighter,
	Cruiser,
	Battleship,
	Battlecruiser,
	Bomber,
	Destroyer,
	Deathstar,
	SmallCargo,
	LargeCargo,
	ColonyShip,
	Recycler,
	EspionageProbe,
	SolarSatellite,
	Crawler,
	Reaper,
	Pathfinder,
}

// Buildings array of all buildings/facilities objects
var Buildings = []Building{
	MetalMine,
	CrystalMine,
	DeuteriumSynthesizer,
	SolarPlant,
	FusionReactor,
	SolarSatellite,
	MetalStorage,
	CrystalStorage,
	DeuteriumTank,
	ShieldedMetalDen,
	UndergroundCrystalDen,
	SeabedDeuteriumDen,
	RoboticsFactory,
	Shipyard,
	ResearchLab,
	AllianceDepot,
	MissileSilo,
	NaniteFactory,
	Terraformer,
	SpaceDock,
	LunarBase,
	SensorPhalanx,
	JumpGate,
}

// PlanetBuildings arrays of planet specific buildings
var PlanetBuildings = []Building{
	MetalMine,
	CrystalMine,
	DeuteriumSynthesizer,
	SolarPlant,
	FusionReactor,
	SolarSatellite,
	MetalStorage,
	CrystalStorage,
	DeuteriumTank,
	ShieldedMetalDen,
	UndergroundCrystalDen,
	SeabedDeuteriumDen,
	RoboticsFactory,
	Shipyard,
	ResearchLab,
	AllianceDepot,
	MissileSilo,
	NaniteFactory,
	Terraformer,
	SpaceDock,
}

// MoonBuildings arrays of moon specific buildings
var MoonBuildings = []Building{
	SolarSatellite,
	MetalStorage,
	CrystalStorage,
	DeuteriumTank,
	RoboticsFactory,
	Shipyard,
	LunarBase,
	SensorPhalanx,
	JumpGate,
}

// Technologies array of all technologies objects
var Technologies = []Technology{
	EnergyTechnology,
	LaserTechnology,
	IonTechnology,
	HyperspaceTechnology,
	PlasmaTechnology,
	CombustionDrive,
	ImpulseDrive,
	HyperspaceDrive,
	EspionageTechnology,
	ComputerTechnology,
	Astrophysics,
	IntergalacticResearchNetwork,
	GravitonTechnology,
	WeaponsTechnology,
	ShieldingTechnology,
	ArmourTechnology,
}