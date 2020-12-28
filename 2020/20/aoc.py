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
        self.size = len(image)
        self.image = image
        self.orientation = 0
        self.set_edges()

        # collect all edges
        self.edges = []
        self.edges.append(self.top)
        self.edges.append(self.right)
        self.edges.append(self.bottom)
        self.edges.append(self.left)
        self.edges.append(self.top[::-1])
        self.edges.append(self.right[::-1])
        self.edges.append(self.bottom[::-1])
        self.edges.append(self.left[::-1])

        self.top_edge = 0

    def get_all_edges(self):
        return self.edges
    
    def set_edges(self):
        self.top = self.image[0]
        self.bottom = self.image[-1]
        self.left = ''.join([str(elem) for elem in list(zip(*self.image))[0]])
        self.right = ''.join([str(elem) for elem in list(zip(*self.image))[-1]])

    def get_centre(self):
        centre = []
        for line in range(1,len(self.image)-1):
            crop = self.image[line][1:-1]
            centre.append(crop)
        return centre

    def print(self):
        for a in self.image:
            print(a)

    def cycle(self):
        if self.orientation < 4:
            self.rotate()
            self.orientation += 1
        elif self.orientation == 4:
            self.rotate()
            self.flip()
            self.orientation += 1
        elif self.orientation > 4 and self.orientation < 7:
            self.rotate()
            self.orientation += 1
        elif self.orientation == 7:
            self.rotate()
            self.flip()
            self.orientation = 0

    def rotate(self):
        rot_image = []
        for i in range(0,self.size):
            temp = ''.join([str(elem) for elem in list(zip(*self.image))[i]])[::-1]
            rot_image.append(temp)
        self.image = rot_image
        self.set_edges()

    def flip(self):
        # horizontal flip is just reverse each line
        flip_image = []
        for line in self.image:
            temp = line[::-1]
            flip_image.append(temp)
        self.image = flip_image
        self.set_edges()

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

def part1(puzzle):
    edge_dict = defaultdict(lambda: list())
    tiles = puzzle.canvas
    # Get all possible edges for every tile and map the ids to edges.
    # A joining edge will have exactly 2 tiles and boundary edges will only have 1
    # a corner tile will have exactly two boundary edges
    for tile in tiles:
        edges = tile.get_all_edges()
        for edge in edges:
            edge_dict[edge].append(tile.id)

    # Collect all unpaired edges
    single_edges = defaultdict(lambda:0)
    for edge in edge_dict:
        if len(edge_dict[edge]) == 1:
            single_edges[edge] = edge_dict[edge]
    # count how many single edges each tile has
    tile_histogram = defaultdict(lambda:0) 
    for edge in single_edges:
        tile_id = single_edges[edge][0]
        tile_histogram[tile_id] += 1
    prod = 1
    corners = []
    for tile in tile_histogram:
        if tile_histogram[tile] == 4:
            corners.append(tile)
            prod *= tile
    print("Corner tiles: {}".format(corners))
    print("Part 1: {}".format(prod))

    return corners, edge_dict

def find_next_tile(current_tile, edge_dict, tile_dict):
    print("Looking for {}".format(current_tile.right))
    next_tiles = edge_dict[current_tile.right]
    next_tiles.remove(current_tile.id)
    next_tile = tile_dict[next_tiles[0]]
    print("Finding edge on tile {}".format(next_tile.id))
    # and rotate it into place
    while next_tile.left != current_tile.right:
        next_tile.cycle()
    print("Found match {}".format(next_tile.left))
    return next_tile

def find_first_row(starting_tile, length, edge_dict, tile_dict, placed_tiles):
    line = []
    current_tile = starting_tile
    print(edge_dict[current_tile.right])
    line.append(current_tile)
    placed_tiles.add(current_tile)
    print("Starting tile {}".format(current_tile.id))
    while len(line) < length:
        next_tile = find_next_tile(current_tile, edge_dict, tile_dict)
        current_tile = next_tile
        line.append(current_tile)
        placed_tiles.add(current_tile)
    if DEBUG:
        ids = ""
        for t in line:
            ids = ids + " {}".format(t.id)
        print(ids)

    return line, placed_tiles

