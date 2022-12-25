package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const data = `root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32`

func TestPart1(t *testing.T) {
	assert.Equal(t, "152", part1(data), "Failed testing part 1")
}
