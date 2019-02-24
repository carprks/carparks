package service_test

import (
	"github.com/stretchr/testify/assert"
	"main/service"
	"testing"
)

func TestCreateModel_CreateCarPark(t *testing.T) {
	tests := []struct {
		request   service.CarPark
		expect    string
		err       error
		errString string
	}{
		{
			request: service.CarPark{
				Spaces: 5,
				Name:   "Tester",
				PostCode: "bb12 6np",
			},
			expect: "success",
			err:    nil,
		},
		{
			request: service.CarPark{
				Name: "Tester",
				PostCode: "bb12 6np",
			},
			expect:    "failure",
			errString: "no spaces",
		},
		{
			request: service.CarPark{
				Spaces: 1,
				PostCode: "bb12 6np",
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
