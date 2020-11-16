package loucore

import (
	"github.com/joaopedrosgs/loucore/ent"
	"log"
	"os"
	"testing"
)

type TestEnv struct {
	accounts      []*ent.User
	cities        []*ent.City
	constructions []*ent.Construction
	queue         []*ent.QueueItem
}

var testEnv = TestEnv{
	accounts:      make([]*ent.User, 0),
	cities:        make([]*ent.City, 0),
	constructions: make([]*ent.Construction, 0),
	queue:         make([]*ent.QueueItem, 0),
}

func TestMain(m *testing.M) {
	err := SetupStorage()
	if err != nil {
		log.Fatal("Failed to SetupStorage: " + err.Error())
	}
	err = LoadModules("./modules/")
	if err != nil {
		log.Fatal("Failed to LoadModules because: " + err.Error())
	}
	if len(Modules.Structures) == 0 {
		log.Fatal("Silent error when loading modules")
	}
	err = CreateTestEnviroment()
	if err != nil {
		log.Fatal("Failed to TestMain because: " + err.Error())
	}
	code := m.Run()
	os.Exit(code)
}

func CreateTestEnviroment() error {
	account1, _ := CreateAccount("name", "email@email.com", "password")
	account2, _ := CreateAccount("name2", "email2@email.com", "password")
	account3, _ := CreateAccount("to_be_deleted", "to_be_deleted@email.com", "password")

	city1, _ := CreateCityWithOwner(1, 1, account1.ID)
	city2, _ := CreateCityWithOwner(3, 1, account1.ID)
	city3, _ := CreateCityWithOwner(3, 3, account2.ID)
	city4, _ := CreateCityWithOwner(3, 5, account2.ID)
	city5, _ := CreateCityWithOwner(3, 7, account2.ID)

	construction1, _ := CreateConstruction(city1.ID, 1, 1, WOOD)
	construction2, _ := CreateConstruction(city1.ID, 6, 6, LAKE)
	construction3, _ := CreateConstruction(city2.ID, 1, 2, IRON)
	construction4, _ := CreateConstruction(city2.ID, 1, 3, WOODCUTTERS_HUT)
	construction5, _ := CreateConstruction(city2.ID, 1, 4, WOOD)
	construction6, _ := CreateConstruction(city3.ID, 1, 3, WOODCUTTERS_HUT)
	construction7, _ := CreateConstruction(city3.ID, 1, 4, WOOD)
	construction8, _ := CreateConstruction(city4.ID, 1, 3, WOODCUTTERS_HUT)
	construction9, _ := CreateConstruction(city5.ID, 1, 3, WAREHOUSE)

	queue1, _ := CreateQueueItem(account1.ID, city1.ID, construction1.ID, 1)
	queue2, _ := CreateQueueItem(account1.ID, city1.ID, construction2.ID, 1)
	queue3, _ := CreateQueueItem(account1.ID, city2.ID, construction3.ID, 1)
	queue4, _ := CreateQueueItem(account1.ID, city2.ID, construction4.ID, 1)
	queue5, _ := CreateQueueItem(account1.ID, city2.ID, construction3.ID, 1)
	queue6, _ := CreateQueueItem(account2.ID, city3.ID, construction7.ID, 1)
	queue7, _ := CreateQueueItem(account2.ID, city3.ID, construction7.ID, 1)
	queue8, _ := CreateQueueItem(account2.ID, city3.ID, construction7.ID, 1)
	queue9, _ := CreateQueueItem(account2.ID, city4.ID, construction8.ID, 1)
	queue10, _ := CreateQueueItem(account2.ID, city4.ID, construction8.ID, 1)
	queue11, _ := CreateQueueItem(account2.ID, city4.ID, construction8.ID, 1)
	queue12, _ := CreateQueueItem(account2.ID, city4.ID, construction8.ID, 1)
	queue13, _ := CreateQueueItem(account2.ID, city4.ID, construction8.ID, 1)
	queue14, _ := CreateQueueItem(account2.ID, city4.ID, construction8.ID, 1)
	queue15, _ := CreateQueueItem(account2.ID, city4.ID, construction8.ID, 1)
	queue16, _ := CreateQueueItem(account2.ID, city4.ID, construction8.ID, 1)
	queue17, _ := CreateQueueItem(account2.ID, city4.ID, construction8.ID, 1)
	queue18, _ := CreateQueueItem(account2.ID, city4.ID, construction8.ID, 1)

	queue19, _ := CreateQueueItem(account2.ID, city5.ID, construction9.ID, 1)
	queue20, _ := CreateQueueItem(account2.ID, city5.ID, construction9.ID, 1)

	_ = CompleteQueueItem(queue19.ID)
	_ = CompleteQueueItem(queue20.ID)

	_, _ = UpdateAffectedBy(construction1.ID)
	_, _ = UpdateAffectedBy(construction2.ID)
	_, _ = UpdateAffectedBy(construction3.ID)
	_, _ = UpdateAffectedBy(construction4.ID)
	_, _ = UpdateAffectedBy(construction5.ID)

	_, _ = CalculateProduction(construction4.ID)

	_, _ = UpdateCityStorage(city1.ID)
	_, _ = UpdateCityStorage(city2.ID)
	_, _ = UpdateCityStorage(city5.ID)

	_, _ = UpdateCityProduction(city1.ID)
	_, _ = UpdateCityProduction(city2.ID)

	testEnv.accounts = append(testEnv.accounts, account1)
	testEnv.accounts = append(testEnv.accounts, account2)
	testEnv.accounts = append(testEnv.accounts, account3)

	testEnv.cities = append(testEnv.cities, city1)
	testEnv.cities = append(testEnv.cities, city2)
	testEnv.cities = append(testEnv.cities, city3)
	testEnv.cities = append(testEnv.cities, city4)
	testEnv.cities = append(testEnv.cities, city5)

	testEnv.constructions = append(testEnv.constructions, construction1)
	testEnv.constructions = append(testEnv.constructions, construction2)
	testEnv.constructions = append(testEnv.constructions, construction3)
	testEnv.constructions = append(testEnv.constructions, construction4)
	testEnv.constructions = append(testEnv.constructions, construction5)
	testEnv.constructions = append(testEnv.constructions, construction6)
	testEnv.constructions = append(testEnv.constructions, construction7)
	testEnv.constructions = append(testEnv.constructions, construction8)
	testEnv.constructions = append(testEnv.constructions, construction9)

	testEnv.queue = append(testEnv.queue, queue1)
	testEnv.queue = append(testEnv.queue, queue2)
	testEnv.queue = append(testEnv.queue, queue3)
	testEnv.queue = append(testEnv.queue, queue4)
	testEnv.queue = append(testEnv.queue, queue5)
	testEnv.queue = append(testEnv.queue, queue6)
	testEnv.queue = append(testEnv.queue, queue7)
	testEnv.queue = append(testEnv.queue, queue8)
	testEnv.queue = append(testEnv.queue, queue9)
	testEnv.queue = append(testEnv.queue, queue10)
	testEnv.queue = append(testEnv.queue, queue11)
	testEnv.queue = append(testEnv.queue, queue12)
	testEnv.queue = append(testEnv.queue, queue13)
	testEnv.queue = append(testEnv.queue, queue14)
	testEnv.queue = append(testEnv.queue, queue15)
	testEnv.queue = append(testEnv.queue, queue16)
	testEnv.queue = append(testEnv.queue, queue17)
	testEnv.queue = append(testEnv.queue, queue18)

	return nil
}
