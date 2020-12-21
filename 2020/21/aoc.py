#!/usr/bin/env python

import os
import re
import time
from collections import defaultdict

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
DEBUG = False

def part1(ingredient_to_possible_allergens, recipes_with_allergen, recipes, ingredient_inventory):
    # each allergen is contained in only one ingredient
    # so each ingredient has a set of potential allergens
    allergen_free = []

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
            allergen_free.append(ingredient)

    # take the list of allergen_free ingredients and find occurrences in recipes
    total = 0
    for i in allergen_free:
        total += ingredient_inventory[i]
    print("{} allergen free ingredients".format(len(allergen_free)))
    print("Part 1: {} allergen free ingredients used".format(total))
    return allergen_free

def part2(ingredient_to_possible_allergens, recipes_with_allergen, recipes, allergen_free):
    print("{}".format(len(ingredient_to_possible_allergens)))
    for ingredient in allergen_free:
        del ingredient_to_possible_allergens[ingredient]
    print("{}".format(len(ingredient_to_possible_allergens)))
    for ing in ingredient_to_possible_allergens:
        print("{} -> {}".format(ing, ingredient_to_possible_allergens[ing]))
    # Then take the output, do the sorting and pruning by hand and get the answer that way

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))
    recipes = []
    # ingredient to possible allergens
    ingredient_to_possible_allergens = defaultdict(lambda: set())

    # allergen to list of recipes with it
    recipes_with_allergen = defaultdict(lambda: list())

    all_allergens = set()

    ingredient_inventory = defaultdict(lambda: 0)

    with open(DATA) as file:
        for i, line in enumerate(file):
            match = re.split("contains", line)
            ingredients = set(re.findall('\w+', match[0]))
            recipes.append(ingredients)
            allergens = set(re.findall('\w+', match[1]))
            for a in allergens:
                all_allergens.add(a)
                recipes_with_allergen[a].append(i)
            for i in ingredients:
                ingredient_inventory[i] += 1
                ingredient_to_possible_allergens[i] |= allergens
    print("{} recipes, {} ingredients, {} allergens".format(len(recipes), len(ingredient_to_possible_allergens), len(all_allergens)))

    allergen_free = []
    time1 = time.perf_counter()
    allergen_free = part1(ingredient_to_possible_allergens, recipes_with_allergen, recipes, ingredient_inventory)
    time2 = time.perf_counter()
    print("{} seconds".format(time2-time1))

    part2(ingredient_to_possible_allergens, recipes_with_allergen, recipes, allergen_free)

main()
