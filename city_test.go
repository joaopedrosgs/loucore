package loucore

import "testing"

func TestCreateCity(t *testing.T) {
	city, err := CreateCity(1, 1)
	if err != nil {
		t.Errorf("Failed to CreateCity: %v", err)
		return
	}
	if city == nil {
		t.Errorf("Failed silently to CreateCity")
	}
	construction, _ := GetStructuresFromCity(city.ID)
	if len(construction) == 0 || construction[0].Type != 4 {
		t.Errorf("Townhall not created correctly")
	}
	err = DeleteCity(city.ID)
	if err != nil {
		t.Errorf("Failed to DeleteCity: %v", err)
		return
	}

}
func TestDeleteCity(t *testing.T) {
	city, err := CreateCity(1, 1)
	if err != nil {
		t.Errorf("Failed to CreateCity: %v", err)
		return
	}
	if city == nil {
		t.Errorf("Failed silently to CreateCity")
	}
	err = DeleteCity(city.ID)
	if err != nil {
		t.Errorf("Failed to DeleteCity: %v", err)
		return
	}
}
func TestCreateCityWithOwner(t *testing.T) {
	user, _ := CreateAccount("test", "email@test.com", "pass")
	city, err := CreateCityWithOwner(1, 2, user.ID)
	if err != nil {
		t.Errorf("Failed to CreateCity: %v", err)
		return
	}
	if city == nil {
		t.Errorf("Failed silently to CreateCity")
	}
	cities, _ := GetCitiesUserWithStructuresByID(user.ID)
	if len(cities) == 0 {
		t.Errorf("User has 0 cities")
	}
	if len(cities[0].Edges.Constructions) == 0 {
		t.Errorf("Townhall not created correctly")
	}
	_ = DeleteAccount(user.ID)
	err = DeleteCity(city.ID)
	if err != nil {
		t.Errorf("Failed to DeleteCity: %v", err)
		return
	}
}

func TestCreateCityWithOwnerRandom(t *testing.T) {
	user, _ := CreateAccount("test", "email@test.com", "pass")
	city, err := CreateCityWithOwnerRandom(user.ID)
	if err != nil {
		t.Errorf("Failed to CreateCity: %v", err)
		return
	}
	if city == nil {
		t.Errorf("Failed silently to CreateCity")
	}
	cities, _ := GetCitiesUserWithStructuresByID(user.ID)
	if len(cities) == 0 {
		t.Errorf("User has 0 cities")
	}
	if len(cities[0].Edges.Constructions) == 0 {
		t.Errorf("Townhall not created correctly")
	}
	_ = DeleteAccount(user.ID)
	err = DeleteCity(city.ID)
	if err != nil {
		t.Errorf("Failed to DeleteCity: %v", err)
		return
	}
}

func TestUpdateCityProduction(t *testing.T) {
	city, err := CreateCity(5,5)
	if err != nil {
		t.Errorf("Failed to CreateCity: %v", err)
		return
	}
	err=UpdateCityProduction(city.ID)
	if err != nil {
		t.Errorf("Failed to UpdateCityProduction: %v", err)
		return
	}
	updatedCity,err:=GetCityById(city.ID)
	if err != nil {
		t.Errorf("Failed to updatedCity: %v", err)
		return
	}
	if updatedCity.WoodProduction!=300 {
		t.Errorf("Production not calculated correctly, woodProduction should be 300 and is %v", updatedCity.WoodProduction)
	}
	err = DeleteCity(city.ID)
	if err != nil {
		t.Errorf("Failed to DeleteCity: %v", err)
		return
	}


}