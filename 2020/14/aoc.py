#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
DEBUG = True

def part1(data):
    # init
    max_memory = 0
    for inst in data:
        if re.search(r'mem', inst):
            cmd = re.split(r'=', inst)
            index = int(re.split(r'(\d+)', str(cmd[0]))[1])
            if index > max_memory:
                max_memory = index
    memory = [0] * (max_memory+1)
    for inst in data:
        if re.search(r'mask', inst):
            mask = re.split(r'=', inst)[1].strip()
            or_mask = ''
            and_mask = ''
            for b in mask:
                if b == 'X':
                    or_mask += '0'
                    and_mask += '1'
                else:
                    or_mask += b
                    and_mask += b
        elif re.search(r'mem', inst):
            if DEBUG:
                print("{}".format(inst.strip()))
            cmd = re.split(r'=', inst)
            index = int(re.split(r'(\d+)', str(cmd[0]))[1])
            value = int(cmd[1].strip())
            # apply mask
            if DEBUG:
                print("Before write: {}".format(memory[index]))
                print("or_mask {}".format(bin(int(or_mask, 2))))
                print("and_mask {}".format(bin(int(and_mask, 2))))
            value = value | (int(or_mask, 2))
            value = value & int(and_mask, 2)
            memory[index] = value
            if DEBUG:
                print("After write: {}".format(memory[index]))
    total = 0
    for a in memory:
        total += a
    print(total)


def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(DATA) as file:
        data = file.read().splitlines()

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
