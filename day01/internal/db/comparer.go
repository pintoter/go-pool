package db

import (
	"Day01/internal/entity"
	"fmt"
)

func CompareDB(oldRecipe, newRecipe entity.Recipe) {
	for _, newCake := range newRecipe.Cake {
		flag := true

		for _, oldCake := range oldRecipe.Cake {
			if newCake.Name == oldCake.Name {
				flag = false
				break
			}
		}

		if flag {
			fmt.Printf("ADDED cake \"%s\"\n", newCake.Name)
		}
	}

	for _, oldCake := range oldRecipe.Cake {
		flag := true

		for _, newCake := range newRecipe.Cake {
			if oldCake.Name == newCake.Name {
				if oldCake.Time != newCake.Time {
					fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldCake.Name, newCake.Time, oldCake.Time)
				}
				compareIngridients(oldCake.Ingridients, newCake.Ingridients, oldCake.Name)

				flag = false
				break
			}
		}

		if flag {
			fmt.Printf("REMOVED cake \"%s\"\n", oldCake.Name)
		}
	}
}

func compareIngridients(oldIngridients, newIngridients []entity.Ingridient, cakeName string) {
	for _, newIngridient := range newIngridients {
		flag := true

		for _, oldIngridient := range oldIngridients {
			if newIngridient.Name == oldIngridient.Name {
				flag = false
				break
			}
		}

		if flag {
			fmt.Printf("ADDED ingredient \"%s\" for cake  \"%s\"\n", newIngridient.Name, cakeName)
		}
	}

	for _, oldIngridient := range oldIngridients {
		flag := true

		for _, newIngridient := range newIngridients {
			if oldIngridient.Name == newIngridient.Name {
				flag = false
				break
			}
		}

		if flag {
			fmt.Printf("REMOVED ingredient \"%s\" for cake  \"%s\"\n", oldIngridient.Name, cakeName)
		}
	}

	for _, oldIngridient := range oldIngridients {
		for _, newIngridient := range newIngridients {
			if oldIngridient.Name == newIngridient.Name {
				compareComponentsIngridient(oldIngridient, newIngridient, cakeName)
				break
			}
		}
	}
}

func compareComponentsIngridient(oldIngridient, newIngridient entity.Ingridient, cakeName string) {
	if oldIngridient.Unit != newIngridient.Unit && newIngridient.Unit != "" && oldIngridient.Unit != "" {
		fmt.Printf("CHANGED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n", oldIngridient.Name, cakeName, newIngridient.Unit, oldIngridient.Unit)
	}

	if oldIngridient.Count != newIngridient.Count {
		fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n", oldIngridient.Name, cakeName, newIngridient.Count, oldIngridient.Count)
	}

	if oldIngridient.Unit == "" && newIngridient.Unit != "" {
		fmt.Printf("ADDED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n", oldIngridient.Name, cakeName, newIngridient.Unit, oldIngridient.Unit)
	}

	if oldIngridient.Unit != "" && newIngridient.Unit == "" {
		fmt.Printf("REMOVED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n", oldIngridient.Name, cakeName, newIngridient.Unit, oldIngridient.Unit)
	}
}
