package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	agentTypesPkg "github.com/Kyberworks/go-eve-sde/models/agentTypes"
	agentsInSpacePkg "github.com/Kyberworks/go-eve-sde/models/agentsInSpace"
	ancestriesPkg "github.com/Kyberworks/go-eve-sde/models/ancestries"
	bloodlinesPkg "github.com/Kyberworks/go-eve-sde/models/bloodlines"
	blueprintsPkg "github.com/Kyberworks/go-eve-sde/models/blueprints"
	categoriesPkg "github.com/Kyberworks/go-eve-sde/models/categories"
	certificatesPkg "github.com/Kyberworks/go-eve-sde/models/certificates"
	characterAttributesPkg "github.com/Kyberworks/go-eve-sde/models/characterAttributes"
	cloneGradesPkg "github.com/Kyberworks/go-eve-sde/models/cloneGrades"
	compressibleTypesPkg "github.com/Kyberworks/go-eve-sde/models/compressibleTypes"
	contrabandTypesPkg "github.com/Kyberworks/go-eve-sde/models/contrabandTypes"
	controlTowerResourcesPkg "github.com/Kyberworks/go-eve-sde/models/controlTowerResources"
	corporationActivitiesPkg "github.com/Kyberworks/go-eve-sde/models/corporationActivities"
	dbuffCollectionsPkg "github.com/Kyberworks/go-eve-sde/models/dbuffCollections"
	dogmaAttributeCategoriesPkg "github.com/Kyberworks/go-eve-sde/models/dogmaAttributeCategories"
	dogmaAttributesPkg "github.com/Kyberworks/go-eve-sde/models/dogmaAttributes"
	dogmaEffectsPkg "github.com/Kyberworks/go-eve-sde/models/dogmaEffects"
	dogmaUnitsPkg "github.com/Kyberworks/go-eve-sde/models/dogmaUnits"
	dynamicItemAttributesPkg "github.com/Kyberworks/go-eve-sde/models/dynamicItemAttributes"
	factionsPkg "github.com/Kyberworks/go-eve-sde/models/factions"
	freelanceJobSchemasPkg "github.com/Kyberworks/go-eve-sde/models/freelanceJobSchemas"
	graphicsPkg "github.com/Kyberworks/go-eve-sde/models/graphics"
	groupsPkg "github.com/Kyberworks/go-eve-sde/models/groups"
	iconsPkg "github.com/Kyberworks/go-eve-sde/models/icons"
	landmarksPkg "github.com/Kyberworks/go-eve-sde/models/landmarks"
	latestPkg "github.com/Kyberworks/go-eve-sde/models/latest"
	mapAsteroidBeltsPkg "github.com/Kyberworks/go-eve-sde/models/mapAsteroidBelts"
	mapConstellationsPkg "github.com/Kyberworks/go-eve-sde/models/mapConstellations"
	mapMoonsPkg "github.com/Kyberworks/go-eve-sde/models/mapMoons"
	mapPlanetsPkg "github.com/Kyberworks/go-eve-sde/models/mapPlanets"
	mapRegionsPkg "github.com/Kyberworks/go-eve-sde/models/mapRegions"
	mapSolarSystemsPkg "github.com/Kyberworks/go-eve-sde/models/mapSolarSystems"
	mapStargatesPkg "github.com/Kyberworks/go-eve-sde/models/mapStargates"
	mapStarsPkg "github.com/Kyberworks/go-eve-sde/models/mapStars"
	marketGroupsPkg "github.com/Kyberworks/go-eve-sde/models/marketGroups"
	masteriesPkg "github.com/Kyberworks/go-eve-sde/models/masteries"
	metaGroupsPkg "github.com/Kyberworks/go-eve-sde/models/metaGroups"
	npcCharactersPkg "github.com/Kyberworks/go-eve-sde/models/npcCharacters"
	npcCorporationDivisionsPkg "github.com/Kyberworks/go-eve-sde/models/npcCorporationDivisions"
	npcCorporationsPkg "github.com/Kyberworks/go-eve-sde/models/npcCorporations"
	npcStationsPkg "github.com/Kyberworks/go-eve-sde/models/npcStations"
	planetResourcesPkg "github.com/Kyberworks/go-eve-sde/models/planetResources"
	planetSchematicsPkg "github.com/Kyberworks/go-eve-sde/models/planetSchematics"
	racesPkg "github.com/Kyberworks/go-eve-sde/models/races"
	skinLicensesPkg "github.com/Kyberworks/go-eve-sde/models/skinLicenses"
	skinMaterialsPkg "github.com/Kyberworks/go-eve-sde/models/skinMaterials"
	skinsPkg "github.com/Kyberworks/go-eve-sde/models/skins"
	sovereigntyUpgradesPkg "github.com/Kyberworks/go-eve-sde/models/sovereigntyUpgrades"
	stationOperationsPkg "github.com/Kyberworks/go-eve-sde/models/stationOperations"
	stationServicesPkg "github.com/Kyberworks/go-eve-sde/models/stationServices"
	translationLanguagesPkg "github.com/Kyberworks/go-eve-sde/models/translationLanguages"
	typeBonusPkg "github.com/Kyberworks/go-eve-sde/models/typeBonus"
	typeDogmaPkg "github.com/Kyberworks/go-eve-sde/models/typeDogma"
	typeMaterialsPkg "github.com/Kyberworks/go-eve-sde/models/typeMaterials"
	typesPkg "github.com/Kyberworks/go-eve-sde/models/types"
	sdePkg "github.com/Kyberworks/go-eve-sde/models/_sde"
)

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
	typeMap := map[string]reflect.Type{
		"agentTypes":                reflect.TypeOf(agentTypesPkg.AgentTypes{}),
		"agentsInSpace":             reflect.TypeOf(agentsInSpacePkg.AgentsInSpace{}),
		"ancestries":                reflect.TypeOf(ancestriesPkg.Ancestries{}),
		"bloodlines":                reflect.TypeOf(bloodlinesPkg.Bloodlines{}),
		"blueprints":                reflect.TypeOf(blueprintsPkg.Blueprints{}),
		"categories":                reflect.TypeOf(categoriesPkg.Categories{}),
		"certificates":              reflect.TypeOf(certificatesPkg.Certificates{}),
		"characterAttributes":       reflect.TypeOf(characterAttributesPkg.CharacterAttributes{}),
		"cloneGrades":               reflect.TypeOf(cloneGradesPkg.CloneGrades{}),
		"compressibleTypes":         reflect.TypeOf(compressibleTypesPkg.CompressibleTypes{}),
		"contrabandTypes":           reflect.TypeOf(contrabandTypesPkg.ContrabandTypes{}),
		"controlTowerResources":     reflect.TypeOf(controlTowerResourcesPkg.ControlTowerResources{}),
		"corporationActivities":     reflect.TypeOf(corporationActivitiesPkg.CorporationActivities{}),
		"dbuffCollections":          reflect.TypeOf(dbuffCollectionsPkg.DbuffCollections{}),
		"dogmaAttributeCategories":  reflect.TypeOf(dogmaAttributeCategoriesPkg.DogmaAttributeCategories{}),
		"dogmaAttributes":           reflect.TypeOf(dogmaAttributesPkg.DogmaAttributes{}),
		"dogmaEffects":              reflect.TypeOf(dogmaEffectsPkg.DogmaEffects{}),
		"dogmaUnits":                reflect.TypeOf(dogmaUnitsPkg.DogmaUnits{}),
		"dynamicItemAttributes":     reflect.TypeOf(dynamicItemAttributesPkg.DynamicItemAttributes{}),
		"factions":                  reflect.TypeOf(factionsPkg.Factions{}),
		"freelanceJobSchemas":       reflect.TypeOf(freelanceJobSchemasPkg.FreelanceJobSchemas{}),
		"graphics":                  reflect.TypeOf(graphicsPkg.Graphics{}),
		"groups":                    reflect.TypeOf(groupsPkg.Groups{}),
		"icons":                     reflect.TypeOf(iconsPkg.Icons{}),
		"landmarks":                 reflect.TypeOf(landmarksPkg.Landmarks{}),
		"latest":                    reflect.TypeOf(latestPkg.Latest{}),
		"mapAsteroidBelts":          reflect.TypeOf(mapAsteroidBeltsPkg.MapAsteroidBelts{}),
		"mapConstellations":         reflect.TypeOf(mapConstellationsPkg.MapConstellations{}),
		"mapMoons":                  reflect.TypeOf(mapMoonsPkg.MapMoons{}),
		"mapPlanets":                reflect.TypeOf(mapPlanetsPkg.MapPlanets{}),
		"mapRegions":                reflect.TypeOf(mapRegionsPkg.MapRegions{}),
		"mapSolarSystems":           reflect.TypeOf(mapSolarSystemsPkg.MapSolarSystems{}),
		"mapStargates":              reflect.TypeOf(mapStargatesPkg.MapStargates{}),
		"mapStars":                  reflect.TypeOf(mapStarsPkg.MapStars{}),
		"marketGroups":              reflect.TypeOf(marketGroupsPkg.MarketGroups{}),
		"masteries":                 reflect.TypeOf(masteriesPkg.Masteries{}),
		"metaGroups":                reflect.TypeOf(metaGroupsPkg.MetaGroups{}),
		"npcCharacters":             reflect.TypeOf(npcCharactersPkg.NpcCharacters{}),
		"npcCorporationDivisions":   reflect.TypeOf(npcCorporationDivisionsPkg.NpcCorporationDivisions{}),
		"npcCorporations":           reflect.TypeOf(npcCorporationsPkg.NpcCorporations{}),
		"npcStations":               reflect.TypeOf(npcStationsPkg.NpcStations{}),
		"planetResources":           reflect.TypeOf(planetResourcesPkg.PlanetResources{}),
		"planetSchematics":          reflect.TypeOf(planetSchematicsPkg.PlanetSchematics{}),
		"races":                     reflect.TypeOf(racesPkg.Races{}),
		"skinLicenses":              reflect.TypeOf(skinLicensesPkg.SkinLicenses{}),
		"skinMaterials":             reflect.TypeOf(skinMaterialsPkg.SkinMaterials{}),
		"skins":                     reflect.TypeOf(skinsPkg.Skins{}),
		"sovereigntyUpgrades":       reflect.TypeOf(sovereigntyUpgradesPkg.SovereigntyUpgrades{}),
		"stationOperations":         reflect.TypeOf(stationOperationsPkg.StationOperations{}),
		"stationServices":           reflect.TypeOf(stationServicesPkg.StationServices{}),
		"translationLanguages":      reflect.TypeOf(translationLanguagesPkg.TranslationLanguages{}),
		"typeBonus":                 reflect.TypeOf(typeBonusPkg.TypeBonus{}),
		"typeDogma":                 reflect.TypeOf(typeDogmaPkg.TypeDogma{}),
		"typeMaterials":             reflect.TypeOf(typeMaterialsPkg.TypeMaterials{}),
		"types":                     reflect.TypeOf(typesPkg.Types{}),
		"_sde":                      reflect.TypeOf(sdePkg.Sde{}),
	}

	files, _ := os.ReadDir(".")
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".jsonl") {
			base := strings.TrimSuffix(file.Name(), ".jsonl")
			typ, ok := typeMap[base]
			if !ok {
				fmt.Printf("Skipping %s\n", base)
				continue
			}
			data, _ := os.ReadFile(file.Name())
			lines := strings.Split(strings.TrimSpace(string(data)), "\n")
			sliceType := reflect.SliceOf(typ)
			slice := reflect.MakeSlice(sliceType, 0, len(lines))
			for _, line := range lines {
				if line == "" {
					continue
				}
				elemPtr := reflect.New(typ)
				json.Unmarshal([]byte(line), elemPtr.Interface())
				slice = reflect.Append(slice, elemPtr.Elem())
			}
			outFile := fmt.Sprintf("models/%s/data.go", base)
			os.MkdirAll(fmt.Sprintf("models/%s", base), 0755)
			f, _ := os.Create(outFile)
			fmt.Fprintf(f, "package %s\n\n", base)
			fmt.Fprintf(f, "var %sData = %#v\n", capitalize(base), slice.Interface())
			f.Close()
			fmt.Printf("Generated %s\n", outFile)
		}
	}
}
