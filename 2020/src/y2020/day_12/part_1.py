from collections import namedtuple

DIRECTIONS = ['E', 'S', 'W', 'N']

Position = namedtuple('Position', ['x', 'y', 'direction'])
Instruction = namedtuple('Instruction', ['direction', 'value'])


def parse_instruction(instruction):
    direction = instruction[:1]
    value = int(instruction[1:])
    return Instruction(direction, value)


def move(position, instruction):
    if instruction.direction in DIRECTIONS:
        return move_direction(position, instruction.direction, instruction.value)
    if instruction.direction == 'F':
        return move_direction(position, position.direction, instruction.value)

    return turn(position, instruction.direction, instruction.value)


def move_direction(position, direction, value):
    if direction == 'E':
        return Position(position.x + value, position.y, position.direction)
    if direction == 'S':
        return Position(position.x, position.y - value, position.direction)
    if direction == 'W':
        return Position(position.x - value, position.y, position.direction)
    if direction == 'N':
        return Position(position.x, position.y + value, position.direction)

    raise Exception('Bad direction {}'.format(direction))


def turn(position, direction, value):
    if direction == 'L':
        next_drection_index = (DIRECTIONS.index(position.direction) - value // 90 + 4) % 4
        next_drection = DIRECTIONS[next_drection_index]
        return Position(position.x, position.y, next_drection)
    if direction == 'R':
        next_drection_index = (DIRECTIONS.index(position.direction) + value // 90) % 4
        next_drection = DIRECTIONS[next_drection_index]
        return Position(position.x, position.y, next_drection)

    raise Exception('Bad turn {}'.format(direction))


def manhattan_distance(position):
    return abs(position.x) + abs(position.y)


def solution(data):
    instructions = [parse_instruction(x) for x in data.splitlines()]
    current_position = Position(0, 0, 'E')
    for instruction in instructions:
        current_position = move(current_position, instruction)
    return manhattan_distance(current_position)
