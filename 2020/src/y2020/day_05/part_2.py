ROWS = 128
COLUMNS = 8


def parse_row(raw_seat):
    raw_row = raw_seat[:7]
    row_index = 0
    for i in range(len(raw_row)):
        if raw_row[i] == 'F':
            continue
        row_index += ROWS // (2 ** (i + 1))
    return row_index


def parse_column(raw_seat):
    raw_column = raw_seat[-3:]
    column_index = 0
    for i in range(len(raw_column)):
        if raw_column[i] == 'L':
            continue
        column_index += COLUMNS // (2 ** (i + 1))
    return column_index


def evaluate_id(row, column):
    return row * 8 + column


def parse_id(raw_seat):
    row_index = parse_row(raw_seat)
    col_index = parse_column(raw_seat)
    return evaluate_id(row_index, col_index)


def solution(data):
    seats = [x for x in data.split('\n')]
    available_seats = {x for x in range(evaluate_id(127, 7))}
    occupied_seats = set()
    for seat in seats:
        seat_id = parse_id(seat)
        available_seats.remove(seat_id)
        occupied_seats.add(seat_id)

    available_seats = {x for x in available_seats if x + 1 in occupied_seats and x - 1 in occupied_seats}

    available_seats = [
        {'row': x // 8, 'col': x % 8, 'id': evaluate_id(x // 8, x % 8)}
        for x in available_seats
        if x // 8 > 0 and x // 8 < 127
    ]
    return available_seats
