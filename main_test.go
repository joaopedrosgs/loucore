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
	code := m.Run()
	os.Exit(code)
}
