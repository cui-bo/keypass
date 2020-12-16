package Models

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserUnmarshalJSON(t *testing.T) {
	data := []struct {
		title   string
		payload []byte
		result  User
		err     error
	}{
		{
			"A",
			[]byte(`{"name":"Bob", "phone":"0123456789", "address":"France"}`),
			User{Name: "Bob", Phone: "0123456789", Address:"France"},
			nil,
		},
		{
			"B",
			[]byte(`{"name":"Bob", "phone":"0123456789", "address":"France"}`),
			User{Name: "Bob", Phone: "0123456789", Address:"France"},
			errors.New("json: cannot unmarshal number into Go struct field name of type string"),
		},
	}

	for _, d := range data {
		var u User
		if err := json.Unmarshal(d.payload, &u); err != nil {
			if err.Error() == d.err.Error() {
				continue
			}
			t.Errorf("try to map payload data with unmarshal %v", err)
			continue
		}
		assert.Equal(t, u.Name, d.result.Name, "they should have the same name.")
		assert.Equal(t, u.Phone, d.result.Phone, "they should have the same phone.")
	}
}
