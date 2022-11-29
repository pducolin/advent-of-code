from src.y2020.day_13.part_2 import find_departure_time

EXAMPLES = [
    {'example': '17,x,13,19', 'expected': 3417},
    {'example': '67,7,59,61', 'expected': 754018},
    {'example': '67,x,7,59,61', 'expected': 779210},
    {'example': '67,7,x,59,61', 'expected': 1261476},
    {'example': '1789,37,47,1889', 'expected': 1202161486},
]


def test_solution():
    for example in EXAMPLES:
        assert find_departure_time(example['example']) == example['expected']
