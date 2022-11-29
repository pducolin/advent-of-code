from src.y2020.day_10.part_1 import solution

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

EXAMPLE_2 = """28
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


def test_solution_1():
    assert solution(EXAMPLE_1) == 7 * 5


def test_solution_2():
    assert solution(EXAMPLE_2) == 22 * 10
