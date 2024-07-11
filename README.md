# cookbookAPI
Cook Book API built in Go and GraphQL (gqlgen)

# Tools used
- Golang for business logic functionality
- GraphQL for API communication
- **gqlgen**: schema-first tool for generating Go/GraphQL server infra and backend business logic automatically

# Core Functionalities (CRUD)
- create a new recipe
- list existing recipes: should serve images, ingredients, steps
- modify or delete existing recipes

# Stretch Goals
- integrate with a RDBMS (PostgreSQL?)
- user reviews and ratings
- refactoring & make types more logical
- response type for client calls

# Example Use Case:
**1. Create recipes:**
```graphql
mutation {
  createRecipe(
    id: "1"
    title: "Pasta Primavera"
    description: "Delicious pasta with fresh vegetables"
    ingredients: ["Pasta", "Broccoli", "Carrots", "Bell peppers", "Parmesan cheese"]
    steps: ["Cook pasta al dente", "Sauté vegetables", "Mix together and serve with Parmesan"]
  ) {
    message
    recipe {
      id
      title
      description
      ingredients
      steps
    }
  }
}

mutation {
  createRecipe(
    id: "2"
    title: "Chicken Alfredo"
    description: "Creamy pasta with grilled chicken"
    ingredients: ["Fettuccine", "Chicken breast", "Cream", "Garlic", "Parmesan cheese"]
    steps: ["Grill chicken", "Cook pasta", "Prepare sauce with cream, garlic, and cheese", "Combine and serve"]
  ) {
    message
    recipe {
      id
      title
      description
      ingredients
      steps
    }
  }
}

mutation {
  createRecipe(
    id: "3"
    title: "Caprese Salad"
    description: "Fresh salad with tomatoes, mozzarella, and basil"
    ingredients: ["Tomatoes", "Mozzarella cheese", "Fresh basil", "Olive oil", "Balsamic vinegar"]
    steps: ["Slice tomatoes and mozzarella", "Arrange on a plate with basil leaves", "Drizzle with olive oil and balsamic vinegar", "Season with salt and pepper"]
  ) {
    message
    recipe {
      id
      title
      description
      ingredients
      steps
    }
  }
}
```

**2. List All Recipes (or only some of them by providing ids):**
```graphql
query {
  recipes {
    message
    recipe {
      id
      title
      description
      ingredients
      steps
    }
  }
}
```

**3. Update Recipes:**
```graphql
mutation {
  updateRecipe(
    id: "1"
    title: "Updated Pasta Primavera"
    description: "New description for pasta dish"
    ingredients: ["Pasta", "Broccoli", "Carrots", "Bell peppers", "Parmesan cheese", "Cherry tomatoes"]
    steps: ["Cook pasta al dente", "Sauté vegetables", "Add cherry tomatoes", "Mix together and serve with Parmesan"]
  ) {
    message
    recipe {
      id
      title
      description
      ingredients
      steps
    }
  }
}

mutation {
  updateRecipe(
    id: "2"
    title: "Spaghetti Carbonara"
    description: "Traditional Italian pasta with eggs and pancetta"
    ingredients: ["Spaghetti", "Eggs", "Pancetta", "Parmesan cheese", "Black pepper"]
    steps: ["Cook spaghetti", "Fry pancetta", "Mix eggs and cheese", "Combine and serve with black pepper"]
  ) {
    message
    recipe {
      id
      title
      description
      ingredients
      steps
    }
  }
}
```
**4. List Updated Recipes:** see step 2

**5. Delete Recipe:**
```graphql
mutation {
  deleteRecipe(id: "3") {
    message
    recipe {
      id
      title
      description
      ingredients
      steps
    }
  }
}
```
... et.c.