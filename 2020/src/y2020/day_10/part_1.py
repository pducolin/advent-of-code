def solution(data):
    adapters = [int(x) for x in data.splitlines()] + [0]
    adapters.sort()
    last_adapter = adapters[-1]
    adapters.append(last_adapter + 3)
    counters = {1: 0, 2: 0, 3: 0}
    for adapter_index in range(1, len(adapters)):
        current_adapter = adapters[adapter_index]
        previous_index = adapter_index - 1
        previous_adapter = adapters[previous_index]
        counters[current_adapter - previous_adapter] += 1
    return counters[1] * counters[3]
