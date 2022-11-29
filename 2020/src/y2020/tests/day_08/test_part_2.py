from src.y2020.day_08.part_2 import solution

EXAMPLE = """nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6"""


def test_solution():
    assert solution(EXAMPLE) == 8
