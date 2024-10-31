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
