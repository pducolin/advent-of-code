def solution(data):
    counter = 0
    lines = [x for x in data.split('\n')]
    for line in lines:
        rule, password = line.split(':')
        min_max, letter = rule.split()
        rule_min, rule_max = [int(x) for x in min_max.split('-')]
        letter_in_password_count = password.count(letter)
        if letter_in_password_count >= rule_min and letter_in_password_count <= rule_max:
            counter += 1
    return counter
