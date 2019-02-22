package location_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"main/location"
	"testing"
)

func TestLocation_GetLocation(t *testing.T) {
	testPostCode1 := "BB11 1PX"
	testPostCode2 := "BB11 ZPX"

	tests := []struct {
		request location.Location
		expect  location.Location
		err     error
	}{
		{
			request: location.Location{
				PostCode: testPostCode1,
			},
			expect: location.Location{
				PostCode:  testPostCode1,
				Longitude: -2.2458968,
				Latitude:  53.79071,
				Street:    "Brick St",
			},
			err: nil,
		},
		{
			request: location.Location{
				PostCode: testPostCode2,
			},
			expect: location.Location{},
			err:    errors.New("invalid postcode"),
		},
	}

	for _, test := range tests {
		resp, err := test.request.GetLocation()
		if err != nil {
			assert.Equal(t, test.err, err)
		}
		assert.Equal(t, test.expect, resp)
	}
}
