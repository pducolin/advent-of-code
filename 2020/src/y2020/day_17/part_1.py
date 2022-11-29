ACTIVE = '#'
INACTIVE = '.'

DEBUG = False


def check_is_active(cube, x, y, z):
    cube_size = len(cube)
    if x not in range(cube_size) or y not in range(cube_size) or z not in range(cube_size):
        return False
    return cube[z][y][x] == ACTIVE


def count_active_neighbours(cube, x, y, z):
    counter_active = 0
    counter = 0
    for z_offset in range(-1, 2):
        for y_offset in range(-1, 2):
            for x_offset in range(-1, 2):
                if x_offset == 0 and y_offset == 0 and z_offset == 0:
                    continue
                counter += 1
                # visit neighbour
                if check_is_active(cube, x + x_offset, y + y_offset, z + z_offset):
                    counter_active += 1
    return counter_active


def count_active_points(cube):
    cube_size = len(cube)
    counter = 0
    for z in range(cube_size):
        for y in range(cube_size):
            for x in range(cube_size):
                if check_is_active(cube, x, y, z):
                    counter += 1
    return counter


def create_cube(cube_size):
    cube = []
    for z_index in range(cube_size):
        cube.append([list(INACTIVE * cube_size) for _ in range(cube_size)])
    return cube


def initialize_cube(data):
    initial_state = [list(x) for x in data.splitlines()]
    cube_size = len(initial_state)
    cube = create_cube(cube_size)
    cube[1] = initial_state
    return cube


def solution(data, iterations=6):
    cube = initialize_cube(data)

    for i in range(iterations):
        cube_size = len(cube)
        if DEBUG:
            print(f'Iteration {i + 1}: cube')
            for z in range(cube_size):
                print(f'z: {z - cube_size // 2}')
                for y in range(cube_size):
                    print(cube[z][y])
        new_cube_size = cube_size + 2
        new_cube = create_cube(new_cube_size)
        for z in range(new_cube_size):
            for y in range(new_cube_size):
                for x in range(new_cube_size):
                    # visit current point
                    is_active = check_is_active(cube, x - 1, y - 1, z - 1)
                    # visit neighbours
                    # cube vs new_cube have an offset of 1 in each plan
                    count_active = count_active_neighbours(cube, x - 1, y - 1, z - 1)
                    # apply rules
                    if is_active and count_active not in range(2, 4):
                        new_cube[z][y][x] = INACTIVE
                    elif not is_active and count_active == 3:
                        new_cube[z][y][x] = ACTIVE
                    else:
                        new_cube[z][y][x] = ACTIVE if is_active else INACTIVE
        cube = new_cube

    return count_active_points(cube)
