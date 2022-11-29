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


def parse_id(raw_seat):
    row_index = parse_row(raw_seat)
    col_index = parse_column(raw_seat)
    return row_index * 8 + col_index


def solution(data):
    seats = [x for x in data.split('\n')]
    max_seat_id = -1
    for seat in seats:
        seat_id = parse_id(seat)
        if max_seat_id < seat_id:
            max_seat_id = seat_id
    return max_seat_id
