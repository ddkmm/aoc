#!/usr/bin/env python

import os
from collections import defaultdict

DIRPATH = os.path.dirname(os.path.realpath(__file__))
FILE = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
TEST2 = os.path.join(DIRPATH, 'testinput2.txt')

# Given a list of deltas, calculate how many different yet
# unique and equivalent ways each series of 1s can be
# represented using 1, 2, or 3.
def permute(deltas):
    length = len(deltas)
    if length == 1:
        count = 1
    elif length == 2:
        count = 2
    elif length == 3:
        count = 4
    elif length == 4:
        count = 7
    else:
        count = 0

    return count


def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(FILE) as file:
        data = file.read().splitlines()

    data = [int(i) for i in data]
    print("{} adapters".format(len(data)))
    data.append(0) # add the charging outlet
    data.sort()
    data.append(data[-1] + 3) # add the device

    diffs = defaultdict(lambda: 0)
    deltas = []
    for i in range(1, len(data)):
        diff = data[i] - data[i-1]
        deltas.append(diff)
        diffs[diff] += 1
    print(data)
    print("Part 1: {} * {} = {}".format(diffs[1], diffs[3], diffs[1] * diffs[3]))
    print(deltas)
    testset = []
    total = 1
    for step in deltas:
        if step == 1:
            testset.append(step)
        elif testset:
            temp = permute(testset)
            total *= temp
            print("permute {} = {}.".format(testset, temp))
            print("New total is {}".format(total))
            testset.clear()

    print(total)

main()
