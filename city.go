package loucore

import (
	"context"
	"github.com/joaopedrosgs/loucore/ent"
	"github.com/joaopedrosgs/loucore/ent/city"
	"github.com/joaopedrosgs/loucore/ent/construction"
	"github.com/joaopedrosgs/loucore/ent/user"
	"math"
	"time"
)

func CreateCity(x, y int) (*ent.City, error) {
	c, err := client.City.Create().SetX(x).SetY(y).SetWoodProduction(300).Save(context.Background())
	if err != nil {
		return nil, err
	}
	_, err = client.Construction.Create().SetX(10).SetY(10).SetProduction(300).SetType(4).SetCity(c).Save(context.Background())
	return c, err
}

func CreateCityWithOwner(x, y, ownerId int) (*ent.City, error) {
	c, err := client.City.Create().SetX(x).SetY(y).SetWoodProduction(5).SetOwnerID(ownerId).Save(context.Background())
	if err != nil {
		return nil, err
	}
	_, err = client.Construction.Create().SetX(10).SetY(10).SetType(4).SetCity(c).Save(context.Background())
	return c, err
}

func CreateCityWithOwnerRandom(ownerId int) (*ent.City, error) {
	x, y := RandomCityCoordinates()
	return CreateCityWithOwner(x, y, ownerId)
}

func CitieInRange(x, y, r int) ([]ent.City, error) {
	var c []ent.City

	err := client.City.
		Query().
		Where(city.XLTE(x+r), city.XGTE(x-r), city.YLTE(y+r), city.YGTE(y-r)).
		Select(city.FieldName, city.FieldPoints, city.FieldX, city.FieldY).
		Scan(context.Background(), &c)

	return c, err
}
func GetCitiesUserByID(userId int) ([]*ent.City, error) {

	return client.City.Query().Where(city.HasOwnerWith(user.IDEQ(userId))).All(context.Background())

}
func GetCityById(cityId int) (*ent.City, error) {

	return client.City.Get(context.Background(),cityId)

}
func GetCitiesUserWithStructuresByID(userId int) ([]*ent.City, error) {

	return client.City.Query().WithConstructions().WithQueue().Where(city.HasOwnerWith(user.IDEQ(userId))).All(context.Background())

}
func Min(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func UpdateCityResources(cityId int) error {
	city, err := client.City.Get(context.Background(), cityId)
	if err != nil {
		return err
	}

	secondsSinceLastUpdate := time.Now().Sub(city.LastUpdated).Round(time.Second).Seconds()

	newFoodStored := math.Min(city.FoodStored+city.FoodProduction*secondsSinceLastUpdate, city.FoodLimit)
	newStoneStored:= math.Min(city.StoneStored+city.StoneProduction*secondsSinceLastUpdate, city.StoneLimit)
	newWoodStored := math.Min(city.WoodStored+city.WoodProduction*secondsSinceLastUpdate, city.WoodLimit)
	newIronStored := math.Min(city.IronStored+city.IronProduction*secondsSinceLastUpdate, city.IronLimit)

	_,err=city.Update().
		SetWoodStored(newWoodStored).
		SetStoneStored(newStoneStored).
		SetFoodStored(newFoodStored).
		SetIronStored(newIronStored).
		Save(context.Background())

	return err
}

func UpdateCityProduction(cityId int) error {
	var production []struct {
		Type string `json:"type"`
		Sum  float64    `json:"sum"`
	}
	client.Construction.
		Query().
		Where(
			construction.HasCityWith(city.ID(cityId)),
			construction.ProductionGT(0),
			construction.TypeGTE(4),
			construction.TypeLTE(8)).
		GroupBy(construction.FieldType).
		Aggregate(
			ent.Sum(construction.FieldProduction)).
		Scan(context.Background(), &production)
	var wood = 0.0
	var stone = 0.0
	var iron = 0.0
	var food = 0.0
	for _, prod := range production {
		switch prod.Type {
		case "5", "4": // wood
			wood += prod.Sum
			break
		case "6": // stone
			stone = prod.Sum
			break
		case "7": // food
			food = prod.Sum
			break
		case "8": // iron
			wood = prod.Sum
			break
		}
	}
	return client.City.UpdateOneID(cityId).
		SetWoodProduction(wood).
		SetStoneProduction(stone).
		SetIronProduction(iron).
		SetFoodProduction(food).
		Exec(context.Background())
}
