package db

import (
	"Day01/internal/entity"
	"encoding/json"
	"encoding/xml"
)

type DBxml struct {
	Recipes entity.Recipe
}

func NewDBxml() *DBxml {
	return &DBxml{}
}

func (xr *DBxml) Read(data []byte) (entity.Recipe, error) {
	if err := xml.Unmarshal(data, &xr.Recipes); err != nil {
		return entity.Recipe{}, err
	}

	return xr.Recipes, nil
}

func (xr *DBxml) Write(recipe entity.Recipe) (string, error) {
	data, err := json.MarshalIndent(recipe, "", "    ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}
