#!/usr/bin/env python

import os
import re
import time

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'testinput.txt')
TEST2 = os.path.join(DIRPATH, 'test2.txt')
DEBUG = False

def check_rule(number, rule):
    digits = re.split(r'\W', "-".join(rule))
    if ((int(digits[0]) <= number <= int(digits[1])) or
        (int(digits[2]) <= number <= int(digits[3]))):
        return True
    return False

def process_input(data):
    rule_list = []
    data_start = 0
    for i, line in enumerate(data):
        temp = re.split(r':', line.strip())
        if len(temp) == 2:
            if re.search(r'^your', temp[0]):
                continue
            elif re.search(r'^nearby', temp[0]):
                data_start = i+1
            else:
                # this is a rule
                fields = re.findall(r'\d+-\d+', temp[1].strip())
                rule_list.append(fields)
    return data_start, rule_list

def find_bad_tickets(data, data_start, rule_list):
    total = 0
    tickets = []
    for line in range(data_start, len(data)):
        badTicket = False
        for number in re.split(r',',data[line]):
            number = int(number)
            res = False
            for rule in rule_list:
                res = (res or check_rule(number, rule))
            if not res:
                total += number
                badTicket = True
        if not badTicket:
            tickets.append(data[line])
    return total

def process_input2(data):
    rule_list = []
    data_start = 0
    your_ticket = 0
    for i, line in enumerate(data):
        temp = re.split(r':', line.strip())
        if len(temp) == 2:
            if re.search(r'^your', temp[0]):
                your_ticket = i+1
                continue
            elif re.search(r'^nearby', temp[0]):
                data_start = i+1
            else:
                # this is a rule
                fields = re.findall(r'\d+-\d+', temp[1].strip())
                rule_list.append(fields)
    return rule_list, your_ticket, data_start

def filter_bad_tickets(data, data_start, your_ticket, rule_list):
    tickets = []
    mine = list(map(int, re.split(r',', data[your_ticket])))
    tickets.append(mine)
    bad = 0

    for line in range(data_start, len(data)):
        badTicket = False
        for number in re.split(r',',data[line]):
            number = int(number)
            res = False
            for rule in rule_list:
                res = (res or check_rule(number, rule))
            if not res:
                badTicket = True
        if not badTicket:
            tickets.append(list(map(int, re.split(r',', data[line]))))
        else:
            bad += 1

    print("Bad tickets: {}, good tickets: {}".format(bad, len(tickets)))
    return tickets, mine

def part1(data):
    data_start, rule_list = process_input(data)
    total = find_bad_tickets(data, data_start, rule_list) 
    print("Part 1: {}".format(total))

def part2(data):
    rule_list, your_ticket, data_start = process_input2(data)
    tickets, mine = filter_bad_tickets(data, data_start, your_ticket, rule_list) 
    ticket = range(0, len(tickets))
    fields = range(0, len(tickets[0]))
    assert(len(rule_list) == len(fields))
    # for each rule, eliminate any fields which contain invalid data
    rule_table = []
    for index, rule in enumerate(rule_list):
        column = []
        for field in fields:
            if DEBUG:
                print("Checking rule {} ({}) against field {}".format(index + 1, rule, field))
            positives = 0
            for ticket in tickets:
                value = ticket[field]
                if DEBUG:
                    print("Ticket {}: value {}".format(ticket, value))
                if check_rule(value, rule):
                    positives += 1
            if DEBUG:
                print("Rule {} matches {}/{} tickets for field {}".format(index+1, positives, len(tickets), field))
            column.append(positives)
        if DEBUG:
            print("\nRule {}: {}\n".format(index+1, column))
        rule_table.append(column)

    count = [0] * len(rule_list)
    for a in rule_table:
        for i, v in enumerate(a):
            if v == len(tickets):
                count[i] += 1
    print("Possible fields: {}".format(count))
    answers = []
    for i, _ in enumerate(count):
        field = count.index(i+1)
        for rule_index, a in enumerate(rule_table):
            if a[field] == len(tickets):
                if (rule_index < 6):
                    print("Rule {} is field position {}".format(rule_index, field))
                    answers.append(field)
                rule_table[rule_index] = [0]*len(rule_list)

    prod = 1
    for a in answers:
        prod *= mine[a]
    print(prod)

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(DATA) as file:
        data = file.read().splitlines()

    time1 = time.perf_counter()
    part1(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

    time1 = time.perf_counter()
    part2(data)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
