package engine

import (
	"archive/zip"
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func downloadZip(url string) ([]*zip.File, error) {
	resp, err := http.Get(JsonCollectionURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Unable to connect to the website")
	}
	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(buff)
	read, err := zip.NewReader(r, resp.ContentLength)
	return read.File, err
}
