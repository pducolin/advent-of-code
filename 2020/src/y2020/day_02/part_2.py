def solution(data):
    counter = 0
    lines = [x for x in data.split('\n')]
    for line in lines:
        rule, password = line.split(':')
        indexes, letter = rule.split()
        index_a, index_b = [int(x) for x in indexes.split('-')]
        letters_at_indexes = password[index_a] + password[index_b]
        if letters_at_indexes.count(letter) == 1:
            counter += 1
    return counter
