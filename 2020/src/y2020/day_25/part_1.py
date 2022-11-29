DEBUG = False


def dbg_print(s):
    if DEBUG:
        print(s)


DIVISOR = 20201227


def encrypt(value, subject_number):
    value *= subject_number
    return value % DIVISOR


def evaluate_loop_size(public_key, subject_number):
    value = 1
    loop_size = 0
    while value != public_key:
        value = encrypt(value, subject_number)
        loop_size += 1
    return loop_size


def evaluate_encryption_key(other_public_key, loop_size):
    value = 1
    for _ in range(loop_size):
        value = encrypt(value, other_public_key)
    return value


def solution(data):
    card_public_key, door_public_key = [int(key) for key in data.splitlines()]
    card_loop_size = evaluate_loop_size(card_public_key, 7)
    encryption_key = evaluate_encryption_key(door_public_key, card_loop_size)
    return encryption_key
