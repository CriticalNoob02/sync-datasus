package test

import (
	"os"
	"testing"

	"github.com/CriticalNoob02/sync-datasus/pkg/util"
)

func TestMakeDir(t *testing.T) {
	name := "testes1"
	util.MakeDir(name)

	_, err := os.ReadDir(name)
	if err != nil {
		t.Logf("Err: %s", err)
		t.Error()
	}

	_, err = os.Open(name)
	if err != nil {
		t.Logf("Err: %s", err)
		t.Error()
	}
	os.Remove(name)
}
