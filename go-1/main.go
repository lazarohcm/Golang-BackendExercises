package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type content struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func main() {

	dir := os.Args[1]
	jsonify(dir)
}

func jsonify(dir string) error {
	var contents []content

	err := filepath.Walk(dir, func(readingPath string, info os.FileInfo, err error) error {
		d, err := os.Stat(readingPath)
		if !d.Mode().IsRegular() {
			var files []string

			files, e := filepath.Glob(readingPath + "/*")
			if e != nil {
				log.Fatal(e)
			}
			for _, f := range files {
				if filepath.Ext(f) != "" {
					filename, file := filepath.Split(f)
					print(filename)
					contents = append(contents,
						content{file, "./" + f})

				}
			}

		}
		return nil
	})

	b, err := json.Marshal(contents)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	d1 := []byte(b)
	err = ioutil.WriteFile("./files.json", d1, 0644)

	return nil
}
