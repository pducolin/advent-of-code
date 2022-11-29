def calculate_fuel(value):
    count = 0
    while True:
        value = value // 3 - 2
        if value < 0:
            break
        count += value
    return count


def solution(data):
    return sum([calculate_fuel(int(x)) for x in data.splitlines()])


def test_calculate_fuel():
    assert calculate_fuel(1969) == 966
