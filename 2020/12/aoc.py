#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
FILE = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
DEBUG = False

def m_distance(x_coord, y_coord):
    if x_coord < 0:
        x_coord = -1 * x_coord
    if y_coord < 0:
        y_coord = -1 * y_coord

    return x_coord + y_coord

COMPASS = ('N', 'E', 'S', 'W')

class Ferry:
    facing = 'E'
    position = {'X': 0, 'Y': 0}
    compass = {'N': 0, 'E': 1, 'S': 2, 'W': 3}

    def init(self):
        self.facing = 'E'
        self.position['X'] = 0
        self.position['Y'] = 0

    def set_facing(self, new_facing):
        self.facing = new_facing

    def do_turn(self, turn_inst):
        self.set_facing(self.turn(turn_inst))
        if DEBUG:
            print("New facing: {}".format(self.facing))

    def turn(self, turn_inst):
        instruction = re.split(r'(\d+)', turn_inst)
        turns = int(int(instruction[1]) / 90)

        # get current heading as a starting point
        current = self.compass[self.facing]
        if (turns != 2 and instruction[0] == 'L'):
            # Add 180 if we need to turn Left
            turns += 2
        turns = current + turns
        turns = turns % 4
        return COMPASS[turns]

    def test_turn(self):
        self.facing = 'E'
        assert self.turn('R90') == 'S'

        self.facing = 'E'
        assert self.turn('L90') == 'N'

        self.facing = 'E'
        assert self.turn('L180') == 'W'

        self.facing = 'E'
        assert self.turn('R180') == 'W'

        self.facing = 'W'
        assert self.turn('L180') == 'E'

        self.facing = 'W'
        assert self.turn('R180') == 'E'

        self.facing = 'E'
        assert self.turn('R270') == 'N'

        self.facing = 'E'
        assert self.turn('L270') == 'S'

        self.facing = 'N'
        assert self.turn('R90') == 'E'

        self.facing = 'N'
        assert self.turn('L90') == 'W'

    def test(self):
        self.test_turn()

    def move(self, move_inst):
        inst = re.split(r'(\d+)', move_inst)
        if inst[0] == 'F':
            heading = self.facing
        else:
            heading = inst[0]

        if heading == 'N':
            self.position['Y'] += int(inst[1])
        elif heading == 'S':
            self.position['Y'] -= int(inst[1])
        elif heading == 'W':
            self.position['X'] -= int(inst[1])
        elif heading == 'E':
            self.position['X'] += int(inst[1])

    def move_to_waypoint(self, move_inst, waypoint):
        inst = re.split(r'(\d+)', move_inst)
        assert inst[0] == 'F'
        multiplier = int(inst[1])
        for temp in waypoint.offset:
            cmd = "{}{}".format(temp,waypoint.offset[temp])
            for _ in range(0,multiplier):
                self.move(cmd)


class Waypoint:
    offset = {'N': 1, 'E': 10, 'S': 0, 'W': 0}

    def translate(self, inst):
        if DEBUG:
            print("Translate waypoint {}".format(inst))
        instruction = re.split(r'(\d+)', inst)
        turns = int(int(instruction[1]) / 90)
        new_offset = self.offset.copy()

        if (turns != 2 and instruction[0] == 'L'):
            # Add 180 if we need to turn Left
            turns += 2
        # translate values
        for i in range(0,4):
            new_offset[COMPASS[(turns + i) % 4]] = self.offset[COMPASS[i]]
        if DEBUG:
            print("{} becomes {}".format(self.offset, new_offset))
        self.offset = new_offset

        return 0

    def move(self, move_inst):
        if DEBUG:
            print("Move waypoint {}".format(move_inst))
        inst = re.split(r'(\d+)', move_inst)
        assert inst[0] != 'F'
        heading = inst[0]

        self.offset[heading] += int(inst[1])
        if DEBUG:
            print("New waypoint position: {}".format(self.offset))

        return 0

def part1(nav_data):
    ferry = Ferry()
    ferry.init()
    if DEBUG:
        ferry.test()
        ferry.init()

    for inst in nav_data:
        if re.search('^R|L',inst):
            ferry.do_turn(inst)
        else:
            ferry.move(inst)
    print("Distance travelled: {}".format(m_distance(ferry.position['X'], ferry.position['Y'])))


def part2(nav_data):
    ferry = Ferry()
    waypoint = Waypoint()
    ferry.init()
    if DEBUG:
        ferry.test()
        ferry.init()

    for inst in nav_data:
        if re.search('^R|L',inst):
            waypoint.translate(inst)
        elif re.search('^N|E|S|W',inst):
            waypoint.move(inst)
        else:
            ferry.move_to_waypoint(inst, waypoint)

    print("Distance travelled: {}".format(m_distance(ferry.position['X'], ferry.position['Y'])))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(FILE) as file:
        data = file.read().splitlines()

    print("Part 1")
    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

    print("Part 2")
    time1 = time.perf_counter()
    part2(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
