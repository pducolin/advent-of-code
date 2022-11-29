from src.y2020.day_22.part_2 import solution

EXAMPLE = """Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10"""


def test_solution():
    assert solution(EXAMPLE) == 291
