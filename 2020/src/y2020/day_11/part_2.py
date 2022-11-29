from copy import deepcopy

FLOOR = '.'
SEAT_EMPTY = 'L'
SEAT_OCCUPIED = '#'


def is_empty(seats, x, y):
    return seats[y][x] == SEAT_EMPTY


def is_occupied(seats, x, y):
    return seats[y][x] == SEAT_OCCUPIED


def is_floor(seats, x, y):
    return seats[y][x] == FLOOR


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
    neighbours_count = {SEAT_EMPTY: 0, SEAT_OCCUPIED: 0}
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

    x_center = x
    y_center = y
    # check NE
    while can_go_NE(seats, x, y):
        x = x - 1
        y = y - 1
        if is_floor(seats, x, y):
            continue
        ne = seats[y][x]
        neighbours_count[ne] += 1
        break

    x = x_center
    y = y_center

    # check N
    while can_go_N(seats, x, y):
        y = y - 1
        if is_floor(seats, x, y):
            continue
        n = seats[y][x]
        neighbours_count[n] += 1
        break

    x = x_center
    y = y_center

    # check NW
    while can_go_NW(seats, x, y):
        x = x + 1
        y = y - 1
        if is_floor(seats, x, y):
            continue
        nw = seats[y][x]
        neighbours_count[nw] += 1
        break

    x = x_center
    y = y_center

    # check W
    while can_go_W(seats, x, y):
        x = x + 1
        if is_floor(seats, x, y):
            continue
        w = seats[y][x]
        neighbours_count[w] += 1
        break

    x = x_center
    y = y_center

    # check SW
    while can_go_SW(seats, x, y):
        x = x + 1
        y = y + 1
        if is_floor(seats, x, y):
            continue
        sw = seats[y][x]
        neighbours_count[sw] += 1
        break

    x = x_center
    y = y_center

    # check S
    while can_go_S(seats, x, y):
        y = y + 1
        if is_floor(seats, x, y):
            continue
        s = seats[y][x]
        neighbours_count[s] += 1
        break

    x = x_center
    y = y_center

    # check SE
    while can_go_SE(seats, x, y):
        y = y + 1
        x = x - 1
        if is_floor(seats, x, y):
            continue
        se = seats[y][x]
        neighbours_count[se] += 1
        break

    x = x_center
    y = y_center

    # check E
    while can_go_E(seats, x, y):
        x = x - 1
        if is_floor(seats, x, y):
            continue
        e = seats[y][x]
        neighbours_count[e] += 1
        break

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
            # If a seat is occupied (#) and five or more seats adjacent to it are also occupied,
            # the seat becomes empty.
            elif is_occupied(seats, x, y) and neighbours_count[SEAT_OCCUPIED] >= 5:
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
