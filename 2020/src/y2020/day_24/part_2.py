DIRECTION_TO_COORDINATES = {
    'e': {'x': -2, 'y': 0},
    'se': {'x': -1, 'y': -1},
    'sw': {'x': 1, 'y': -1},
    'w': {'x': 2, 'y': 0},
    'nw': {'x': 1, 'y': 1},
    'ne': {'x': -1, 'y': 1},
}

WHITE = False
BLACK = True


class HexTile:
    x = 0
    y = 0
    color = WHITE

    def __init__(self, x, y):
        self.x = x
        self.y = y

    @classmethod
    def from_instructions(cls, instructions):
        tile = cls(0, 0)
        for direction in instructions:
            tile._walk(direction)
        return tile

    def _walk(self, direction):
        self.x += DIRECTION_TO_COORDINATES[direction]['x']
        self.y += DIRECTION_TO_COORDINATES[direction]['y']

    def flip(self):
        self.color = not self.color

    @property
    def is_black(self):
        return self.color == BLACK

    @property
    def is_white(self):
        return self.color == WHITE

    @property
    def neighbours(self):
        neighbours = []
        for direction in DIRECTION_TO_COORDINATES:
            n = HexTile(self.x, self.y)
            n._walk(direction)
            neighbours.append(n)
        return neighbours

    @property
    def hash(self):
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


class TilesMap:
    tiles = {}
    current_day = 0

    def __init__(self, data):
        for line in data.splitlines():
            instructions = parse_instructions(line)
            hex_tile = HexTile.from_instructions(instructions)
            # add tile and neighbours to tiles hexgrid
            for tile in hex_tile.neighbours + [hex_tile]:
                if tile.hash not in self.tiles:
                    self.tiles[tile.hash] = tile
            # flip target tile
            self.tiles[hex_tile.hash].flip()

    @property
    def black_tiles_count(self):
        return len([tile for _, tile in self.tiles.items() if tile.is_black])

    def iterate_to_next_day(self):
        to_flip = []
        tiles_to_visit = list(self.tiles.values())
        while len(tiles_to_visit) > 0:
            tile = tiles_to_visit.pop(0)
            should_flip, to_add = self._visit_neighbours(tile)
            if should_flip:
                to_flip.append(tile)
                if tile.is_white:
                    for x in to_add:
                        self.tiles[x.hash] = x
                        tiles_to_visit.append(x)
        for tile in to_flip:
            tile.flip()
        self.current_day += 1

    def _visit_neighbours(self, tile):
        neighbours = []
        to_add = []
        for n in tile.neighbours:
            if n.hash in self.tiles:
                neighbours.append(self.tiles[n.hash])
            else:
                neighbours.append(n)
                to_add.append(n)
        if tile.is_white and len([x for x in neighbours if x.is_black]) == 2:
            # add neighbours of black tiles
            for n in to_add:
                self.tiles[n.hash] = n
            return True, to_add
        if tile.is_black:
            black_neighbours_count = len([x for x in neighbours if x.is_black])
            if black_neighbours_count == 0 or black_neighbours_count > 2:
                return True, []
        return False, []


def solution(data):
    tile_map = TilesMap(data)
    for _ in range(100):
        tile_map.iterate_to_next_day()
    return tile_map.black_tiles_count
