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
	if p.Save.Resources[Food].Amount != getDefaultSave(&p).Resources[Food].Amount {
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
