from part_1 import solution as solution_1
from part_2 import solution as solution_2


def load_input(filename):
    with open(filename) as f:
        return f.read()


if __name__ == "__main__":
    data = load_input('input.txt')
    print('🎉 Result for part 1 is {}'.format(solution_1(data)))
    print('🎉 Result for part 2 is {}'.format(solution_2(data)))
