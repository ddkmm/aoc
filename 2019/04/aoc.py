#!/usr/bin/env python

import os
import re
import time
from collections import defaultdict

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DEBUG = True

def part1(data):
    total = 0
    for num in data:
        digits = list(str(num)) 
        previous = 0
        res = True
        repeats = 0 
        for i, d in enumerate(digits):
            if i == 0:
                previous = d
            if d < previous:
                res = False
                break
            if i != 0 and d == previous:
                repeats += 1
            previous = d
        if res and repeats != 0:
            total += 1
    print("Part 1: {}".format(total))

def part2(data):
    total = 0
    for num in data:
        digits = list(str(num)) 
        res = True
        double = False
        repeats = set()
        for i, _ in enumerate(digits):
            if i != 0:
                if digits[i] < digits[i-1]:
                    res = False
                    break
                if digits[i] == digits[i-1]:
                    double = True
                    repeats.add(digits[i])

        if res and double:
            freq = defaultdict(lambda: list())
            for i, digit in enumerate(digits):
                freq[digit].append(i)
            val = False
            for d in repeats:
                if len(freq[d]) == 2:
                    val = True
            if val:
                total += 1
    print("Part 2: {}".format(total))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))
    data = range(145852,616942)

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))
    time1 = time.perf_counter()
    part2(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
