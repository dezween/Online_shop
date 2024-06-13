package tests

import (
	"bytes"
	"encoding/json"
	"go-shop/controllers"
	"go-shop/models"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Setup
	gin.SetMode(gin.TestMode)
	models.InitDB()

	// Clear database initially
	// clearDatabase()

	// Run tests
	code := m.Run()

	os.Exit(code)
}

// func clearDatabase() {
// 	models.DB.Exec("DELETE FROM users")
// 	models.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
// }

func TestRegisterUser(t *testing.T) {
	r := setupRouter()

	user := map[string]string{
		"login":       "testuser1",
		"password":    "testpassword1",
		"first_name":  "Test",
		"second_name": "User",
		"phone":       "1234567890",
		"email":       "test1@example.com",
	}

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.NotNil(t, response["data"])

	// Cleanup user after test
	defer models.DB.Exec("DELETE FROM users WHERE login = ?", user["login"])
}

func TestGetAllUsers(t *testing.T) {
	r := setupRouter()

	// Create a new user for testing
	user := models.User{
		Login:      "testuser2",
		Password:   "testpassword2",
		FirstName:  "Test",
		SecondName: "User",
		Phone:      "1234567891",
		Email:      "test2@example.com",
	}

	models.DB.Create(&user)
	defer models.DB.Exec("DELETE FROM users WHERE login = ?", user.Login)

	req, _ := http.NewRequest("GET", "/users", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.NotNil(t, response["data"])
}

func TestGetUserByID(t *testing.T) {
	r := setupRouter()

	// Create a new user for testing
	user := models.User{
		Login:      "testuser3",
		Password:   "testpassword3",
		FirstName:  "Test",
		SecondName: "User",
		Phone:      "1234567892",
		Email:      "test3@example.com",
	}

	models.DB.Create(&user)
	defer models.DB.Exec("DELETE FROM users WHERE login = ?", user.Login)

	req, _ := http.NewRequest("GET", "/users/"+strconv.Itoa(int(user.ID)), nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.NotNil(t, response["data"])
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controllers.Register)
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:id", controllers.GetUserByID)
	return r
}
