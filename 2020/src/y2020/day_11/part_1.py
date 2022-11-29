from copy import deepcopy

FLOOR = '.'
SEAT_EMPTY = 'L'
SEAT_OCCUPIED = '#'


def is_empty(seats, x, y):
    return seats[y][x] == SEAT_EMPTY


def is_occupied(seats, x, y):
    return seats[y][x] == SEAT_OCCUPIED


def can_go_N(seats, x, y):
    return y > 0


def can_go_E(seats, x, y):
    return x > 0


def can_go_W(seats, x, y):
    width = len(seats[0])
    return x < width - 1


def can_go_S(seats, x, y):
    height = len(seats)
    return y < height - 1


def can_go_NE(seats, x, y):
    return can_go_N(seats, x, y) and can_go_E(seats, x, y)


def can_go_NW(seats, x, y):
    return can_go_N(seats, x, y) and can_go_W(seats, x, y)


def can_go_SE(seats, x, y):
    return can_go_S(seats, x, y) and can_go_E(seats, x, y)


def can_go_SW(seats, x, y):
    return can_go_S(seats, x, y) and can_go_W(seats, x, y)


def count_neighbours(seats, x, y):
    neighbours_count = {SEAT_EMPTY: 0, SEAT_OCCUPIED: 0, FLOOR: 0}
    #     NE     N    NW
    #    y - 1|y - 1|y - 1
    #    x - 1|  x  |x + 1
    #    -----------------
    #  E   y  |  y  |  y    W
    #    x - 1|  x  |x + 1
    #    -----------------
    #    y + 1|y + 1|y + 1  SW
    #    x - 1|  x  |x + 1
    #      SE    S

    # check NE
    if can_go_NE(seats, x, y):
        ne = seats[y - 1][x - 1]
        neighbours_count[ne] += 1

    # check N
    if can_go_N(seats, x, y):
        n = seats[y - 1][x]
        neighbours_count[n] += 1

    # check NW
    if can_go_NW(seats, x, y):
        nw = seats[y - 1][x + 1]
        neighbours_count[nw] += 1

    # check W
    if can_go_W(seats, x, y):
        w = seats[y][x + 1]
        neighbours_count[w] += 1

    # check SW
    if can_go_SW(seats, x, y):
        sw = seats[y + 1][x + 1]
        neighbours_count[sw] += 1

    # check S
    if can_go_S(seats, x, y):
        s = seats[y + 1][x]
        neighbours_count[s] += 1

    # check SE
    if can_go_SE(seats, x, y):
        se = seats[y + 1][x - 1]
        neighbours_count[se] += 1

    # check E
    if can_go_E(seats, x, y):
        e = seats[y][x - 1]
        neighbours_count[e] += 1

    return neighbours_count


def execute_pass(seats):
    width = len(seats[0])
    height = len(seats)
    result_seats = deepcopy(seats)
    for x in range(width):
        for y in range(height):
            neighbours_count = count_neighbours(seats, x, y)
            # If a seat is empty (L) and there are no occupied seats adjacent to it,
            #  the seat becomes occupied.
            if is_empty(seats, x, y) and neighbours_count[SEAT_OCCUPIED] == 0:
                row = result_seats[y][:x] + SEAT_OCCUPIED
                if can_go_W(seats, x, y):
                    x_w = x + 1
                    row += result_seats[y][x_w:]
                result_seats[y] = row
            # If a seat is occupied (#) and four or more seats adjacent to it are also occupied,
            # the seat becomes empty.
            elif is_occupied(seats, x, y) and neighbours_count[SEAT_OCCUPIED] >= 4:
                row = result_seats[y][:x] + SEAT_EMPTY
                if can_go_W(seats, x, y):
                    x_w = x + 1
                    row += result_seats[y][x_w:]
                result_seats[y] = row
    return result_seats


def solution(data):
    seats = data.splitlines()
    while True:
        processed_seats = execute_pass(seats)
        if processed_seats == seats:
            break
        seats = processed_seats

    seats = ''.join(seats)
    # How many seats end up occupied?
    return seats.count(SEAT_OCCUPIED)
