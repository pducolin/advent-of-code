from src.y2020.day_16.part_2 import solution

EXAMPLE = """class: 0-1 or 4-19
departure a: 0-5 or 8-19
departure b: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9"""


def test_solution():
    assert solution(EXAMPLE) == 143
