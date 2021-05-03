package main

import (
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

	urlTempate := "https://github.com/exastro-suite/it-automation/releases/download/v%s/exastro-it-automation-%s.tar.gz"
	version := "1.7.0"
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

	downloadUrl := fmt.Sprintf(urlTemplate, version, version)

	request, err := http.NewRequest("GET", downloadUrl, nil)
	if err != nil {
		log.Print(err)
		return
	}

	tlsTransport := http.DefaultTransport.(*http.Transport).Clone()
	tlsTransport.TLSClientConfig.InsecureSkipVerify = true

	client := &http.Client{
		Transport: tlsTransport,
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
