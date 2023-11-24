package main

import (
	"fmt"
	"reflect"
)

type UnknownPlant struct {
	FlowerType string
	LeafType   string
	Color      int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
	FlowerColor int
	LeafType    string
	Height      int `unit:"inches"`
}

func describePlant(a interface{}) {
	typeOfA := reflect.TypeOf(a)
	typeName := typeOfA.Name()

	if typeOfA.Kind() != reflect.Struct {
		fmt.Println(("unsupported type:"), typeOfA)
		return
	}

	valueOfA := reflect.ValueOf(a)

	for i := 0; i < typeOfA.NumField(); i++ {
		fieldName := typeOfA.Field(i).Name
		fmt.Printf("%s", fieldName)

		var tagName, tagValue string
		if typeName == "UnknownPlant" && fieldName == "Color" {
			tagValue = typeOfA.Field(i).Tag.Get("color_scheme")
			tagName = "color_scheme"

		} else if typeName == "AnotherUnknownPlant" && fieldName == "Height" {
			tagValue = typeOfA.Field(i).Tag.Get("unit")
			tagName = "unit"
		}

		if tagValue != "" && tagName != "" {
			fmt.Printf("(%s=%s)", tagName, tagValue)
		}

		fmt.Printf(":%v\n", valueOfA.Field(i).Interface())
	}
}
