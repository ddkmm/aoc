#!/usr/bin/env python

import os
import re

DIRPATH = os.path.dirname(os.path.realpath(__file__))
FILEPATH = os.path.join(DIRPATH, 'input.txt')
TESTPATH = os.path.join(DIRPATH, 'testinput.txt')

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(FILEPATH) as file:
        data = file.readlines()
    print(data)

main()
