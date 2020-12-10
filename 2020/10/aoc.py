#!/usr/bin/env python

import os
from collections import defaultdict
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
FILE = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
TEST2 = os.path.join(DIRPATH, 'testinput2.txt')
DEBUG = False

# Given a list of deltas, calculate how many different yet
# unique and equivalent ways each series of 1s can be
# represented using 1, 2, or 3.
def permute(deltas):
    length = len(deltas)
    if length == 1:
        # (1)
        count = 1
    elif length == 2:
        # (11, 2)
        count = 2
    elif length == 3:
        # (111, 12, 21)
        count = 4
    elif length == 4:
        # (1111, 112, 211, 121, 22, 31, 13)
        count = 7
    else:
        # we shouldn't hit this but if we do, we'll zero out
        count = 0
    return count

def part1(data, debug):
    diffs = defaultdict(lambda: 0)
    deltas = []
    for i in range(1, len(data)):
        diff = data[i] - data[i-1]
        deltas.append(diff)
        diffs[diff] += 1
    if debug:
        print(data)
    print("Part 1: {} * {} = {}".format(diffs[1], diffs[3], diffs[1] * diffs[3]))
    if debug:
        print(deltas)
    return deltas

def part2(deltas, debug):
    testset = []
    total = 1
    for step in deltas:
        if step == 1:
            testset.append(step)
        elif testset:
            temp = permute(testset)
            total *= temp
            if debug:
                print("permute {} = {}.".format(testset, temp))
                print("New total is {}".format(total))
            testset.clear()
    print("Part 2: {}".format(total))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(FILE) as file:
        data = file.read().splitlines()

    data = [int(i) for i in data]
    print("{} adapters".format(len(data)))
    data.append(0) # add the charging outlet
    data.sort()
    data.append(data[-1] + 3) # add the device

    time1 = time.perf_counter()
    deltas = part1(data, DEBUG)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

    print()

    time1 = time.perf_counter()
    part2(deltas, DEBUG)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
