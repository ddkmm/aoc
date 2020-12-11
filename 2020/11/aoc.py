#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
FILE = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
DEBUG = True

row_width = 0
col_num = 0

def rule1(seats, occupied):
    print("{} occupied seats after rule 1".format(occupied))
    if DEBUG:
        print(seats)
    return seats, occupied

def rule2(seats, occupied):
    print("{} occupied seats after rule 2".format(occupied))
    if DEBUG:
        print(seats)
    return seats, occupied

def rule3(seats, occupied):
    print("{} occupied seats after rule 3".format(occupied))
    if DEBUG:
        print(seats)
    return seats, occupied

def part1(seats):
    occupied = 0
    seats,occupied = rule1(seats, occupied)
    seats,occupied = rule2(seats, occupied)
    seats,occupied = rule3(seats, occupied)


def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(TEST) as file:
        seats = file.read().splitlines()
    col_num = len(seats)
    row_width = len(seats[0])

    time1 = time.perf_counter()
    part1(seats)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
