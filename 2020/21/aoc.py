#!/usr/bin/env python

import os
import re
import time
from collections import defaultdict

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
DEBUG = False

def part1(ingredient_to_possible_allergens, recipes_with_allergen, recipes):
    # each allergen is contained in only one ingredient
    # so each ingredient has a set of potential allergens
    no_allergens = []

    for ingredient in ingredient_to_possible_allergens:
        if DEBUG:
            print("Ingredient {}".format(ingredient))
        possible_allergens = ingredient_to_possible_allergens[ingredient]
        impossible = set()
        
        for allergen in possible_allergens:
            if DEBUG:
                print("Checking for {}".format(allergen))
            temp = recipes_with_allergen[allergen]
            for index in temp:
                if ingredient not in recipes[index]:
                    impossible.add(allergen)
                    break
        
        # remove the impossible ingredient from the possible ones
        possible_allergens -= impossible
        if not possible_allergens:
            no_allergens.append(ingredient)
    print("Part 1: {} ingredients have no allergens".format(len(no_allergens)))



def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))
    recipes = []
    # ingredient to possible allergens
    ingredient_to_possible_allergens = defaultdict(lambda: set())

    # allergen to list of recipes with it
    recipes_with_allergen = defaultdict(lambda: list())

    with open(DATA) as file:
        for i, line in enumerate(file):
            match = re.split("contains", line)
            ingredients = set(re.findall('\w+', match[0]))
            recipes.append(ingredients)
            allergens = set(re.findall('\w+', match[1]))
            for a in allergens:
                recipes_with_allergen[a].append(i)
            for i in ingredients:
                ingredient_to_possible_allergens[i] |= allergens

    time1 = time.perf_counter()
    part1(ingredient_to_possible_allergens, recipes_with_allergen, recipes)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

main()
