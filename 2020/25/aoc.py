#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
DEBUG = False

def find_loop(subject_number, card_key, door_key):
    value = 1
    handshake = 20201227
    found_card = False
    found_door = False
    i = 1
    while not (found_card and found_door):
        value *= subject_number
        value = value % handshake
        if value == card_key:
            found_card = True
            card_loop = i
            print("Card loop found: {}".format(card_loop))
        if value == door_key:
            found_door = True
            door_loop = i
            print("Door loop found: {}".format(door_loop))
        if DEBUG:
            print("Loop {}: {}".format(i, value))
        i += 1
    return card_loop, door_loop

def transform(subject_number, loop):
    value = 1
    handshake = 20201227
    for _ in range(loop):
        value *= subject_number
        value = value % handshake
    return value

def part1(data):
    card_public = int(data[0])
    door_public = int(data[1])
    card_loop = 0
    door_loop = 0
    subject_number = 7
    card_loop, door_loop = find_loop(subject_number, card_public, door_public) 
    print("Card loop: {}, door loop: {}".format(card_loop, door_loop))
    card_private = transform(card_public, door_loop)
    door_private = transform(door_public, card_loop)

    print(card_private)
    print(door_private)

    if DEBUG:
        print(data)

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(DATA) as file:
        data = file.read().splitlines()

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
