from collections import namedtuple

DIRECTIONS = ['E', 'S', 'W', 'N']

Position = namedtuple('Position', ['x', 'y'])
Instruction = namedtuple('Instruction', ['direction', 'value'])


def parse_instruction(instruction):
    direction = instruction[:1]
    value = int(instruction[1:])
    return Instruction(direction, value)


def move(waypoint_position, ship_position, instruction):
    if instruction.direction in DIRECTIONS:
        return move_direction(waypoint_position, instruction.direction, instruction.value), ship_position
    if instruction.direction == 'F':
        return waypoint_position, move_to_waypoint(ship_position, waypoint_position, instruction.value)

    return turn(waypoint_position, instruction.direction, instruction.value), ship_position


def move_direction(position, direction, value):
    if direction == 'E':
        return Position(position.x + value, position.y)
    if direction == 'S':
        return Position(position.x, position.y - value)
    if direction == 'W':
        return Position(position.x - value, position.y)
    if direction == 'N':
        return Position(position.x, position.y + value)

    raise Exception('Bad direction {}'.format(direction))


def turn(position, direction, value):
    rotation = value
    # rotation of x degrees counter-clockwise equals a rotation of 360º-x
    # taking value modulo 360 to support values > 360
    if direction == 'L':
        rotation = 360 - value % 360
    while rotation > 0:
        position = Position(position.y, -position.x)
        rotation -= 90
    return position


def move_to_waypoint(ship_position, waypoint_position, value):
    return Position(ship_position.x + waypoint_position.x * value, ship_position.y + waypoint_position.y * value)


def manhattan_distance(position):
    return abs(position.x) + abs(position.y)


def solution(data):
    instructions = [parse_instruction(x) for x in data.splitlines()]
    current_ship_position = Position(0, 0)
    current_waypoint_position = Position(10, 1)
    for instruction in instructions:
        current_waypoint_position, current_ship_position = move(
            current_waypoint_position, current_ship_position, instruction
        )
    return manhattan_distance(current_ship_position)
