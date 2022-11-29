from src.y2020.day_24.part_1 import HexTile, parse_instructions, solution, walk_from_center

EXAMPLE = """sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew"""


def test_hex_tile():
    hex_tile = HexTile()
    assert hex_tile.x == 0
    assert hex_tile.y == 0


def test_hex_tile_position_to_string():
    hex_tile = HexTile()
    assert hex_tile.position_to_string == '0:0'


def test_hex_tile_walk():
    hex_tile = HexTile()
    hex_tile.walk('nw')
    assert hex_tile.x == 1
    assert hex_tile.y == 1
    assert hex_tile.position_to_string == '1:1'


def test_walk_from_center():
    instructions = ['nw', 'w', 'sw', 'e', 'e']
    assert walk_from_center(instructions) == '0:0'


def test_parse_instructions():
    line = 'nwwswee'
    assert parse_instructions(line) == ['nw', 'w', 'sw', 'e', 'e']


def test_solution():
    assert solution(EXAMPLE) == 10
