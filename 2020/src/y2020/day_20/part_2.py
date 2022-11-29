import re

DEBUG = False

SEA_MONSTER = """                  # 
#    ##    ##    ###
 #  #  #  #  #  #   """


SEA_MONSTER_PATTERN = [
    r'(?=([#\.]{18}#))',
    r'(?=(#[#\.]{4}##[#\.]{4}##[#\.]{4}###))',
    r'(?=([#\.]#[#\.]{2}#[#\.]{2}#[#\.]{2}#[#\.]{2}#[#\.]{2}#[#\.]{3}))',
]


def dbg_print(s):
    if DEBUG:
        print(s)


DIRECTIONS = ['N', 'W', 'S', 'E']


def rotate_data(data):
    rotate_data = []
    for i in range(len(data[0])):
        rotate_data.append(''.join([x[i] for x in data])[::-1])
    return rotate_data


def flip_east_with_west_data(data):
    for i in range(len(data)):
        data[i] = data[i][::-1]


class ImageTile:
    tile_id = -1
    tile_data = []
    neighbours = {}

    def __init__(self, tile_id, tile_data):
        self.tile_id = tile_id
        self.tile_data = tile_data
        for direction in DIRECTIONS:
            self.neighbours[direction] = None

    def rotate(self):
        self.tile_data = rotate_data(self.tile_data)
        # rotate all neighbours
        old_north_tile = self.neighbours['N']
        self.neighbours['N'] = self.neighbours['W']
        self.neighbours['W'] = self.neighbours['S']
        self.neighbours['S'] = self.neighbours['E']
        self.neighbours['E'] = old_north_tile

    def flip_east_with_west(self):
        for i in range(len(self.tile_data)):
            self.tile_data[i] = self.tile_data[i][::-1]
        old_east = self.neighbours['E']
        self.neighbours['E'] = self.neighbours['W']
        self.neighbours['W'] = old_east

    def flip_north_with_south(self):
        self.tile_data.reverse()
        old_north = self.neighbours['N']
        self.neighbours['N'] = self.neighbours['S']
        self.neighbours['S'] = old_north

    def remove_borders(self):
        # remove north border
        self.tile_data = self.tile_data[1:]
        # remove south border
        self.tile_data = self.tile_data[:-1]
        # remove west
        self.tile_data = [x[1:] for x in self.tile_data]
        # remove east
        self.tile_data = [x[:-1] for x in self.tile_data]

    @property
    def neighbours_count(self):
        return len([x for _, x in self.neighbours.items() if x is not None])

    @property
    def is_corner(self):
        return self.neighbours_count == 2

    @property
    def border_north(self):
        return self.tile_data[0]

    @property
    def border_south(self):
        return self.tile_data[-1]

    @property
    def border_west(self):
        return ''.join(x[0] for x in self.tile_data)

    @property
    def border_east(self):
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
    for other_tile_id, other_tile in tiles.items():
        if other_tile_id == tile_id:
            continue
        if other_tile.border_south == tile.border_north:
            return other_tile_id
        elif other_tile.border_south[::-1] == tile.border_north:
            return other_tile_id
        elif other_tile.border_west[::-1] == tile.border_north:
            return other_tile_id
        elif other_tile.border_west == tile.border_north:
            return other_tile_id
        elif other_tile.border_north[::-1] == tile.border_north:
            return other_tile_id
        elif other_tile.border_north == tile.border_north:
            return other_tile_id
        elif other_tile.border_east == tile.border_north:
            return other_tile_id
        elif other_tile.border_east[::-1] == tile.border_north:
            return other_tile_id
    return None


def evaluate_neighbours(tiles):
    corners = []
    neighbours = {}
    for tile_id in tiles.keys():
        neighbours[tile_id] = {}
        # find 4 neighbours
        for direction in DIRECTIONS:
            n = find_top_neighbour(tiles, tile_id)
            neighbours[tile_id][direction] = n
            tiles[tile_id].rotate()
        dbg_print(f'Tile {tile_id} has neighbours: {neighbours[tile_id]}')
        if len([x for _, x in neighbours[tile_id].items() if x is not None]) == 2:
            corners.append(tile_id)
    return corners, neighbours


