from src.y2020.day_05.part_1 import parse_column, parse_id, parse_row

EXAMPLES = [
    {'in': 'BFFFBBFRRR', 'expected_row': 70, 'expected_column': 7, 'expected_id': 567},
    {'in': 'FFFBBBFRRR', 'expected_row': 14, 'expected_column': 7, 'expected_id': 119},
    {'in': 'BBFFBBFRLL', 'expected_row': 102, 'expected_column': 4, 'expected_id': 820},
]


def test_raw():
    for example in EXAMPLES:
        assert parse_row(example['in']) == example['expected_row']


def test_col():
    for example in EXAMPLES:
        assert parse_column(example['in']) == example['expected_column']


def test_id():
    for example in EXAMPLES:
        assert parse_id(example['in']) == example['expected_id']
