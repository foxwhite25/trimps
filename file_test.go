package trimps

import (
	"os"
	"testing"
)

func TestLoadEmptyAndSave(t *testing.T) {
	f, err := os.CreateTemp("", "test")
	_, err = f.WriteString("{}")
	if err != nil {
		t.Error(f)
	}
	f, err = os.Open(f.Name())
	if err != nil {
		t.Error(err)
	}
	p, err := LoadPlayer(f)
	if err != nil {
		t.Error(err)
	}
	if p.Save.Resources[Food].Amount != getDefaultSave(p).Resources[Food].Amount {
		t.Error("Default value does not work")
	}
	f, err = os.CreateTemp("", "test")
	if err != nil {
		t.Error(f)
	}
	err = SavePlayer(p, f)
	if err != nil {
		t.Error(err)
	}
}

func TestBasicFunction(t *testing.T) {
	f, err := os.CreateTemp("", "test")
	_, err = f.WriteString("{}")
	if err != nil {
		t.Error(f)
	}
	f, err = os.Open(f.Name())
	if err != nil {
		t.Error(err)
	}
	p, err := LoadPlayer(f)
	if err != nil {
		t.Error(err)
	}
	if !p.Save.ChangePlayerAction(Farming) {
		t.Error("Could not change to farming")
	}
	p.Tick(101)
	if !p.Save.ChangePlayerAction(Lumbering) {
		t.Error("Could not change to Lumber")
	}
	p.Tick(101)
	if !p.Save.Buildings[Trap].Buy(1) {
		t.Error("Could not buy trap")
	}
	p.Tick(51)
	for _, s := range p.PendingMessage {
		println(s)
	}
}
