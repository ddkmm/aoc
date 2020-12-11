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
WALL = 'X'
ME = 'O'

def get_seat(seats, col, row):
    # Returns seat value
    if (row < 0) or (col < 0):
        return WALL
    elif (row >= len(seats[0]) or (col >= len(seats))):
        return WALL
    else:
        return seats[col][row]

def check_adjacent(seats, col, row, rule, isVisible):
    # Returns True if the rule criteria has been met
    # rule == 1
    #   no adjacent seats are occupied
    # rule == 2
    #   at least 4 adjacent seats are occupied
    count = 0

    # These are our 8 directions to look in.
    # If isVisible flag is not set,
    #   we only look at the adjacent seat
    # otherwise, we keep looking until we find
    #   a seat or run out of seats to look at
    row_range = [row - 1, row, row + 1]
    col_range = [col - 1, col, col + 1]

    if isVisible:
        threshold = 4
    else:
        threshold = 3

    for col_val in col_range:
        for row_val in row_range:
            origin = False
            if ((row_val == row) and (col_val == col)):
                # Skip the seat we're checking around
                continue
            if not isVisible:
                # just check once in each direction
                seat_val = get_seat(seats, col_val, row_val)
                if seat_val is OCCUPIED:
                    if rule == 1:
                        return False
                    elif rule == 2:
                        count += 1
                        if count > threshold:
                            return True
            else:
                # look in a direction until we hit something
                #   not floor 
                temp_col = col_val
                temp_row = row_val
                seat_val = get_seat(seats, temp_col, temp_row)
                while (seat_val != WALL and not origin):
                    if seat_val is OCCUPIED:
                        if rule == 1:
                            return False
                        elif rule == 2:
                            count += 1
                            if count > threshold:
                                return True
                    if col_val == 0 and row_val == 0:
                        origin = True
                        continue
                    temp_col += col_val
                    temp_row += row_val
                    seat_val = get_seat(seats, temp_col, temp_row)
    if rule == 1:
        return True
    else:
        return False

def rule1(seats, new_seats, isVisible):
    # If a seat is empty (L) and
    # there are no occupied seats adjacent to it,
    # the seat becomes occupied.
    count = 0
    for col in range(0, len(seats)):
        for row in range(0, len(seats[0])):
            if (seats[col][row] is EMPTY and
                check_adjacent(seats, col, row, 1, isVisible)):
                new_seats[col] = new_seats[col][:row] + OCCUPIED + new_seats[col][row+1:]
                count += 1
    return new_seats, count

def rule2(seats, new_seats, isVisible):
    # If a seat is occupied (#) and
    # four or more seats adjacent to it are also occupied,
    # the seat becomes empty.
    count = 0
    for col in range(0, len(seats)):
        for row in range(0, len(seats[0])):
            if (seats[col][row] is OCCUPIED and
                check_adjacent(seats, col, row, 2, isVisible)):
                new_seats[col] = new_seats[col][:row] + EMPTY + new_seats[col][row+1:]
                count += -1
    return new_seats, count

def apply_rules(seats, occupied, isVisible):
    change = 999 
    iteration = 0
    while change != 0:
        iteration += 1
        change = 0
        new_seats = seats.copy()
        new_seats, change1 = rule1(seats, new_seats, isVisible)
        change += change1
        seats, change2 = rule2(seats, new_seats, isVisible)
        change += change2

        occupied += change
        if DEBUG:
            print("Iteration {}: {} seats occupied".format(iteration, occupied))
    return seats, occupied

def part1(seats):
    print("Part 1")
    occupied = 0
    seats,occupied = apply_rules(seats, occupied, False)

def part2(seats):
    print("Part 2")
    occupied = 0
    seats,occupied = apply_rules(seats, occupied, True)

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(TEST) as file:
        seats = file.read().splitlines()

    time1 = time.perf_counter()
    part1(seats)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

    time1 = time.perf_counter()
    part2(seats)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
