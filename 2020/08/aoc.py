#!/usr/bin/env python

import os
import re
from collections import defaultdict

dir_path = os.path.dirname(os.path.realpath(__file__))
FILEPATH = os.path.join(dir_path, 'input.txt')
TESTPATH = os.path.join(dir_path, 'testinput.txt')

def machine(op, val, int_ptr, acc):
    if re.search('nop', op):
        int_ptr += 1
    elif re.search('acc', op):
        acc += val
        int_ptr += 1
    elif re.search('jmp', op):
        int_ptr += val
    return op, val, int_ptr, acc

def run1(program):
    int_ptr = 0
    runlog = defaultdict(lambda: 0)
    acc = 0
    while runlog[int_ptr] != 1:
        runlog[int_ptr] += 1
        inst = program[int_ptr]
        op = re.split(' ', inst)[0]
        val = int(re.split(' ', inst)[1])
        op, val, int_ptr, acc = machine(op, val, int_ptr, acc)
    print ("Part 1: {}".format(acc))

def run_modded(program):
    int_ptr = 0
    runlog = defaultdict(lambda: 0)
    acc = 0
    stop = len(program)
    while runlog[int_ptr] != 1:
        if int_ptr == stop:
            print ("Part 2: {}".format(acc))
            break
        assert(runlog[int_ptr] != 1)
        runlog[int_ptr] += 1
        inst = program[int_ptr]
        op = re.split(' ', inst)[0]
        val = int(re.split(' ', inst)[1])
        op, val, int_ptr, acc = machine(op, val, int_ptr, acc)
    assert(False)

def run2(program):
    x = range(0, len(program))
    for i in x:
        cmd =  program[i]
        if re.search('nop', cmd):
            back_up = cmd
            cmd = str(cmd).replace('nop', 'jmp')
            program[i] = cmd
            try:
                run_modded(program)
            except AssertionError:
                program[i] = back_up
            continue
        elif re.search('jmp', cmd):
            back_up = cmd
            cmd = str(cmd).replace('jmp', 'nop')
            program[i] = cmd
            try:
                run_modded(program)
            except AssertionError:
                program[i] = back_up
            continue

def main():
    print("Day {}".format(os.path.split(dir_path)[1]))

    with open(FILEPATH) as f:
        program = f.read().splitlines()
    run1(program)
    run2(program)


main()
