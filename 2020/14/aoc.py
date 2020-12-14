#!/usr/bin/env python

import os
import re
import time
from collections import defaultdict

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
TEST2 = os.path.join(DIRPATH, 'testinput2.txt')
DEBUG = True

def calculate_memory_size(data):
    max_memory = 0
    for inst in data:
        if re.search(r'mem', inst):
            cmd = re.split(r'=', inst)
            index = int(re.split(r'(\d+)', str(cmd[0]))[1])
            if index > max_memory:
                max_memory = index
    return max_memory * 2

def make_mask(mask):
    or_mask = ''
    and_mask = ''
    for b in mask:
        if b == 'X':
            or_mask += '0'
            and_mask += '1'
        else:
            or_mask += b
            and_mask += b
    return or_mask, and_mask

def apply_mask(value, or_mask, and_mask):
    if DEBUG:
        print("Before write: {}".format(value))
        print("or_mask {}".format(bin(int(or_mask, 2))))
        print("and_mask {}".format(bin(int(and_mask, 2))))
    value = value | (int(or_mask, 2))
    value = value & int(and_mask, 2)
    if DEBUG:
        print("After write: {}".format(value))
    return value

def replace_x(x_bits, prog):
    for a in x_bits:
        prog = re.split('X', prog, 1)
        prog = prog[0] + a + prog[1]
    print("{}".format(int(prog, 2)))
    return prog

def apply_mask2(value, mask, memory, index):
    value_string = "{0:036b}".format(index)
    if DEBUG:
        print("{}".format(value_string))
        print("{}".format(mask))
    value_list = list(value_string)
    x_bit_size = 0
    for i, b in enumerate(mask):
        if b == '1':
            value_list[i] = b
        if b == 'X':
            value_list[i] = b
            x_bit_size += 1 
    value_string = ""
    value_string = value_string.join(value_list)

    # Translate X bits into number of permutations
    float_list = []
    for n in range(0,2**x_bit_size):
        n_string = "{0:0b}".format(n)
        n_string = n_string.zfill(x_bit_size)
        float_list.append(list(n_string))
    print(float_list)
    for bitmask in float_list:
        new_index = int(replace_x(bitmask, value_string),2)
        memory[new_index] = value



    if DEBUG:
        print("{0:036b}".format(value))
    return value

def part1(data):
    # initialise
    max_memory = calculate_memory_size(data)
    memory = [0] * max_memory

    # Run the docking program
    for inst in data:
        if re.search(r'mask', inst):
            mask = re.split(r'=', inst)[1].strip()
            or_mask, and_mask = make_mask(mask)
        elif re.search(r'mem', inst):
            if DEBUG:
                print("{}".format(inst.strip()))
            cmd = re.split(r'=', inst)
            index = int(re.split(r'(\d+)', str(cmd[0]))[1])
            value = int(cmd[1].strip())
            # apply mask
            if DEBUG:
                print("Before write: {}".format(memory[index]))
            memory[index] = apply_mask(value, or_mask, and_mask)
    total = 0

    # Calculate puzzle answer
    for a in memory:
        total += a
    print("Part 1: {}".format(total))

def part2(data):
    # initialise
    memory = defaultdict(lambda: 0)

    # Run the docking program
    for inst in data:
        if re.search(r'mask', inst):
            mask = re.split(r'=', inst)[1].strip()
        elif re.search(r'mem', inst):
            if DEBUG:
                print("{}".format(inst.strip()))
            cmd = re.split(r'=', inst)
            index = int(re.split(r'(\d+)', str(cmd[0]))[1])
            value = int(cmd[1].strip())
            # apply mask
            if DEBUG:
                print("Before write: {}".format(memory[index]))
            apply_mask2(value, mask, memory, index)
    total = 0

    # Calculate puzzle answer
    for a in memory:
        total += memory[a]
    print("Part 2: {}".format(total))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(DATA) as file:
        data = file.read().splitlines()

    time1 = time.perf_counter()
#    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

    time1 = time.perf_counter()
    part2(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))


main()