def build_image(corners, tiles, neighbours):
    # build image
    # [[corner_NW, TILE, TILE, ... ,CORNER_NE],
    # ...
    # [corner_SW, TILE, TILE, ... ,CORNER_SE]]
    all_image = []
    all_image.append([])
    current_row = 0
    current_tile_id = corners.pop()
    while len(tiles) > 0:
        tile = tiles.pop(current_tile_id)
        tile.neighbours = neighbours[tile.tile_id]
        # NW corner
        if current_row == 0 and len(all_image[current_row]) == 0:
            # rotate tile until N and W neighbours are None
            rotation_count = 0
            while tile.neighbours['N'] is not None or tile.neighbours['W'] is not None:
                rotation_count += 1
                if rotation_count > 3:
                    error_msg = f'Could not find rotation for tile {tile.tile_id}'
                    error_msg += f' while at [{current_row}:{len(all_image[current_row])}]'
                    raise Exception(error_msg)
                tile.rotate()
            all_image[current_row].append(tile)
            current_tile_id = tile.neighbours['E']
            continue
        # W border tile
        if len(all_image[current_row]) == 0:
            # rotate or flip until it has W border None
            previous_tile = all_image[current_row - 1][0]
            rotation_count = 0
            while tile.neighbours['N'] != previous_tile.tile_id:
                tile.rotate()
                rotation_count += 1
                if rotation_count > 3:
                    error_msg = f'Could not find rotation for tile {tile.tile_id}'
                    error_msg += f' while at [{current_row}:{len(all_image[current_row])}]'
                    raise Exception(error_msg)
            if tile.border_north != previous_tile.border_south:
                tile.flip_east_with_west()
                if tile.border_north != previous_tile.border_south:
                    dbg_print('Should not happen!')
                    raise Exception(f'Could not find orientation for tile {tile.tile_id}')
            all_image[current_row].append(tile)
            current_tile_id = tile.neighbours['E']
            continue
        # rotate or flip until it matches previous tile
        previous_tile = all_image[current_row][-1]
        rotation_count = 0
        while tile.neighbours['W'] != previous_tile.tile_id:
            rotation_count += 1
            if rotation_count > 3:
                error_msg = f'Could not find rotation for tile {tile.tile_id}'
                error_msg += f' while at [{current_row}:{len(all_image[current_row])}]'
                raise Exception(error_msg)
            tile.rotate()
        if tile.border_west != previous_tile.border_east:
            tile.flip_north_with_south()
            if tile.border_west != previous_tile.border_east:
                dbg_print('Should not happen!')
                raise Exception(f'Could not find orientation for tile {tile.tile_id}')
        all_image[current_row].append(tile)
        current_tile_id = tile.neighbours['E']
        if current_tile_id is None:
            # change row
            current_tile_id = all_image[current_row][0].neighbours['S']
            all_image.append([])
            current_row += 1
    return all_image[:-1]


def remove_borders(grid_of_tiles):
    for row in grid_of_tiles:
        for tile in row:
            tile.remove_borders()


def get_image(grid_of_tiles):
    image = []
    tile_height = len(grid_of_tiles[0][0].tile_data)
    tot_row_count = len(grid_of_tiles) * tile_height
    for _ in range(tot_row_count):
        image.append('')
    for grid_row_index in range(len(grid_of_tiles)):
        row = grid_of_tiles[grid_row_index]
        for tile in row:
            for tile_row_index in range(len(tile.tile_data)):
                tile_data_row = tile.tile_data[tile_row_index]
                image[grid_row_index * tile_height + tile_row_index] += tile_data_row
    return image


def parse_sea_monster_indexes():
    rows = []
    for row in SEA_MONSTER.splitlines():
        rows.append([])
        for i in range(len(row)):
            ch = row[i]
            if ch == '#':
                rows[-1].append(i)
    return rows


def find_and_mark_sea_monsters(image):
    counter = 0
    for row_index in range(len(image) - 2):
        # find 3 rows that match sea monster pattern
        # at same index
        row_0 = image[row_index]
        row_1 = image[row_index + 1]
        row_2 = image[row_index + 2]
        matches_0 = [match.start(1) for match in re.finditer(SEA_MONSTER_PATTERN[0], row_0)]
        if len(matches_0) == 0:
            continue
        matches_1 = [match.start(1) for match in re.finditer(SEA_MONSTER_PATTERN[1], row_1)]
        if len(matches_1) == 0:
            continue
        matches_2 = [match.start(1) for match in re.finditer(SEA_MONSTER_PATTERN[2], row_2)]
        if len(matches_2) == 0:
            continue
        intersection = set(matches_0) & set(matches_1) & set(matches_2)
        if len(intersection) == 0:
            continue
        # found match
        while len(intersection) > 0:
            counter += 1
            start_index_offset = intersection.pop()
            for i in range(3):
                # change row_0
                for sm_index in sea_monster_indexes[i]:
                    row = list(image[row_index + i])
                    row[start_index_offset + sm_index] = '0'
                    image[row_index + i] = ''.join(row)
    return image, counter


sea_monster_indexes = parse_sea_monster_indexes()


def count_hashes(image):
    counter = 0
    for row in image:
        counter += row.count('#')
    return counter


def solution(data):
    tiles = parse_tiles_data(data)
    corners, neighbours = evaluate_neighbours(tiles)
    dbg_print(f'Corners: {corners}')
    grid_of_tiles = build_image(corners, tiles, neighbours)
    remove_borders(grid_of_tiles)
    image = get_image(grid_of_tiles)
    for _ in range(2):
        for _ in range(4):
            image, counter = find_and_mark_sea_monsters(image)
            if counter > 0:
                break
            image = rotate_data(image)
        flip_east_with_west_data(image)
    return count_hashes(image)
