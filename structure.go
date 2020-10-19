package loucore

import (
	"context"
	"github.com/joaopedrosgs/loucore/ent"
	"github.com/joaopedrosgs/loucore/ent/city"
	"github.com/joaopedrosgs/loucore/ent/construction"
)

func CreateConstruction(ownerId, cityId, x, y, cType int) (*ent.Construction, error) {
	return client.Construction.
		Create().
		SetOwnerID(ownerId).
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
func SetAffectedBy(structureId int) error {
	structure, err := client.Construction.Query().Where(construction.IDEQ(structureId)).WithCity().Only(context.Background())
	if err != nil {
		return err
	}
	neighbors, err := client.Construction.Query().Where(
		construction.HasCityWith(
			city.ID(structure.Edges.City.ID)),
		construction.XIn(structure.X-1, structure.X+1),
		construction.YIn(structure.Y-1, structure.Y+1),
		construction.TypeIn(Modules.Structures[structure.Type].AffectedBy...)).
		Limit(8).All(context.Background())
	if err != nil {
		return err
	}
	_, err = client.Construction.UpdateOne(structure).AddAffectedBy(neighbors...).Save(context.Background())

	return err
}

func CalculateProduction(structureId int) (float64, error) {
	structure, err := client.Construction.Query().WithAffectedBy().Where(construction.IDEQ(structureId)).Only(context.Background())
	if err != nil {
		return 0, err
	}
	structure.Modifier = 1.0
	for _, affectingStructure := range structure.Edges.AffectedBy {
		structure.Modifier += Modules.Structures[affectingStructure.Type].Bonus[affectingStructure.Level]
	}
	structure.Production = Modules.Structures[structure.Type].Bonus[structure.Level]*structure.Modifier

	return structure.Production, client.Construction.UpdateOne(structure).Exec(context.Background())
}
