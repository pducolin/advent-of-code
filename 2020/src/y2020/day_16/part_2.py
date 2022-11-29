def parse_rules(rules_raw):
    rules = {}
    for rule_raw in rules_raw.splitlines():
        [rule_name, rule_ranges] = [x.strip() for x in rule_raw.split(':')]
        rules[rule_name] = set()
        for r in rule_ranges.split('or'):
            [r_min, r_max] = [int(x) for x in r.split('-')]
            rules[rule_name].update(range(r_min, r_max + 1))
    return rules


def parse_ticket(ticket_raw):
    return [int(x.strip()) for x in ticket_raw.split(',')]


def is_valid_ticket(rules, ticket):
    found_counter = 0
    for value in ticket:
        for _, rule_range in rules.items():
            if value in rule_range:
                found_counter += 1
                break
    return found_counter == len(ticket)


def count_multiple_labels(possible_labels):
    counter = 0
    for labels in possible_labels:
        if len(labels) > 1:
            counter += 1
    return counter


def find_labels(rules, valid_tickets):
    possible_labels = []
    for i in range(len(valid_tickets[0])):
        possible_labels.append(set())
        values_at_i = set([x[i] for x in valid_tickets])
        for key, rule_range in rules.items():
            if values_at_i.issubset(rule_range):
                possible_labels[i].add(key)
    # find rules with only one label
    while count_multiple_labels(possible_labels) > 0:
        single_labels = [list(x)[0] for x in possible_labels if len(x) == 1]
        for single_label in single_labels:
            for i in range(len(possible_labels)):
                labels = possible_labels[i]
                if len(labels) > 1 and single_label in labels:
                    possible_labels[i].remove(single_label)
    return [list(x)[0] for x in possible_labels]


def solution(data):
    [rules_raw, my_ticket_raw, nearby_tickets_raw] = data.split('\n\n')
    rules = parse_rules(rules_raw)
    my_ticket = parse_ticket(my_ticket_raw.splitlines()[1].strip())
    nearby_tickets = []
    for ticket in nearby_tickets_raw.splitlines()[1:]:
        nearby_tickets.append(parse_ticket(ticket.strip()))
    valid_tickets = [x for x in nearby_tickets if is_valid_ticket(rules, x)]
    if is_valid_ticket(rules, my_ticket):
        valid_tickets.append(my_ticket)
    labels = find_labels(rules, valid_tickets)
    counter = 1
    for i in range(len(my_ticket)):
        label = labels[i]
        value = my_ticket[i]
        if label.startswith('departure'):
            counter *= value
    return counter
