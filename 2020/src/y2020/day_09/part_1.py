def has_sum_in_previous_numbers(target_index, lenght, numbers):
    from_index = max(0, target_index - lenght)
    for i in range(from_index, target_index):
        if numbers[i] > numbers[target_index]:
            continue
        complementar = numbers[target_index] - numbers[i]
        next_index = i + 1
        if complementar in set(numbers[from_index:i]) or complementar in set(numbers[next_index:target_index]):
            return True
    return False


def solution(data, lenght=25):
    numbers = [int(x) for x in data.splitlines()]
    for i in range(lenght, len(numbers)):
        if not has_sum_in_previous_numbers(i, lenght, numbers):
            return numbers[i]
