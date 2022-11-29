from datetime import date

from importlib import import_module

import click


def load_input(filename):
    with open(filename) as f:
        return f.read()


@click.command()
@click.option('--year', default=date.today().year, help='Year to execute.')
@click.option('--day', default=date.today().day, help='Day to execute.')
def execute_day(year,day):
    """Execute Advent of code solution for a specific day."""
    part_1 = import_module('src.y{:04d}.day_{:02d}.part_1'.format(year,day))
    part_2 = import_module('src.y{:04d}.day_{:02d}.part_2'.format(year,day))
    data = load_input('src/y{:04d}/day_{:02d}/input.txt'.format(year, day))
    print('🎉 Result for year {}, day {}, part 1 is {}'.format(year, day, part_1.solution(data)))
    print('🎉 Result for year {}, day {}, part 2 is {}'.format(year, day, part_2.solution(data)))


if __name__ == '__main__':
    execute_day()
