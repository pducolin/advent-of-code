from src.y2020.day_12.part_2 import solution

EXAMPLE = """F10
N3
F7
R90
F11"""


def test_solution():
    assert solution(EXAMPLE) == 286
