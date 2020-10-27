package loucore

import (
	"testing"
)

func TestCreateConstruction(t *testing.T) {
	city, _ := CreateCity(0, 0)

	construction, err := CreateConstruction(city.ID, 1, 1, WAREHOUSE)
	if err != nil {
		t.Errorf("Failed to CreateConstruction: %v", err)
		return
	}
	if construction == nil {
		t.Errorf("CreateConstruction failed silently: %v", err)
		return
	}
	_ = DeleteCity(city.ID)

}
func TestCreateConstructionShouldFailBecauseTwoBuildingAtSameSpot(t *testing.T) {
	city, _ := CreateCity(0, 0)

	construction, err := CreateConstruction(city.ID, 1, 1, WAREHOUSE)
	if err != nil {
		t.Errorf("Failed to CreateConstruction: %v", err)
		return
	}
	if construction == nil {
		t.Errorf("CreateConstruction failed silently: %v", err)
		return
	}

	construction2, err := CreateConstruction(city.ID, 1, 1, WOODCUTTERS_HUT)
	if err == nil {
		t.Errorf("CreateConstruction did not fail: %v", err)
		return
	}
	if construction2 != nil {
		t.Errorf("CreateConstruction failed but returned a structure: %v", err)
		return
	}
	_ = DeleteCity(city.ID)

}

func TestCalculateProduction(t *testing.T) {
	city, _ := CreateCity(0, 0)

	construction, _ := CreateConstruction(city.ID, 1, 1, WOODCUTTERS_HUT)

	v, err := CalculateProduction(construction.ID)
	if err != nil {
		t.Errorf("Failed to CalculateProduction: %v", err)
		return
	}

	if v != 5 {
		t.Errorf("Production calculated wrong, got %v and expected 5", v)
	}

	_, _ = CreateConstruction(city.ID, 1, 0, WOOD)

	_ = UpdateAffectedBy(construction.ID)

	v, err = CalculateProduction(construction.ID)
	if err != nil {
		t.Errorf("Failed to CalculateProduction: %v", err)
		return
	}

	if v != 6.5 {
		t.Errorf("Production calculated wrong, got %v and expected 6.5", v)

	}

	_ = DeleteCity(city.ID)
}
