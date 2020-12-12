#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
FILE = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
DEBUG = False

def Manhattan_distance(x,y):
    if x < 0:
        x = -1 * x
    if y < 0:
        y = -1 * y

    return x + y

COMPASS = ('N', 'E', 'S', 'W')

class Ferry:
    heading = 'E'
    position = (0,0)
    compass = {'N': 0, 'E': 1, 'S': 2, 'W': 3}

    def Init(self):
        self.heading = 'E'
        self.position = (0,0)

    def Turn(self, new_heading):
        instruction = re.split('(\d+)', new_heading)
        turns = int(int(instruction[1]) / 90)

        # get current heading as a starting point
        # and add our new turns
        current = self.compass[self.heading]
        
        if instruction[0] is 'R':
            turns = current + turns
            turns = turns % 4
        else:
            turns = -1 * turns
            turns = current + turns

        self.heading = COMPASS[turns]
        if DEBUG:
            print("New heading: {}".format(self.heading))

    def Test_Turn(self):
        self.heading = 'E'
        self.Turn('R90')
        assert(self.heading == 'S')

        self.heading = 'E'
        self.Turn('L90')
        assert(self.heading == 'N')

        self.heading = 'E'
        self.Turn('L180')
        assert(self.heading == 'W')

        self.heading = 'E'
        self.Turn('R180')
        assert(self.heading == 'W')

        self.heading = 'E'
        self.Turn('R270')
        assert(self.heading == 'N')

        self.heading = 'E'
        self.Turn('L270')
        assert(self.heading == 'S')

        self.heading = 'N'
        self.Turn('R90')
        assert(self.heading == 'E')

        self.heading = 'N'
        self.Turn('L90')
        assert(self.heading == 'W')

    def Test(self):
        self.Test_Turn()

    def Move(self, course):
        print("Plot course for {}".format(course))


def part1(nav_data):
    ferry = Ferry()
    ferry.Init()
    if DEBUG:
        ferry.Test()
        ferry.Init()

    for inst in nav_data:
        print(inst)
        if re.search('^R|L',inst):
            ferry.Turn(inst)
        else:
            ferry.Move(inst)

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(FILE) as file:
        data = file.read().splitlines()

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
