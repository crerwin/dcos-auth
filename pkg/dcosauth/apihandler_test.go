package dcosauth

import (
	"testing"
)

func TestParseLoginResponse(t *testing.T) {
	type response struct {
		token        string `json:"token"`
		anotherthing string `json:"anotherthing"`
	}
	cases := []struct {
		input []byte
		want  string
	}{
		{[]byte(`{"token": "b", "anotherthing": "cc"}`), "b"},
	}

	for _, c := range cases {
		got, _ := parseLoginResponse(c.input)
		if got != c.want {
			t.Errorf("parseLoginResponse(%v) == %v, wanted %v", c.input, got, c.want)
		}
	}
}
