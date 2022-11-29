from collections import namedtuple

BusSchedule = namedtuple('BusSchedule', ['bus_id', 'delta'])


def parse_bus_schedule(data):
    bus_ids = [int(x) if x != 'x' else x for x in data.split(',')]
    buses = []
    for i in range(len(bus_ids)):
        bus_id = bus_ids[i]
        if bus_id == 'x':
            continue
        buses.append(BusSchedule(bus_id, i % bus_id))
    return buses


def find_max_bus_id(buses):
    max_bus_id = 0
    for _, bus in buses.items():
        if bus.bus_id > max_bus_id:
            max_bus_id = bus.bus_id
    return max_bus_id


def wait_time(t, bus_id):
    return (bus_id - t % bus_id) % bus_id


def find_common_time(bus_a, bus_b, start_time):
    t = start_time
    delta = bus_a.bus_id
    while True:
        if t < bus_a.bus_id or t < bus_b.bus_id:
            t += delta
            continue
        if wait_time(t, bus_b.bus_id) != bus_b.delta:
            t += delta
            continue
        return t


def find_departure_time(data):
    buses = parse_bus_schedule(data)
    t = 0
    while True:
        # take first 2 buses
        bus_a = buses.pop(0)
        bus_b = buses.pop(0)
        # find first time where bus_b arrives bus_b.delta after bus_a
        common_time = find_common_time(bus_a, bus_b, t)
        if len(buses) == 0:
            return common_time
        buses.insert(0, BusSchedule(bus_a.bus_id * bus_b.bus_id, 0))
        t = common_time


def solution(data):
    data = data.splitlines()
    return find_departure_time(data[1])
