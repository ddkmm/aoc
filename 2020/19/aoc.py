#!/usr/bin/env python

import os
import re
import time
from collections import defaultdict
from itertools import count

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
DATA2 = os.path.join(DIRPATH, 'input2.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
TEST2 = os.path.join(DIRPATH, 'test2.txt')
DEBUG = False

counter = count()

def get_rule(rules_dict, rule):
    if (rule == 'a' or rule == 'b'):
        return rule

    output = ""
    if '|' not in rule:
        matches = re.findall(r'\d+', rule)
        for m in matches:
            output += get_rule(rules_dict, rules_dict[m])
        return output
    else:
        rule_parts = rule.split('|')
        split_out = "({}|{})".format(get_rule(rules_dict, rule_parts[0]), get_rule(rules_dict, rule_parts[1]))
        return split_out

def get_rule2(rules_dict, rule):
    # rule is a or b, so just return it
    if (rule == 'a' or rule == 'b'):
        return rule
    
    output = ""
    if '|' not in rule:
        # This rule has only a single part, so call the function on each rule in it
        # adding the return value to the string
        matches = re.findall(r'\d+', rule)
        for m in matches:
            output += get_rule2(rules_dict, rules_dict[m])
        return output
    else:
        # otherwise, the rule contains a split, so we call the function on either half
        # of the split, combining them with a split in the resultant rule string
        rule_parts = rule.split('|')
        split_out = "("
        for rp in range(0,len(rule_parts)-1):
            split_out += "{}|".format(get_rule2(rules_dict, rule_parts[rp]))
        split_out += "{})".format(get_rule2(rules_dict, rule_parts[-1]))
        return split_out

def part1(rules, mesg):
    my_rule = get_rule(rules, '0')
    temp = "^" + my_rule + "$"
    print("Regex is {} characters long".format(len(temp)))

    total = 0
    for m in mesg:
        if (re.match(temp, m)):
            total += 1
            if DEBUG:
                print(m)
    print("Part 1: {}".format(total))

def part2(rules, mesg):
    my_rule = get_rule2(rules, '0')
    temp = "^" + my_rule + "$"
    print("Regex is {} characters long".format(len(temp)))

    total = 0
    for m in mesg:
        if (re.match(temp, m)):
            total += 1
            if DEBUG:
                print(m)
    print("Part 2: {}".format(total))

def process_input(input_stream):
    with open(input_stream) as file:
        data = file.read().splitlines()

    rules_dict = {}
    mesg = []
    for line in data:
        if re.search(r':', line):
            rule = re.split(r':', line)
            if re.search('a', rule[1]):
                rule_a = rule[0]
                rules_dict[rule_a] = 'a'
            elif re.search('b', rule[1]):
                rule_b = rule[0]
                rules_dict[rule_b] = 'b'
            else:
                rules_dict[rule[0]] = rule[1]
        elif re.search(r'a|b', line):
            mesg.append(line)
    return rules_dict, mesg

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    rules_dict, mesg = process_input(DATA)
    time1 = time.perf_counter()
    part1(rules_dict, mesg)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

    rules_dict, mesg = process_input(DATA2)
    time1 = time.perf_counter()
    part2(rules_dict, mesg)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
