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

BUS_OFFSET = 0
BUS_PERIOD = 1
BUS_MEMORY = 2
BUS_MATCH = 3

def get_bus_sequence(timetable):
    data = []
    final_bus = len(timetable) - 1
    e_timetable = enumerate(timetable)
    # make a list of bus data
    # each bus data contains:
    #   offset
    #   period
    #   memory
    #   match
    for offset, bus in e_timetable:
        if bus != 'x':
            bus_data = []
            # offset
            bus_data.append(offset)
            # period
            bus_data.append(int(bus))
            # memory
            bus_data.append(int(bus) - offset)
            # match
            bus_data.append(False)
            data.append(bus_data)

    # Start counting
    current_tick = 1
    finished = False
    while not finished:
        target = 0

        for bus_data in data:
            keep_going = False
            nth_bus = 0
            if bus_data[BUS_OFFSET] == 0:
                target = current_tick * bus_data[BUS_PERIOD]
                keep_going = True
            else:
                nth_bus += 1
                current_val = bus_data[BUS_MEMORY]
                while True:
                    if current_val > target:
                        bus_data[BUS_MATCH] = False
                        keep_going = False
                        break
                    elif current_val < target:
                        bus_data[BUS_MEMORY] = current_val
                        current_val += bus_data[BUS_PERIOD]
                    else:
                        assert current_val == target
                        bus_data[BUS_MEMORY] = current_val
                        bus_data[BUS_MATCH] = True
                        if nth_bus == final_bus:
                            finished = True
                        keep_going = True
                        break
            if keep_going == False:
                break
        current_tick += 1
        if finished:
            break

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
