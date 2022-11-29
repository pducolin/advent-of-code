from src.y2020.day_01.part_1 import solution

EXAMPLE = """1721
979
366
299
675
1456"""


def test_solution():
    assert solution(EXAMPLE) == 514579
