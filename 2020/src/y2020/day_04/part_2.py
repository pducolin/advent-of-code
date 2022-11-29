import re


def is_valid_height(value):
    unit = value[-2:]
    val = value[:-2]
    return val.isdigit() and (
        (unit == 'cm' and int(val) >= 150 and int(val) <= 193) or (unit == 'in' and int(val) >= 59 and int(val) <= 76)
    )


def is_valid_hcl(value):
    pattern = r'#{1}[\da-f]{6}'
    return len(value) == 7 and re.search(pattern, value)


def is_valid_ecl(value):
    return value in {'gry', 'amb', 'hzl', 'blu', 'brn', 'grn', 'oth'}


def is_valid_pid(value):
    return len(value) == 9 and re.search(r'\d{9}', value)


def is_valid_byr(value):
    return is_valid_year(value, 1920, 2002)


def is_valid_iyr(value):
    return is_valid_year(value, 2010, 2020)


def is_valid_eyr(value):
    return is_valid_year(value, 2020, 2030)


def is_valid_year(value, min, max):
    return len(value) == 4 and value.isdigit() and int(value) >= min and int(value) <= max


RULES = {
    'byr': is_valid_byr,
    'iyr': is_valid_iyr,
    'eyr': is_valid_eyr,
    'hgt': is_valid_height,
    'hcl': is_valid_hcl,
    'ecl': is_valid_ecl,
    'pid': is_valid_pid,
}


def is_valid_fields(passport):
    return len(set(RULES) - set(passport)) == 0


def is_valid_passport(passport):
    if not is_valid_fields(passport):
        return False

    for key, val in passport.items():
        if key in RULES and not RULES[key](val):
            return False

    return True


def solution(data):
    groups = [x for x in data.split('\n\n')]
    counter = 0
    for group in groups:
        passport = {}
        for element in [x for x in group.replace('\n', ' ').split(' ')]:
            [field, value] = element.split(':')
            passport[field] = value
        if is_valid_passport(passport):
            counter += 1
    return counter
