#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
DEBUG = True

def do_math(stack):
    a = stack.pop()
    op = stack.pop()
    b = stack.pop()
    if op == '+':
        stack.append(str(int(b)+int(a)))
    elif op == '*':
        stack.append(str(int(b)*int(a)))
    return stack

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
            if len(calc_stack) > 1:
                op = calc_stack[-1]
                if op == '+' or op == '*':
                    calc_stack.append(a)
                    calc_stack = do_math(calc_stack)
                else:
                    calc_stack.append(a)
        elif a == ')':
            # pop back everything into a list until we get the first matching (
            #   pass it back into calc_stack and carry on
            temp = 0
            c = calc_stack.pop()
            while c != '(':
                temp = c
                c = calc_stack.pop()
            if len(calc_stack) > 1:
                op = calc_stack[-1]
                if op == '+' or op == '*':
                    calc_stack.append(temp)
                    calc_stack = do_math(calc_stack)
                else:
                    calc_stack.append(temp)
            else:
                calc_stack.append(temp)

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

    with open(DATA) as file:
        data = file.read().splitlines()

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
