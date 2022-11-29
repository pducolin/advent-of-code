from collections import namedtuple

BusSchedule = namedtuple('BusSchedule', ['bus_id', 'wait_time'])


def solution(data):
    data = data.splitlines()
    my_time = int(data[0])
    bus_ids = [int(x) for x in data[1].split(',') if x != 'x']
    best_bus = None
    for bus in bus_ids:
        wait_time = bus - my_time % bus
        if not best_bus or best_bus.wait_time > wait_time:
            best_bus = BusSchedule(bus, wait_time)
    return best_bus.bus_id * best_bus.wait_time
