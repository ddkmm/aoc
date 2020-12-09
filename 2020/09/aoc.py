#!/usr/bin/env python

import os

DIRPATH = os.path.dirname(os.path.realpath(__file__))
FILEPATH = os.path.join(DIRPATH, 'input.txt')
TESTPATH = os.path.join(DIRPATH, 'testinput.txt')
PREAMBLE = 25
#PREAMBLE = 5

def part1(code):
    print("Part 1")
    head = 0
    tail = head + PREAMBLE - 1
    target = tail + 1

    while target != len(code):
        left = head
        right = head + 1
        while left != len(code):
            compare_left = code[left]
            compare_right = code[right]
            if compare_left + compare_right != code[target]:
                if right < tail:
                    right += 1
                else:
                    left += 1
                    right = left + 1
                    if right > len(code)-1:
                        match = False
                        break
            else:
                match = True
                break
        if not match:
            print("{}".format(code[target]))
            return code[target]
        target += 1
        head = target - PREAMBLE
        tail = head + PREAMBLE - 1


def part2(code, target):
    print("Part 2")
    head = 0
    tail = head + PREAMBLE - 1

    start = 0
    while True:
        answers = []
        left = head
        right = head + 1
        compare_left = code[left]
        total = compare_left
        answers.append(compare_left)

        while left != len(code):
            compare_right = code[right]
            total += compare_right
            answers.append(compare_right)
            if total < target:
                if right < tail:
                    right += 1
                else:
                    left += 1
                    right = left + 1
                    if right > len(code)-1:
                        match = False
                        total = 0
                        answers.clear()
                        break
            elif total == target:
                match = True
                break
            else:
                match = False
                total = 0
                answers.clear()
                break

        if match:
            print("{} + {} = {}".format(min(answers), max(answers), min(answers) + max(answers)))
            break

        start += 1
        head = start
        tail = head + PREAMBLE - 1

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

#    with open(TESTPATH) as file:
    with open(FILEPATH) as file:
        code = file.read().splitlines()
    code = [int(i) for i in code]

    part2(code, part1(code))

main()
