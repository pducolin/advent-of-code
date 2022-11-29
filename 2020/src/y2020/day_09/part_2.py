try:
    from src.y2020.day_09.part_1 import has_sum_in_previous_numbers
except ModuleNotFoundError:
    from part_1 import has_sum_in_previous_numbers


def find_sum_range(sum_value, numbers):
    from_index = 0
    to_index = 1
    while to_index < len(numbers):
        current_sum = sum(numbers[from_index:to_index])
        if current_sum == sum_value:
            return from_index, to_index
        if current_sum > sum_value and from_index < to_index:
            from_index += 1
        else:
            to_index += 1
    return from_index, to_index


def solution(data, lenght=25):
    numbers = [int(x) for x in data.splitlines()]
    for i in range(lenght, len(numbers)):
        if not has_sum_in_previous_numbers(i, lenght, numbers):
            invalid_number = numbers[i]
            from_index, to_index = find_sum_range(invalid_number, numbers[:i])
            return min(numbers[from_index:to_index]) + max(numbers[from_index:to_index])
