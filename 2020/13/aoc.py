#!/usr/bin/env python

import os
import re
import time
from collections import defaultdict

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
DEBUG = True

def get_next_bus(bus_id, timestamp):
    current_bus_time = 0
    while current_bus_time < timestamp:
        current_bus_time += bus_id
    wait_time = current_bus_time - timestamp
    print("wait {} * bus {} = {}".format(wait_time, bus_id, bus_id * wait_time))

def get_bus_sequence(timetable):
    e_timetable = enumerate(timetable)
    bus_time = 0
    res = False
    while not res:
        buses = defaultdict(lambda: 0)
        for offset, bus in e_timetable:
            if bus != 'x':
                buses[offset] = bus_time*int(bus) + offset
        dict_items = buses.items()
        sorted_items = sorted(dict_items)
        for offset, bus in sorted_items:
            if offset == 0:
                try:
                    previous = int(bus)
                except:
                    print(bus)
            else:
                delta = int(bus) - previous
                if delta != offset:
                    print(delta, offset)
                    break
                else:
                    print(bus)
                    continue

def part1(data):
    print("Part 1")
    timestamp = int(data[0])
    timetable = re.findall(r'[0-9]+', data[1])
    for bus in timetable:
        get_next_bus(int(bus), timestamp)

def part2(data):
    print("Part 2")
    timetable = re.split(r',', data[1])
    get_bus_sequence(timetable)

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(TEST) as file:
        data = file.read().splitlines()

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

    time1 = time.perf_counter()
    part2(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
