from src.y2020.day_17.part_1 import create_cube, initialize_cube, solution

EXAMPLE = """.#.
..#
###"""


def test_create_cube():
    assert create_cube(1) == [[['.']]]
    assert create_cube(1)[0][0][0] == '.'
    assert create_cube(3) == [
        [['.', '.', '.'], ['.', '.', '.'], ['.', '.', '.']],
        [['.', '.', '.'], ['.', '.', '.'], ['.', '.', '.']],
        [['.', '.', '.'], ['.', '.', '.'], ['.', '.', '.']],
    ]


def test_initialize_cube():
    assert initialize_cube(EXAMPLE) == [
        [['.', '.', '.'], ['.', '.', '.'], ['.', '.', '.']],
        [['.', '#', '.'], ['.', '.', '#'], ['#', '#', '#']],
        [['.', '.', '.'], ['.', '.', '.'], ['.', '.', '.']],
    ]


def test_solution():
    assert solution(EXAMPLE, 1) == 11
    assert solution(EXAMPLE, 2) == 21
    assert solution(EXAMPLE) == 112
