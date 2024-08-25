package service

import (
	"io"
	"os"
	"time"

	"github.com/CriticalNoob02/sync-datasus/pkg/util"

	"github.com/jlaffaye/ftp"
)

func FtpLogin(username string, password string, url string) *ftp.ServerConn {
	conn, err := ftp.Dial(url, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		util.Logger.Fatal("Opss!", "err", err.Error())
	}

	err = conn.Login(username, password)
	if err != nil {
		util.Logger.Fatal("Opss!", "err", err.Error())
	}

	return conn
}

func FtpList(conn *ftp.ServerConn, path string) []string {
	files, err := conn.List(path)
	if err != nil {
		util.Logger.Fatal("Opss!", "err", err.Error())
	}

	var names []string
	for _, file := range files {
		names = append(names, file.Name)
	}
	return names
}

func FtpDownloadFile(conn *ftp.ServerConn, localFile string, remoteFile string) {
	f, err := os.Create(localFile)
	if err != nil {
		util.Logger.Fatal("Opss!", "err", err.Error())
	}
	defer f.Close()

	resp, err := conn.Retr(remoteFile)
	if err != nil {
		util.Logger.Fatal("Opss!", "err", err.Error())
	}
	defer resp.Close()

	_, err = io.Copy(f, resp)
	if err != nil {
		util.Logger.Fatal("Opss!", "err", err.Error())
	}
}
