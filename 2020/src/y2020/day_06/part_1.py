def solution(data):
    groups = [x for x in data.split('\n\n')]
    counter = 0
    for group in groups:
        answers = set()
        for answer in ''.join(group.split('\n')):
            answers.add(answer)
        counter += len(answers)
    return counter
