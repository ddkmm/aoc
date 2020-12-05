#!/usr/bin/env python

import os

dir_path = os.path.dirname(os.path.realpath(__file__))
FILEPATH = os.path.join(dir_path, 'input.txt')
TESTPATH = os.path.join(dir_path, 'testinput.txt')

def main():
    print("Day {}".format(os.path.split(dir_path)[1]))

    with open(FILEPATH) as f:
        input = f.readlines()
    print (input)

main()
