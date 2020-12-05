#!/usr/bin/env python

import re
import os

dir_path = os.path.dirname(os.path.realpath(__file__))
FILEPATH = os.path.join(dir_path, 'input.txt')
TESTPATH = os.path.join(dir_path, 'testinput.txt')

valid_entry = [ "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid" ]

# byr (Birth Year) - four digits; at least 1920 and at most 2002.
def validate_byr(entry):
    isValid = 1

    if not str(entry).isdigit:
        isValid = 0
    elif len(entry) != 4:
        isValid = 0
    elif int(entry) < 1920:
        isValid = 0
    elif int(entry) > 2002:
        isValid = 0
    
    return isValid

def test_validate_byr():
    assert(not validate_byr('123'))
    assert(not validate_byr('12345'))
    assert(not validate_byr('1919'))
    assert(validate_byr('1920'))
    assert(not validate_byr('2003'))
    assert(validate_byr('2002'))

# iyr (Issue Year) - four digits; at least 2010 and at most 2020.
def validate_iyr(entry):
    isValid = 1

    if not str(entry).isdigit:
        isValid = 0
    elif len(entry) != 4:
        isValid = 0
    elif int(entry) < 2010:
        isValid = 0
    elif int(entry) > 2020:
        isValid = 0

    return isValid

def test_validate_iyr():
    assert(not validate_iyr('123'))
    assert(not validate_iyr('12345'))
    assert(not validate_iyr('2009'))
    assert(validate_iyr('2010'))
    assert(not validate_iyr('2021'))
    assert(validate_iyr('2020'))

# eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
def validate_eyr(entry):
    isValid = 1

    if not str(entry).isdigit:
        isValid = 0
    elif len(entry) != 4:
        isValid = 0
    elif int(entry) < 2020:
        isValid = 0
    elif int(entry) > 2030:
        isValid = 0

    return isValid

def test_validate_eyr():
    assert(not validate_eyr('123'))
    assert(not validate_eyr('12345'))
    assert(not validate_eyr('2019'))
    assert(validate_eyr('2020'))
    assert(not validate_eyr('2031'))
    assert(validate_eyr('2030'))

# hgt (Height) - a number followed by either cm or in:
# If cm, the number must be at least 150 and at most 193.
# If in, the number must be at least 59 and at most 76.
def validate_hgt(entry):
    isValid = 1

    out = re.split('(\d+)', entry)
    if len(out) != 3:
        isValid = 0
    if re.search('cm', out[2]):
        if int(out[1]) < 150:
            isValid = 0
        elif int(out[1]) > 193: 
            isValid = 0
    elif re.search('in', out[2]):
        if int(out[1]) < 59:
            isValid = 0
        elif int(out[1]) > 76: 
            isValid = 0
    else:
        isValid = 0

    return isValid

def test_validate_hgt():
    assert(validate_hgt('150cm'))
    assert(validate_hgt('193cm'))
    assert(validate_hgt('59in'))
    assert(validate_hgt('76in'))
    assert(not validate_hgt('77in'))
    assert(not validate_hgt('58in'))
    assert(not validate_hgt('149cm'))
    assert(not validate_hgt('194cm'))
    assert(not validate_hgt('150'))
    assert(not validate_hgt('59'))
    assert(not validate_hgt('150xx'))

# hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
def validate_hcl(entry):
    isValid = 1

    if len(entry) != 7:
        isValid = 0

    if entry[0] != '#':
        isValid = 0
    
    for a in entry[1:]:
        if not re.search("[0-9,a-f]",a):
            isValid = 0
    
    return isValid

def test_validate_hcl():
    assert(validate_hcl("#ffffff"))
    assert(validate_hcl("#123456"))
    assert(validate_hcl("#1f3f5f"))
    assert(not validate_hcl("#fffffff"))
    assert(not validate_hcl("ffffff"))
    assert(not validate_hcl("#fffff"))
    assert(not validate_hcl("#fxffff"))

# ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
def validate_ecl(entry):
    isValid = 1

    hair = ("amb", "blu", "brn", "gry", "grn", "hzl", "oth")
    if len(entry) != 3:
        isValid = 0

    if entry not in hair:
        isValid = 0

    return isValid

def test_validate_ecl():
    assert(validate_ecl('amb'))
    assert(validate_ecl('blu'))
    assert(validate_ecl('brn'))
    assert(validate_ecl('gry'))
    assert(validate_ecl('grn'))
    assert(validate_ecl('hzl'))
    assert(validate_ecl('oth'))
    assert(not validate_ecl('othoth'))
    assert(not validate_ecl(''))
    assert(not validate_ecl(' '))

# pid (Passport ID) - a nine-digit number, including leading zeroes.
def validate_pid(entry):
    isValid = 1
    if len(entry) != 9:
        isValid = 0

    for a in entry:
        if not re.search("[0-9]",a):
            isValid = 0

    return isValid

def test_validate_pid():
    too_long = "1234567890"
    too_short = "13345678"
    non_number = 'a12345678'

    assert(not validate_pid(too_long))
    assert(not validate_pid(too_short))
    assert(not validate_pid(non_number))

def test_validate():
    test_validate_byr()
    test_validate_iyr()
    test_validate_eyr()
    test_validate_hgt()
    test_validate_hcl()
    test_validate_ecl()
    test_validate_pid()


def validate(entry, value, debug=False):
    isValid = 1

    if entry in "byr":
        if not validate_byr(value):
            if debug: print ("Failed byr")
            isValid = 0
    elif entry in "iyr":
        if not validate_iyr(value):
            if debug: print ("Failed iyr")
            isValid = 0
    elif entry in "eyr":
        if not validate_eyr(value):
            if debug: print ("Failed eyr")
            isValid = 0
    elif entry in "hgt":
        if not validate_hgt(value):
            if debug: print ("Failed hgt")
            isValid = 0
    elif entry in "hcl":
        if not validate_hcl(value):
            if debug: print ("Failed hcl")
            isValid = 0
    elif entry in "ecl":
        if not validate_ecl(value):
            if debug: print ("Failed ecl")
            isValid = 0
    elif entry in "pid":
        if not validate_pid(value):
            if debug: print ("Failed pid")
            isValid = 0

    return isValid

def process_passport(current_entry):
    isValid = 1
    if not (len(current_entry) == 7 or len(current_entry) == 8):
        return 0
    for a in valid_entry:
        try:
            current_entry[a]
        except KeyError:
            return 0

    return isValid

def process_and_validate_passport(current_entry):
    if not (len(current_entry) == 7 or len(current_entry) == 8):
        return 0
    for a in valid_entry:
        try:
            value = current_entry[a]
            if not validate(a, value):
                return 0
        except KeyError:
            return 0

    return 1
    
def main():
    print("Day {}".format(os.path.split(dir_path)[1]))

    current_entry = dict()
    sample = []
    total = 0
    validated_total = 0
    test_validate()

    with open(FILEPATH) as passports:
        while True:
            line = passports.readline()
            if line == "\n":
                if process_passport(current_entry):
                    total += 1
                    if process_and_validate_passport(current_entry):
                        validated_total += 1
                        #print (current_entry)
                current_entry.clear()
                sample.clear()
                continue
            if not line:
                if process_passport(current_entry):
                    total += 1
                    if process_and_validate_passport(current_entry):
                        validated_total += 1
                        #print (current_entry)
                break
            record = line.split()
            for entry in record:
                sample.append(tuple(entry.split(":")))
            current_entry = dict(sample)
    print ("Total: {}".format(total))
    print ("Validated Total: {}".format(validated_total))

main()