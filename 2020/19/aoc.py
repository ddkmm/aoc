#!/usr/bin/env python

import os
import re
import time
from collections import defaultdict

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
DEBUG = False

def process_rule_dict(rules_dict):
    rule_0 = rules_dict['0']

    sep = '|'
    while True:
        print("Rule 0 = {}".format(rule_0))
        
        # Search for the first digit in the rule
        match = re.search(r'\d+', rule_0)
        if not match:
            break
        i = match.group()
        print("Substituting rule {}: {}".format(i, rules_dict[i]))
        # for each rule number, split on it and generate two new rules,
        # take the rule to substitute in and split on '|'.
        # Then use each as a substitution into the original rule
        rule_0_splits = rule_0.split('|')
        if DEBUG:
            print("Split rule0 on |: {}".format(rule_0_splits))
        new_rule_0 = []
        for r0 in rule_0_splits:
            # This is each | separated part of rule 0. Probably could recurse it.
            original_splits = re.split(i, r0)
            if DEBUG:
                print("Split each part on {}: {}".format(i, original_splits))
            if len(original_splits) == 1:
#                new_rule_0.append(r0)
                continue
            if not i:
                break
            sub_rule_splits = rules_dict[i].split('|')
            if DEBUG:
                print("Splitting rule {} on |: {}".format(i, sub_rule_splits))
            # permute the subs into the original_splits
            new_rules = []
            for sp in sub_rule_splits:
                new_rules.append(str(original_splits[0] + sp + original_splits[1]))
            new_rule_0.append(sep.join(new_rules))
        if new_rule_0:
            rule_0 = sep.join(new_rule_0)

    if DEBUG:
        print(rule_0)

    return rule_0

def apply_rule(rule_0, mesg):
    # the rule dictionary contains white space which needs to be cleaned up
    rule = rule_0.replace(" ", "")
    print(rule)
    sep = "|"
    new_rule_parts = []
    rule_parts = rule.split('|')
    for r in rule_parts:
        new_rule_parts.append(str("\\b" + r + "\\b"))
    rule = sep.join(new_rule_parts)

    print(rule)
    regex = re.compile(rule)
    print(regex)
    total = 0
    for line in mesg:
        if re.search(regex,line):
            print(line)
            total += 1

    return total

def part1(rules, mesg):
    # substitute everything for rule 0
    rule_0 = process_rule_dict(rules)
    # apply rule 0
    print("Part 1: {}".format(apply_rule(rule_0, mesg)))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(TEST) as file:
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

    # replace a and b
    for key in rules_dict:
        value = rules_dict[key]
        value = re.sub(rule_a, 'a', value)
        value = re.sub(rule_b, 'b', value)
        rules_dict[key] = value

    time1 = time.perf_counter()
    part1(rules_dict, mesg)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
