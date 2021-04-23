package user

import (
	"bootcamp/entity"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	url := "/user"

	r := gin.Default()
	r.POST(
		url,
		CreateUserHandler(),
	)

	payload := entity.RegistrationUserEntity{
		FullName: "farhan ramadhana",
		Email:    "frans@kata.ai",
		Password: "password12345",
		Role:     "admin",
	}

	// PREPARING THE PAYLOAD
	marshalled, err := json.Marshal(payload)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	// LETS SEND THE PAYLOAD GO!
	req, err := http.NewRequest(
		http.MethodPost,
		url,
		strings.NewReader(string(marshalled)),
	)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	// rr IS THE RECORDER, WHICH LATER BE COMPARED AGAINST THE EXPECTED RESPONSE
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		t.Errorf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "success create user data", responseMap["status"])
}

func TestSplitFullName(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		result := SplitFullName("farhan")
		assert.Equal(t, "farhan", result[0])
	})
	t.Run("two", func(t *testing.T) {
		result := SplitFullName("farhan ramadhana")
		assert.Equal(t, "farhan", result[0])
		assert.Equal(t, "", result[1])
		assert.Equal(t, "ramadhana", result[2])
	})
	t.Run("three", func(t *testing.T) {
		result := SplitFullName("farhan muhammad ramadhana")
		assert.Equal(t, "farhan", result[0])
		assert.Equal(t, "muhammad", result[1])
		assert.Equal(t, "ramadhana", result[2])
	})
	t.Run("four", func(t *testing.T) {
		result := SplitFullName("stona stoni stino stani")
		assert.Equal(t, "stona", result[0])
		assert.Equal(t, "stoni stino", result[1])
		assert.Equal(t, "stani", result[2])
	})
}

func TestCreateInitialName(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		result := CreateInitialName("farhan")
		assert.Equal(t, "F", result)
	})
	t.Run("two", func(t *testing.T) {
		result := CreateInitialName("farhan ramadhana")
		assert.Equal(t, "FR", result)
	})
	t.Run("three", func(t *testing.T) {
		result := CreateInitialName("farhan ramadhana ganteng")
		assert.Equal(t, "FRG", result)
	})
	t.Run("four", func(t *testing.T) {
		result := CreateInitialName("farhan ramadhana ganteng banget")
		assert.Equal(t, "FRGB", result)
	})
}

func TestGenerateUserName(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		name := []string{"stona","",""}
		result := GenerateUserName(name)
		assert.Equal(t, "stona", result)
	})
	t.Run("two", func(t *testing.T) {
		name := []string{"stona", "", "stoni"}
		result := GenerateUserName(name)
		assert.Equal(t, "stona.stoni", result)
	})
	t.Run("three", func(t *testing.T) {
		name := []string{"stona", "stoni", "stino"}
		result := GenerateUserName(name)
		assert.Equal(t, "stona.stino", result)
	})
	t.Run("four", func(t *testing.T) {
		name := []string{"stona", "", "stoni"}
		result := GenerateUserName(name)
		assert.Equal(t, "stona.stoni1", result)
	})
	t.Run("five", func(t *testing.T) {
		name := []string{"stona", "", "stoni"}
		result := GenerateUserName(name)
		assert.Equal(t, "stona.stoni2", result)
	})
	t.Run("six", func(t *testing.T) {
		name := []string{"stona", "", "stoni"}
		result := GenerateUserName(name)
		assert.Equal(t, "stona.stoni3", result)
	})
}
