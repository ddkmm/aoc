#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
FILE = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
DEBUG = True
FLOOR = '.'
EMPTY = 'L'
OCCUPIED = '#'

def get_seat(seats, col, row):
    # Returns seat value
    if (row < 0) or (col < 0):
        return FLOOR
    elif (row >= len(seats[0]) or (col >= len(seats))):
        return FLOOR
    else:
        return seats[col][row]

def check_adjacent(seats, col, row, rule):
    # Returns True if the rule criteria has been met
    # rule == 1
    #   no adjacent seats are occupied
    # rule == 2
    #   at least 4 adjacent seats are occupied
    count = 0
    if rule == 1:
        res = True
    elif rule == 2:
        res = False
    # left and right
    row_range = [row - 1, row, row + 1]
    col_range = [col - 1, col, col + 1]
    for col_val in col_range:
        for row_val in row_range:
            if ((row_val == row) and (col_val == col)):
                # Skip the seat we're checking around
                continue
            seat_val = get_seat(seats, col_val, row_val)
            if seat_val is OCCUPIED:
                if rule == 1:
                    res = False
                    break
                elif rule == 2:
                    count += 1
                    if count > 3:
                        res = True
                        break

    return res

def rule1(seats, new_seats):
    # If a seat is empty (L) and
    # there are no occupied seats adjacent to it,
    # the seat becomes occupied.
    count = 0
    for col in range(0, len(seats)):
        for row in range(0, len(seats[0])):
            if (seats[col][row] is EMPTY and
                check_adjacent(seats, col, row, 1)):
                new_seats[col] = new_seats[col][:row] + OCCUPIED + new_seats[col][row+1:]
                count += 1
    return new_seats, count

def rule2(seats, new_seats):
    # If a seat is occupied (#) and
    # four or more seats adjacent to it are also occupied,
    # the seat becomes empty.
    count = 0
    for col in range(0, len(seats)):
        for row in range(0, len(seats[0])):
            if (seats[col][row] is OCCUPIED and
                check_adjacent(seats, col, row, 2)):
                new_seats[col] = new_seats[col][:row] + EMPTY + new_seats[col][row+1:]
                count += -1
    return new_seats, count

def apply_rules(seats, occupied):
    change = 999 
    iteration = 0
    while change != 0:
        iteration += 1
        change = 0
        new_seats = seats.copy()
        new_seats, change1 = rule1(seats, new_seats)
        change += change1
        seats, change2 = rule2(seats, new_seats)
        change += change2

        occupied += change
        if DEBUG:
            print("Iteration {}: {} seats occupied".format(iteration, occupied))
    return seats, occupied


def part1(seats):
    occupied = 0
    seats,occupied = apply_rules(seats, occupied)


def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(FILE) as file:
        seats = file.read().splitlines()

    time1 = time.perf_counter()
    part1(seats)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
