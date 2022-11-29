from src.y2020.day_07.part_2 import parse_rules, solution

EXAMPLE_1 = """light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags."""

EXAMPLE_2 = """shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags."""


def test_parse_rules():
    assert len(parse_rules(EXAMPLE_1)['shiny gold']) == len(
        [
            {'color': 'dark olive', 'count': 1},
            {'color': 'vibrant plum', 'count': 2},
        ]
    )
    assert len(parse_rules(EXAMPLE_2)['shiny gold']) == len([{'color': 'dark red', 'count': 2}])


def test_solution_1():
    assert solution(EXAMPLE_1) == 32


def test_solution_2():
    assert solution(EXAMPLE_2) == 126
