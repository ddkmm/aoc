#!/usr/bin/env python

import os

dir_path = os.path.dirname(os.path.realpath(__file__))
FILEPATH = os.path.join(dir_path, 'input.txt')
TESTPATH = os.path.join(dir_path, 'testinput.txt')

def part1(input):
    forms = {}
    total = 0
    for a in input:
        if a != '\n':
            for b in a.strip():
                forms[b] = 1
        else:
            total += (len(forms))
            forms.clear()

    # One last time after file ends
    total += (len(forms))
    forms.clear()
    print ("Part 1: {}". format(total))

def part2(input):
    forms = [] 
    total = 0
    i = 0
    for a in input:
        if a != '\n':
            forms.append({})
            for b in a.strip():
                forms[i][b] = 1
            i += 1
        else:
            if len(forms) == 1:
                total += (len(forms[0]))
            else:
                largest_dict = {}
                for a_dict in forms:
                    if len(a_dict) > len(largest_dict):
                        largest_dict = a_dict
                for val in largest_dict:
                    isPresent = True
                    for d in forms:
                        try:
                            d[val]
                        except:
                            isPresent = False
                    if isPresent:
                        total += 1
            forms = []
            i = 0

    # One last time after file ends
    largest_dict = {}
    for a_dict in forms:
        if len(a_dict) > len(largest_dict):
            largest_dict = a_dict
    for val in largest_dict:
        isPresent = True
        for d in forms:
            try:
                d[val]
            except:
                isPresent = False
        if isPresent:
            total += 1

    print ("Part 2: {}". format(total))

def main():
    print("Day {}".format(os.path.split(dir_path)[1]))

    with open(FILEPATH) as f:
#    with open(TESTPATH) as f:
        input = f.readlines()
    part1(input)
    part2(input)

main()
