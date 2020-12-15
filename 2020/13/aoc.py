#!/usr/bin/env python

import os
import re
import time
from collections import defaultdict

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
DEBUG = True

def get_next_bus(bus_id, timestamp):
    current_bus_time = 0
    while current_bus_time < timestamp:
        current_bus_time += bus_id
    wait_time = current_bus_time - timestamp
    print("wait {} * bus {} = {}".format(wait_time, bus_id, bus_id * wait_time))

# Somebody else's solution
def find_convergence(timetable):
    e_timetable = enumerate(timetable)

    current_tick = 0
    tick_size = 1
    for offset, bus in e_timetable:
       	if bus == 'x':
            continue
        bus = int(bus)
        needed_remainder = (bus - offset) % bus
        print("{} = {} - {} % {}".format(needed_remainder, bus, offset, bus))

        while current_tick % bus - needed_remainder != 0:
            print("{} % {} = {} - {} != 0".format(current_tick, bus, current_tick % bus, needed_remainder))
            current_tick += tick_size

        print(f"{current_tick:15d} % {bus:3d} = {current_tick % bus:3d}")
        tick_size = tick_size * bus 


def temp():
    b = [7,13]
    i = 0
    j = 0
    while b[0] * i != (b[1] * j) + 1:
        if b[0]*i > b[1]*j + 1:
            j += 1
        elif b[0]*i < b[1]*j+1:
            i += 1
    print(b[0]*i, b[1]*j + 1)

        

# chinese remainder theorem
def do_crt(data):
    # need to make equations in the form
    # x = b[i] (mod(n[i]))
    # x = offset (mod(bus_id))
    # for test input
    # 7,13,x,x,59,x,31,19
    # x =   mod(7)
    # x = 1 mod(13)
    # x = 4 mod(59)
    # x = 6 mod(31)
    # x = 7 mod(19)
    # table to solve
    # b[i]*N[i]*x[i]
    # b[] = (0, 1, 4, 6, 7) 
    # n[] = (7, 13, 59, 31, 19)
    # N = 7 * 13 * 59 * 31 * 19 
    # N[i] = N/n[i]
    # x is found by iteration
    b_list = []
    n_list = []
    x_list = []
    N = 1
    for b, n in enumerate(data):
        if n != 'x':
            b_list.append(b)
            n_list.append(int(n))
            N *= int(n)
            x_list.append(b**(int(n)-2)%int(n))
    total = 0
    for i in range(0, len(n_list)):
        total += b_list[i] * (int(N/n_list[i])) * x_list[i]

  # x = b**(n-2) % n
  # x =   mod(7)
    x = 0**(7-2) % 7 
  # x = 1 mod(13)
    x = 1**(13-2) % 13 
  # x = 4 mod(59)
    x = 4**(59-2) % 59 
  # x = 6 mod(31)
    x = 6**(31-2) % 31 
  # x = 7 mod(19)
    x = 7**(19-2) % 19 

    return total 


def part1(data):
    print("Part 1")
    timestamp = int(data[0])
    timetable = re.findall(r'[0-9]+', data[1])
    for bus in timetable:
        get_next_bus(int(bus), timestamp)

def part2(data):
    print("Part 2")
    timetable = re.split(r',', data[1])
    print(do_crt(timetable))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(TEST) as file:
        data = file.read().splitlines()
    temp()

    time1 = time.perf_counter()
#    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

    time1 = time.perf_counter()
    part2(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
