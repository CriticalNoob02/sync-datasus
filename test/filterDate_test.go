package test

import (
	"testing"

	"github.com/CriticalNoob02/sync-datasus/pkg/util"
)

func TestFilterDate(t *testing.T) {
	name := "2101Abelhossauro"
	date := "20/01"
	mockDate := util.DataFilterStruct{
		MonthPosition: [2]int{0, 2},
		YearPosition:  [2]int{2, 4},
	}

	got, err := util.FilterDate(date, name, mockDate)
	want := true

	if err != nil {
		t.Logf("Err: %s", err)
		t.Error()
	}
	if got != want {
		t.Logf("1- Valor esperado: %t, Valor retornado: %t", want, got)
		t.Fail()
	}

	name = "1901Abelhossauro"

	got, err = util.FilterDate(date, name, mockDate)
	want = false

	if err != nil {
		t.Logf("Err: %s", err)
		t.Error()
	}
	if got != want {
		t.Logf("1- Valor esperado: %t, Valor retornado: %t", want, got)
		t.Fail()
	}

	t.Log("Passou!")
}
