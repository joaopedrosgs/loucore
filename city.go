package loucore

import (
	"context"
	"fmt"
	"github.com/ojrac/opensimplex-go"
	"math"
	"math/rand"
	"time"

	"github.com/joaopedrosgs/loucore/ent"
	"github.com/joaopedrosgs/loucore/ent/city"
	"github.com/joaopedrosgs/loucore/ent/construction"
	"github.com/joaopedrosgs/loucore/ent/queueitem"
	"github.com/joaopedrosgs/loucore/ent/user"
)

func CreateCity(x, y int) (*ent.City, error) {
	if x < 0 || y < 0 {
		return nil, fmt.Errorf("negative coordinates are not allowed")
	}
	c, err := client.City.Create().SetX(x).SetY(y).SetWoodLimit(1000).SetWoodProduction(5).Save(context.Background())
	if err != nil {
		return nil, err
	}
	_, err = client.Construction.Create().SetX(10).SetY(10).SetProduction(5).SetType(4).SetCity(c).Save(context.Background())
	return c, err
}
func CreateCityWithResources(x, y int) (*ent.City, error) {
	type pos struct {
		x int
		y int
	}
	noise := opensimplex.New(rand.Int63())
	createdNodes := make(map[int][]pos)
	chosenNodes := make(map[int][]pos)
	expectedValues := map[int]int{0: 18, 1: 18, 2: 18, 3: 2}

	c, err := CreateCity(x, y)
	if err != nil {
		return nil, err
	}

	for cityX := 0; cityX < 21; cityX++ {
		for cityY := 0; cityY < 21; cityY++ {
			if cityY == 10 && cityX == 10 {
				continue
			}
			xFloat := float64(cityX) / 5.0
			yFloat := float64(cityY) / 5.0
			randType := int(math.Abs(noise.Eval2(xFloat, yFloat))*4) + 1
			createdNodes[randType] = append(createdNodes[randType], pos{cityX, cityY})
		}
	}

	for nodeType := 0; nodeType < 4; nodeType++ {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for _, i := range r.Perm(len(createdNodes[nodeType])) {
			if expectedValues[nodeType] > len(chosenNodes[nodeType]) {
				chosenNodes[nodeType] = append(chosenNodes[nodeType], createdNodes[nodeType][i])
			}
		}
	}
	for nodeType := 0; nodeType < 4; nodeType++ {
		for _, node := range chosenNodes[nodeType] {
			_, _ = CreateConstruction(c.ID, node.x, node.y, nodeType)
		}
	}

	return c, err
}

func DeleteCity(cityID int) error {
	return client.City.DeleteOneID(cityID).Exec(context.Background())
}
func DeleteCityByCoord(cityX int, cityY int) (int, error) {
	return client.City.Delete().Where(city.XEQ(cityX), city.YEQ(cityY)).Exec(context.Background())
}
func CreateCityWithOwner(x, y, ownerID int) (*ent.City, error) {
	c, err := client.City.Create().SetX(x).SetY(y).SetWoodProduction(5).SetOwnerID(ownerID).Save(context.Background())
	if err != nil {
		return nil, err
	}
	_, err = client.Construction.Create().SetX(10).SetY(10).SetProduction(5).SetType(4).SetCity(c).Save(context.Background())
	return c, err
}

func CreateCityWithOwnerRandom(ownerID int) (*ent.City, error) {
	x, y := RandomCityCoordinates()
	return CreateCityWithOwner(x, y, ownerID)
}

func CitiesInRange(x, y, r int) ([]ent.City, error) {
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

	return client.City.Get(context.Background(), cityId)
}
func GetCitiesUserWithStructuresByID(userId int) ([]*ent.City, error) {

	return client.City.Query().WithConstructions().WithQueue().Where(city.HasOwnerWith(user.IDEQ(userId))).All(context.Background())
}

func UpdateCityResourcesUntil(cityId int, t time.Time) (*ent.City, error) {
	c, err := client.City.Get(context.Background(), cityId)
	if err != nil {
		return nil, err
	}

	secondsSinceLastUpdate := t.Sub(c.LastUpdated).Round(time.Second).Hours()

	newFoodStored := math.Min(c.FoodStored+c.FoodProduction*secondsSinceLastUpdate, c.FoodLimit)
	newStoneStored := math.Min(c.StoneStored+c.StoneProduction*secondsSinceLastUpdate, c.StoneLimit)
	newWoodStored := math.Min(c.WoodStored+c.WoodProduction*secondsSinceLastUpdate, c.WoodLimit)
	newIronStored := math.Min(c.IronStored+c.IronProduction*secondsSinceLastUpdate, c.IronLimit)

	return c.Update().
		SetWoodStored(newWoodStored).
		SetStoneStored(newStoneStored).
		SetFoodStored(newFoodStored).
		SetIronStored(newIronStored).
		Save(context.Background())
}

func UpdateCityStorage(cityId int) (*ent.City, error) {
	c, err := client.City.Query().WithConstructions(func(query *ent.ConstructionQuery) {
		query.Where(construction.TypeIn(10, 4))
	}).Where(city.IDEQ(cityId)).Only(context.Background())
	if err != nil {
		return nil, fmt.Errorf("(UpdateCityStorage) failed to get city with storage related constructions: %s", err.Error())
	}
	maxStorage := 0.0
	for _, c := range c.Edges.Constructions {
		maxStorage += Modules.Structures[c.Type].Bonus[c.Level] * c.Modifier
	}

	return c.Update().SetFoodLimit(maxStorage).SetIronLimit(maxStorage).SetStoneLimit(maxStorage).SetWoodLimit(maxStorage).Save(context.Background())
}

func UpdateCityProduction(cityId int) ([]float64, error) {
	var production []struct {
		Type string  `json:"type"`
		Sum  float64 `json:"sum"`
	}
	err := client.Construction.
		Query().
		Where(
			construction.HasCityWith(city.ID(cityId)),
			construction.TypeGTE(4),
			construction.TypeLTE(8)).
		GroupBy(construction.FieldType).
		Aggregate(
			ent.Sum(construction.FieldProduction)).
		Scan(context.Background(), &production)
	if err != nil {
		return nil, err
	}

	resources := make([]float64, 4)
	for _, prod := range production {
		switch prod.Type {
		case "5", "4": // wood
			resources[0] += prod.Sum
			break
		case "6": // stone
			resources[1] += prod.Sum
			break
		case "7": // food
			resources[2] += prod.Sum
			break
		case "8": // iron
			resources[3] += prod.Sum
			break
		}
	}
	return resources, client.City.UpdateOneID(cityId).
		SetWoodProduction(resources[0]).
		SetStoneProduction(resources[1]).
		SetIronProduction(resources[2]).
		SetFoodProduction(resources[3]).
		Exec(context.Background())
}
func AdvanceInTime(cityId int, until time.Time) error {
	c, err := client.City.Query().WithQueue(func(query *ent.QueueItemQuery) {
		query.Order(ent.Asc(queueitem.FieldPosition)).Limit(10)
	}).Where(city.ID(cityId)).First(context.Background())
	if err != nil {
		return err
	}
	start := c.QueueStartedAt
	for _, q := range c.Edges.Queue {
		completion := start.Add(time.Duration(q.Duration))
		if completion.Before(until) {
			err = CompleteQueueItem(q.ID)
			if err != nil {
				return err
			}
			start = completion
		} else {
			return nil
		}

		_, err = UpdateCityStorage(cityId)
		if err != nil {
			return err
		}

		_, err = UpdateCityResourcesUntil(cityId, completion)
		if err != nil {
			return err
		}

	}
	return nil

}
