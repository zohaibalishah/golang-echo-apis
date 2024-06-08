package main

import (
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	// e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":3000"))
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// e.GET("/users/:id", getUser)
func deleteUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	// err := deleteUserFromDB(id)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "User deleted successfully", "id": id, "status": 1})
}

// e.POST("/save", save)
func saveUser(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")

	// Validate name and email
	if name == "" || email == "" {
		response := map[string]interface{}{"error": "Invalid name or email", "status": 0}
		return c.JSON(http.StatusBadRequest, response)
	}

	userID := rand.Int()

	// Return result in JSON format
	result := map[string]interface{}{
		"id":     userID,
		"name":   name,
		"email":  email,
		"status": 1,
	}
	return c.JSON(http.StatusOK, result)
}
