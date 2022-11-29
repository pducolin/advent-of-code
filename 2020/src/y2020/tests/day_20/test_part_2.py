from src.y2020.day_20.part_2 import flip_east_with_west_data, get_image, ImageTile, rotate_data, solution

EXAMPLE = """Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###..."""


def test_solution():
    assert solution(EXAMPLE) == 273


DATA = """123
456
789"""


def test_image_tile_data():
    tile = ImageTile(1, [x.strip() for x in DATA.splitlines()])
    assert tile.tile_data == ['123', '456', '789']
    assert tile.neighbours_count == 0
    assert tile.is_corner is False
    assert tile.border_north == '123'
    assert tile.border_east == '369'
    assert tile.border_south == '789'
    assert tile.border_west == '147'


def test_image_tile_data_with_neighbours():
    tile = ImageTile(1, [x.strip() for x in DATA.splitlines()])
    tile.neighbours = {'N': 3, 'E': None, 'S': None, 'W': 2}
    assert tile.tile_data == ['123', '456', '789']
    assert tile.neighbours == {'N': 3, 'E': None, 'S': None, 'W': 2}
    assert tile.neighbours_count == 2
    assert tile.is_corner is True


def test_remove_borders_ImageTile():
    tile = ImageTile(1, [x.strip() for x in DATA.splitlines()])
    tile.remove_borders()
    assert tile.tile_data == ['5']


def test_rotate():
    tile = ImageTile(1, [x.strip() for x in DATA.splitlines()])
    tile.rotate()
    assert tile.tile_data == ['741', '852', '963']


def test_rotate_with_neighbours():
    tile = ImageTile(1, [x.strip() for x in DATA.splitlines()])
    tile.neighbours = {'N': 0, 'E': 1, 'S': 2, 'W': 3}
    tile.rotate()
    assert tile.tile_data == ['741', '852', '963']
    assert tile.neighbours == {'N': 3, 'E': 0, 'S': 1, 'W': 2}


def test_flip_east_with_west():
    tile = ImageTile(1, [x.strip() for x in DATA.splitlines()])
    tile.neighbours = {'N': 0, 'E': 1, 'S': 2, 'W': 3}
    tile.flip_east_with_west()
    assert tile.tile_data == ['321', '654', '987']
    assert tile.neighbours == {'N': 0, 'E': 3, 'S': 2, 'W': 1}


def test_flip_north_with_south():
    tile = ImageTile(1, [x.strip() for x in DATA.splitlines()])
    tile.neighbours = {'N': 0, 'E': 1, 'S': 2, 'W': 3}
    tile.flip_north_with_south()
    assert tile.tile_data == ['789', '456', '123']
    assert tile.neighbours == {'N': 2, 'E': 1, 'S': 0, 'W': 3}


def test_rotate_data():
    data = """123
456"""
    data = data.splitlines()
    assert rotate_data(data) == ['41', '52', '63']


def test_flip_east_with_west_data():
    data = """123
456"""
    data = data.splitlines()
    assert data == ['123', '456']
    flip_east_with_west_data(data)
    assert data == ['321', '654']
