#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
DEBUG = True

def check_rule(number, rule):
    assert(len(rule) == 2)
    digits = re.split(r'\W', "-".join(rule))
    if ((int(digits[0]) <= number <= int(digits[1])) or
        (int(digits[2]) <= number <= int(digits[3]))):
        return True
    return False

def part1(data):
    rule_list = []
    data_start = 0
    for i, line in enumerate(data):
        temp = re.split(r':', line.strip())
        if len(temp) == 2:
            if re.search(r'^your', temp[0]):
                continue
            elif re.search(r'^nearby', temp[0]):
                data_start = i+1
            else:
                # this is a rule
                fields = re.findall(r'\d+-\d+', temp[1].strip())
                rule_list.append(fields)

    # apply rules
    total = 0
    for line in range(data_start, len(data)):
        for number in re.split(r',',data[line]):
            number = int(number)
            res = False
            for rule in rule_list:
                res = (res or check_rule(number, rule))
            if not res:
                total += number
    print(total)

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(DATA) as file:
        data = file.read().splitlines()

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
