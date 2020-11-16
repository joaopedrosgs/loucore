package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	WOOD = iota
	LAKE
	STONE
	IRON
	TOWNHALL
	WOODCUTTERS_HUT
	QUARRY
	FARM
	IRON_MINE
	COTTAGE
	WAREHOUSE
)

type StructureBlueprint struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	Buildable    bool           `json:"buildable"`
	Media        string         `json:"media"`
	Info         string         `json:"info"`
	Bonus        [10]float64    `json:"bonus"`
	BonusType    string         `json:"bonus_type"`
	ResourceCost [10][2]float64 `json:"resource_cost"`
	Score        [10]int        `json:"score"`
	Affects      []int          `json:"affects"`
	AffectedBy   []int          `json:"affectedBy"`
}

func (s *StructureBlueprint) DoesAffect(i int) bool {
	for v := range s.Affects {
		if v == i {
			return true
		}
	}
	return false
}

var Modules struct {
	Structures []StructureBlueprint
}

func LoadModules(location string) error {
	jsonFile, err := os.Open(filepath.Join(location, "structures.json"))
	if err != nil {
		return fmt.Errorf("(LoadModules) Failed to Open file: %v", err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("(LoadModules) Failed to ReadAll file: %v", err)
	}

	err = json.Unmarshal(byteValue, &Modules.Structures)
	if err != nil {
		return fmt.Errorf("(LoadModules) Failed to Unmarshal file content: %v", err)
	}

	err = jsonFile.Close()
	if err != nil {
		return fmt.Errorf("(LoadModules) Failed to close file: %v", err)
	}

	return nil
}
