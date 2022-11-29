ACTIVE = '#'
INACTIVE = '.'

DEBUG = False


def check_is_active(hypercube, x, y, z, w):
    cube_size = len(hypercube)
    if x not in range(cube_size) or y not in range(cube_size) or z not in range(cube_size) or w not in range(cube_size):
        return False
    return hypercube[w][z][y][x] == ACTIVE


def count_active_neighbours(hypercube, x, y, z, w):
    counter_active = 0
    counter = 0
    for w_offset in range(-1, 2):
        for z_offset in range(-1, 2):
            for y_offset in range(-1, 2):
                for x_offset in range(-1, 2):
                    if x_offset == 0 and y_offset == 0 and z_offset == 0 and w_offset == 0:
                        continue
                    counter += 1
                    # visit neighbour
                    if check_is_active(hypercube, x + x_offset, y + y_offset, z + z_offset, w + w_offset):
                        counter_active += 1
    return counter_active


def count_active_points(hypercube):
    hypercube_size = len(hypercube)
    counter = 0
    for w in range(hypercube_size):
        for z in range(hypercube_size):
            for y in range(hypercube_size):
                for x in range(hypercube_size):
                    if check_is_active(hypercube, x, y, z, w):
                        counter += 1
    return counter


def create_hypercube(hypercube_size):
    hypercube = []
    for w_index in range(hypercube_size):
        cube = []
        for z_index in range(hypercube_size):
            cube.append([list(INACTIVE * hypercube_size) for _ in range(hypercube_size)])
        hypercube.append(cube)
    return hypercube


def initialize_hyper_cube(data):
    initial_state = [list(x) for x in data.splitlines()]
    hypercube_size = len(initial_state)
    hypercube = create_hypercube(hypercube_size)
    hypercube[1][1] = initial_state
    return hypercube


def solution(data, iterations=6):
    hypercube = initialize_hyper_cube(data)

    for i in range(iterations):
        hypercube_size = len(hypercube)
        if DEBUG:
            print(f'Iteration {i + 1}: hypercube')
            for z in range(hypercube_size):
                print(f'z: {z - hypercube_size // 2}')
                for y in range(hypercube_size):
                    print(hypercube[z][y])
        new_hypercube_size = hypercube_size + 2
        new_hypercube = create_hypercube(new_hypercube_size)
        for w in range(new_hypercube_size):
            for z in range(new_hypercube_size):
                for y in range(new_hypercube_size):
                    for x in range(new_hypercube_size):
                        # visit current point
                        is_active = check_is_active(hypercube, x - 1, y - 1, z - 1, w - 1)
                        # visit neighbours
                        # cube vs new_cube have an offset of 1 in each plan
                        count_active = count_active_neighbours(hypercube, x - 1, y - 1, z - 1, w - 1)
                        # apply rules
                        if is_active and count_active not in range(2, 4):
                            new_hypercube[w][z][y][x] = INACTIVE
                        elif not is_active and count_active == 3:
                            new_hypercube[w][z][y][x] = ACTIVE
                        else:
                            new_hypercube[w][z][y][x] = ACTIVE if is_active else INACTIVE
        hypercube = new_hypercube

    return count_active_points(hypercube)
