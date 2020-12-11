#!/usr/bin/env python

import os
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
    if (row < 0 or col < 0):
        return WALL
    if (row >= len(seats[0]) or col >= len(seats)):
        return WALL
    return seats[col][row]

def advance(col, row, direct):
    if direct == 0:
        col += -1
        row += -1
    elif direct == 1:
        col += -1
    elif direct == 2:
        col += -1
        row += 1
    elif direct == 3:
        row += -1
    elif direct == 4:
        row += 1
    elif direct == 5:
        col += 1
        row += -1
    elif direct == 6:
        col += 1
    elif direct == 7:
        col += 1
        row += 1

    return col, row

def check_adjacent(seats, col, row, rule, first_visible):
    # Returns True if the rule criteria has been met
    # rule == 1
    #   no adjacent seats are occupied
    # rule == 2
    #   at least 4 adjacent seats are occupied
    count = 0

    # These are our 8 directions to look in.
    # If first_visible flag is not set,
    #   we only look at the adjacent seat
    # otherwise, we keep looking until we find
    #   a seat or run out of seats to look at
    row_range = [row - 1, row, row + 1]
    col_range = [col - 1, col, col + 1]

    if first_visible:
        threshold = 4
    else:
        threshold = 3
    direct = 0

    for col_val in col_range:
        for row_val in row_range:
            if ((row_val == row) and (col_val == col)):
                # Skip the seat we're checking around
                continue
            if not first_visible:
                # just check once in each direction
                seat_val = get_seat(seats, col_val, row_val)
                if seat_val is OCCUPIED:
                    if rule == 1:
                        return False
                    if rule == 2:
                        count += 1
                        if count > threshold:
                            return True
            else:
                # look in a direction until we hit something
                #   not floor
                temp_col = col_val
                temp_row = row_val
                seat_val = get_seat(seats, temp_col, temp_row)
                while seat_val == FLOOR:
                    temp_col, temp_row = advance(temp_col, temp_row, direct)
                    seat_val = get_seat(seats, temp_col, temp_row)
                if seat_val is OCCUPIED:
                    if rule == 1:
                        return False
                    if rule == 2:
                        count += 1
            direct += 1

    if rule == 1:
        return True
    if count > threshold:
        return True
    return False

def rule1(seats, new_seats, first_visible):
    # If a seat is empty (L) and
    # there are no occupied seats adjacent to it,
    # the seat becomes occupied.
    count = 0
    for col in range(0, len(seats)):
        for row in range(0, len(seats[0])):
            if (seats[col][row] is EMPTY and
                    check_adjacent(seats, col, row, 1, first_visible)):
                new_seats[col] = new_seats[col][:row] + OCCUPIED + new_seats[col][row+1:]
                count += 1
    return new_seats, count

def rule2(seats, new_seats, first_visible):
    # If a seat is occupied (#) and
    # four or more seats adjacent to it are also occupied,
    # the seat becomes empty.
    count = 0
    for col in range(0, len(seats)):
        for row in range(0, len(seats[0])):
            if (seats[col][row] is OCCUPIED and
                    check_adjacent(seats, col, row, 2, first_visible)):
                new_seats[col] = new_seats[col][:row] + EMPTY + new_seats[col][row+1:]
                count += -1
    return new_seats, count

def apply_rules(seats, occupied, first_visible):
    change = 999
    iteration = 0
    while change != 0:
        iteration += 1
        change = 0
        new_seats = seats.copy()
        new_seats, change1 = rule1(seats, new_seats, first_visible)
        change += change1
        seats, change2 = rule2(seats, new_seats, first_visible)
        change += change2

        occupied += change
        if DEBUG:
            print("Iteration {}: {} seats occupied".format(iteration, occupied))
    return seats, occupied

def part1(seats):
    print("Part 1")
    occupied = 0
    seats, occupied = apply_rules(seats, occupied, False)

def part2(seats):
    print("Part 2")
    occupied = 0
    seats, occupied = apply_rules(seats, occupied, True)

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(FILE) as file:
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
