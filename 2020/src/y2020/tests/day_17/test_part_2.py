from src.y2020.day_17.part_2 import solution

EXAMPLE = """.#.
..#
###"""


def test_solution():
    assert solution(EXAMPLE) == 848
