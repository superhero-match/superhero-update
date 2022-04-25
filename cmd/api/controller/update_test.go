/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/superhero-match/superhero-update/internal/producer/model"
)

type MockProducer interface {
	Close() error
	UpdateSuperhero(s model.Superhero) error
}

type mockProducer struct {
	updateSuperhero func(s model.Superhero) error
}

func (m *mockProducer) Close() error {
	return nil
}

func (m *mockProducer) UpdateSuperhero(s model.Superhero) error {
	return m.updateSuperhero(s)
}

func mockPublishUpdateSuperhero(s model.Superhero) error {
	err := s.Validate()
	if err != nil {
		return err
	}

	var sb bytes.Buffer

	err = json.NewEncoder(&sb).Encode(s)
	if err != nil {
		return fmt.Errorf("encoder error")
	}

	return nil
}

type MockService interface {
	Close() error
	UpdateSuperhero(s model.Superhero) error
}

type mockService struct {
	mProducer mockProducer
}

func (srv *mockService) Close() error {
	return srv.mProducer.Close()
}

func (srv *mockService) UpdateSuperhero(s model.Superhero) error {
	return srv.mProducer.UpdateSuperhero(s)
}

func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func MockGet(c *gin.Context) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
}

func TestController_UpdateSuperhero(t *testing.T) {
	mockProd := mockProducer{
		updateSuperhero: mockPublishUpdateSuperhero,
	}

	mService := &mockService{
		mProducer: mockProd,
	}

	logger, err := zap.NewProduction()
	if err != nil {
		t.Fatal(err)
	}

	defer logger.Sync()

	mockController := &Controller{
		Service:    mService,
		Logger:     logger,
		TimeFormat: "2006-01-02T15:04:05",
	}

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	MockJsonPost(
		ctx,
		map[string]interface{}{
			"id":                    "test-id",
			"lookingForGender":      2,
			"age":                   28,
			"lookingForAgeMin":      25,
			"lookingForAgeMax":      45,
			"lookingForDistanceMax": 50,
			"distanceUnit":          "km",
			"lat":                   0.123456789,
			"lon":                   0.123456789,
			"country":               "Test Country",
			"city":                  "Test City",
			"superPower":            "Unit Testing",
			"accountType":           "FREE",
		},
	)

	mockController.UpdateSuperhero(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)
}

func TestController_UpdateSuperheroInvalidRequest(t *testing.T) {
	mockProd := mockProducer{
		updateSuperhero: mockPublishUpdateSuperhero,
	}

	mService := &mockService{
		mProducer: mockProd,
	}

	logger, err := zap.NewProduction()
	if err != nil {
		t.Fatal(err)
	}

	defer logger.Sync()

	mockController := &Controller{
		Service:    mService,
		Logger:     logger,
		TimeFormat: "2006-01-02T15:04:05",
	}

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	MockJsonPost(
		ctx,
		map[string]interface{}{
			"lookingForGender":      2,
			"age":                   28,
			"lookingForAgeMin":      25,
			"lookingForAgeMax":      45,
			"lookingForDistanceMax": 50,
			"distanceUnit":          "km",
			"lat":                   0.123456789,
			"lon":                   0.123456789,
			"country":               "Test Country",
			"city":                  "Test City",
			"superPower":            "Unit Testing",
			"accountType":           "FREE",
		},
	)

	mockController.UpdateSuperhero(ctx)
	assert.EqualValues(t, http.StatusInternalServerError, w.Code)
}
