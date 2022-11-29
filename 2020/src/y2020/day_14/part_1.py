import re


def parse_mask(instruction):
    # mask = value
    pattern = r'mask = ([X01]+)'
    mask_string = re.search(pattern, instruction).group(1)
    mask = {}
    for i in range(len(mask_string)):
        c = mask_string[i]
        if c != 'X':
            mask[i] = c
    return mask


def parse_assignment(instruction):
    # mem[line] = value
    pattern = r'mem\[(\d+)\] = (\d+)'
    m = re.search(pattern, instruction)
    line = int(m.group(1))
    value = int(m.group(2))
    return line, value


def to_binary_string(value):
    return "{0:036b}".format(value)


def from_binary_string(bin_string):
    return int(bin_string, 2)


def replace_at(s, index, replacement):
    after_index = index + 1
    return s[:index] + replacement + s[after_index:]


def apply_mask(value, mask):
    bin_string_value = to_binary_string(value)
    for mask_item in mask.keys():
        bin_string_value = replace_at(bin_string_value, mask_item, mask[mask_item])
    return from_binary_string(bin_string_value)


def apply_instructions(instructions):
    mask = {}
    memory = {}
    for instruction in instructions:
        if 'mask' in instruction:
            mask = parse_mask(instruction)
            continue
        line, value = parse_assignment(instruction)
        value = apply_mask(value, mask)
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
