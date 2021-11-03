package domains

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	m = Map{
		"google.com":     1,
		"www.google.com": 2,
		".google.com":    3,
		".osnews.com":    4,
		"cnn.com":        5,
		// This ensures we have the ability to handle all traffic from a given
		// tld in a certain way.
		".ir": 5,
	}

	tree = NewTreeFromMap(Map(m))
)

func TestBestMatch(t *testing.T) {
	check := func(domain string, expected interface{}) {
		result, _ := tree.BestMatch(domain)
		assert.Equal(t, expected, result)
	}

	check("random.ir", 5)
	check("www.google.com", 2)
	check("www2.google.com", 3)
	check("sub.www2.google.com", 3)
	check("google.com", 1)
	check("osnews.com", 4)
	check("www.osnews.com", 4)
	check("cnn.com", 5)
	check("sub.cnn.com", nil)
	check("e.com", nil)
	check("google.com.cn", nil)
}

func TestGet(t *testing.T) {
	check := func(domain string, expected interface{}) {
		result, _ := tree.Get(domain)
		assert.Equal(t, expected, result)
	}

	check("www.google.com", 2)
	check("google.com", 1)
	check(".google.com", 3)
	check("www2.google.com", nil)
}

func TestWalk(t *testing.T) {
	found := make(Map, 0)
	tree.Walk(func(domain string, value interface{}) bool {
		found[domain] = value
		return true
	})

	assert.EqualValues(t, m, found)
}

func TestToMap(t *testing.T) {
	assert.EqualValues(t, m, tree.ToMap())
}
