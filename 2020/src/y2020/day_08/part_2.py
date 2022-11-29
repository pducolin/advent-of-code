from collections import namedtuple

Command = namedtuple('Command', ['operation', 'argument'])


class ProgramFixer:
    def __init__(self, data):
        self.lines = [
            {'command': line, 'swapped': False, 'executed': False} for line in ProgramFixer._parse_program(data)
        ]
        self.next_line_index = 0
        self.accumulator = 0

    def reset(self):
        self.next_line_index = 0
        self.accumulator = 0
        for line in self.lines:
            line['swapped'] = False
            line['executed'] = False

    def fix(self):
        swapped_index = 0
        while swapped_index < len(self.lines):
            while self.lines[swapped_index]['command'].operation == 'acc':
                swapped_index += 1
            self.lines[swapped_index]['swapped'] = True
            if self.run():
                # successfully executed
                break
            self.reset()
            swapped_index += 1

    def run(self):
        while True:
            self.execute_line()
            if self.next_line_index == len(self.lines):
                return True
            if self.lines[self.next_line_index]['executed']:
                return False

    def execute_line(self):
        current_command = self.lines[self.next_line_index]['command']
        is_swapped = self.lines[self.next_line_index]['swapped']
        self.lines[self.next_line_index]['executed'] = True
        value = int(current_command.argument)

        operation = current_command.operation
        if is_swapped:
            operation = ProgramFixer.swap_operation(operation)

        if operation == 'acc':
            self._accumulate(value)

        if operation == 'jmp':
            self._jump(value)

        if operation == 'nop':
            self._nop()

    def _jump(self, value):
        self.next_line_index += value

    def _accumulate(self, value):
        self.accumulator += value
        self.next_line_index += 1

    def _nop(self):
        self.next_line_index += 1

    @staticmethod
    def _parse_program(data):
        lines = []
        for line in data:
            operation, argument = line.split(' ')
            lines.append(Command(operation, argument))
        return lines

    @staticmethod
    def swap_operation(operation):
        operation_swaps = {'jmp': 'nop', 'nop': 'jmp', 'acc': 'acc'}
        return operation_swaps[operation]


def solution(data):
    p = ProgramFixer(data.splitlines())
    p.fix()
    return p.accumulator
