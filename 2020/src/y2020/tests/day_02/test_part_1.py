from src.y2020.day_02.part_1 import solution

EXAMPLE = """1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc"""


def test_solution():
    assert solution(EXAMPLE) == 2
