#!/usr/bin/env python

import os
import re
from collections import defaultdict
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
FILEPATH = os.path.join(DIRPATH, 'input.txt')
TESTPATH = os.path.join(DIRPATH, 'testinput.txt')

def machine(opcode, val, int_ptr, acc, runlog):
    if runlog[int_ptr] != 0:
        assert False
    runlog[int_ptr] += 1
    if re.search('nop', opcode):
        int_ptr += 1
    elif re.search('acc', opcode):
        acc += val
        int_ptr += 1
    elif re.search('jmp', opcode):
        int_ptr += val
    return opcode, val, int_ptr, acc

def run(program):
    int_ptr = 0
    runlog = defaultdict(lambda: 0)
    acc = 0
    while runlog[int_ptr] != 1:
        inst = program[int_ptr]
        opcode = re.split(' ', inst)[0]
        val = int(re.split(' ', inst)[1])
        opcode, val, int_ptr, acc = machine(opcode, val, int_ptr, acc, runlog)
    print("Part 1: {}".format(acc))

def run_modded(program):
    int_ptr = 0
    runlog = defaultdict(lambda: 0)
    acc = 0
    stop = len(program)
    while runlog[int_ptr] != 1:
        if int_ptr == stop:
            print("Part 2: {}".format(acc))
            break
        inst = program[int_ptr]
        opcode = re.split(' ', inst)[0]
        val = int(re.split(' ', inst)[1])
        opcode, val, int_ptr, acc = machine(opcode, val, int_ptr, acc, runlog)
    assert False

def run2(program):
    x = range(0, len(program))
    for i in x:
        cmd = program[i]
        if re.search('nop|jmp', cmd):
            back_up = cmd
            if re.search('nop', cmd):
                cmd = str(cmd).replace('nop', 'jmp')
            else:
                cmd = str(cmd).replace('jmp', 'nop')
            program[i] = cmd
            try:
                run_modded(program)
            except AssertionError:
                program[i] = back_up
            continue

def run3(program):
    int_ptr = 0
    stack = []
    acc_stack = []
    runlog = defaultdict(lambda: 0)
    acc = 0
    rewind = False
    while int_ptr < len(program):
        inst = program[int_ptr]
        opcode = re.split(' ', inst)[0]
        val = int(re.split(' ', inst)[1])

        if re.search('nop|jmp', opcode):
            if not stack:
                if rewind:
                    rewind = False
                else:
                    if opcode in 'nop':
                        opcode = str(opcode).replace('nop', 'jmp')
                    else:
                        opcode = str(opcode).replace('jmp', 'nop')
                    # save system state
                    stack.append(int_ptr)
                    acc_stack.append(acc)
                    runlog_back = runlog.copy()
        try:
            opcode, val, int_ptr, acc = machine(opcode, val, int_ptr, acc, runlog)
        except AssertionError:
            rewind = True
            # restore system state
            int_ptr = stack.pop()
            acc = acc_stack.pop()
            runlog = runlog_back.copy()

    print("Part 2: {}".format(acc))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(FILEPATH) as f:
        program = f.read().splitlines()

    # Part 1
    run(program)

    # Part 2
    time1 = time.perf_counter()
    run2(program)
    time2 = time.perf_counter()
    run3(program)
    time3 = time.perf_counter()
    print("Brute force: {} seconds".format(time2-time1))
    print("Stack unwind: {} seconds".format(time3-time2))


main()
