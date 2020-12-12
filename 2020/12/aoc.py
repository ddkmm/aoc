#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
FILE = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
DEBUG = True

def Manhattan_distance(x,y):
    if x < 0:
        x = -1 * x
    if y < 0:
        y = -1 * y

    return x + y

COMPASS = ('N', 'E', 'S', 'W')
    
class Ferry:
    facing = 'E'
    position = {'X': 0, 'Y': 0}
    compass = {'N': 0, 'E': 1, 'S': 2, 'W': 3}

    def Init(self):
        self.facing = 'E'
        self.position['X'] = 0
        self.position['Y'] = 0

    def SetFacing(self, new_facing):
        self.facing = new_facing
    
    def DoTurn(self, turn_inst):
        self.SetFacing(self.Turn(turn_inst))
        if DEBUG:
            print("New facing: {}".format(self.facing))
        
    def Turn(self, turn_inst):
        instruction = re.split('(\d+)', turn_inst)
        turns = int(int(instruction[1]) / 90)

        # get current heading as a starting point
        current = self.compass[self.facing]
        if (turns != 2 and instruction[0] == 'L'):
            # Add 180 if we need to turn Left
            turns += 2
        turns = current + turns
        turns = turns % 4
        return COMPASS[turns]

    def Test_Turn(self):
        self.facing = 'E'
        assert(self.Turn('R90') == 'S')

        self.facing = 'E'
        assert(self.Turn('L90') == 'N')

        self.facing = 'E'
        assert(self.Turn('L180') == 'W')

        self.facing = 'E'
        assert(self.Turn('R180') == 'W')

        self.facing = 'W'
        assert(self.Turn('L180') == 'E')

        self.facing = 'W'
        assert(self.Turn('R180') == 'E')

        self.facing = 'E'
        assert(self.Turn('R270') == 'N')

        self.facing = 'E'
        assert(self.Turn('L270') == 'S')

        self.facing = 'N'
        assert(self.Turn('R90') == 'E')

        self.facing = 'N'
        assert(self.Turn('L90') == 'W')

    def Test(self):
        self.Test_Turn()

    def Move(self, move_inst):
        print("Plot course for {}".format(move_inst))
        inst = re.split('(\d+)', move_inst)
        if inst[0] == 'F':
            heading = self.facing
        else:
            heading = inst[0]
        print("Current position: {}, {} and heading {}".format(self.position['X'], self.position['Y'], heading))

        if heading == 'N':
            self.position['Y'] += int(inst[1])
        elif heading == 'S':
            self.position['Y'] -= int(inst[1])
        elif heading == 'W':
            self.position['X'] -= int(inst[1])
        elif heading == 'E':
            self.position['X'] += int(inst[1])

        print("New position: {}, {}".format(self.position['X'], self.position['Y']))

    def MoveToWaypoint(self, move_inst, waypoint):
        print("Plot course {} to waypoint".format(move_inst))
        inst = re.split('(\d+)', move_inst)
        assert inst[0] == 'F'
        multiplier = int(inst[1])
        for temp in waypoint.offset:
            cmd = "{}{}".format(temp,waypoint.offset[temp])
            for _ in range(0,multiplier):
                self.Move(cmd)


class Waypoint:
    offset = {'N': 1, 'E': 10, 'S': 0, 'W': 0}
    
    def Translate(self, inst):
        print("Translate waypoint {}".format(inst))
        instruction = re.split('(\d+)', inst)
        turns = int(int(instruction[1]) / 90)
        new_offset = self.offset.copy()

        if (turns != 2 and instruction[0] == 'L'):
            # Add 180 if we need to turn Left
            turns += 2
        # translate values
        for i in range(0,4):
            try:
                new_offset[COMPASS[(turns + i) % 4]] = self.offset[COMPASS[i]]
            except:
                print("bang")
        print("{} becomes {}".format(self.offset, new_offset))
        self.offset = new_offset

        return 0
    
    def Move(self, move_inst):
        print("Move waypoint {}".format(move_inst))
        inst = re.split('(\d+)', move_inst)
        assert inst[0] != 'F'
        heading = inst[0]

        self.offset[heading] += int(inst[1])
        print("New waypoint position: {}".format(self.offset))

        return 0

def part1(nav_data):
    ferry = Ferry()
    ferry.Init()
    if DEBUG:
        ferry.Test()
        ferry.Init()

    for inst in nav_data:
        print(inst)
        if re.search('^R|L',inst):
            ferry.DoTurn(inst)
        else:
            ferry.Move(inst)
    print("Distance travelled: {}".format(Manhattan_distance(ferry.position['X'], ferry.position['Y'])))


def part2(nav_data):
    ferry = Ferry()
    waypoint = Waypoint()
    ferry.Init()
    if DEBUG:
        ferry.Test()
        ferry.Init()

    for inst in nav_data:
        print(inst)
        if re.search('^R|L',inst):
            waypoint.Translate(inst)
        elif re.search('^N|E|S|W',inst):
            waypoint.Move(inst)
        else:
            ferry.MoveToWaypoint(inst, waypoint)

    print("Distance travelled: {}".format(Manhattan_distance(ferry.position['X'], ferry.position['Y'])))

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
