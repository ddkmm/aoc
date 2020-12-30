#!/usr/bin/env python

import os
import re
import time
import math

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
DEBUG = True

def calculate_fuel(mass):
    fuel = math.floor(mass/3) - 2
    if fuel < 0:
        return 0
    fuel = fuel + calculate_fuel(fuel)
    return fuel

def part1(data):
    total = 0
    for mass in data:
       total += math.floor(int(mass)/3) - 2
    print("Part 1: {}".format(total))

def part2(data):
    total = 0
    for mass in data:
       total = total + calculate_fuel(int(mass)) 
    print("Part 2: {}".format(total))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(DATA) as file:
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
