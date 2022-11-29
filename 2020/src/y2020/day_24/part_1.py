DIRECTION_TO_COORDINATES = {
    'e': {'x': -2, 'y': 0},
    'se': {'x': -1, 'y': -1},
    'sw': {'x': 1, 'y': -1},
    'w': {'x': 2, 'y': 0},
    'nw': {'x': 1, 'y': 1},
    'ne': {'x': -1, 'y': 1},
}


class HexTile:
    x = 0
    y = 0

    def walk(self, direction):
        self.x += DIRECTION_TO_COORDINATES[direction]['x']
        self.y += DIRECTION_TO_COORDINATES[direction]['y']

    @property
    def position_to_string(self):
        return ':'.join([str(coord) for coord in [self.x, self.y]])


def parse_instructions(line):
    instructions = []
    while len(line) > 0:
        if line[0] in {'e', 'w'}:
            instructions.append(line[0])
            line = line[1:]
        else:
            instructions.append(line[:2])
            line = line[2:]
    return instructions


def walk_from_center(instructions):
    hex_tile = HexTile()
    for direction in instructions:
        hex_tile.walk(direction)
    return hex_tile.position_to_string


def solution(data):
    flipped_tiles = set()
    for line in data.splitlines():
        instructions = parse_instructions(line)
        tile_id = walk_from_center(instructions)
        if tile_id in flipped_tiles:
            flipped_tiles.remove(tile_id)
        else:
            flipped_tiles.add(tile_id)
    return len(flipped_tiles)
