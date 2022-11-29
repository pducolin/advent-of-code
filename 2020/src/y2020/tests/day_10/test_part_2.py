from src.y2020.day_10.part_2 import solution

EXAMPLE_1 = """16
10
15
5
1
11
7
19
6
12
4"""

EXAMPLE_2 = """1
2
3
6
7
8
11
12
15"""

EXAMPLE_3 = """28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3"""

EXAMPLE_4 = """1
2
3
4
6
7"""

EXAMPLE_5 = """1
2
3
5"""


def test_solution_1():
    assert solution(EXAMPLE_1) == 8


def test_solution_2():
    assert solution(EXAMPLE_2) == 8


def test_solution_3():
    assert solution(EXAMPLE_3) == 19208


def test_solution_4():
    assert solution(EXAMPLE_4) == 18


def test_solution_5():
    assert solution(EXAMPLE_5) == 6
