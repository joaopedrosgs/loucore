package loucore

import (
	"context"
	"github.com/joaopedrosgs/loucore/ent"
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

func CalculateProduction(structureId int) (int, error) {
	structure, err := client.Construction.Query().WithAffectedBy().Where(construction.IDEQ(structureId)).Only(context.Background())
	if err != nil {
		return 0, err
	}
	structure.Modifier = 0

	for _, modifier := range structure.Edges.AffectedBy {
		structure.Modifier += modifier.Modifier
	}
	structure.Production = (structure.RawProduction * structure.Modifier) / 100

	return structure.Production, client.Construction.UpdateOne(structure).Exec(context.Background())
}
