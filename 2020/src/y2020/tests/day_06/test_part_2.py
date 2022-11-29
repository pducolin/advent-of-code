from src.y2020.day_06.part_2 import solution

EXAMPLE = """abc

a
b
c

ab
ac

a
a
a
a

b"""


def test_solution():
    assert solution(EXAMPLE) == 6
