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

package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/superhero-match/superhero-update/internal/producer/model"
)

var shouldGenerateEncodeError = false

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

	var encoderValue interface{}
	encoderValue = s

	if shouldGenerateEncodeError {
		encoderValue = make(chan int)
	}

	err = json.NewEncoder(&sb).Encode(encoderValue)
	if err != nil {
		return fmt.Errorf("encoder error")
	}

	return nil
}

func TestProducer_UpdateSuperhero(t *testing.T) {
	mockProd := &mockProducer{
		updateSuperhero: mockPublishUpdateSuperhero,
	}

	mockService := &service{
		Producer: mockProd,
	}

	tests := []struct {
		superhero               model.Superhero
		willGenerateEncodeError bool
		shouldReturnError       bool
		expected                error
	}{
		{
			superhero: model.Superhero{
				ID:                    "test-id",
				Email:                 "test@test.com",
				Name:                  "John Doe 1",
				SuperheroName:         "superJoe1",
				MainProfilePicURL:     "https://www.test-url.com",
				Gender:                1,
				LookingForGender:      2,
				Age:                   30,
				LookingForAgeMin:      25,
				LookingForAgeMax:      45,
				LookingForDistanceMax: 50,
				DistanceUnit:          "km",
				Lat:                   0.123456789,
				Lon:                   0.123456789,
				Birthday:              "1988-01-10",
				Country:               "Test Country",
				City:                  "Test City",
				SuperPower:            "Unit Testing",
				AccountType:           "FREE",
				IsDeleted:             false,
				DeletedAt:             "",
				IsBlocked:             false,
				BlockedAt:             "",
				UpdatedAt:             "",
				CreatedAt:             "",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       false,
			expected:                nil,
		},
		{
			superhero: model.Superhero{
				ID:                    "",
				Email:                 "test@test.com",
				Name:                  "John Doe 2",
				SuperheroName:         "superJoe2",
				MainProfilePicURL:     "https://www.test-url.com",
				Gender:                1,
				LookingForGender:      2,
				Age:                   30,
				LookingForAgeMin:      25,
				LookingForAgeMax:      45,
				LookingForDistanceMax: 50,
				DistanceUnit:          "km",
				Lat:                   0.123456789,
				Lon:                   0.123456789,
				Birthday:              "1988-01-10",
				Country:               "Test Country",
				City:                  "Test City",
				SuperPower:            "Unit Testing",
				AccountType:           "FREE",
				IsDeleted:             false,
				DeletedAt:             "",
				IsBlocked:             false,
				BlockedAt:             "",
				UpdatedAt:             "",
				CreatedAt:             "",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("field validation error"),
		},
		{
			superhero: model.Superhero{
				ID:                    "test-id",
				Email:                 "test@test.com",
				Name:                  "John Doe 3",
				SuperheroName:         "superJoe3",
				MainProfilePicURL:     "https://www.test-url.com",
				Gender:                1,
				LookingForGender:      2,
				Age:                   30,
				LookingForAgeMin:      25,
				LookingForAgeMax:      45,
				LookingForDistanceMax: 50,
				DistanceUnit:          "",
				Lat:                   0.123456789,
				Lon:                   0.123456789,
				Birthday:              "1988-01-10",
				Country:               "Test Country",
				City:                  "Test City",
				SuperPower:            "Unit Testing",
				AccountType:           "FREE",
				IsDeleted:             false,
				DeletedAt:             "",
				IsBlocked:             false,
				BlockedAt:             "",
				UpdatedAt:             "",
				CreatedAt:             "",
			},
			willGenerateEncodeError: false,
			shouldReturnError:       true,
			expected:                fmt.Errorf("field validation error"),
		},
		{
			superhero: model.Superhero{
				ID:                    "test-id",
				Email:                 "test@test.com",
				Name:                  "John Doe 4",
				SuperheroName:         "superJoe4",
				MainProfilePicURL:     "https://www.test-url.com",
				Gender:                1,
				LookingForGender:      2,
				Age:                   30,
				LookingForAgeMin:      25,
				LookingForAgeMax:      45,
				LookingForDistanceMax: 50,
				DistanceUnit:          "km",
				Lat:                   0.123456789,
				Lon:                   0.123456789,
				Birthday:              "1988-01-10",
				Country:               "Test Country",
				City:                  "Test City",
				SuperPower:            "Unit Testing",
				AccountType:           "FREE",
				IsDeleted:             false,
				DeletedAt:             "",
				IsBlocked:             false,
				BlockedAt:             "",
				UpdatedAt:             "",
				CreatedAt:             "",
			},
			willGenerateEncodeError: true,
			shouldReturnError:       true,
			expected:                fmt.Errorf("encoder error"),
		},
	}

	for _, test := range tests {
		shouldGenerateEncodeError = false

		if test.willGenerateEncodeError {
			shouldGenerateEncodeError = true
		}

		err := mockService.UpdateSuperhero(test.superhero)
		fmt.Println("err: ", err)
		fmt.Println("test.shouldReturnError: ", test.shouldReturnError)
		if test.shouldReturnError && err == nil {
			t.Fatal(err)
		}

		if test.shouldReturnError == false && err != nil {
			t.Fatal(err)
		}
	}
}
