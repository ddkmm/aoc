#!/usr/bin/env python

import os
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = "123487596"
TEST = "389125467"
DEBUG = False

def increment_pointer(cups, pointer):
    pointer += 1
    pointer = pointer % len(cups)
    return pointer

def get_largest(cups):
    largest = 0
    for cup in cups:
        if cup > largest:
            largest = cup
    return largest

def part1(cups):
    current_pointer = 0
    for i in range(100):
        current_cup = cups[current_pointer]
        if DEBUG:
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
        if DEBUG:
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

        if DEBUG:
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
    answer = "".join(output)
    print("Part 1: {}".format(answer))

class Cup:
    def __init__(self, label):
        self.label = label
        self.next = None

# A better solution using a linked list. Makes part 2 possible.
def part1_opt(data):
    cups = {}
    max_label = 9

    head = data[0]
    for cup in data:
        new_cup = Cup(cup)
        cups[cup] = new_cup

    for i, cup in enumerate(data[:-1]):
        next_cup = cups[data[i+1]]
        cups[cup].next = next_cup

    cups[data[-1]].next = cups[head]

    # get starting cup
    current_cup = cups[data[0]]
    for i in range(100):
        if DEBUG:
            print("-- move {} --".format(i+1))
            labels = ""
            cup = current_cup
            while cup.next.label != current_cup.label:
                labels = labels + " {}".format(cup.label)
                cup = cup.next
            labels = labels + " {}".format(cup.label)
            print("{}".format(labels))
            print("current cup {}".format(current_cup.label))
        removed = [current_cup.next.label, current_cup.next.next.label,
                   current_cup.next.next.next.label]
        if DEBUG:
            print("pick up: [{}, {}. {}]".format(removed[0], removed[1], removed[2]))
        # update the old loop
        current_cup.next = current_cup.next.next.next.next

        # get destination cup
        dest_cup = current_cup.label - 1
        if dest_cup == 0:
            dest_cup = max_label
        while dest_cup in removed:
            dest_cup -= 1
            if dest_cup == 0:
                dest_cup = max_label
        if DEBUG:
            print("destination: {}".format(dest_cup))

        # move cups
        cups[removed[2]].next = cups[dest_cup].next
        cups[dest_cup].next = cups[removed[0]]
        # increment index to next current cup
        current_cup = current_cup.next

    output = []
    nxt = cups[1].next
    while nxt.label != 1:
        output.append(str(nxt.label))
        nxt = cups[nxt.label].next
    answer = "".join(output)
    print("Part 1: {}".format(answer))

def part2(data, total_cups):
    cups = {}
    max_label = 1000000
    rounds = 10000000

    extra_cups = []
    for i in range(10,total_cups+1):
        extra_cups.append(i)
    data = data + extra_cups

    head = data[0]
    for cup in data:
        new_cup = Cup(cup)
        cups[cup] = new_cup

    for i, cup in enumerate(data[:-1]):
        next_cup = cups[data[i+1]]
        cups[cup].next = next_cup

    cups[data[-1]].next = cups[head]

    # get starting cup
    current_cup = cups[data[0]]
    for i in range(rounds):
        if DEBUG:
            print("-- move {} --".format(i+1))
            labels = ""
            cup = current_cup
            while cup.next.label != current_cup.label:
                labels = labels + " {}".format(cup.label)
                cup = cup.next
            labels = labels + " {}".format(cup.label)
            print("{}".format(labels))
            print("current cup {}".format(current_cup.label))
        removed = [current_cup.next.label, current_cup.next.next.label,
                   current_cup.next.next.next.label]
        if DEBUG:
            print("pick up: [{}, {}. {}]".format(removed[0], removed[1], removed[2]))
        # update the old loop
        current_cup.next = current_cup.next.next.next.next

        # get destination cup
        dest_cup = current_cup.label - 1
        if dest_cup == 0:
            dest_cup = max_label
        while dest_cup in removed:
            dest_cup -= 1
            if dest_cup == 0:
                dest_cup = max_label
        if DEBUG:
            print("destination: {}".format(dest_cup))

        # move cups
        cups[removed[2]].next = cups[dest_cup].next
        cups[dest_cup].next = cups[removed[0]]
        # increment index to next current cup
        current_cup = current_cup.next

    print("Part 2: {}".format(cups[1].next.next.label * cups[1].next.next.next.label))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    cups = [int(i) for i in list(TEST)]

    time1 = time.perf_counter()
    part1(cups)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

    cups = [int(i) for i in list(TEST)]
    time1 = time.perf_counter()
    part1_opt(cups)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

    cups = [int(i) for i in list(DATA)]
    time1 = time.perf_counter()
    part2(cups, 1000000)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
