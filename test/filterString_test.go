package test

import (
	"testing"

	"github.com/CriticalNoob02/sync-datasus/pkg/util"
)

func TestFilterString_start(t *testing.T) {
	name := "AbelhossauroRex"

	got := util.FilterString("Abe", name, "start")
	want := true

	if got != want {
		t.Logf("1- Valor esperado: %t, Valor retornado: %t", want, got)
		t.Fail()
	}

	got = util.FilterString("sauro", name, "start")
	want = false

	if got != want {
		t.Logf("2- Valor esperado: %t, Valor retornado: %t", want, got)
		t.Fail()
	}
}

func TestFilterString_end(t *testing.T) {
	name := "AbelhossauroRex"

	got := util.FilterString("Rex", name, "end")
	want := true

	if got != want {
		t.Logf("1- Valor esperado: %t, Valor retornado: %t", want, got)
		t.Fail()
	}

	got = util.FilterString("sauro", name, "end")
	want = false

	if got != want {
		t.Logf("2- Valor esperado: %t, Valor retornado: %t", want, got)
		t.Fail()
	}
}

func TestFilterString_contain(t *testing.T) {
	name := "AbelhossauroRex"

	got := util.FilterString("ssauro", name, "contain")
	want := true

	if got != want {
		t.Logf("1- Valor esperado: %t, Valor retornado: %t", want, got)
		t.Fail()
	}

	got = util.FilterString("Abelha", name, "contain")
	want = false

	if got != want {
		t.Logf("2- Valor esperado: %t, Valor retornado: %t", want, got)
		t.Fail()
	}
}

func TestFilterString_containAny(t *testing.T) {
	name := "AbelhossauroRex"

	got := util.FilterString("lhos", name, "containAny")
	want := true

	if got != want {
		t.Logf("1- Valor esperado: %t, Valor retornado: %t", want, got)
		t.Fail()
	}

	got = util.FilterString("J22", name, "containAny")
	want = false

	if got != want {
		t.Logf("2- Valor esperado: %t, Valor retornado: %t", want, got)
		t.Fail()
	}
}
