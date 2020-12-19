#!/usr/bin/env python

import os
import re
import time
from collections import defaultdict

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
DEBUG = True

def process_rules(rules, rule_a, rule_b):
    raw_rules = dict(enumerate(rules))
    del raw_rules[rule_a]
    del raw_rules[rule_b]

    processed_rules = defaultdict(lambda: '.')
    processed_rules[rule_a] = 'a'
    processed_rules[rule_b] = 'b'

    # first pass, replace a and b
    raw_clean_up = []
    for raw_key in raw_rules:
        value = raw_rules[raw_key]
        value = re.sub(str(rule_a), 'a', value)
        value = re.sub(str(rule_b), 'b', value)
        print(value)
        m = re.findall(r'\d+', value)
        if not m:
            # This rule has no more numbers
            processed_rules[raw_key] = value
            raw_clean_up.append(raw_key)
        else:
            # update the rule
            raw_rules[raw_key] = value

    for key in raw_clean_up:
        del raw_rules[key]
    raw_clean_up.clear()

    # keep substituting in letter rules for number rules
    while raw_rules:
        for proc_key in processed_rules:
            for raw_key in raw_rules:
                regex = re.compile("({})".format(str(proc_key)))
                m = re.findall(regex, raw_rules[raw_key])
                if m:
                    value = re.sub(regex, processed_rules[proc_key], raw_rules[raw_key])
                    print(value)
                    # update the rule
                    raw_rules[raw_key] = value
                    m2 = re.findall(r'\d+', value)
                    if not m2:
                        # This rule has no more numbers
                        raw_clean_up.append(raw_key)
        if raw_clean_up:
            key = raw_clean_up.pop()
            processed_rules[key] = raw_rules[key]
            del raw_rules[key]
        
    return processed_rules

def apply_rule(rule_0, mesg):
    # the rule dictionary contains white space which needs to be cleaned up
    print(rule_0)
    rule = rule_0.replace(" ", "")
    regex = re.compile(rule)
    print(rule)

def part1(rules, mesg, rule_a, rule_b):
    # substitute everything for rule 0
    processed_rules = process_rules(rules, rule_a, rule_b)
    # apply rule 0
    apply_rule(processed_rules[0], mesg)

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(TEST) as file:
        data = file.read().splitlines()

    rules = []
    mesg = []
    for line in data:
        if re.search(r':', line):
            rule = re.split(r':', line)
            if re.search('a', rule[1]):
                rule_a = int(rule[0])
                rules.append('a')
            elif re.search('b', rule[1]):
                rule_b = int(rule[0])
                rules.append('b')
            else:
                rules.append(rule[1])
        elif re.search(r'a|b', line):
            mesg.append(line)

    time1 = time.perf_counter()
    part1(rules, mesg, rule_a, rule_b)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
