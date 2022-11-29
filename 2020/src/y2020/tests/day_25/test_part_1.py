from src.y2020.day_25.part_1 import evaluate_encryption_key, evaluate_loop_size, solution

EXAMPLE = """5764801
17807724"""


def test_evaluate_loop_size_card():
    assert evaluate_loop_size(5764801, 7) == 8


def test_evaluate_loop_size_door():
    assert evaluate_loop_size(17807724, 7) == 11


def test_evaluate_encryption_key_door():
    assert evaluate_encryption_key(5764801, 11) == 14897079


def test_evaluate_encryption_key_card():
    assert evaluate_encryption_key(17807724, 8) == 14897079


def test_solution():
    assert solution(EXAMPLE) == 14897079
