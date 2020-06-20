package loucore

import (
	"context"
	"github.com/joaopedrosgs/loucore/ent"
	"github.com/joaopedrosgs/loucore/ent/city"
	"github.com/joaopedrosgs/loucore/ent/construction"
	"github.com/joaopedrosgs/loucore/ent/user"
)

func CitieInRange(x, y, r int) ([]ent.City, error) {
	var c []ent.City

	err := client.City.
		Query().
		Where(city.XLTE(x+r), city.XGTE(x-r), city.YLTE(y+r), city.YGTE(y-r)).
		Select(city.FieldName, city.FieldPoints, city.FieldX, city.FieldY).
		Scan(context.Background(), &c)

	return c, err
}
func CitiesUserByID(userId int) ([]*ent.City, error) {

	return client.City.Query().Where(city.HasOwnerWith(user.IDEQ(userId))).All(context.Background())

}
func CitiesUserWithStructuresByID(userId int) ([]*ent.City, error) {

	return client.City.Query().WithConstructions().WithQueue().Where(city.HasOwnerWith(user.IDEQ(userId))).All(context.Background())

}

func CityUpdateProduction(cityId int) error {
	var production []struct {
		Type string `json:"type"`
		Sum  int    `json:"sum"`
	}
	client.Construction.
		Query().
		Where(
			construction.HasCityWith(city.IDEQ(cityId)),
			construction.ProductionGT(0)).
		GroupBy(construction.FieldType).
		Aggregate(
			ent.Sum(construction.FieldProduction)).
		Scan(context.Background(), &production)

	return client.City.UpdateOneID(cityId).
		SetWoodProduction(production[0].Sum).
		SetStoneProduction(production[1].Sum).
		SetIronProduction(production[2].Sum).
		SetFoodProduction(production[3].Sum).
		Exec(context.Background())
}
