import os
import re
import time
import numpy as np

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
DEBUG = True
ACTIVE = '#'
INACTIVE = '.'

# If a cube is active and
#   exactly 2 or 3 of its neighbors are also active,
#   the cube remains active.
# Otherwise, the cube becomes inactive.

# If a cube is inactive but
#   exactly 3 of its neighbors are active,
#   the cube becomes active.
# Otherwise, the cube remains inactive.

class Cube:
    x = 0
    y = 0
    z = 0

    def __init__(self, x = 0, y = 0, z = 0):
        self.x = x 
        self.y = y
        self.z = z

    def __eq__(self, other):
        if isinstance(other, Cube):
            return self.x == other.x and self.y == other.y and self.z == other.z

    def __str__(self):
        return "({}, {}, {})".format(self.x, self.y, self.z)

    def __hash__(self):
        return hash(str(self))

    def get_pos(self):
        return self.x, self.y, self.z

    def set_pos(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z

def part1(data):
    world = set()
    if DEBUG:
        print(data)
    for y_val, line in enumerate(data):
        for x_val, cube in enumerate(line):
            if cube == ACTIVE:
                world.add(Cube(x_val, y_val, 0))
    print(len(world))
    tc = Cube(0,0,0)
    if tc in world:
        print("Yes")
    else:
        print("No")
    td = Cube(1,0,0)
    print(str(td))
    if td in world:
        print("Yes")
    else:
        print("No")

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(TEST) as file:
        data = file.read().splitlines()

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
