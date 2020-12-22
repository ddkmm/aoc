#!/usr/bin/env python

import os
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
DEBUG = False

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
    if DEBUG:
        print("{} rounds".format(round_number))
    if p1:
        return 1, p1
    return 2, p2

def part1(p1, p2):
    _, winner = do_combat(p1, p2)
    total = calculate_score(winner)
    print("Part 1: {}".format(total))

def do_rec_combat(p1, p2):
    round_number = 0
    played = set()
    hands = (tuple(p1), tuple(p2))

    while p1 and p2:
        round_number += 1
        if hands in played:
            return 1, None
        played.add(hands)
        card1 = p1.pop(0)
        card2 = p2.pop(0)
        len1 = len(p1)
        len2 = len(p2)
        if card1 <= len1 and card2 <= len2:
            if DEBUG:
                print("Round {}, sub-game with {} and {} cards".format(round_number, card1, card2))
            # play recursive combat
            sub_p1 = p1[:card1]
            sub_p2 = p2[:card2]
            wp, _ = do_rec_combat(sub_p1, sub_p2)
            if wp == 1:
                p1.append(card1)
                p1.append(card2)
            else:
                p2.append(card2)
                p2.append(card1)
        else:
            # manually finish this round
            if card1 > card2:
                p1.append(card1)
                p1.append(card2)
            else:
                p2.append(card2)
                p2.append(card1)
        hands = (tuple(p1), tuple(p2))

    if p1:
        return 1, p1

    return 2, p2

def part2(p1, p2):
    _, winner = do_rec_combat(p1, p2)
    total = calculate_score(winner)
    print("Part 2: {}".format(total))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(DATA) as file:
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
