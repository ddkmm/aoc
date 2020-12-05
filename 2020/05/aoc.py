#!/usr/bin/env python

import os
import math

dir_path = os.path.dirname(os.path.realpath(__file__))
FILEPATH = os.path.join(dir_path, 'input.txt')
TESTPATH = os.path.join(dir_path, 'testinput.txt')

min_row = 0
max_row = 0
min_col = 0
max_col = 0

def reset():
    global min_row, max_row, min_col, max_col
    min_row = 0
    max_row = 127 
    min_col = 0 
    max_col = 7 

def main():
    print("Day {}".format(os.path.split(dir_path)[1]))

    with open(FILEPATH) as f:
        passes = f.read()

    global min_row, max_row, min_col, max_col
    seats = [0] * 1024

    highest = 0
    reset()
    for bp in passes:
        if bp == 'F':
            max_row = math.floor((min_row + max_row) / 2)
        elif bp == 'B':
            min_row = math.ceil((min_row + max_row) / 2)
        elif bp == 'L':
            max_col = math.floor((min_col + max_col) / 2)
        elif bp == 'R':
            min_col = math.ceil((min_col + max_col) / 2)
        elif bp == '\n':
            assert(min_row == max_row)
            assert(min_col == max_col)
            current = min_row*8 + min_col
            # keep track of non-empty seats
            seats[current] = 1
            if current > highest:
                highest = current
            reset()

    # Final processing for the last bording pass
    assert(min_row == max_row)
    assert(min_col == max_col)
    current = min_row*8 + min_col
    seats[current] = 1
    if current > highest:
        highest = current

    # Part 1
    print("Highest seat ID is {}".format(highest))

    # Part 2
    a = 0
    while a < 1024:
        if seats[a] == 1:
            if seats[a+1] == 0 and seats[a+2] != 0:
                print("My seat is {}".format(a+1))
        a += 1

main()