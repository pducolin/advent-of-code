def solution(data):
    total = 2020
    complementars = set()
    for input_line in data.split():
        value = int(input_line)
        if value in complementars:
            result = value * (total - value)
            return result
        complementars.add(total - value)
