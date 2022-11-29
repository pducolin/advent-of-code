import re


def parse_mask(instruction):
    # mask = value
    # example
    # mask = 000000000000000000000000000000X1001X
    # pattern: mask<space>=<space>(Group containing 1 or more '0', '1' or 'X')
    pattern = r'mask = ([X01]+)'
    # group(0): 'mask = 000000000000000000000000000000X1001X'
    # group(1): '000000000000000000000000000000X1001X'
    return re.search(pattern, instruction).group(1)


def parse_assignment(instruction):
    # mem[line] = value
    # example
    # mem[24] = 100
    # pattern: mem[(Group containing 1 or more digist)]<space>=<space>(Group containing 1 or more digits)
    pattern = r'mem\[(\d+)\] = (\d+)'
    m = re.search(pattern, instruction)
    # group(0): mem[24] = 100
    # group(1): 24
    # group(2): 100
    line = int(m.group(1))
    value = int(m.group(2))
    return line, value


def to_binary_string(value):
    return "{0:036b}".format(value)


def from_binary_string(bin_string):
    if 'X' not in bin_string:
        return [int(bin_string, 2)]
    to_replace = [bin_string]
    lines = []
    while len(to_replace) > 0:
        replacing = to_replace.pop(0)
        x_index = replacing.find('X')
        replaced_with_0 = replace_at(replacing, x_index, '0')
        replaced_with_1 = replace_at(replacing, x_index, '1')
        if 'X' in replaced_with_0:
            to_replace.append(replaced_with_0)
            to_replace.append(replaced_with_1)
        else:
            lines.append(replaced_with_0)
            lines.append(replaced_with_1)
    return [int(x, 2) for x in lines]


def replace_at(s, index, replacement):
    after_index = index + 1
    return s[:index] + replacement + s[after_index:]


def apply_mask(value, mask):
    bin_string = to_binary_string(value)
    for i in range(len(mask)):
        c = mask[i]
        if c == '0':
            continue
        bin_string = replace_at(bin_string, i, c)
    return from_binary_string(bin_string)


def apply_instructions(instructions):
    mask = {}
    memory = {}
    for instruction in instructions:
        if 'mask' in instruction:
            mask = parse_mask(instruction)
            continue
        line, value = parse_assignment(instruction)
        lines = apply_mask(line, mask)
        for line in lines:
            memory[line] = value
    return memory


def sum_memory(memory):
    counter = 0
    for _, item in memory.items():
        counter += item
    return counter


def solution(data):
    instructions = data.splitlines()
    memory = apply_instructions(instructions)
    return sum_memory(memory)
