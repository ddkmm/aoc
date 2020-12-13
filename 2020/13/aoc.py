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
    final_bus = len(timetable) - 1

    # Start counting
    current_tick = 1
    finished = False
    while not finished:
        e_timetable = enumerate(timetable)
        times = defaultdict(lambda: 0)
        target = 0

        for offset, bus in e_timetable:
            keep_going = False
            if offset == 0:
                target = current_tick * int(bus)
                times[offset] = target
                keep_going = True
            elif bus != 'x':
                for i in range(1,current_tick):
                    current_val = i * int(bus)
                    current_target = target + offset
                    if current_val < current_target:
                        continue
                    elif i * int(bus) == current_target:
                        times[offset] = current_target
                        keep_going = True
                        if offset == final_bus:
                            finished == True
                    else:
                        break
            if keep_going == False:
                break
        current_tick += 1

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
