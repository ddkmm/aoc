#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
DEBUG = True

def do_add(program, ptr):
    arg1_loc = program[ptr+1]
    arg2_loc = program[ptr+2]
    dest_loc = program[ptr+3]
    arg1 = program[arg1_loc]
    arg2 = program[arg2_loc] 
    program[dest_loc] = arg1 + arg2
    ptr += 4

    return program, ptr

def do_mul(program, ptr):
    arg1_loc = program[ptr+1]
    arg2_loc = program[ptr+2]
    dest_loc = program[ptr+3]
    arg1 = program[arg1_loc]
    arg2 = program[arg2_loc] 
    program[dest_loc] = arg1 * arg2
    ptr += 4

    return program, ptr

def run(program):
    ptr = 0
    op = 0
    while op != 99:
       op = program[ptr]
       if op == 1:
           program, ptr = do_add(program, ptr)
       if op == 2:
           program, ptr = do_mul(program, ptr)

    return program

def part1(data):
    data = re.split(r',', data)
    program = [int(i) for i in data] 
    # setup
    program[1] = 12
    program[2] = 2
    program = run(program)
    print("Part 1: {}".format(program[0]))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(DATA) as file:
        data = file.readline().strip()

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
