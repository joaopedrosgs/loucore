package loucore

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := SetupStorage()
	if err != nil {
		log.Fatal("Failed to SetupStorage")
	}
	err = LoadModules("./modules/")
	if err != nil {
		log.Fatal("Failed to LoadModules because: "+err.Error())
	}
	if len(Modules.Structures) == 0 {
		log.Fatal("Silent error when loading modules")
	}
	code := m.Run()
	os.Exit(code)
}
