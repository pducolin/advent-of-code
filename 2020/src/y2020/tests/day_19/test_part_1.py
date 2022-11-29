from src.y2020.day_19.part_1 import parse_rules_dict, solve_rule, solution

EXAMPLE_1 = """0: 1 2
1: \"a\"
2: 1 3 | 3 1
3: \"b\""""


def test_parse_rules():
    expected = '(a)((a)(b)|(b)(a))'
    rules_dict = parse_rules_dict(EXAMPLE_1.splitlines())
    regex_dict = {}
    result = solve_rule(rules_dict, regex_dict, 0)
    assert result == expected


EXAMPLE_2 = """0: 45 1 5
1: 2 3 | 3 2
2: 45 45 | 5 5
3: 45 5 | 5 45
45: \"a\"
5: \"b\"

ababbb
bababa
abbbab
aaabbb
aaaabbb"""


def test_solution():
    assert solution(EXAMPLE_2) == 2
