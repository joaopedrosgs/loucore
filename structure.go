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
func SetAffectedBy(structureId int) {
}

func CalculateProduction(structureId int) (int, error) {
	structure, err := client.Construction.Query().WithAffectedBy().Where(construction.IDEQ(structureId)).Only(context.Background())
	if err != nil {
		return 0, err
	}
	structure.Modifier = 1

	structure.Production = Modules.Structures[structure.Type].Bonus[structure.Level]

	return structure.Production, client.Construction.UpdateOne(structure).Exec(context.Background())
}
