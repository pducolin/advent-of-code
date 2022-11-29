DEBUG = False


def dbg_print(s):
    if DEBUG:
        print(s)


class Player:
    def __init__(self, identifier, deck):
        self.identifier = identifier
        self.deck = deck

    def deal_card(self):
        if not self.has_cards:
            return None
        return self.deck.pop(0)

    @property
    def deck_cards_count(self):
        return len(self.deck)

    @property
    def has_cards(self):
        return self.deck_cards_count > 0

    def win_cards(self, cards):
        self.deck.extend(cards)

    @property
    def deck_value(self):
        value = 0
        for i in range(1, len(self.deck) + 1):
            value += i * self.deck[-i]
        return value

    @property
    def deck_printable(self):
        return ','.join([str(card) for card in self.deck])

    @staticmethod
    def parse_player(data):
        identifier = data[0].split(' ')[1][:-1]
        deck = []
        for card in data[1:]:
            deck.append(int(card))
        return (identifier, deck)


def play_recursive_match(player_1, player_2, current_game=1):
    dbg_print('')
    dbg_print(f'==== Game {current_game} ===')
    dbg_print('')
    game_history = [set(), set()]
    current_round = 1
    while player_1.has_cards and player_2.has_cards:
        dbg_print(f'-- Round {current_round} (Game {current_game}) --')
        history_entry_1 = player_1.deck_printable
        history_entry_2 = player_2.deck_printable
        dbg_print(f'Player 1\'s deck: {player_1.deck_printable}')
        dbg_print(f'Player 2\'s deck: {player_2.deck_printable}')
        if history_entry_1 in game_history[0] or history_entry_2 in game_history[1]:
            dbg_print(f'Infinite loop, player 1 wins game {current_game}')
            return player_1, True
        game_history[0].add(history_entry_1)
        game_history[1].add(history_entry_2)
        player_1_card = player_1.deal_card()
        player_2_card = player_2.deal_card()
        dbg_print(f'Player 1 plays: {player_1_card}')
        dbg_print(f'Player 2 plays: {player_2_card}')
        if player_1_card > player_1.deck_cards_count or player_2_card > player_2.deck_cards_count:
            # at least one player must not have enough cards left in their deck to recurse
            # the winner of the round is the player with the higher-value card
            round_winner = player_1 if player_1_card > player_2_card else player_2
            round_winner_cards = sorted([player_1_card, player_2_card], reverse=True)
        else:
            # both players have at least as many cards remaining in their deck as the value
            # of the card they just drew, the winner of the round is determined by playing a
            # new game of Recursive Combat
            sub_player_1 = Player(player_1.identifier, player_1.deck[:player_1_card])
            sub_player_2 = Player(player_2.identifier, player_2.deck[:player_2_card])
            dbg_print('Playing a sub-game to determine the winner...')
            sub_winner, has_infinite_loop = play_recursive_match(sub_player_1, sub_player_2, current_game + 1)
            dbg_print(f'...anyway, back to game {current_game}.')
            if has_infinite_loop:
                dbg_print(f'Infinite loop in sub game, player 1 wins game {current_game}')
                player_1.win_cards([player_1_card, player_2_card])
                continue
            if sub_winner.identifier == player_1.identifier:
                round_winner = player_1
                round_winner_cards = [player_1_card, player_2_card]
            else:
                round_winner = player_2
                round_winner_cards = [player_2_card, player_1_card]
        round_winner.win_cards(round_winner_cards)
        dbg_print(f"Player {round_winner.identifier} wins round {current_round} of game {current_game}")
        current_round += 1
    game_winner = player_1 if player_1.has_cards else player_2
    dbg_print(f'Player {game_winner.identifier} wins game {current_game}')
    return game_winner, False


def solution(data):
    parsed_players = [Player.parse_player(x.splitlines()) for x in data.split('\n\n')]
    player_1, player_2 = [Player(parsed_player[0], parsed_player[1]) for parsed_player in parsed_players]
    winner, _ = play_recursive_match(player_1, player_2)
    dbg_print(f'Winner is {winner.identifier} with deck {winner.deck_printable}')
    return winner.deck_value
