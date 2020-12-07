#!/usr/bin/env python

import re
import os

dir_path = os.path.dirname(os.path.realpath(__file__))
FILEPATH = os.path.join(dir_path, 'input.txt')
TESTPATH = os.path.join(dir_path, 'testinput.txt')
TEST2PATH = os.path.join(dir_path, 'testinput2.txt')

rules = []

def find_bag1(rules, target, output_set):
    for rule in rules:
        bag_list = re.split(',', rule)
        for i in range(0,len(bag_list)):
            if i == 0 and re.search(target, bag_list[i]):
                if not re.search(target, "shiny gold"):
                    output_set.add(rule)
            elif re.search(target, bag_list[i]):
                find_bag1(rules, bag_list[0], output_set)

def main():
    print("Day {}".format(os.path.split(dir_path)[1]))

    with open(TESTPATH) as f:
        input = f.readlines()

    for a in input:
        rule_string = ""
        b = re.split("contain", a.rstrip())
        assert(len(b) == 2)
        parent = b[0]
        rule_string += str("{} {}, ".format(parent.split()[0], parent.split()[1]))

        children = b[1].split(',')
        for c in children:
            if not re.search("no other bags", c):
                rule_string += str("{} {}, ".format(c.split()[1], c.split()[2]))
        rules.append(rule_string)

    output_set = set() 
    find_bag1(rules, "shiny gold", output_set)
    print("Part 1: {}".format(len(output_set)))


main()
