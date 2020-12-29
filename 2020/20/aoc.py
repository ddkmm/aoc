#!/usr/bin/env python

import os
import re
import time
import math
from collections import defaultdict

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
CROP = os.path.join(DIRPATH, 'crop.txt')
DEBUG = False

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
        elif self.orientation > 4 and self.orientation < 8:
            self.rotate()
            self.orientation += 1
        elif self.orientation == 8:
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
    def __init__(self, image):
        self.image = image
        self.orientation = 0
        self.size = len(image)

    def cycle(self):
        old_orientation = self.orientation
        if self.orientation < 4:
            self.rotate()
            self.orientation += 1
        elif self.orientation == 4:
            self.rotate()
            self.flip()
            self.orientation += 1
        elif self.orientation > 4 and self.orientation < 8:
            self.rotate()
            self.orientation += 1
        elif self.orientation == 8:
            self.rotate()
            self.flip()
            self.orientation = 0
        print("orientation {} to {}".format(old_orientation, self.orientation))

    def rotate(self):
        rot_image = []
        for i in range(0,self.size):
            temp = ''.join([str(elem) for elem in list(zip(*self.image))[i]])[::-1]
            rot_image.append(temp)
        self.image = rot_image

    def flip(self):
        flip_image = []
        for line in self.image:
            temp = line[::-1]
            flip_image.append(temp)
        self.image = flip_image

    def print(self):
        for a in self.image:
            print(a)

def part1(tiles):
    edge_dict = defaultdict(lambda: list())
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

def find_next(current_tile, tiles):
    edge = current_tile.right
    next_tile = None
    total = 0
    for tile in tiles:
        total += 1
        i = 0
        while i < 9:
            if edge != tile.left:
                tile.cycle()
                i += 1
            else:
                next_tile = tile
                i = 10 
        if next_tile:
            break

    return next_tile

def find_next_row_start(current_tile, tiles):
    edge = current_tile.bottom
    next_tile = None
    for tile in tiles:
        i = 0
        while i < 9:
            if edge != tile.top:
                tile.cycle()
                i += 1
            else:
                next_tile = tile
                i = 10 
        if next_tile:
            break
    return next_tile

def find_row(current_tile, length, tiles):
    line = []
    line.append(current_tile)
    tiles.remove(current_tile)
    while len(line) < length:
        next_tile = find_next(current_tile, tiles)
        assert(next_tile)
        current_tile = next_tile
        line.append(current_tile)
        tiles.remove(current_tile)

    return line, tiles

def print_puzzle_ids(layout):
    for line in layout:
        ids = ""
        for t in line:
            ids = ids + " {}".format(t.id)
        print(ids)

def part2(tiles, tile_dict, corners, edge_dict, length):
    # start with first corner
    layout = []
    # find tile matching right edge
    corner = corners[0]
    current_tile = tile_dict[corner]

    while len(edge_dict[current_tile.right]) != 2 or len(edge_dict[current_tile.bottom]) != 2:
        current_tile.cycle()

    line = []
    new_tiles = tiles.copy()

    line, new_tiles = find_row(current_tile, length, new_tiles)
    layout.append(line)

    for i in range(length-1):
        current_tile = layout[i][0]
        next_tile = find_next_row_start(current_tile, new_tiles)
        current_tile = next_tile
        line, new_tiles = find_row(current_tile, length, new_tiles)
        layout.append(line)

    if DEBUG:
        print_puzzle_ids(layout)

    return layout

def border_remover(layout):
    cropped = []
    for line in layout:
        cropped_line = []
        for tile in line:
            cropped_line.append(tile.get_centre())
        cropped.append(cropped_line)

    # for each line, combine each row of each tile together
    temp = []
    cropped_length = len(cropped[0][0])
    for line in cropped:
        for i in range(cropped_length):
            dave = []
            for tile in line:
                strip = tile[i]
                for c in strip:
                    dave.append(c)
            temp.append(dave)
    f = open(CROP, "w")
    for line in temp:
        sep = ""
        f.write(sep.join(line))
        f.write("\n")
    f.close()

def search(monsters):
    total = 0
    total = find_monster(monsters)
    while total == 0:
        monsters.cycle()
        total = find_monster(monsters)
    return total

def find_monster(monsters):
    monster_body = r'(?=(#.{4}##.{4}##.{4}###))'
    mb = re.compile(monster_body)
    total = 0
    image = monsters.image
    rows = range(1, len(image) - 1)
    bodies = 0
    heads = 0
    for row in rows:
        matched_bodies = mb.finditer(image[row])
        for m in matched_bodies:
            bodies += 1
            if DEBUG:
                print("Body {} found on line {} between {}".format(bodies, row, m.span()))
            if image[row-1][m.span()[0]+18] != '#':
                print(image[row-1])
                continue
            else:
                heads += 1
                if DEBUG:
                    print("Head {} found at {}, {}".format(heads, row-1, m.span()[0]+18))
                    print("Looking for feet on line {}".format(row+1))
                feet_start = m.span()[0]+1
                found = True
                for i in range(0, 18, 3):
                    if DEBUG:
                        print("{}, {}: {}".format(row+1, feet_start+i, image[row+1][feet_start+i]))
                    if image[row+1][feet_start+i] != '#':
                        found = False
                if found:
                    total += 1
                    if DEBUG:
                        print("Monster {} found".format(total))

    return total

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

    length = int(math.sqrt(len(tiles)))
    tile_dict = {}
    for tile in tiles:
        tile_dict[tile.id] = tile

    time1 = time.perf_counter()
    corners, edge_dict = part1(tiles)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

    time1 = time.perf_counter()
    layout = part2(tiles, tile_dict, corners, edge_dict, length)
    border_remover(layout)
    with open(CROP) as file:
        data = file.read().splitlines()
    total_marks = 0
    for line in data:
        total_marks += line.count('#')
    monsters = Puzzle(data)
    monsters_found = search(monsters)
    # Each monster contains 15 # symbols
    remove = 15 * monsters_found
    print("{} monsters found = {} marks".format(monsters_found, remove))
    print("{} - {} = {}".format(total_marks, remove, total_marks - remove))
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
