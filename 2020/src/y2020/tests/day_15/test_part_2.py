from src.y2020.day_15.part_2 import solution

EXAMPLE = """0,3,6: 175594
1,3,2: 2578
2,1,3: 3544142
1,2,3: 261214
2,3,1: 6895259
3,2,1: 18
3,1,2: 362"""


def test_solution():
    examples = [x.split(':') for x in EXAMPLE.splitlines()]
    for example in examples[:1]:
        assert solution(example[0]) == int(example[1])
