from src.y2020.day_14.part_1 import solution

EXAMPLE = """mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0"""


def test_solution():
    assert solution(EXAMPLE) == 165
