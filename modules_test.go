package loucore

import (
	"testing"
)

func TestLoadModules(t *testing.T) {
	err := LoadModules("./modules/")
	if err != nil {
		t.Errorf("Failed to LoadModules because: "+err.Error())
	}
	if len(Modules.Structures) == 0 {
		t.Errorf("Silent error when loading modules")
	}
}
