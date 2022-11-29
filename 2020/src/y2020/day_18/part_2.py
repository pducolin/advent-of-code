from dataclasses import dataclass

DEBUG = False


@dataclass
class Operator:
    priority: int
    executor: callable


OPERATORS = {
    '+': Operator(2, lambda a, b: a + b),
    '*': Operator(1, lambda a, b: a * b),
}


def dbg_print(s):
    if DEBUG:
        print(s)


def parse_lexemes(expression):
    expression = expression.replace(' ', '')
    lexemes = []
    index = 0
    while index < len(expression):
        ch = expression[index]
        if ch.isdigit():
            number, index = parse_number(expression, index)
            lexemes.append(number)
            continue
        elif ch in OPERATORS or ch in {'(', ')'}:
            lexemes.append(ch)
        index += 1
    return lexemes


def evaluate(operation, a, b):
    return operation.executor(a, b)


def execute_operations(ops, args, criteria):
    dbg_print(f'Execute operations, ops: {ops}, args: {args}')
    while len(ops) > 0 and criteria():
        b = args.pop()
        a = args.pop()
        op = ops.pop()
        dbg_print(f'Execute operation: {a} {op} {b}')
        args.append(evaluate(OPERATORS[op], a, b))
    dbg_print(f'Args: {args}')


def parse_number(expression, index):
    number_str = ''
    while index < len(expression) and expression[index].isdigit():
        number_str += expression[index]
        index += 1

    dbg_print(f'Parsed number {number_str}')
    return int(number_str), index


def evaluate_expression(expression):
    dbg_print('expression: {}'.format(expression))
    lexemes = parse_lexemes(expression)
    ops = []
    args = []
    dbg_print('lexemes: {}'.format(''.join([str(x) for x in lexemes])))
    for lex in lexemes:
        if isinstance(lex, int):
            dbg_print(f'Appending {str(lex)} to args')
            args.append(lex)
            continue
        if lex == '(':
            dbg_print(f'Appending {lex} to ops')
            ops.append(lex)
            continue
        if lex == ')':
            execute_operations(ops, args, lambda: ops[-1] != '(')
            ops.pop()
            continue
        if lex in OPERATORS:
            priority = OPERATORS[lex].priority
            execute_operations(ops, args, lambda: ops[-1] != '(' and OPERATORS[ops[-1]].priority >= priority)
            dbg_print(f'Appending {lex} to ops')
            ops.append(lex)
    execute_operations(ops, args, lambda: True)
    dbg_print(f'Args: {args}')
    return args[-1] if len(args) == 1 else None


def solve_expression_list(expression_list):
    return sum([evaluate_expression(x) for x in expression_list])


def solution(data):
    return solve_expression_list(data.splitlines())
