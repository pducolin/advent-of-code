EXAMPLE = '1,9,10,3,2,3,11,0,99,30,40,50'


def parse_program(data):
    return [int(x) for x in data.split(',')]


def execute_program(instructions):
    offset = 0
    while True:
        current_instruction = instructions[offset]
        if current_instruction == 99:
            break

        if current_instruction != 1 and current_instruction != 2:
            raise Exception('Invalid operation {}'.format(current_instruction))

        value_1 = instructions[instructions[offset + 1]]
        value_2 = instructions[instructions[offset + 2]]
        target_index = instructions[offset + 3]
        if current_instruction == 1:
            instructions[target_index] = value_1 + value_2
        else:
            instructions[target_index] = value_1 * value_2

        offset += 4


def solution(data):
    instructions = parse_program(data)
    instructions[1] = 12
    instructions[2] = 2
    execute_program(instructions)
    return instructions[0]


def test_execute_program():
    instructions = parse_program(EXAMPLE)
    execute_program(instructions)
    assert ','.join([str(x) for x in instructions]) == '3500,9,10,70,2,3,11,0,99,30,40,50'
