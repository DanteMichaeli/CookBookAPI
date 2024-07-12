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
func idExists(id string) error {
	fileName := fmt.Sprintf("%s.json", id)
	filePath := filepath.Join(recipesDir, fileName)

	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // File does not exist, which is expected
		}
		return fmt.Errorf("error checking recipe file: %w", err)
	}

	return fmt.Errorf("recipe with id %s already exists", id)
}

// checks if a recipe with the given id does NOT exist:
func idNotExist(id string) error {
	fileName := fmt.Sprintf("%s.json", id)
	filePath := filepath.Join(recipesDir, fileName)

	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("recipe with id %s does not exist:", id)
		}
		return fmt.Errorf("error checking recipe file: %w", err)
	}

	return nil
}

// encodes a recipeInput file to a JSON file
func encodeRecipe(recipe *model.Recipe) ([]byte, error) {
	recipeJSON, err := json.Marshal(recipe)
	if err != nil {
		return nil, fmt.Errorf("error encoding recipe: %w", err)
	}

	return recipeJSON, nil
}

// writes JSON file to the recipes directory
func writeToDir(id string, recipeJSON []byte) error {
	fileName := fmt.Sprintf("%s.json", id)
	filePath := filepath.Join(recipesDir, fileName)
	err := os.WriteFile(filePath, recipeJSON, 0644)
	if err != nil {
		return fmt.Errorf("error writing recipe to file: %w", err)
	}

	return nil
}

// decodes from a JSON file to a recipe struct
func decodeRecipe(id string) (*model.Recipe, error) {
	fileName := fmt.Sprintf("%s.json", id)
	filePath := filepath.Join(recipesDir, fileName)

	recipeFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading recipe file: %w", err)
	}

	recipe := &model.Recipe{}
	err = json.Unmarshal(recipeFile, recipe)
	if err != nil {
		return nil, fmt.Errorf("error decoding recipe: %w", err)
	}

	return recipe, nil
}

// lists all recipes if no id is provided
func listAll() ([]*model.Recipe, error) {
	files, err := os.ReadDir(recipesDir)
	if err != nil {
		return nil, fmt.Errorf("error reading recipes directory: %w", err)
	}

	recipes := []*model.Recipe{}
	for _, file := range files {
		fileName := file.Name()
		if filepath.Ext(fileName) == ".json" {
			fileName = fileName[:len(fileName)-len(".json")]
		}

		recipePtr, err := decodeRecipe(fileName)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipePtr)
	}

	return recipes, nil
}

// lists the recipes whose IDs are provided
func listWithID(id []string) ([]*model.Recipe, error) {
	recipes := []*model.Recipe{}
	for _, recipeID := range id {
		recipePtr, err := decodeRecipe(recipeID)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipePtr)
	}

	return recipes, nil
}
