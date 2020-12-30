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

def do_add(program, ptr, p1, p2, p3):
    arg1_loc = program[ptr+1]
    if p1 == 0:
        arg1 = program[arg1_loc]
    else:
        arg1 = arg1_loc
        
    arg2_loc = program[ptr+2]
    if p2 == 0:
        arg2 = program[arg2_loc] 
    else:
        arg2 = arg2_loc

    dest_loc = program[ptr+3]
    program[dest_loc] = arg1 + arg2
    ptr += 4 

    return program, ptr

def do_mul(program, ptr, p1, p2, p3):
    arg1_loc = program[ptr+1]
    if p1 == 0:
        arg1 = program[arg1_loc]
    else:
        arg1 = arg1_loc
        
    arg2_loc = program[ptr+2]
    if p2 == 0:
        arg2 = program[arg2_loc] 
    else:
        arg2 = arg2_loc

    dest_loc = program[ptr+3]
    program[dest_loc] = arg1 * arg2
    ptr += 4 

    return program, ptr

def get_in(program, ptr, input, p1, p2, p3):
    dest_loc = program[ptr+1]
    program[dest_loc] = input 
    ptr += 2 

    return program, ptr

def do_out(program, ptr, p1, p2, p3):
    arg1 = program[ptr+1]
    if p1 == 0:
        out = program[arg1]
    else:
        out = arg1
    print(out)
    ptr += 2 

    return program, ptr

def process_op(opcode):
    opcode = list(str(opcode))
    if len(opcode) == 1:
        op = "0{}".format(opcode[0]) 
        p1 = 0
        p2 = 0
        p3 = 0
    elif len(opcode) == 3:
        op = "{}{}".format(opcode[-2],opcode[-1])
        p1 = int(opcode[0])
        p2 = 0
        p3 = 0
    elif len(opcode) == 4:
        op = "{}{}".format(opcode[-2],opcode[-1])
        p1 = int(opcode[1])
        p2 = int(opcode[0])
        p3 = 0
    elif len(opcode) == 5:
        op = "{}{}".format(opcode[-2],opcode[-1])
        p1 = int(opcode[2])
        p2 = int(opcode[1])
        p3 = int(opcode[0]) 

    return op, p1, p2, p3

def run(program, input):
    ptr = 0
    while True:
        opcode = program[ptr]
        if opcode == 99:
            break
        op, p1, p2, p3 = process_op(opcode)
        if op == '01':
            program, ptr = do_add(program, ptr, p1, p2, p3)
        if op == '02':
            program, ptr = do_mul(program, ptr, p1, p2, p3)
        if op == '03':
            program, ptr = get_in(program, ptr, input, p1, p2, p3)
        if op == '04':
            program, ptr = do_out(program, ptr, p1, p2, p3)

    return program

def part1(data):
    data = re.split(r',', data)
    program = [int(i) for i in data] 
    print("Part 1: Input 1")
    program = run(program, 1)

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(DATA) as file:
        data = file.readline().strip()

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
