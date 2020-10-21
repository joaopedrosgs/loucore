package loucore

import (
	"testing"
)

func TestCalculateProduction(t *testing.T) {
	city, err:= CreateCity(0,0)
	if err != nil {
		t.Errorf("Failed to CreateCity: %v", err)
		return
	}

	construction, err := CreateConstruction(city.ID, 1,1,WOODCUTTERS_HUT)
	if err != nil {
		t.Errorf("Failed to CreateConstruction: %v", err)
		return
	}

	v, err:= CalculateProduction(construction.ID)
	if err != nil {
		t.Errorf("Failed to CalculateProduction: %v", err)
		return
	}

	if v!=5 {
		t.Errorf("Production calculated wrong, got %v and expected 5", v)
	}

	_, err= CreateConstruction(city.ID, 1,0, WOOD)
	if err != nil {
		t.Errorf("Failed to CreateConstruction: %v", err)
		return
	}

	err=UpdateAffectedBy(construction.ID)
	if err != nil {
		t.Errorf("Failed to UpdateAffectedBy: %v", err)
		return
	}

	v, err= CalculateProduction(construction.ID)
	if err != nil {
		t.Errorf("Failed to CalculateProduction: %v", err)
		return
	}

	if v!=6.5 {
		t.Errorf("Production calculated wrong, got %v and expected 6.5", v)

	}


	err = DeleteCity(city.ID)
	if err != nil {
		t.Errorf("Failed to DeleteCity: %v", err)
		return
	}
}