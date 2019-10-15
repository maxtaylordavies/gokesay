package config

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Pokemon struct {
	Number int `json:"number"`
	Pokemon string `json:"pokemon"`
	Form string `json:"form,omitempty"`
	Say string `json:"say"`
}

func CreateJSON() error {
	var pokelist []Pokemon

	root := "config/images/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path == "config/images/" || path == "config/images/.DS_Store" {
			return nil
		}

		r, err := os.Open(path)
		if err != nil {
			return err
		}
		defer r.Close()

		data, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}

		path = strings.TrimPrefix(path, "config/images/")
		path = strings.TrimSuffix(path, ".png")
		slc := strings.Split(path, "-")

		num, err := strconv.Atoi(slc[0])
		if err != nil {
			return err
		}

		p := Pokemon{
			Number: num,
			Pokemon: slc[1],
			Say: base64.StdEncoding.EncodeToString(data),
		}
		pokelist = append(pokelist, p)

		return nil
	})
	if err != nil {
		return err
	}

	f, _ := json.MarshalIndent(pokelist, "", " ")
	_ = ioutil.WriteFile("static/pokemon.json", f, 0644)

	return nil
}
