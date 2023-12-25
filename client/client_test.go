package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetAll(t *testing.T) {
	testCases := []struct {
		testName string
		seedData []Unicorn
		expectErr bool
	}{
		{
			testName: "more items",
			seedData: []Unicorn{
				{
					Id: "65895754831c5703e89e3c89",
					Name:  "one",
					Age:   10,
					Colour: "red",
				},
				{
					Id: "65895759831c5703e89e3c8a",
					Name:  "two",
					Age:   10,
					Colour: "red",
				},
				{
					Id: "6589575e831c5703e89e3c8b",
					Name:  "three",
					Age:   10,
					Colour: "red",
				},
			},
			expectErr: false,
		},
	}
	

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://crudcrud.com/api", "f7b6da402e194650b3ce879659c04a50")

			items, err := client.GetAll()
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			if tc.seedData == nil {
				tc.seedData = []Unicorn{}
			}
			assert.Equal(t, tc.seedData, *items)
		})
	}
}

func TestClient_GetItem(t *testing.T) {
	testCases := []struct {
		testName     string
		itemId     string
		seedData     map[string]Unicorn
		expectErr    bool
		expectedResp *Unicorn
	}{
		{
			testName: "item exists",
			itemId: "65895754831c5703e89e3c89",
			seedData: map[string]Unicorn{
				"unicorn1": {
					Id: "65895754831c5703e89e3c89",
					Name:   "unicorn1",
					Age:    11,
					Colour: "blue",
				},
			},
			expectErr: false,
			expectedResp: &Unicorn{
				Id: "65895754831c5703e89e3c89",
				Name:   "one",
				Age:    10,
				Colour: "red",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://crudcrud.com/api", "f7b6da402e194650b3ce879659c04a50")

			item, err := client.GetItem(tc.itemId)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, item)
		})
	}
}

func TestClient_NewItem(t *testing.T) {
	testCases := []struct {
		testName  string
		newItem   *Unicorn
		seedData  map[string]Unicorn
		expectErr bool
	}{
		{
			testName: "success",
			newItem: &Unicorn{
				Name:   "four111",
				Age:    10,
				Colour: "red",
			},
			seedData:  nil,
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://crudcrud.com/api", "f7b6da402e194650b3ce879659c04a50")

			err := client.NewItem(tc.newItem)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			})
	}
}

func TestClient_UpdateItem(t *testing.T) {
	testCases := []struct {
		testName	string
		testId 		string
		updatedItem *Unicorn
		expectErr   bool
	}{
		{
			testName: "item exists",
			testId: "65896274831c5703e89e3cba",
			updatedItem: &Unicorn{
				Name:   "four",
				Age:    10,
				Colour: "red",
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://crudcrud.com/api", "f7b6da402e194650b3ce879659c04a50")

			err := client.UpdateItem(tc.testId, tc.updatedItem)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
		})
	}
}

func TestClient_DeleteItem(t *testing.T) {
	testCases := []struct {
		testName  string
		itemId  string
		expectErr bool
	}{
		{
			testName: "item exists",
			itemId: "6589627b831c5703e89e3cbb",
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://crudcrud.com/api", "f7b6da402e194650b3ce879659c04a50")

			err := client.DeleteItem(tc.itemId)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			_, err = client.GetItem(tc.itemId)
			assert.Error(t, err)
		})
	}
}
