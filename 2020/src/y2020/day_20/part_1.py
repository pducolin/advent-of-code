from functools import reduce

DEBUG = False


def dbg_print(s):
    if DEBUG:
        print(s)


class ImageTile:
    tile_id = -1
    tile_data = []

    def __init__(self, tile_id, tile_data):
        self.tile_id = tile_id
        self.tile_data = tile_data

    def flip_vertically(self, from_tile_id=-1):
        for i in range(len(self.tile_data)):
            self.tile_data[i] = self.tile_data[i][::-1]

    def flip_horizontally(self, from_tile_id=-1):
        self.tile_data.reverse()

    def rotate(self, from_id=-1):
        rotate_data = []
        for i in range(len(self.tile_data)):
            rotate_data.append(''.join([x[i] for x in self.tile_data])[::-1])
        self.tile_data = rotate_data

    @property
    def top_border(self):
        return self.tile_data[0]

    @property
    def bottom_border(self):
        return self.tile_data[-1]

    @property
    def left_border(self):
        return ''.join(x[0] for x in self.tile_data)

    @property
    def right_border(self):
        return ''.join(x[-1] for x in self.tile_data)


def parse_tile(data):
    tile_id = int(data[0].split(' ')[1][:-1])
    tile = data[1:]
    return ImageTile(tile_id, tile)


def parse_tiles_data(data):
    tiles = {}
    for tile_data in data.split('\n\n'):
        dbg_print(f'Parsing [{tile_data}]')
        tile = parse_tile(tile_data.splitlines())
        tiles[tile.tile_id] = tile
    return tiles


def find_top_neighbour(tiles, tile_id):
    tile = tiles[tile_id]
    neighbours = []
    for other_tile_id, other_tile in tiles.items():
        if other_tile_id == tile_id:
            continue
        if other_tile.bottom_border == tile.top_border:
            neighbours.append(other_tile_id)
        elif other_tile.bottom_border[::-1] == tile.top_border:
            neighbours.append(other_tile_id)
        elif other_tile.right_border[::-1] == tile.top_border:
            neighbours.append(other_tile_id)
        elif other_tile.right_border == tile.top_border:
            neighbours.append(other_tile_id)
        elif other_tile.top_border[::-1] == tile.top_border:
            neighbours.append(other_tile_id)
        elif other_tile.top_border == tile.top_border:
            neighbours.append(other_tile_id)
        elif other_tile.left_border == tile.top_border:
            neighbours.append(other_tile_id)
        elif other_tile.left_border[::-1] == tile.top_border:
            neighbours.append(other_tile_id)
    return neighbours


def solution(data):
    tiles = parse_tiles_data(data)
    corners = []
    for tile_id in tiles.keys():
        # find 4 neighbours
        neighbours = []
        for i in range(4):
            neighbours.extend(find_top_neighbour(tiles, tile_id))
            tiles[tile_id].rotate()
        dbg_print(f'Tile {tile_id} has neighbours: {neighbours}')
        if len(neighbours) == 2:
            corners.append(tile_id)
    # find tiles with 2 neighbours
    product = reduce((lambda x, y: x * y), corners)
    return product