def find_row(starting_tile, length, edge_dict, tile_dict, placed_tiles, top = True, prev_line = None):
    line = []
    current_tile = starting_tile
    print(edge_dict[current_tile.right])
    line.append(current_tile)
    placed_tiles.add(current_tile)
    print("Starting tile {}".format(current_tile.id))
    while len(line) < length:
        next_tile = find_next_tile(current_tile, edge_dict, tile_dict)
        current_tile = next_tile
        line.append(current_tile)
        placed_tiles.add(current_tile)
    if DEBUG:
        ids = ""
        for t in line:
            ids = ids + " {}".format(t.id)
        print(ids)

    return line, placed_tiles


    return line, placed_tiles


def print_puzzle_ids(layout):
    for line in layout:
        ids = ""
        for t in line:
            ids = ids + " {}".format(t.id)
        print(ids)

def part2(tile_dict, corners, edge_dict, length):
    # start with first corner
    layout = []
    placed_tiles = set()
    # find tile matching right edge
    corner = corners[1]
    current_tile = tile_dict[corner]

    while len(edge_dict[current_tile.right]) != 2:
        print(edge_dict[current_tile.right])
        current_tile.cycle()

    print("Get first row")
    line, placed_tiles = find_first_row(current_tile, length, edge_dict, tile_dict, placed_tiles)
    layout.append(line)

    # Start next row
    current_tile = layout[0][0]
    # did we just make the top row or the bottom row?
    if len(edge_dict[current_tile.bottom]) == 2:
        next_tiles = edge_dict[current_tile.bottom]
        top = True
    else:
        next_tiles = edge_dict[current_tile.top]
        top = False

    next_tiles.remove(current_tile.id)
    next_tile = tile_dict[next_tiles[0]] 
    current_tile = next_tile
    print("Get second row")
    next_line, placed_tiles = find_row(current_tile, length, edge_dict, tile_dict, placed_tiles, top, layout[0])
    if top:
        layout.append(next_line)
        current_tile = layout[1][0]
        next_tiles = edge_dict[current_tile.bottom]
    else:
        layout.insert(0, next_line)
        current_tile = layout[0][0]
        next_tiles = edge_dict[current_tile.top]
    next_tiles.remove(current_tile.id)
    next_tile = tile_dict[next_tiles[0]] 
    current_tile = next_tile
    line = []
    print("Get third row")
    try:
        line, placed_tiles = find_row(current_tile, length, edge_dict, tile_dict, placed_tiles)
    except:
        current_tile.flip()
        line, placed_tiles = find_row(current_tile, length, edge_dict, tile_dict, placed_tiles)
        
    if top:
        layout.append(line)
    else:
        layout.insert(0, line)

    print_puzzle_ids(layout)

    pass

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    # Read data and generate tiles 
    tiles = []
    input = DATA 
    print("Using {}".format(input))
    with open(input) as file:
        for line in file:
            match = re.search(r'\d+', line)
            if match:
                tile_id = int(match.group())
                image = []
            elif line != '\n':
                image.append(line.strip())
            else:
                new_tile = Tile(tile_id, image)
                tiles.append(new_tile)
        new_tile = Tile(tile_id, image)
        tiles.append(new_tile)

    puzzle = Puzzle(tiles)
    length = puzzle.size
    tile_dict = {}
    for tile in tiles:
        tile_dict[tile.id] = tile

    time1 = time.perf_counter()
    corners, edge_dict = part1(puzzle)
    time2 = time.perf_counter()
    part2(tile_dict, corners, edge_dict, length)
    print("{} seconds".format(time2-time1))

main()
