import time
from datetime import datetime, timedelta

DEBUG = False


def dbg_print(s):
    if DEBUG:
        print(s)


class Cup:
    def __init__(self, label):
        self.label = label
        self.next_cup = None

    def set_next_cup(self, next_cup):
        self.next_cup = next_cup


ONE_MILLION = 1000000
TEN_MILLIONS = 10000000


class CupCircle:
    def __init__(self, sequence):
        cup = None
        self.cup_index_mapping = {}
        previous_cup = None
        for ch in sequence:
            label = int(ch)
            cup = Cup(label)
            self.cup_index_mapping[label] = cup
            if previous_cup is None:
                self.current_cup = cup
            else:
                previous_cup.next_cup = cup
            previous_cup = cup
        for label in range(len(sequence) + 1, ONE_MILLION + 1):
            cup = Cup(label)
            self.cup_index_mapping[label] = cup
            previous_cup.next_cup = cup
            previous_cup = cup
        self.cup_index_mapping[ONE_MILLION].next_cup = self.current_cup
        self.pick = []

    def move(self, iteration_index):
        self._pick_cups()
        self._place_cups(iteration_index)
        self._move_current()

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

    def _place_cups(self, iteration_index):
        destination_label = self._find_destination_label()
        destination_cup = self.cup_index_mapping[destination_label]
        after_destination = destination_cup.next_cup
        destination_cup.next_cup = self.pick[0]
        self.pick[-1].next_cup = after_destination

    def _find_destination_label(self):
        destination_label = self.current_cup.label - 1
        if destination_label < 1:
            destination_label = ONE_MILLION
        while destination_label in [cup.label for cup in self.pick]:
            destination_label -= 1
            if destination_label < 1:
                destination_label = ONE_MILLION
        return destination_label

    def _move_current(self):
        self.current_cup = self.current_cup.next_cup

    @property
    def printable_circle(self, iteration_index):
        ret = ''
        current = self.cup_index_mapping[1].next_cup
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

    @property
    def stars(self):
        star_cup = self.cup_index_mapping[1].next_cup
        return star_cup.label * star_cup.next_cup.label


def solution(data, rounds=TEN_MILLIONS):
    cup_circle = CupCircle(data)
    start = time.time()
    last_iter_check = 0
    for i in range(rounds):
        now = time.time()
        if now - start > 5:
            iterations = i - last_iter_check
            speed = iterations // 5
            remaininig_time = (rounds - iterations) // speed
            end_time = datetime.now() + timedelta(seconds=remaininig_time)
            dbg_print(f'Round {i + 1}/{rounds}, should be over at {end_time}')
            last_iter_check = i
            start = now
        cup_circle.move(i)
    return cup_circle.stars
