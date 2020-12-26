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

def adjacent_tiles(coord, floor):
    black_adj = set() 
    white_adj = set() 
    for s in STEPS:
        offset_x = coord[0] + STEPS[s][0]
        offset_y = coord[1] + STEPS[s][1]
        offset_z = coord[2] + STEPS[s][2]
        new_coord = Coord(offset_x, offset_y, offset_z)
        if not floor[new_coord.get()]:
            black_adj.add(new_coord.get()) 
        else:
            white_adj.add(new_coord.get()) 

    return black_adj, white_adj

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
            walker.move(STEPS[step])
        floor[walker.get()] = not floor[walker.get()]
        if not floor[walker.get()]:
            black.add(walker.get())
        else:
            black.discard(walker.get())

    print(len(black))
    return floor, black


def part2(floor, black):
    days = 5

    for day in range(days):
        new_floor = floor.copy()
        working_floor = floor.copy()
        white_neighbours = defaultdict(lambda:0) 

        for tile in floor:
            adj_black, adj_white = adjacent_tiles(tile, working_floor)
            if not floor[tile]:
                # this is a black tile, increment the neighbouring black tile count
                # for any adjacent white tiles
                for t in adj_white:
                    white_neighbours[t] += 1
                # apply black tile rule
                if len(adj_black) == 0 or len(adj_black) > 2:
                    new_floor[tile] = True
                    black.discard(tile)
            else:
                # this is a white tile and we need to add neighbouring black count
                white_neighbours[tile] += len(adj_black)

        # Now find how many white tiles have exactly 2 black adjacent tiles
        # white tile rule
        for tile in white_neighbours:
            if white_neighbours[tile] == 2: 
                new_floor[tile] = False
                black.add(tile)

        print("Day {}: {}".format(day+1, len(black)))
        floor = new_floor

    print(len(black))

def get_neighbours(coord):
    neighbours = []
    for s in STEPS:
        offset_x = coord[0] + STEPS[s][0]
        offset_y = coord[1] + STEPS[s][1]
        offset_z = coord[2] + STEPS[s][2]
        new_coord = (offset_x, offset_y, offset_z)
        neighbours.append(new_coord)
    return neighbours

def count_black(tile, black_set):
    black = 0
    neighbours = get_neighbours(tile)
    for t in neighbours:
        if t in black_set:
            black += 1
    return black

def part2b(black):
    days = 100

    for day in range(days):
        new_black = black.copy()
        stack = set()
        visited = set()

        # Add tiles to consider to our stack
        for tile in black:
            stack.add(tile)
            neighbours = get_neighbours(tile)
            stack.update(neighbours)
        while len(stack) > 0:
            tile = stack.pop()
            if tile in visited:
                continue
            visited.add(tile)

            neighbours = count_black(tile, black)
            if tile in black:
                if neighbours == 0 or neighbours > 2:
                    new_black.discard(tile)
            else:
                if neighbours == 2:
                    new_black.add(tile)
        print("Day {}: {}".format(day+1, len(new_black)))
        black = new_black

    print(len(black))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(DATA) as file:
        data = file.read().splitlines()

    time1 = time.perf_counter()
    new_floor, black_set = part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

    time1 = time.perf_counter()
    part2b(black_set)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
