package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.`

func TestPart1(t *testing.T) {
	assert.Equal(t, "33", part1(data), "Failed testing part 1")
}

func TestFindMaxGeodes(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping part 2 tests")
	}
	bp1 := NewBlueprint("Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.")
	assert.Equal(t, 9, bp1.FindMaxGeodes(24), "Failed testing blueprint 1 with 24 minutes")
	assert.Equal(t, 56, bp1.FindMaxGeodes(32), "Failed testing blueprint 1 with 32 minutes")
	bp2 := NewBlueprint("Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.")
	assert.Equal(t, 12, bp2.FindMaxGeodes(24), "Failed testing blueprint 2 with 24 minutes")
	assert.Equal(t, 62, bp2.FindMaxGeodes(32), "Failed testing blueprint 2 with 24 minutes")
}
