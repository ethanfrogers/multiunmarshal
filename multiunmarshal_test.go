package multiunmarshal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalMultiple(t *testing.T) {
	// TODO: make this a table test

	type waterway struct {
		Canals []string `json:"canals"`
	}

	inputs := [][]byte{
		[]byte(`{"canals": ["panama", "suez"]}`),
		[]byte(`{"canals": []}`),
	}

	var w1 []waterway
	if err := UnmarshalMultiple(inputs, &w1); err != nil {
		t.Fatal(err.Error())
	}
	assert.EqualValues(t, []waterway{
		{Canals: []string{"panama", "suez"}},
		{Canals: []string{}},
	}, w1)

	var w2 []*waterway
	if err := UnmarshalMultiple(inputs, &w2); err != nil {
		t.Fatal(err.Error())
	}

	assert.EqualValues(t, []*waterway{
		{Canals: []string{"panama", "suez"}},
		{Canals: []string{}},
	}, w2)
}
