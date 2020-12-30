#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
DEBUG = True
NOUN = 1
VERB = 2

def do_add(program, ptr, step):
    arg1_loc = program[ptr+1]
    arg2_loc = program[ptr+2]
    dest_loc = program[ptr+3]
    arg1 = program[arg1_loc]
    arg2 = program[arg2_loc] 
    program[dest_loc] = arg1 + arg2
    ptr += step

    return program, ptr

def do_mul(program, ptr, step):
    arg1_loc = program[ptr+1]
    arg2_loc = program[ptr+2]
    dest_loc = program[ptr+3]
    arg1 = program[arg1_loc]
    arg2 = program[arg2_loc] 
    program[dest_loc] = arg1 * arg2
    ptr += step 

    return program, ptr

def run(program, step):
    ptr = 0
    op = 0
    while op != 99:
       op = program[ptr]
       if op == 1:
           program, ptr = do_add(program, ptr, step)
       if op == 2:
           program, ptr = do_mul(program, ptr, step)

    return program

def part1(data):
    data = re.split(r',', data)
    program = [int(i) for i in data] 
    # setup
    program[1] = 12
    program[2] = 2
    program = run(program, 4)
    print("Part 1: {}".format(program[0]))

def part2(data):
    data = re.split(r',', data)
    init_program = [int(i) for i in data] 
    # setup
    for noun in range(99):
        for verb in range(99):
            program = init_program.copy()
            program[NOUN] = noun
            program[VERB] = verb
            program = run(program, 4)
            output = program[0]
            if output == 19690720:
                print("Answer n: {}, v: {}".format(noun, verb))
                print("Part 2: {}".format(100*noun + verb))
                break

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(DATA) as file:
        data = file.readline().strip()

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))
    time1 = time.perf_counter()
    part2(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
