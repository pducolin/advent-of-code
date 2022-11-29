from src.y2020.day_18.part_1 import evaluate_expression, solve_expression_list

EXAMPLES = [
    {'expression': '1 + 2 * 3 + 4 * 5 + 6', 'expected': 71},
    {'expression': '1 + (2 * 3) + (4 * (5 + 6))', 'expected': 51},
    {'expression': '2 * 3 + (4 * 5)', 'expected': 26},
    {'expression': '5 + (8 * 3 + 9 + 3 * 4 * 3)', 'expected': 437},
    {'expression': '5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))', 'expected': 12240},
    {'expression': '((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2', 'expected': 13632},
]


def test_evaluate_expression():
    for example in EXAMPLES:
        assert evaluate_expression(example['expression']) == example['expected']


def test_solve_expression_list():
    expected_sum = sum([x['expected'] for x in EXAMPLES])
    assert solve_expression_list([x['expression'] for x in EXAMPLES]) == expected_sum
