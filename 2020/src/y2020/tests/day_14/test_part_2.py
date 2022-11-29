from src.y2020.day_14.part_2 import solution

EXAMPLE = """mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1"""


def test_solution():
    assert solution(EXAMPLE) == 208
