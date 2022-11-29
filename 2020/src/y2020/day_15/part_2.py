def solution(data):
    counters = {}
    starting_numbers = [int(x) for x in data.split(',')]
    for i in range(30000000):
        if len(starting_numbers) > 0:
            current_number = starting_numbers.pop(0)
            counters[current_number] = [i]
            previous_number = current_number
            continue
        last_index = counters[previous_number][-1]
        if len(counters[previous_number]) > 1:
            first_index = counters[previous_number].pop(0)
        else:
            first_index = last_index
        current_number = last_index - first_index
        if current_number not in counters:
            counters[current_number] = []
        counters[current_number].append(i)
        previous_number = current_number

    return previous_number
