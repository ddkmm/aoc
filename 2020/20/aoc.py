#!/usr/bin/env python

import os
import re
import time
import math

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
        self.left = list(zip(*image))[0]
        self.right = list(zip(*image))[-1]

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
        
    def flip(self):
        # swap sides
        temp = self.right
        self.right = self.left
        self.left = temp
        # reverse top and bottom
        self.top.reverse()
        self.bottom.reverse()

    def rotate(self):
        # rotate 90 degrees
        #   t -> r
        #   r.reverse -> b
        #   b -> l
        #   l.reverse -> t
        new_right = self.top
        new_bottom = self.right.reverse()
        new_left = self.bottom
        new_top = self.left.reverse()
        self.top = new_top
        self.right = new_right
        self.bottom = new_bottom
        self.left = new_left

class Puzzle:
    def __init__(self, tiles):
        self.size = int(math.sqrt(len(tiles)))
        canvas = []
        for col in range(0,self.size):
            tile_row = []
            for row in range(0,self.size):
                tile_row.append(tiles[row + 3*col])
            canvas.append(tile_row)
        self.canvas = canvas

    def get_corner_product(self):
        corner_prod = 1
        corner_prod *= self.canvas[0][0].get_id()
        corner_prod *= self.canvas[0][self.size-1].get_id()
        corner_prod *= self.canvas[self.size-1][0].get_id()
        corner_prod *= self.canvas[self.size-1][self.size-1].get_id()
        return corner_prod

def part1(puzzle):
    # Now all we need to do is solve the puzzle...
    print(puzzle.get_corner_product())

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    # Read data and generate tiles 
    tiles = []
    image = []
    with open(TEST) as file:
        for line in file:
            match = re.search(r'\d+', line)
            if match:
                tile_id = int(match.group())
            elif line != '\n':
                image.append(line)
            else:
                new_image = Tile(tile_id, image)
                tiles.append(new_image)
        new_image = Tile(tile_id, image)
        tiles.append(new_image)

    puzzle = Puzzle(tiles)

    time1 = time.perf_counter()
    part1(puzzle)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
