package migration

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIDFromFilename(t *testing.T) {
	input := "V0001__initial.sql"
	output := "0001__initial"

	require.Equal(t,
		output,
		IDFromFilename(input),
	)
}
