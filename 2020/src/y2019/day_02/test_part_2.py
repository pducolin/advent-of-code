EXAMPLE = '1,9,10,3,2,3,11,0,99,30,40,50'


def parse_program(data):
    return [int(x) for x in data.split(',')]


def execute_program(instructions):
    instruction_pointer = 0
    while True:
        operation = instructions[instruction_pointer]
        if operation == 99:
            break

        if operation != 1 and operation != 2:
            raise Exception('Invalid operation ({})'.format(operation))

        parameter_pointer_a = instructions[instruction_pointer + 1]
        parameter_pointer_b = instructions[instruction_pointer + 2]
        parameter_a = instructions[parameter_pointer_a]
        parameter_b = instructions[parameter_pointer_b]
        target_pointer = instructions[instruction_pointer + 3]
        if operation == 1:
            instructions[target_pointer] = parameter_a + parameter_b
        else:
            instructions[target_pointer] = parameter_a * parameter_b

        instruction_pointer += 4


def solution(data):
    original_instructions = parse_program(data)
    for noun in range(0, 99):
        for verb in range(0, 99):
            instructions = original_instructions.copy()
            instructions[1] = noun
            instructions[2] = verb
            try:
                execute_program(instructions)
                if instructions[0] == 19690720:
                    return noun * 100 + verb
            except Exception:
                continue
