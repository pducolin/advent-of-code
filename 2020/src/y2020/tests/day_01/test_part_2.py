from src.y2020.day_01.part_2 import solution_space_const, solution_space_N

EXAMPLE = """1721
979
366
299
675
1456"""


def test_solution_space_const():
    assert solution_space_const(EXAMPLE) == 241861950


def test_solution_space_N():
    assert solution_space_N(EXAMPLE) == 241861950
