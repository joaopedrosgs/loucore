package loucore

import (
	"context"
	"errors"

	"github.com/joaopedrosgs/loucore/ent"
	"github.com/joaopedrosgs/loucore/ent/city"
	"github.com/joaopedrosgs/loucore/ent/construction"
)

func CreateConstructionWithOwner(ownerId, cityId, x, y, cType int) (*ent.Construction, error) {
	return client.Construction.
		Create().
		SetOwnerID(ownerId).
		SetCityID(cityId).
		SetY(y).
		SetX(x).
		SetType(cType).
		Save(context.Background())

}
func CreateConstruction(cityId, x, y, cType int) (*ent.Construction, error) {
	return client.Construction.
		Create().
		SetCityID(cityId).
		SetY(y).
		SetX(x).
		SetType(cType).
		Save(context.Background())

}
func DeleteUpgrade(upgradeId int) error {
	return client.QueueItem.DeleteOneID(upgradeId).Exec(context.Background())
}

func GetStructuresFromCity(cityId int) ([]*ent.Construction, error) {
	return client.Construction.Query().Where(construction.HasCityWith(city.IDEQ(cityId))).All(context.Background())
}
func UpdateAffectedBy(structureId int) error {

	structure, err := client.Construction.Query().Where(construction.IDEQ(structureId)).WithCity().Only(context.Background())
	if err != nil {
		return err
	}
	if structure == nil {
		return errors.New("Structure not found")
	}
	neighbors, err := structure.Edges.City.QueryConstructions().Where(
		construction.XIn(structure.X-1, structure.X, structure.X+1),
		construction.YIn(structure.Y-1, structure.Y, structure.Y+1),
		construction.TypeIn(Modules.Structures[structure.Type].AffectedBy...)).
		Limit(8).All(context.Background())
	if err != nil {
		return err
	}
	_, err = client.Construction.UpdateOne(structure).ClearAffectedBy().AddAffectedBy(neighbors...).Save(context.Background())

	return err
}

func CalculateProduction(structureId int) (float64, error) {
	structure, err := client.Construction.Query().WithAffectedBy().Where(construction.IDEQ(structureId)).Only(context.Background())
	if err != nil {
		return 0, err
	}
	modifier := 100.0
	for _, affectingStructure := range structure.Edges.AffectedBy {
		modifier += Modules.Structures[affectingStructure.Type].Bonus[affectingStructure.Level]
	}
	production := Modules.Structures[structure.Type].Bonus[structure.Level] * modifier / 100.0

	return production, client.Construction.UpdateOne(structure).SetModifier(modifier).SetProduction(production).Exec(context.Background())
}
