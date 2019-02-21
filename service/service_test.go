package service_test

import (
	"github.com/stretchr/testify/assert"
	"main/service"
	"testing"
)

func TestCreateModel_CreateCarPark(t *testing.T) {
	tests := []struct {
		request   service.CreateModel
		expect    string
		err       error
		errString string
	}{
		{
			request: service.CreateModel{
				Spaces: 5,
				Name:   "Tester",
			},
			expect: "success",
			err:    nil,
		},
		{
			request: service.CreateModel{
				Name: "Tester",
			},
			expect:    "failure",
			errString: "no spaces",
		},
		{
			request: service.CreateModel{
				Spaces: 1,
			},
			expect:    "failure",
			errString: "name missing",
		},
	}

	for _, test := range tests {
		response, err := test.request.CreateCarPark()
		if err != nil {
			assert.EqualError(t, err, test.errString)
		}
		assert.Equal(t, test.expect, response)
	}
}
