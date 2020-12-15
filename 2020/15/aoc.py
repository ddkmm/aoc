#!/usr/bin/env python

import os
import re
import time
from collections import defaultdict

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DEBUG = True

def part1(numbers):
    
    # If that was the first time the number has been spoken,
    #   the current player says 0.
    # Otherwise, the number had been spoken before,
    #   the current player announces how many turns apart
    #   the number is from when it was previously spoken.

    # Test data: (0, 3, 6)
    # Turn    1  2  3  4  5  6  7  8  9 10
    # Number  0  3  6  0  3  3  1  0  4  0 
    # number_history is a dictionary
    #   key is the number
    #   value is a tuple of the last two turns it was spoken
    #       (prev_turn, prev_prev_turn)
    number_history = defaultdict(lambda: (-1, -1))
    for i, n in enumerate(numbers):
        number_history[n] = (i + 1, -1)
        # {0:(1, -1)}, {3:(2, -1)}, {6:(3, -1)}

    end_turn = 2020 
    current_turn = len(numbers) + 1
    spoken_number = numbers[-1]
    while current_turn <= end_turn:
        history = number_history[spoken_number]
        if history[1] == -1:
            # number was spoken for the first time
            current_number = 0
            history = number_history[current_number]
            new_history = (current_turn, history[0])
            number_history[current_number] = new_history
            print("Turn {}, number {}".format(current_turn, current_number))
            spoken_number = current_number
        else:
            # This number was already spoken, find history difference
            current_number = history[0] - history[1]
            history = number_history[current_number]
            new_history = (current_turn, history[0])
            number_history[current_number] = new_history
            print("Turn {}, number {}".format(current_turn, current_number))
            spoken_number = current_number
        current_turn += 1


def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    data = [8,13,1,0,18,9]
    test = [0,3,6]

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
