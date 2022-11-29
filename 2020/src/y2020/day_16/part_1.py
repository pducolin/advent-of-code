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


def solution(data):
    [rules_raw, my_ticket_raw, nearby_tickets_raw] = data.split('\n\n')
    rules = parse_rules(rules_raw)
    # my_ticket = parse_ticket(my_ticket_raw.splitlines()[1].strip())
    nearby_tickets = []
    for ticket in nearby_tickets_raw.splitlines()[1:]:
        nearby_tickets.append(parse_ticket(ticket.strip()))
    ret = 0
    for ticket in nearby_tickets:
        for value in ticket:
            found = False
            for _, rule_range in rules.items():
                if value in rule_range:
                    found = True
                    break
            if not found:
                ret += value
    return ret
