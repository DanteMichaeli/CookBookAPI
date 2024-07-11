package graph

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/DanteMichaeli/CookBookAPI/graph/model"
)

const recipesDir = "recipes"

// checks if a recipe with the given id exists in the recipes directory
func idCheck(id string) error {
	fileName := fmt.Sprintf("%s.json", id)
	filePath := fmt.Sprintf("%s/%s", recipesDir, fileName)

	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("recipe with id %s not found", id)
		}
		return fmt.Errorf("error checking recipe file: %w", err)
	}

	return nil
}

// encodes a recipeInput file to a JSON file
func encodeRecipe(recipe *model.RecipeInput) ([]byte, error) {
	recipeJSON, err := json.Marshal(recipe)
	if err != nil {
		return nil, fmt.Errorf("error encoding recipe: %w", err)
	}

	return recipeJSON, nil
}

// writes JSON file to the recipes directory
func writeToDir(ID string, recipeJSON []byte) error {
	fileName := fmt.Sprintf("%s.json", ID)
	filePath := filepath.Join(recipesDir, fileName)
	err := os.WriteFile(filePath, recipeJSON, 0644)
	if err != nil {
		return fmt.Errorf("error writing recipe to file: %w", err)
	}

	return nil
}

// decodes from a JSON file to a recipe struct
