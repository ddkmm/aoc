#!/usr/bin/env python

import os
import re
import time
from collections import defaultdict

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
DEBUG = False

STEPS = {
  'e': (1, -1, 0),
  'ne': (1, 0, -1),
  'nw': (0, 1, -1),
  'w': (-1, 1, 0),
  'sw': (-1, 0, 1),
  'se': (0, -1, 1)
}

class Coord:
    def __init__(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z

    def move(self, coord):
        self.x += coord[0]
        self.y += coord[1]
        self.z += coord[2]
    
    def get(self):
        return (self.x, self.y, self.z)

def part1(data):
    # centre tile
    floor = defaultdict(lambda: True)
    origin = Coord(0,0,0)
    floor[origin.get()] = True 
    black = set() 

    regex = '(e|w|ne|nw|se|sw)'
    for line in data:
        steps = re.findall(regex,line)
        # always start from centre
        walker = Coord(0,0,0) 
        if DEBUG:
            print(steps)
        
        for step in steps:
            if DEBUG:
                print("Go {}".format(step))
            walker.move(STEPS[step])
        floor[walker.get()] = not floor[walker.get()]
        if not floor[walker.get()]:
            black.add(walker.get())
        else:
            black.discard(walker.get())

    print(len(black))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(DATA) as file:
        data = file.read().splitlines()

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
