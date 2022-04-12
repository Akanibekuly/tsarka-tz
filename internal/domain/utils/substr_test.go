package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSt_FindSubstr(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{
			in:  "12345",
			out: "12345",
		},
		{
			in:  "abacc",
			out: "bac",
		},
		{
			in:  "somethinf12311sdasda",
			out: "somethinf123",
		},
	}

	for i, c := range cases {
		require.Equal(t, FindSubstr(c.in), c.out, i+1)
	}
}
