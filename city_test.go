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

}
