package pkg

import (
	"math"
	"testing"
)

func TestCalculateProduction(t *testing.T) {
	type args struct {
		structureId int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{"Production on WOOD", args{testEnv.constructions[0].ID}, 0, false},
		{"Production on LAKE", args{testEnv.constructions[1].ID}, 0, false},
		{"Production on IRON", args{testEnv.constructions[2].ID}, 0, false},
		{"Production on WOODCUTTERS_HUT", args{testEnv.constructions[3].ID}, 6.5, false},
		{"Production on WOOD", args{testEnv.constructions[4].ID}, 0, false},
		{"TestCalculateProduction should fail because the structure doesn't exists", args{-1}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateProduction(tt.args.structureId)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateProduction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateProduction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateConstruction(t *testing.T) {
	type args struct {
		cityId int
		x      int
		y      int
		cType  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Perfect Case", args{testEnv.cities[0].ID, 20, 20, WOODCUTTERS_HUT}, false},
		{"Two buildings at same spot should fail", args{testEnv.cities[0].ID, 20, 20, IRON_MINE}, true},
		{"Building off limits 1", args{testEnv.cities[0].ID, 22, 22, IRON_MINE}, true},
		{"Building off limits 2", args{testEnv.cities[0].ID, 22, -1, IRON_MINE}, true},
		{"Building off limits 3", args{testEnv.cities[0].ID, math.MaxInt32, math.MaxInt32, IRON_MINE}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateConstruction(tt.args.cityId, tt.args.x, tt.args.y, tt.args.cType)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateConstruction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	for _, tt := range tests {
		_, _ = DeleteConstructionByCityIDAndCoordinates(tt.args.cityId, tt.args.x, tt.args.y)
	}
}

func TestUpdateAffectedBy(t *testing.T) {
	type args struct {
		structureId int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{

		{"WOOD should not be affected", args{testEnv.constructions[0].ID}, 0, false},
		{"LAKE should not be affected", args{testEnv.constructions[1].ID}, 0, false},
		{"IRON  should not be affected", args{testEnv.constructions[2].ID}, 0, false},
		{"WOODCUTTERS_HUT should be affected by one wood", args{testEnv.constructions[3].ID}, 1, false},
		{"WOOD should not be affected", args{testEnv.constructions[4].ID}, 0, false},
		{"Should fail since the construction doesn't exists", args{6000}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateAffectedBy(tt.args.structureId)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAffectedBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateAffectedBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStructuresFromCity(t *testing.T) {
	type args struct {
		cityId int
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
		wantErr bool
	}{
		{"Common case 1", args{testEnv.cities[0].ID}, 3, false},
		{"Common case 2", args{testEnv.cities[1].ID}, 4, false},
		{"Should return 0 since the city doesn't exists", args{9999}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStructuresFromCity(tt.args.cityId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStructuresFromCity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("GetStructuresFromCity() got = %v, with len=%v want %v", got, len(got), tt.wantLen)
			}
		})
	}
}
