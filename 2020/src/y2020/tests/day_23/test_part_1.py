from src.y2020.day_23.part_1 import CupCircle, solution

EXAMPLE = """389125467"""


def test_cup_circle():
    cup_circle = CupCircle(EXAMPLE)
    assert cup_circle.current_cup.label == 3
    assert cup_circle.printable_circle == '25467389'
    assert len(cup_circle.pick) == 0
    assert cup_circle.printable_destination == '2'


def test_cup_circle_move_sub():
    cup_circle = CupCircle(EXAMPLE)
    cup_circle._pick_cups()
    assert cup_circle.printable_pick == '891'
    cup_circle._place_cups()
    cup_circle._move_current()
    assert cup_circle.current_cup.label == 2
    assert cup_circle.printable_circle == '54673289'
    assert len(cup_circle.pick) == 0
    assert cup_circle.printable_destination == '1'


def test_cup_circle_move():
    cup_circle = CupCircle(EXAMPLE)
    cup_circle.move()
    assert cup_circle.current_cup.label == 2
    assert cup_circle.printable_circle == '54673289'
    assert len(cup_circle.pick) == 0
    assert cup_circle.printable_destination == '1'


def test_short_solution():
    assert solution(EXAMPLE, 10) == '92658374'


def test_solution():
    assert solution(EXAMPLE) == '67384529'
