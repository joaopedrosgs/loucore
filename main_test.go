package loucore

import (
	"fmt"
	"github.com/joaopedrosgs/loucore/ent"
	"log"
	"os"
	"testing"
)

type TestEnv struct {
	accounts []*ent.User
	cities []*ent.City
	constructions []*ent.Construction
	queue []*ent.QueueItem
}

var testEnv = TestEnv{
	accounts:      make([]*ent.User,0),
	cities:         make([]*ent.City,0),
	constructions:  make([]*ent.Construction,0),
	queue:          make([]*ent.QueueItem,0),
}

func TestMain(m *testing.M) {
	err := SetupStorage()
	if err != nil {
		log.Fatal("Failed to SetupStorage")
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
	account,err:=CreateAccount("name", "email@email.com","password")
	if err != nil {
		return fmt.Errorf("Failed to CreateAccount because: " + err.Error())
	}
	city1,err:=CreateCityWithOwner(1,1,account.ID )
	if err != nil {
		return fmt.Errorf("Failed to CreateCityWithOwner because: " + err.Error())
	}
	construction1,err:=CreateConstruction(city1.ID, 1,1,0)
	if err != nil {
		return fmt.Errorf("Failed to CreateConstruction because: " + err.Error())
	}
	construction2, err := CreateConstruction(city1.ID, 6, 6, 1)
	if err != nil {
		return fmt.Errorf("Failed to CreateConstruction because: " + err.Error())
	}
	construction3, err := CreateConstruction(city1.ID, 1, 2, 3)
	if err != nil {
		return fmt.Errorf("Failed to CreateConstruction because: " + err.Error())
	}
	construction4,err := CreateConstruction(city1.ID, 1, 3, 6)
	if err != nil {
		return fmt.Errorf("Failed to CreateConstruction because: " + err.Error())
	}
	queue1,err:=CreateQueueItem(account.ID, city1.ID, construction1.ID, 1)
	if err != nil {
		return fmt.Errorf("Failed to CreateQueueItem because: " + err.Error())
	}
	testEnv.accounts = append(testEnv.accounts,account)
	testEnv.cities = append(testEnv.cities,city1 )
	testEnv.constructions = append(testEnv.constructions,construction1 )
	testEnv.constructions = append(testEnv.constructions,construction2 )
	testEnv.constructions = append(testEnv.constructions,construction3 )
	testEnv.constructions = append(testEnv.constructions,construction4 )
	testEnv.queue = append(testEnv.queue,queue1 )
	return nil
}