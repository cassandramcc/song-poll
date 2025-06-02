package data

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDataFromCSV(t *testing.T) {
	artists := GetDataFromCSV("test.csv")
	expTime := time.Time(time.Date(2024, time.May, 1, 0, 0, 0, 0, time.UTC))

	require.Len(t, artists, 3)
	assert.Equal(t, "Juche", artists[0].Name)
	assert.Equal(t, "Skeler", artists[1].Name)
	assert.Equal(t, "Ytho", artists[2].Name)

	assert.Equal(t, "3GZ8Rfap7VxAOzABiZEXCL", artists[0].URI)
	assert.Equal(t, "7ks4LdnBvp6HUsmVJiKgsB", artists[1].URI)
	assert.Equal(t, "5x0yHFCXMkXydFudgs6o3y", artists[2].URI)

	assert.Equal(t, expTime, artists[0].LastVisted)
	assert.Equal(t, expTime, artists[1].LastVisted)
	assert.Equal(t, expTime, artists[2].LastVisted)
}
