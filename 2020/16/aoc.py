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
#    rule = ('2-13', '16-100')

    # rule = ('1-3', '5-7')
    # ^([1-3]|[5-7])$
    # rule = ('2-13', '16-100')
    # ^([2-9]|1[0-3]|1[6-9]|[2-9][0-9]|100)$
    # rules are of the form 1-13, 16-100
    # for each part of the rule, need to split around the - 
    #   and get the start and end digit for each filter
    regex = "^("
    part = re.split(r'\W', rule[0])
    # simple case is both are single digits
    if (len(part[0]) == 1 and len(part[1]) == 1):
        regex = regex + "[{}-{}]".format(part[0], part[1])
    elif (len(part[0]) == 1 and len(part[1]) != 1):
        digits = re.findall(r'[0-9]', part[1])
        # digits = ("2", "4")
        
    regex = regex + "|"
    part = re.split(r'\W', rule[1])
    if (len(part[0]) == 1 and len(part[1]) == 1):
        regex = regex + "[{}-{}]".format(part[0], part[1])

    regex = regex + ")$"
    return regex

def check_rule(number, rule):
    if re.search(rule, number):
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
