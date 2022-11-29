from collections import namedtuple

Command = namedtuple('Command', ['operation', 'argument'])


class Program:
    def __init__(self, data):
        self.lines = Program._parse_program(data)
        self.next_line_index = 0
        self.accumulator = 0
        self.executed_lines = set()
        self.steps = []

    def reset(self):
        self.next_line_index = 0
        self.accumulator = 0
        self.executed_lines.clear()
        self.steps.clear()

    def run(self):
        while (
            self.next_line_index not in self.executed_lines
            and self.next_line_index > -1
            and self.next_line_index < len(self.lines)
        ):
            self.executed_lines.add(self.next_line_index)
            self.execute_line()
        return self.accumulator

    def execute_line(self):
        current_line = self.lines[self.next_line_index]

        if current_line.operation == 'acc':
            value = int(current_line.argument)
            self.accumulator += value
            self.next_line_index += 1
            self.steps.append('Add {} to acc [{}], go to {}'.format(value, self.accumulator, self.next_line_index))
            return

        if current_line.operation == 'jmp':
            value = int(current_line.argument)
            self.next_line_index += value
            self.steps.append('Jump of {}, acc [{}], go to {}'.format(value, self.accumulator, self.next_line_index))
            return

        if current_line.operation == 'nop':
            self.next_line_index += 1
            self.steps.append('Go to {}'.format(self.next_line_index))
            return

    @staticmethod
    def _parse_program(data):
        lines = []
        for line in data:
            operation, argument = line.split(' ')
            lines.append(Command(operation, argument))
        return lines


def solution(data):
    p = Program(data.split('\n'))
    return p.run()
