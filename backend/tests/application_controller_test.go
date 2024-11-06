package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/R-Thibault/OrgaJobSearch/backend/controllers"
	"github.com/R-Thibault/OrgaJobSearch/backend/models"
	serviceMocks "github.com/R-Thibault/OrgaJobSearch/backend/services/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestApplication_ApplicationCreationSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(serviceMocks.UserServiceInterface)
	mockApplicationService := new(serviceMocks.ApplicationServiceInterface)
	applicationController := controllers.NewApplicationController(mockUserService, mockApplicationService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	applicationDatas := models.Application{
		Url:         "https://fr.indeed.com/?vjk=13d574920531c12a&from=smart-apply&advn=9955624168329073",
		Title:       "Développeur PHP H/F",
		Company:     "Example Ltd",
		Location:    "Lyon",
		Description: "Blabla",
		Salary:      "30k",
		JobType:     "CDI",
		Applied:     true,
	}

	body, _ := json.Marshal(applicationDatas)
	c.Request, _ = http.NewRequest(http.MethodPost, "/application", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")

	c.Set("userUUID", "valid-uuid")
	mockUserService.On("GetUserByUUID", "valid-uuid").Return(&models.User{Model: gorm.Model{
		ID: 1,
	}}, nil)

	mockApplicationService.On("SaveApplication", uint(1), applicationDatas).Return(nil)

	applicationController.SaveApplication(c)

	assert.Equal(t, http.StatusOK, w.Code, "Expected status code 200, but got %v", w.Code)
	assert.Contains(t, w.Body.String(), "Application saved successfully")
	mockApplicationService.AssertExpectations(t)
}

func TestApplication_ApplicationCreationFail(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(serviceMocks.UserServiceInterface)
	mockApplicationService := new(serviceMocks.ApplicationServiceInterface)
	applicationController := controllers.NewApplicationController(mockUserService, mockApplicationService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	applicationDatas := models.Application{
		Url:         "https://fr.indeed.com/?vjk=13d574920531c12a&from=smart-apply&advn=9955624168329073",
		Title:       "Développeur PHP H/F",
		Company:     "Example Ltd",
		Location:    "Lyon",
		Description: "Blabla",
		Salary:      "30k",
		JobType:     "CDI",
		Applied:     true,
	}

	body, _ := json.Marshal(applicationDatas)
	c.Request, _ = http.NewRequest(http.MethodPost, "/application", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")

	c.Set("userUUID", "valid-uuid")
	mockUserService.On("GetUserByUUID", "valid-uuid").Return(&models.User{Model: gorm.Model{
		ID: 1,
	}}, nil)
	mockApplicationService.On("SaveApplication", uint(1), applicationDatas).Return(errors.New("Saving application failed"))

	applicationController.SaveApplication(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Saving application failed")
	mockApplicationService.AssertExpectations(t)
}

func TestApplication_ApplicationUpdateSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(serviceMocks.UserServiceInterface)
	mockApplicationService := new(serviceMocks.ApplicationServiceInterface)
	applicationController := controllers.NewApplicationController(mockUserService, mockApplicationService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	applicationDatas := models.Application{
		Model: gorm.Model{
			ID: 1,
		},
		Url:         "https://fr.indeed.com/?vjk=13d574920531c12a&from=smart-apply&advn=9955624168329073",
		Title:       "Développeur PHP H/F",
		Company:     "Example Ltd",
		Location:    "Lyon",
		Description: "Blabla",
		Salary:      "30k",
		JobType:     "CDI",
		Applied:     true,
	}

	body, _ := json.Marshal(applicationDatas)
	c.Request, _ = http.NewRequest(http.MethodPost, "/application", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")

	c.Set("userUUID", "valid-uuid")
	mockUserService.On("GetUserByUUID", "valid-uuid").Return(&models.User{Model: gorm.Model{
		ID: 1,
	}}, nil)
	// Use a concrete instance of *models.Application here
	updatedApplication := &models.Application{
		Model: gorm.Model{
			ID: 1,
		},
		Url:         applicationDatas.Url,
		Title:       applicationDatas.Title,
		Company:     applicationDatas.Company,
		Location:    applicationDatas.Location,
		Description: applicationDatas.Description,
		Salary:      applicationDatas.Salary,
		JobType:     applicationDatas.JobType,
		Applied:     applicationDatas.Applied,
	}
	mockApplicationService.On("UpdateApplication", uint(1), applicationDatas).Return(updatedApplication, nil)

	applicationController.UpdateApplication(c)

	// Parse JSON response to check specific fields
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Verify the response contains the updated application data
	assert.Contains(t, response, "UpdatedApplication")
	assert.Equal(t, float64(1), response["UpdatedApplication"].(map[string]interface{})["ID"])
	assert.Equal(t, "Développeur PHP H/F", response["UpdatedApplication"].(map[string]interface{})["Title"])

	mockApplicationService.AssertExpectations(t)
}

func TestApplication_ApplicationUpdateFail(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := new(serviceMocks.UserServiceInterface)
	mockApplicationService := new(serviceMocks.ApplicationServiceInterface)
	applicationController := controllers.NewApplicationController(mockUserService, mockApplicationService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	applicationDatas := models.Application{
		Model: gorm.Model{
			ID: 1,
		},
		Url:         "https://fr.indeed.com/?vjk=13d574920531c12a&from=smart-apply&advn=9955624168329073",
		Title:       "Développeur PHP H/F",
		Company:     "Example Ltd",
		Location:    "Lyon",
		Description: "Blabla",
		Salary:      "30k",
		JobType:     "CDI",
		Applied:     true,
	}

	body, _ := json.Marshal(applicationDatas)
	c.Request, _ = http.NewRequest(http.MethodPost, "/application", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-type", "application/json")

	c.Set("userUUID", "valid-uuid")
	mockUserService.On("GetUserByUUID", "valid-uuid").Return(&models.User{Model: gorm.Model{
		ID: 1,
	}}, nil)
	mockApplicationService.On("UpdateApplication", uint(1), applicationDatas).Return(nil, errors.New("Error during application update"))

	applicationController.UpdateApplication(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Error during application update")
	mockApplicationService.AssertExpectations(t)
}
