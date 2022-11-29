from src.y2020.day_09.part_1 import solution

EXAMPLE = """35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576"""


def test_solution():
    assert solution(EXAMPLE, 5) == 127
