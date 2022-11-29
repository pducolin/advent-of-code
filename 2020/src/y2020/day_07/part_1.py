import re


def parse_rules(data):
    rules = {}
    raw_rules = [x for x in data.split('\n')]
    # pattern_parent: (group 1: 2 words separated by a space) bags contain
    pattern_parent = r'(\w+ \w+) bags contain '
    # pattern_child: 1 number (group 1: 2 words) (group 2: 'bags' or 'bag')
    pattern_child = r'\d+ (\w+ \w+) (bags|bag)'
    for rule in raw_rules:
        if 'no other bags' in rule:
            continue
        m = re.search(pattern_parent, rule)
        if not m:
            continue
        parent_color = m.group(1)
        children_start_index = len(m.group(0))
        children = [x.strip() for x in rule[children_start_index:].split(',')]
        for child in children:
            m = re.search(pattern_child, child)
            child_color = m.group(1).strip()
            if child_color not in rules:
                rules[child_color] = []
            rules[child_color].append(parent_color)
    return rules


def solution(data):
    rules = parse_rules(data)
    to_check_colors = rules['shiny gold']
    visited = set()
    counter = 0
    while len(to_check_colors) > 0:
        to_check = to_check_colors.pop()
        if to_check in visited:
            continue
        counter += 1
        visited.add(to_check)
        if to_check not in rules:
            continue
        to_check_colors.extend(rules[to_check])

    return counter
