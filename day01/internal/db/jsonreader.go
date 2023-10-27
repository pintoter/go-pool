package db

import (
	"Day01/internal/entity"
	"encoding/json"
	"encoding/xml"
)

type DBjson struct {
	Recipes entity.Recipe
}

func NewDBjson() *DBjson {
	return &DBjson{}
}

func (jr *DBjson) Read(data []byte) (entity.Recipe, error) {
	if err := json.Unmarshal(data, &jr.Recipes); err != nil {
		return entity.Recipe{}, err
	}

	return jr.Recipes, nil
}

func (jr *DBjson) Write(recipe entity.Recipe) (string, error) {
	data, err := xml.MarshalIndent(recipe, "", "    ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}
