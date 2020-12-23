#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = "123487596"
TEST = "389125467"
DEBUG = True

def increment_pointer(cups, pointer):
    pointer += 1
    pointer = pointer % len(cups)
    return pointer

def get_largest(cups):
    largest = 0
    for c in cups:
        if c > largest:
            largest = c
    return largest

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    cups = [int(i) for i in list(DATA)]
    original = cups.copy()
    total = len(cups)

    current_pointer = 0
    for i in range(100):
        current_cup = cups[current_pointer]
        print("-- move {} --".format(i+1))
        print("{}".format(cups))
        print("current cup {}".format(current_cup))
        removed = []
        remove_pointer = increment_pointer(cups, current_pointer)
        for _ in range(3):
            try:
                removed.append(cups.pop(remove_pointer))
            except IndexError:
                remove_pointer = increment_pointer(cups, current_pointer)
                removed.append(cups.pop(remove_pointer))
        print("pick up: {}".format(removed))

        # get destination cup
        dest_cup = current_cup - 1
        dest_pointer = -1
        while dest_cup > 0:
            try:
                dest_pointer = cups.index(dest_cup)
                break
            except ValueError:
                dest_cup -= 1
        if dest_pointer == -1:
            # find the largest cup label instead
            dest_cup = get_largest(cups)
            dest_pointer = cups.index(dest_cup)

        print("destination: {}".format(dest_cup))

        dest_pointer = increment_pointer(cups, dest_pointer) 
        cups.insert(dest_pointer, removed.pop())
        cups.insert(dest_pointer, removed.pop())
        cups.insert(dest_pointer, removed.pop())

        # increment pointer to next current cup
        current_pointer = cups.index(current_cup)
        current_pointer = increment_pointer(cups, current_pointer)

    print_pointer = increment_pointer(cups, cups.index(1))
    output = []
    while cups[print_pointer] != 1:
        output.append(str(cups[print_pointer]))
        print_pointer = increment_pointer(cups, print_pointer)
    print("".join(output))

main()
