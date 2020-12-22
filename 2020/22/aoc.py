#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
DEBUG = True

def calculate_score(winning_hand):
    total = 0
    end = len(winning_hand)
    for i in range(1, end+1):
        card = winning_hand.pop()
        total += card * i
    return total

def do_combat(p1, p2):
    round_number = 0
    while p1 and p2:
        round_number += 1
        card1 = p1.pop(0)
        card2 = p2.pop(0)
        if card1 > card2:
            p1.append(card1)
            p1.append(card2)
        else:
            p2.append(card2)
            p2.append(card1)
    print("{} rounds".format(round_number))
    if p1:
        return p1
    else:
        return p2

def part1(p1, p2):
    total = calculate_score(do_combat(p1, p2))
    print("Part 1: {}".format(total))

def part2(p1, p2):
    winner = do_combat(p1, p2)
    total = calculate_score(winner)
    print("Part 2: {}".format(total))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(TEST) as file:
        data = file.read().splitlines()

    index = data.index("Player 2:")
    player_1 = data[1:index-1]
    player_2 = data[index+1:]

    p1 = [int(i) for i in player_1]
    p2 = [int(i) for i in player_2]

    time1 = time.perf_counter()
    part1(p1,p2)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

    # reset cards
    p1 = [int(i) for i in player_1]
    p2 = [int(i) for i in player_2]
    time1 = time.perf_counter()
    part2(p1,p2)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
