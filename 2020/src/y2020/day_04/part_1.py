VALID_PASSPORT_FIELDS = set(
    """byr
iyr
eyr
hgt
hcl
ecl
pid""".split()
)


def is_valid_passport(passport):
    return len(VALID_PASSPORT_FIELDS - passport) == 0


def solution(data):
    groups = [x for x in data.split('\n\n')]
    counter = 0
    for group in groups:
        passport = set()
        for element in [x for x in group.replace('\n', ' ').split(' ')]:
            passport.add(element.split(':')[0])
        if is_valid_passport(passport):
            counter += 1
    return counter
