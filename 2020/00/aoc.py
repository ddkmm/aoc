#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
FILEPATH = os.path.join(DIRPATH, 'input.txt')
TESTPATH = os.path.join(DIRPATH, 'testinput.txt')
DEBUG = False

def part1(data, debug):
    if debug:
        print(data)

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(FILEPATH) as file:
        data = file.readlines()

    time1 = time.perf_counter()
    part1(data, DEBUG)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
