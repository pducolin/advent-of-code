package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputData string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()

	fmt.Println("Running part", part)

	if part == 1 {
		fmt.Println(part1(inputData))
	} else {
		fmt.Println(part2(inputData))
	}
}

type Point struct {
	x int
	y int
}

func (point *Point) toString() string {
	return fmt.Sprintf("%d,%d", point.x, point.y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (point *Point) isTouching(otherPoint Point) bool {
	return abs(point.x-otherPoint.x) <= 1 && abs(point.y-otherPoint.y) <= 1
}

func (point *Point) stepTo(direction string) {
	if direction == "R" {
		point.x += 1
		return
	}

	if direction == "L" {
		point.x -= 1
		return
	}

	if direction == "U" {
		point.y -= 1
		return
	}

	if direction == "D" {
		point.y += 1
		return
	}
}

func part1(data string) string {
	moves := strings.Split(data, "\n")

	head := &Point{}
	tail := &Point{}

	visitedPositions := map[string]struct{}{}

	visitedPositions[tail.toString()] = struct{}{}

	for _, move := range moves {
		moveParts := strings.Split(move, " ")
		direction := moveParts[0]
		steps, err := strconv.Atoi(moveParts[1])
		if err != nil {
			panic(err)
		}

		// move
		for i := 0; i < steps; i++ {
			head.stepTo(direction)

			if tail.isTouching(*head) {
				continue
			}

			if tail.x == head.x {
				// on the same column
				// move diagonal
				//    T
				//    |
				// ---H---
				//    |

				//    |
				// ---H---
				//    |
				//    T
				if tail.y > head.y {
					tail.stepTo("U")
				} else {
					tail.stepTo("D")
				}
			} else if tail.y == head.y {
				// on the same row
				//    |
				// ---H--T
				//    |

				//    |
				// T--H---
				//    |
				if tail.x > head.x {
					tail.stepTo("L")
				} else {
					tail.stepTo("R")
				}
			} else {
				// move diagonal
				//    |  T
				// ---H---
				//    |

				// T  |
				// ---H---
				// 	  |

				//    |
				// ---H---
				//    |  T

				//    |
				// ---H---
				// T  |

				if tail.x > head.x {
					tail.stepTo("L")
				} else if tail.x < head.x {
					tail.stepTo("R")
				}
				if tail.y < head.y {
					tail.stepTo("D")
				} else if tail.y > head.y {
					tail.stepTo("U")
				}
			}

			visitedPositions[tail.toString()] = struct{}{}
		}
	}

	return strconv.Itoa(len(visitedPositions))
}

func part2(data string) string {
	moves := strings.Split(data, "\n")

	dots := make([]*Point, 10)

	for i := range dots {
		dots[i] = &Point{}
	}

	visitedPositions := map[string]struct{}{}

	visitedPositions[dots[9].toString()] = struct{}{}

	for _, move := range moves {
		moveParts := strings.Split(move, " ")
		direction := moveParts[0]
		steps, err := strconv.Atoi(moveParts[1])
		if err != nil {
			panic(err)
		}

		// move
		for i := 0; i < steps; i++ {
			dots[0].stepTo(direction)

			for otherDotIndex := 1; otherDotIndex < len(dots); otherDotIndex++ {

				if dots[otherDotIndex].isTouching(*dots[otherDotIndex-1]) {
					continue
				}

				if dots[otherDotIndex].x == dots[otherDotIndex-1].x {
					// on the same column
					// move diagonal
					//    T
					//    |
					// ---H---
					//    |

					//    |
					// ---H---
					//    |
					//    T
					if dots[otherDotIndex].y > dots[otherDotIndex-1].y {
						dots[otherDotIndex].stepTo("U")
					} else {
						dots[otherDotIndex].stepTo("D")
					}
				} else if dots[otherDotIndex].y == dots[otherDotIndex-1].y {
					// on the same row
					//    |
					// ---H--T
					//    |

					//    |
					// T--H---
					//    |
					if dots[otherDotIndex].x > dots[otherDotIndex-1].x {
						dots[otherDotIndex].stepTo("L")
					} else {
						dots[otherDotIndex].stepTo("R")
					}
				} else {
					// move diagonal
					//    |  T
					// ---H---
					//    |

					// T  |
					// ---H---
					// 	  |

					//    |
					// ---H---
					//    |  T

					//    |
					// ---H---
					// T  |

					if dots[otherDotIndex].x > dots[otherDotIndex-1].x {
						dots[otherDotIndex].stepTo("L")
					} else if dots[otherDotIndex].x < dots[otherDotIndex-1].x {
						dots[otherDotIndex].stepTo("R")
					}
					if dots[otherDotIndex].y < dots[otherDotIndex-1].y {
						dots[otherDotIndex].stepTo("D")
					} else if dots[otherDotIndex].y > dots[otherDotIndex-1].y {
						dots[otherDotIndex].stepTo("U")
					}
				}
			}

			visitedPositions[dots[9].toString()] = struct{}{}
		}
	}

	return strconv.Itoa(len(visitedPositions))
}
