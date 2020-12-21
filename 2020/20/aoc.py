#!/usr/bin/env python

import os
import re
import time
import math
from collections import defaultdict

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
DEBUG = True

class Tile:
    def __init__(self, id, image):
        self.id = int(id)
        self.image = image
        self.size = len(image)
        self.top = image[0]
        self.bottom = image[-1]
        self.left = ''.join([str(elem) for elem in list(zip(*image))[0]])
        self.right = ''.join([str(elem) for elem in list(zip(*image))[-1]])

    def get_top(self):
        return self.top

    def get_bottom(self):
        return self.bottom

    def get_left(self):
        return self.left

    def get_right(self):
        return self.right

    def get_id(self):
        return self.id

    def get_all_edges(self):
        edges = []
        edges.append(self.top)
        edges.append(self.right)
        edges.append(self.bottom)
        edges.append(self.left)
        edges.append(self.top[::-1])
        edges.append(self.right[::-1])
        edges.append(self.bottom[::-1])
        edges.append(self.left[::-1])
        return edges
        
    def flip(self):
        # swap sides
        temp = self.right
        self.right = self.left
        self.left = temp
        # reverse top and bottom
        self.top[::-1]
        self.bottom[::-1]

    def rotate(self):
        # rotate 90 degrees
        #   t -> r
        #   r.reverse -> b
        #   b -> l
        #   l.reverse -> t
        new_right = self.top
        new_bottom = self.right[::-1]
        new_left = self.bottom
        new_top = self.left[::-1]
        self.top = new_top
        self.right = new_right
        self.bottom = new_bottom
        self.left = new_left

class Puzzle:
    def __init__(self, tiles):
        self.size = int(math.sqrt(len(tiles)))
        self.canvas = tiles

    def get_corner_product(self):
        corner_prod = 1
        corner_prod *= self.canvas[0].get_id()
        corner_prod *= self.canvas[self.size - 1].get_id()
        corner_prod *= self.canvas[-self.size].get_id()
        corner_prod *= self.canvas[-1].get_id()
        return corner_prod

    def solve(self):
        for start_tile in self.canvas:
            for tile in self.canvas:
                if start_tile.get_id == tile.get_id:
                    continue
                print("Compare {} with {}".format(start_tile.get_id(), tile.get_id()))
                match_right(start_tile, tile)
            start_tile.rotate()
            for tile in self.canvas:
                if start_tile.get_id == tile.get_id:
                    continue
                print("Compare {} with {}".format(start_tile.get_id(), tile.get_id()))
                match_right(start_tile, tile)
            start_tile.rotate()
            for tile in self.canvas:
                if start_tile.get_id == tile.get_id:
                    continue
                print("Compare {} with {}".format(start_tile.get_id(), tile.get_id()))
                match_right(start_tile, tile)
            start_tile.rotate()
            for tile in self.canvas:
                if start_tile.get_id == tile.get_id:
                    continue
                print("Compare {} with {}".format(start_tile.get_id(), tile.get_id()))
                match_right(start_tile, tile)

def match_right(tile1, tile2):
    for _ in range(0,3):
        if tile1.get_right() == tile2.get_left():
            print("Match!")
            return True
        tile2.rotate()
    tile2.flip()
    tile2.rotate()
    if tile1.get_right() == tile2.get_left():
        return True
    tile2.rotate()
    tile2.rotate()
    if tile1.get_right() == tile2.get_left():
        return True

def step(start_tile, puzzle):
    for tile in puzzle.canvas:
        if start_tile.get_id == tile.get_id:
            continue
        match_right(start_tile, tile)

def solve_puzzle(start_tile, puzzle):
    for _ in range(0,3):
        if step(start_tile, puzzle):
            print("Match!")
            return True
        start_tile.rotate()
    start_tile.flip()
    start_tile.rotate()
    if step(start_tile, puzzle):
        print("Match!")
        return True
    start_tile.rotate()
    start_tile.rotate()
    if step(start_tile, puzzle):
        print("Match!")
        return True

def part1(puzzle):
    edge_dict = defaultdict(lambda: list())
    tiles = puzzle.canvas
    # Get all possible edges for every tile and map the ids to edges.
    # A joining edge will have exactly 2 tiles and boundary edges will only have 1
    # a corner tile will have exactly two boundary edges
    for tile in tiles:
        edges = tile.get_all_edges()
        for edge in edges:
            temp = edge_dict[edge]
            temp.append(tile.id)
            e1 = {edge: temp}
            edge_dict.update(e1)
        edges.clear()

    edge_count = defaultdict(lambda:0)
    for edge in edge_dict:
        if len(edge_dict[edge]) == 1:
            edge_count[edge_dict[edge][0]] += 1
    prod = 1
    for tile in edge_count:
        if edge_count[tile] == 2:
            prod *= tile
    print("Part 1: {}".format(prod))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    # Read data and generate tiles 
    tiles = []
    image = []
    with open(DATA) as file:
        for line in file:
            match = re.search(r'\d+', line)
            if match:
                tile_id = int(match.group())
            elif line != '\n':
                image.append(line.strip())
            else:
                new_image = Tile(tile_id, image)
                tiles.append(new_image)
                image.clear()
        new_image = Tile(tile_id, image)
        tiles.append(new_image)

    puzzle = Puzzle(tiles)

    time1 = time.perf_counter()
    part1(puzzle)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
