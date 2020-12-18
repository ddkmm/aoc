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

def do_calc1(equation):
    calc_stack = []
    for a in equation:
        if not calc_stack:
            calc_stack.append(a)
        elif a == '+' or a == '*':
            calc_stack.append(a)
        elif a == '(':
            calc_stack.append(a)
        elif a.isdigit():
            if len(calc_stack) > 0:
                op = calc_stack[-1]
                if op == '+' or op == '*':
                    calc_stack.append(a)
                    calc_stack = do_math(calc_stack)
                else:
                    calc_stack.append(a)
        elif a == ')':
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

def add(stack):
    a = stack.pop()
    op = stack.pop()
    b = stack.pop()
    assert(op == '+')
    stack.append(str(int(b)+int(a)))
    return stack

def multiply(stack):
    a = stack.pop()
    op = stack.pop()
    b = stack.pop()
    assert(op == '*')
    stack.append(str(int(b)*int(a)))
    return stack

def do_calc2(equation):
    calc_stack = []
    for a in equation:
        if not calc_stack:
            calc_stack.append(a)
        elif a == '+' or a == '*':
            calc_stack.append(a)
        elif a == '(':
            calc_stack.append(a)
        elif a.isdigit():
            if len(calc_stack) > 0:
                op = calc_stack[-1]
                if op == '+':
                    calc_stack.append(a)
                    calc_stack = add(calc_stack)
                # this is different
                elif op == '*':
                    calc_stack.append(a)
                    if '*' in calc_stack:
                        continue
                    else:
                        calc_stack = multiply(calc_stack)
                else:
                    calc_stack.append(a)
        elif a == ')':
            temp_stack = []
            temp = 0
            c = calc_stack.pop()
            while c != '(':
                temp_stack.append(c)
                c = calc_stack.pop()
            while len(temp_stack) > 1:
                temp_stack = multiply(temp_stack)
            temp = temp_stack.pop()
            if len(calc_stack) > 1:
                op = calc_stack[-1]
                if op == '+':
                    calc_stack.append(temp)
                    calc_stack = add(calc_stack)
                elif op == '*':
                    calc_stack.append(temp)
                    calc_stack = multiply(calc_stack)
                else:
                    calc_stack.append(temp)
            else:
                calc_stack.append(temp)

    while len(calc_stack) > 1:
        calc_stack = multiply(calc_stack)

    return int(calc_stack.pop())

def part1(data):
    reg = re.compile(r'\d+|[\(\)\*\+]')
    sum = 0
    for line in data:
        chars = reg.findall(line)
        answer = do_calc1(chars)
        if DEBUG:
            print("{} = {}".format(line, answer))
        sum += answer
    print("Part 1: {}".format(sum))

def part2(data):
    reg = re.compile(r'\d+|[\(\)\*\+]')
    sum = 0
    for line in data:
        chars = reg.findall(line)
        answer = do_calc2(chars)
        if DEBUG:
            print("{} = {}".format(line, answer))
        sum += answer
    print("Part 2: {}".format(sum))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(TEST) as file:
        data = file.read().splitlines()

#    time1 = time.perf_counter()
#    part1(data)
#    time2 = time.perf_counter()
#    print("{} seconds".format(time2-time1))

    time1 = time.perf_counter()
    part2(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
