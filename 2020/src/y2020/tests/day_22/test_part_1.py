from src.y2020.day_22.part_1 import Player, solution

EXAMPLE = """Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10"""


def test_parse_player():
    data = """Player 1:
9
2
6
3
1"""
    identifier, deck = Player.parse_player(data.splitlines())
    assert identifier == '1'
    assert deck == [9, 2, 6, 3, 1]


def test_init_player():
    data = """Player 1:
9
2
6
3
1"""
    parsed_player_data = Player.parse_player(data.splitlines())
    player = Player(parsed_player_data[0], parsed_player_data[1])
    assert player.identifier == '1'
    assert player.deck == [9, 2, 6, 3, 1]


def test_play():
    player = Player('1', [1, 2, 3, 4])
    assert player.deal_card() == 1


def test_play_loser():
    player = Player('1', [])
    assert player.deal_card() is None


def test_win_cards():
    player = Player('1', [1, 2])
    player.win_cards([2, 4])
    assert player.deck == [1, 2, 4, 2]


def test_deck_value():
    player = Player('1', [1, 2, 3])
    assert player.deck_value == 10


def test_solution():
    assert solution(EXAMPLE) == 306
