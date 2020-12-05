#!/usr/bin/env python

import os

dir_path = os.path.dirname(os.path.realpath(__file__))
FILEPATH = os.path.join(dir_path, 'input.txt')
TESTPATH = os.path.join(dir_path, 'testinput.txt')

def count_trees(trees, rightstep, downstep):
    print("right {}, down {}".format(rightstep, downstep))
    counter = 0
    n = range(downstep,len(trees),downstep)
    wrap = len(trees[0])
    for x in n:
        y = (int(rightstep*x/downstep)) % wrap
        print ("{},{}: {}".format(y,x,trees[x][y]))
        if trees[x][y] == '#':
            counter += 1
    print ("Total trees: {}".format(counter))
    return counter

def main():
    print("Day {}".format(os.path.split(dir_path)[1]))

    with open(FILEPATH) as fp:
        trees = fp.read().splitlines()
    counter = 1
    counter *= count_trees(trees,1,1)
    counter *= count_trees(trees,3,1)
    counter *= count_trees(trees,5,1)
    counter *= count_trees(trees,7,1)
    counter *= count_trees(trees,1,2)
    print ("Total {}".format(counter))

main()