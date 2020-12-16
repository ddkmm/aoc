#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
DEBUG = True

def make_rule(rule):
    assert(len(rule) == 2)
    rule1 = rule[0]
    rule2 = rule[1]
    rule_string = "[{}]|[{}]".format(rule1, rule2)
    return rule_string

def check_rule(number, rule):
    return(re.search(rule, number))

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
    for line in range(data_start, len(data)):
        print(data[line])
        for number in re.split(r',',data[line]):
            res = False
            for rule in rule_list:
                if not check_rule(number, make_rule(rule)):
                    print(number)

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(TEST) as file:
        data = file.read().splitlines()

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
