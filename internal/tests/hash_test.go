package tests

import (
	"testing"

	"github.com/masoudsadeghiDev/url-shortener/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestBase58(t *testing.T) {
	expected := "00BukQL"
	got := util.IntToBase58(123456789)
	assert.EqualValues(t, expected, got)
}
