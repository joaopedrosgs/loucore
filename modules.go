package loucore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type StructureBlueprint struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Buildable bool   `json:"buildable"`
	Media     string `json:"media"`
	Info      string `json:"info"`
	Bonus     struct {
		WoodProduction  [10]int `json:"wood_production,omitempty"`
		IronProduction  [10]int `json:"iron_production,omitempty"`
		StoneProduction [10]int `json:"stone_production,omitempty"`
		FoodProduction  [10]int `json:"food_production,omitempty"`
		GoldProduction  [10]int `json:"gold_production,omitempty"`

		WoodProductionBonus [10]int `json:"wood_production_bonus,omitempty"`
		WoodStorageBonus    [10]int `json:"wood_storage_bonus,omitempty"`
		WoodStorage         [10]int `json:"wood_storage,omitempty"`

		StoneProductionBonus [10]int `json:"stone_production_bonus,omitempty"`
		StoneStorageBonus    [10]int `json:"stone_storage_bonus,omitempty"`
		StoneStorage         [10]int `json:"stone_storage,omitempty"`

		IronProductionBonus [10]int `json:"iron_production_bonus,omitempty"`
		IronStorageBonus    [10]int `json:"iron_storage_bonus,omitempty"`
		IronStorage         [10]int `json:"iron_storage,omitempty"`

		FoodProductionBonus [10]int `json:"food_production_bonus,omitempty"`
		FoodStorageBonus    [10]int `json:"food_storage_bonus,omitempty"`
		FoodStorage         [10]int `json:"food_storage,omitempty"`

		GoldProductionBonus [10]int `json:"gold_production_bonus,omitempty"`

		Storage [10]int `json:"storage,omitempty"`
	} `json:"bonus"`
	ResourceCost [10][4]int `json:"resource_cost"`
	Score        [10]int    `json:"score"`
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
