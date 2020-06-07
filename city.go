package loucore

import (
	"context"
	"log"
	"lou-core/ent/city"
	"lou-core/ent/user"
)

func CitieInRange(x int, y int, r int) interface{} {
	var c []struct {
		Name   string `json:"name"`
		Points int    `json:"points"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
	}
	err := client.City.
		Query().
		Where(city.XLTE(x+r), city.XGTE(x-r), city.YLTE(y+r), city.YGTE(y-r)).
		Select(city.FieldName, city.FieldPoints, city.FieldX, city.FieldY).
		Scan(context.Background(), &c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
func CitiesUserByID(userId int) interface{} {

	cities, err := client.City.Query().Where(city.HasOwnerWith(user.IDEQ(userId))).All(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return cities
}
func CitiesUserWithStructuresByID(userId int) interface{} {

	cities, err := client.City.Query().WithConstructions().WithQueue().Where(city.HasOwnerWith(user.IDEQ(userId))).All(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return cities
}
