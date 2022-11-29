class Cup:
    def __init__(self, label):
        self.label = int(label)
        self.next_cup = None

    def set_next_cup(self, next_cup):
        self.next_cup = next_cup


class CupCircle:
    def __init__(self, sequence):
        previous_cup = None
        cup = None
        self.all_cups = set()
        for label in sequence:
            cup = Cup(label)
            self.all_cups.add(cup.label)
            if previous_cup is None:
                # first cup
                self.current_cup = cup
                self.max = cup.label
                self.min = cup.label
            else:
                previous_cup.set_next_cup(cup)
                if cup.label > self.max:
                    self.max = cup.label
                if cup.label < self.min:
                    self.min = cup.label
            previous_cup = cup
        previous_cup.set_next_cup(self.current_cup)
        self.pick = []

    def _pick_cups(self):
        self.pick = []
        # 5 cups: 1,2,3,4,5
        # circle links:
        # 1 -> 2 -> 3 -> 4 -> 5 -> 1
        # current cup: 1
        # pick 3 cups next to current cup
        for _ in range(3):
            # pick cup
            self.pick.append(self.current_cup.next_cup)
            # re-assign current cup next cup
            self.current_cup.next_cup = self.current_cup.next_cup.next_cup
        # circle:
        # 1 -> 5 -> 1
        # re-assign last picked cup next to None
        # pick: 2 -> 3 -> 4
        self.pick[-1].next_cup = None

    def move(self):
        self._pick_cups()
        self._place_cups()
        self._move_current()

    def _move_current(self):
        self.current_cup = self.current_cup.next_cup

    def _place_cups(self):
        destination_label = self._find_destination_label()
        current = self.current_cup
        while current.label != destination_label:
            current = current.next_cup
        self.pick[-1].next_cup = current.next_cup
        current.next_cup = self.pick[0]
        self.pick.clear()

    def _find_destination_label(self):
        destination_label = self.current_cup.label - 1
        hashed_pick = set()
        for cup in self.pick:
            hashed_pick.add(cup.label)
        while destination_label in hashed_pick or destination_label not in self.all_cups:
            destination_label -= 1
            if destination_label < self.min:
                destination_label = self.max
        return destination_label

    @property
    def printable_circle(self):
        ret = ''
        current = self.current_cup
        # find cup with label == 1
        while current.label != 1:
            current = current.next_cup
        # now current cup has label == 1
        # start from next from 1
        current = current.next_cup
        # print until 1 is reached again
        while current.label != 1:
            ret += str(current.label)
            current = current.next_cup
        return ret

    @property
    def printable_pick(self):
        ret = ''
        for i in range(len(self.pick)):
            ret += str(self.pick[i].label)
        return ret

    @property
    def printable_destination(self):
        return str(self.current_cup.label - 1)


def solution(data, rounds=100):
    cup_circle = CupCircle(data)
    for _ in range(rounds):
        cup_circle.move()
    return cup_circle.printable_circle
