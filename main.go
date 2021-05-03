package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mtakahashiif/gopractice/internal/untar"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	urlTempate := "https://example.com/it-automation-%s.tar.gz"
	version := "1.6.1"
	tarGzDir := "./tmp/download"
	untarDir := "./tmp/untar"

	tarGzFilePath, err := downloadTarGzFile(urlTempate, version, tarGzDir)
	if err != nil {
		log.Print(err)
		return
	}

	untarGzFile(tarGzFilePath, untarDir)
}

func downloadTarGzFile(urlTemplate string, version string, dir string) (tarGzFilePath string, err error) {

	url := fmt.Sprintf(urlTemplate, version)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		return
	}

	insecure := true
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: insecure,
			},
		},
	}

	response, err := client.Do(request)
	if err != nil {
		log.Print(err)
		return
	}

	defer response.Body.Close()

	fileName := fmt.Sprintf("it-automation-%s.tar.gz", version)
	tarGzFilePath = filepath.Join(dir, fileName)

	if err = os.MkdirAll(dir, 0755); err != nil {
		return
	}

	out, err := os.Create(tarGzFilePath)
	if err != nil {
		log.Print(err)
		return
	}

	defer out.Close()

	_, err = io.Copy(out, response.Body)
	if err != nil {
		log.Print(err)
		return
	}

	return
}

func untarGzFile(tarGzFilePath string, dir string) (err error) {
	tarGzFile, err := os.Open(tarGzFilePath)
	if err != nil {
		log.Print(err)
		return
	}

	defer tarGzFile.Close()

	err = untar.Untar(tarGzFile, dir)
	if err != nil {
		log.Print(err)
		return
	}

	return
}
