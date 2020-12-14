#!/usr/bin/env python

import os
import re
import time
from collections import defaultdict

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
TEST2 = os.path.join(DIRPATH, 'testinput2.txt')
DEBUG = False

def make_mask(mask):
    or_mask = ''
    and_mask = ''
    for bit in mask:
        if bit == 'X':
            or_mask += '0'
            and_mask += '1'
        else:
            or_mask += bit
            and_mask += bit
    return or_mask, and_mask

def apply_mask(value, or_mask, and_mask):
    if DEBUG:
        print("Before write: {}".format(value))
        print("or_mask {}".format(bin(int(or_mask, 2))))
        print("and_mask {}".format(bin(int(and_mask, 2))))
    value = value | int(or_mask, 2)
    value = value & int(and_mask, 2)
    if DEBUG:
        print("After write: {}".format(value))
    return value

def replace_x(x_bits, prog):
    for bit in x_bits:
        prog = re.split('X', prog, 1)
        prog = prog[0] + bit + prog[1]
    return prog

def apply_mask2(value, mask, memory, addr):
    addr_string = "{0:036b}".format(addr)
    if DEBUG:
        print("{}".format(addr_string))
        print("{}".format(mask))
    addr_list = list(addr_string)
    x_bit_size = 0

    # process mask
    for i, bit in enumerate(mask):
        if bit == '1':
            addr_list[i] = bit
        elif bit == 'X':
            addr_list[i] = bit
            x_bit_size += 1
    addr_string = ""
    addr_string = addr_string.join(addr_list)

    # Translate X bits into number of permutations
    float_list = []
    for bit in range(0,2**x_bit_size):
        bit_string = "{0:0b}".format(bit)
        bit_string = bit_string.zfill(x_bit_size)
        float_list.append(list(bit_string))
    # apply permutations of X to original address
    for bitmask in float_list:
        new_addr = int(replace_x(bitmask, addr_string),2)
        memory[new_addr] = value
    return value

def part1(data):
    memory = defaultdict(lambda: 0)

    # Run the docking program
    for inst in data:
        if re.search(r'mask', inst):
            mask = re.split(r'=', inst)[1].strip()
            or_mask, and_mask = make_mask(mask)
        else:
            if DEBUG:
                print("{}".format(inst.strip()))
            cmd = re.split(r'=', inst)
            addr = int(re.split(r'(\d+)', str(cmd[0]))[1])
            value = int(cmd[1].strip())
            # apply mask
            if DEBUG:
                print("Before write: {}".format(memory[addr]))
            memory[addr] = apply_mask(value, or_mask, and_mask)
    total = 0

    # Calculate puzzle answer
    for addr in memory:
        total += memory[addr]
    print("Part 1: {}".format(total))

def part2(data):
    memory = defaultdict(lambda: 0)

    # Run the docking program
    for inst in data:
        if re.search(r'mask', inst):
            mask = re.split(r'=', inst)[1].strip()
        else:
            if DEBUG:
                print("{}".format(inst.strip()))
            cmd = re.split(r'=', inst)
            addr = int(re.split(r'(\d+)', str(cmd[0]))[1])
            value = int(cmd[1].strip())
            # apply mask
            if DEBUG:
                print("Before write: {}".format(memory[addr]))
            apply_mask2(value, mask, memory, addr)
    total = 0

    # Calculate puzzle answer
    for addr in memory:
        total += memory[addr]
    print("Part 2: {}".format(total))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(DATA) as file:
        data = file.read().splitlines()

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

    time1 = time.perf_counter()
    part2(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))


main()
