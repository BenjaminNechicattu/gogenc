package strings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverse(t *testing.T) {
	testString := "benjamin"
	expected := "nimajneb"
	resp := Reverse(testString)
	require.Equal(t, expected, resp)
}
