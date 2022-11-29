class Player:
    def __init__(self, identifier, deck):
        self.identifier = identifier
        self.deck = deck

    def deal_card(self):
        if self.has_lost:
            return None
        return self.deck.pop(0)

    @property
    def has_lost(self):
        return len(self.deck) == 0

    def win_cards(self, cards):
        cards.sort(reverse=True)
        self.deck.extend(cards)

    @property
    def deck_value(self):
        value = 0
        for i in range(1, len(self.deck) + 1):
            value += i * self.deck[-i]
        return value

    @staticmethod
    def parse_player(data):
        identifier = data[0].split(' ')[1][:-1]
        deck = []
        for card in data[1:]:
            deck.append(int(card))
        return (identifier, deck)


def play_match(player_1, player_2):
    while not player_1.has_lost and not player_2.has_lost:
        player_1_card = player_1.deal_card()
        player_2_card = player_2.deal_card()
        if player_1_card > player_2_card:
            player_1.win_cards([player_1_card, player_2_card])
        else:
            player_2.win_cards([player_1_card, player_2_card])
    if player_1.has_lost:
        return player_2
    return player_1


def solution(data):
    parsed_players = [Player.parse_player(x.splitlines()) for x in data.split('\n\n')]
    player_1, player_2 = [Player(parsed_player[0], parsed_player[1]) for parsed_player in parsed_players]
    winner = play_match(player_1, player_2)
    return winner.deck_value
