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
    res = False

    # Delay each bus's start time based on the index
    times = []
    for offset, bus in e_timetable:
        if bus == 'x':
            times.append(0)
        else:
            times.append(-1*offset)

    # Start counting
    while True:
        e_timetable = enumerate(timetable)
        for offset, bus in e_timetable:
            if bus != 'x':
                times[offset] += int(bus)
        check_val = times[0]
        if check_val == 1068788:
            print('Bing')
        res = True
        for a in times:
            if a != 0:
                if a != times:
                    res = False
                    break
        if res:
            break
    print(current_time)

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
