package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestPart1(t *testing.T) {
	assert.Equal(t, "95437", part1(data), "Failed testing part 1")
}

func TestPart2(t *testing.T) {
	assert.Equal(t, "24933642", part2(data), "Failed testing part 2")
}
