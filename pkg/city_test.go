package pkg

import (
	"reflect"
	"testing"
	"time"
)

func TestDeleteCity(t *testing.T) {
	city, err := CreateCity(88, 88)
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

func TestAdvanceInTime(t *testing.T) {
	type args struct {
		cityId int
		until  time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Advance 1 hour", args{testEnv.cities[0].ID, time.Now().Add(time.Duration(150) * time.Second)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AdvanceInTime(tt.args.cityId, tt.args.until); (err != nil) != tt.wantErr {
				t.Errorf("AdvanceInTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateCityWithResources(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test", args{14, 16}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateCityWithResources(tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCityWithResources() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
	for _, tt := range tests {
		_, _ = DeleteCityByCoord(tt.args.x, tt.args.y)
	}
}

func TestCreateCity(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Create City", args{99, 99}, false},
		{"Create City should fail because invalid coordinates", args{-99, 99}, true},
		{"Create City should fail because there is a city in this coordinates", args{1, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			city, err := CreateCity(tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if city != nil {
				DeleteCity(city.ID)
			}
		})
	}
}

func TestCreateCityWithOwner(t *testing.T) {
	type args struct {
		x       int
		y       int
		ownerID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Simple case", args{120, 120, testEnv.accounts[0].ID}, false},
		{"CreateCityWithOwner should fail because the owner doesn't exists", args{121, 121, -1}, true},
		{"CreateCityWithOwner should fail because the tile is already taken", args{testEnv.cities[0].X, testEnv.cities[0].Y, testEnv.accounts[0].ID}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateCityWithOwner(tt.args.x, tt.args.y, tt.args.ownerID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCityWithOwner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUpdateCityResourcesUntil(t *testing.T) {
	type args struct {
		cityId int
		t      time.Time
	}
	tests := []struct {
		name          string
		args          args
		wantResources []float64
		wantErr       bool
	}{
		{"Simple case", args{testEnv.cities[0].ID, time.Now().AddDate(0, 0, 1)}, []float64{420, 0, 0, 0}, false},
		{"Simple case", args{testEnv.cities[1].ID, time.Now().AddDate(0, 0, 1)}, []float64{576, 0, 0, 0}, false},

		{"UpdateCityResourcesUntil should fail since the city doesn't exists", args{-1, time.Now().AddDate(0, 0, 1)}, []float64{0, 0, 0, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateCityResourcesUntil(tt.args.cityId, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateCityResourcesUntil() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				if got.WoodStored != tt.wantResources[0] ||
					got.StoneStored != tt.wantResources[1] ||
					got.IronStored != tt.wantResources[2] ||
					got.FoodStored != tt.wantResources[3] {
					t.Errorf("UpdateCityResourcesUntil() %v != %v or %v != %v or %v != %v or %v != %v",
						got.WoodStored, tt.wantResources[0],
						got.StoneStored, tt.wantResources[1],
						got.IronStored, tt.wantResources[2],
						got.FoodStored, tt.wantResources[3])
					return
				}
			}

		})
	}
}

func TestUpdateCityStorage(t *testing.T) {
	type args struct {
		cityId int
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{"Simple case 1", args{testEnv.cities[0].ID}, []float64{5000, 5000, 5000, 5000}, false},
		{"Simple case 2", args{testEnv.cities[1].ID}, []float64{5000, 5000, 5000, 5000}, false},
		{"Simple case 3", args{testEnv.cities[2].ID}, []float64{5000, 5000, 5000, 5000}, false},
		{"Simple case 3", args{testEnv.cities[3].ID}, []float64{5000, 5000, 5000, 5000}, false},
		{"Simple case 3", args{testEnv.cities[4].ID}, []float64{14000, 14000, 14000, 14000}, false},

		{"UpdateCityStorage should fail since the city doesn't exists", args{-1}, []float64{0, 0, 0, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateCityStorage(tt.args.cityId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateCityStorage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				if tt.want[0] != got.WoodLimit ||
					tt.want[1] != got.StoneLimit ||
					tt.want[2] != got.IronLimit ||
					tt.want[3] != got.FoodLimit {
					t.Errorf("UpdateCityStorage()  %v != %v OR  %v != %v OR  %v != %v OR  %v != %v",
						tt.want[0], got.WoodLimit,
						tt.want[1], got.StoneLimit,
						tt.want[2], got.IronLimit,
						tt.want[3], got.FoodLimit)
				}
			}
		})
	}
}

func TestUpdateCityProduction(t *testing.T) {
	type args struct {
		cityId int
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr bool
	}{
		{"Simple case", args{testEnv.cities[0].ID}, []float64{5, 0, 0, 0}, false},
		{"Simple case1", args{testEnv.cities[1].ID}, []float64{11.5, 0, 0, 0}, false},
		{"Simple case2", args{-1}, []float64{0, 0, 0, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateCityProduction(tt.args.cityId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateCityProduction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateCityProduction() got = %v, want %v", got, tt.want)
			}
		})
	}
}
