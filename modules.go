package loucore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type StructureBlueprint struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Buildable    bool       `json:"buildable"`
	Media        string     `json:"media"`
	Info         string     `json:"info"`
	Bonus        [10]int    `json:"bonus"`
	BonusType    string     `json:"bonus_type"`
	ResourceCost [10][2]int `json:"resource_cost"`
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