package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsEmails(t *testing.T) {
	cases := []struct {
		s     string
		valid bool
	}{
		{"a.kanibekuly@gmail.com", true},
		{"test@gmail.com", true},
		{"something", false},
	}

	for i, c := range cases {
		require.Equal(t, isEmailValid(c.s), c.valid, i+1)
	}
}

func TestFindEmails(t *testing.T) {
	cases := []struct {
		in  string
		out []string
	}{
		{
			in:  "Email:  \na.kanibekuly@gmail.com",
			out: []string{"a.kanibekuly@gmail.com"},
		},
		{
			in:  "Email:  \na.kanibekuly@gmail.com asdvasv    Email:    \ttest@gmail.com",
			out: []string{"a.kanibekuly@gmail.com", "test@gmail.com"},
		},
	}

	for i, c := range cases {
		require.Equal(t, FindEmails(c.in), c.out, i+1)
	}
}
