package test

import (
	"testing"

	"github.com/CriticalNoob02/sync-datasus/internal/config"
	"github.com/CriticalNoob02/sync-datasus/pkg/service"
)

func FtpLogin(t *testing.T) {
	conn := service.FtpLogin("anonymous", "anonymous", config.GetFtpUrl())
	want := "/"

	got, err := conn.CurrentDir()
	if err != nil {
		t.Logf("Err: %s", err)
		t.Error()
	}

	if got != want {
		t.Logf("1- Valor esperado: %s, Valor retornado: %s", want, got)
		t.Fail()
	}

}
