#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
DEBUG = True

def do_calc2(equation):
    if len(equation) == 1 and equation[0].isdigit():
        return int(equation[0])

    if equation[0] == '(':
        # reform our character array and find matching pair
        del equation[0]
        equation.remove(')')
        # remove both parentheses

    for index, a in enumerate(equation):
        if a.isdigit():
            left = int(a)
        elif a == '+':
            left += do_calc(equation[index+1:])
            return left
        elif a == '*':
            left *= do_calc(equation[index+1:])
            return left

def do_calc(equation):
    calc_stack = []
    for a in equation:
        if not calc_stack:
            calc_stack.append(a)
        elif a == '+' or a == '*':
            calc_stack.append(a)
        elif a == '(':
            calc_stack.append(a)
        elif a.isdigit():
            if calc_stack[-1] == '+': # n + n 
                calc_stack.pop() # remove operator
                b = calc_stack.pop()
                calc_stack.append(str(int(b)+int(a)))
            elif calc_stack[-1] == '*': # n + n 
                calc_stack.pop() # remove operator
                b = calc_stack.pop()
                calc_stack.append(str(int(b)*int(a)))
            else:
                calc_stack.append(a)
        elif a == ')':
            # pop back everything into a list until we get the first matching (
            #   pass it back into calc_stack and carry on
            temp = []
            c = calc_stack.pop()
            while c != '(':
                temp.append(c)
                c = calc_stack.pop()
            calc_stack.append(temp[0])

    return int(calc_stack.pop())

def part1(data):
    reg = re.compile(r'\d+|[\(\)\*\+]')
    sum = 0
    for line in data:
        chars = reg.findall(line)
        answer = do_calc(chars)
        print("{} = {}".format(line, answer))
        sum += answer
    print(sum)


def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(TEST) as file:
        data = file.read().splitlines()

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
