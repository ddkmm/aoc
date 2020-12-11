#!/usr/bin/env python

import os
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
FILE = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
DEBUG = False
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

# These are our 8 directions to look
DIRECTIONS = {}
DIRECTIONS[0] = [-1, -1]
DIRECTIONS[1] = [-1, 0]
DIRECTIONS[2] = [-1, 1]
DIRECTIONS[3] = [0, -1]
DIRECTIONS[4] = [0, 1]
DIRECTIONS[5] = [1, -1]
DIRECTIONS[6] = [1, 0]
DIRECTIONS[7] = [1, 1]

def check_rule_1(seats, col, row, first_visible):
    # Returns True if no adjacent seats are occupied

    row_range = [row - 1, row, row + 1]
    col_range = [col - 1, col, col + 1]
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
                    return False
            else:
                # look in a direction until we hit something
                #   not floor
                temp_col = col_val
                temp_row = row_val
                seat_val = get_seat(seats, temp_col, temp_row)
                while seat_val == FLOOR:
                    temp_col += DIRECTIONS[direct][0]
                    temp_row += DIRECTIONS[direct][1]
                    seat_val = get_seat(seats, temp_col, temp_row)
                if seat_val is OCCUPIED:
                    return False
            direct += 1

    return True

def check_rule_2(seats, col, row, first_visible):
    # Returns True if at least:
    #   4 adjacent seats are occupied
    #   or
    #   5 visibly adjacent seats are occupied
    row_range = [row - 1, row, row + 1]
    col_range = [col - 1, col, col + 1]
    direct = 0

    count = 0

    if first_visible:
        threshold = 4
    else:
        threshold = 3

    for col_val in col_range:
        for row_val in row_range:
            if ((row_val == row) and (col_val == col)):
                # Skip the seat we're checking around
                continue
            if not first_visible:
                # just check once in each direction
                seat_val = get_seat(seats, col_val, row_val)
                if seat_val is OCCUPIED:
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
                    temp_col += DIRECTIONS[direct][0]
                    temp_row += DIRECTIONS[direct][1]
                    seat_val = get_seat(seats, temp_col, temp_row)
                if seat_val is OCCUPIED:
                    count += 1
                    if count > threshold:
                        return True
            direct += 1

    return False

def rule1(seats, new_seats, first_visible):
    count = 0
    for col in range(0, len(seats)):
        for row in range(0, len(seats[0])):
            if (seats[col][row] is EMPTY and
                    check_rule_1(seats, col, row, first_visible)):
                new_seats[col] = new_seats[col][:row] + OCCUPIED + new_seats[col][row+1:]
                count += 1
    return new_seats, count

def rule2(seats, new_seats, first_visible):
    count = 0
    for col in range(0, len(seats)):
        for row in range(0, len(seats[0])):
            if (seats[col][row] is OCCUPIED and
                    check_rule_2(seats, col, row, first_visible)):
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
    occupied = 0
    seats, occupied = apply_rules(seats, occupied, False)
    print("Part 1: {}".format(occupied))

def part2(seats):
    occupied = 0
    seats, occupied = apply_rules(seats, occupied, True)
    print("Part 2: {}".format(occupied))

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
