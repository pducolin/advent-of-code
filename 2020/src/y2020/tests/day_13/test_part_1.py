from src.y2020.day_13.part_1 import solution

EXAMPLE = """939
7,13,x,x,59,x,31,19"""


def test_solution():
    assert solution(EXAMPLE) == 295
