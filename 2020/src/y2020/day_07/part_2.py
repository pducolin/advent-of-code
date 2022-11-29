import re


def parse_rules(data):
    rules = {}
    raw_rules = [x for x in data.split('\n')]
    # pattern_parent: (group 1: 2 words separated by a space) bags contain
    pattern_parent = r'(\w+ \w+) bags contain '
    # pattern_child: (group 1: 1 number) (group 2: 2 words) (group 3: 'bags' or 'bag')
    pattern_child = r'(\d+) (\w+ \w+) (bags|bag)'
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
            child_count = int(m.group(1).strip())
            child_color = m.group(2).strip()
            if parent_color not in rules:
                rules[parent_color] = []
            rules[parent_color].append({'color': child_color, 'count': child_count})
    return rules


def recurively_add_bags(rules, color):
    counter = 0
    inner_bags = rules.get(color, [])
    for bag in inner_bags:
        counter += bag['count'] + bag['count'] * recurively_add_bags(rules, bag['color'])
    return counter


def solution(data):
    rules = parse_rules(data)
    counter = recurively_add_bags(rules, 'shiny gold')
    return counter
