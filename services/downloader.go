/*
Copyright © 2020 Jamie Phillips

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package services

import (
	"fmt"
	"github.com/apoorvam/goterminal"
	"io"
	"net/http"
	"os"
)

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	writer := goterminal.New(os.Stdout)

	fmt.Fprintf(writer, "Downloading %s bytes...\n", wc.Total)

	writer.Print()
	writer.Clear()
	writer.Reset()
}

func DownloadFile(url string, filepath string) error {
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	err = os.Rename(filepath+".tmp", filepath)
	if err != nil {
		return err
	}

	if err := os.Chmod(filepath, 0700); err != nil {
		return err
	}

	return nil
}
