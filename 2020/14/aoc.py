#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
DEBUG = True

def part1(data):
    if DEBUG:
        print(data)
    # init
    memory = [0] * 36
    for inst in data:
        if re.search(r'mask', inst):
            mask = re.split(r'=', inst)[1].strip()
            for b in mask:
                if b == 'X':
                    or_mask =
            or_mask = 
            and_mask =
            print(mask)
        elif re.search(r'mem', inst):
            cmd = re.split(r'=', inst)
            index = int(re.split(r'(\d+)', str(cmd[0]))[1])
            value = int(cmd[1].strip())
            # apply mask
            memory[index] = value
    print(memory)
            


def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(TEST) as file:
        data = file.read().splitlines()

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
