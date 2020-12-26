import os
import time
import math

DIRPATH = os.path.dirname(os.path.realpath(__file__))
DATA = os.path.join(DIRPATH, 'input.txt')
TEST = os.path.join(DIRPATH, 'test.txt')
DEBUG = False
ACTIVE = '#'
INACTIVE = '.'



class Cube:
    def __init__(self, x, y, z, w):
        self.x = x 
        self.y = y
        self.z = z
        self.w = w

    def __eq__(self, other):
        if isinstance(other, Cube):
            res = self.x == other.x and self.y == other.y and self.z == other.z and self.w == other.w

            return res 

    def __str__(self):
        return "({}, {}, {}, {})".format(self.x, self.y, self.z, self.w)

    def __hash__(self):
        return hash(str(self))

    def get_pos(self):
        return self.x, self.y, self.z, self.w

# RULE 1
# If a cube is active and
#   exactly 2 or 3 of its neighbors are also active,
#   the cube remains active.
 # Otherwise, the cube becomes inactive.
def stay_active(world, cube):
    x0, y0, z0, w0 = cube.get_pos() 
    spread = (-1, 0, 1)
    count = 0
    for w in spread:
        for z in spread:
            for y in spread:
                for x in spread:
                    temp_cube = Cube(x0 + x, y0 + y, z0 + z, w0 + w)
                    if DEBUG:
                        print("neighbour {}".format(str(temp_cube)))
                    if temp_cube == cube:
                        if DEBUG:
                            print("Self, skip")
                        continue
                    elif temp_cube in world:
                        if DEBUG:
                            print("Active")
                        count += 1
                    else:
                        if DEBUG:
                            print("Inactive")

    if 1 < count < 4:
        if DEBUG:
            print("{}: {} active neighbours.".format(str(cube), count))
            print("Remains active")
        # remain active
        return True
    # become inactive
    if DEBUG:
        print("{}: {} active neighbours. {} becomes inactive.".format(str(cube), count, str(cube)))
    return False

# RULE 2
# If a cube is inactive but
#   exactly 3 of its neighbors are active,
#   the cube becomes active.
# Otherwise, the cube remains inactive.
def stay_inactive(world, cube):
    x0, y0, z0, w0 = cube.get_pos() 
    spread = (-1, 0, 1)
    count = 0
    for w in spread:
        for z in spread:
            for y in spread:
                for x in spread:
                    temp_cube = Cube(x0 + x, y0 + y, z0 + z, w0 + w)
                    if DEBUG:
                        print("neighbour {}".format(str(temp_cube)))
                    if temp_cube == cube:
                        if DEBUG:
                            print("Self")
                        continue
                    elif temp_cube in world:
                        if DEBUG:
                            print("Active")
                        count += 1
                    else:
                        if DEBUG:
                            print("Inactive")
    if count == 3:
        if DEBUG:
            print("{}: {} active neighbours. {} becomes active.".format(str(cube), count, str(cube)))
        # become active
        return False

    if DEBUG:
        print("{}: {} active neighbours.".format(str(cube), count))
        print("Stays inactive")
    # remain inactive
    return True

def traverse_world(world, x_bound, y_bound, z_bound):
    new_world = world.copy() 

    for z in range(z_bound[0], z_bound[1]):
#        print("z = {}".format(z))
        for y in range(y_bound[0], y_bound[1]):
            for x in range(x_bound[0], y_bound[1]):
                cube = Cube(x, y, z)
                if cube in world:
                    if DEBUG:
                        print("Checking active {}".format(str(cube)))
                    if not stay_active(world, cube):
                        new_world.remove(cube)
                else:
                    if DEBUG:
                        print("Checking inactive {}".format(str(cube)))
                    if not stay_inactive(world, cube):
                        new_world.add(cube)
#        print(len(new_world))
    return new_world
    
def traverse_world2(world, x_bound, y_bound, z_bound, w_bound):
    new_world = world.copy() 

    for w in range(w_bound[0], w_bound[1]):
        for z in range(z_bound[0], z_bound[1]):
            for y in range(y_bound[0], y_bound[1]):
                for x in range(x_bound[0], y_bound[1]):
                    cube = Cube(x, y, z, w)
                    if cube in world:
                        if DEBUG:
                            print("Checking active {}".format(str(cube)))
                        if not stay_active(world, cube):
                            new_world.remove(cube)
                    else:
                        if DEBUG:
                            print("Checking inactive {}".format(str(cube)))
                        if not stay_inactive(world, cube):
                            new_world.add(cube)
    return new_world

def part1(data):
    world = set()
    if DEBUG:
        print(data)

    x_size = math.floor(len(data[0])/2)
    y_size = math.floor(len(data)/2)

    # Create world
    for y_val, line in enumerate(reversed(data),-1*y_size):
        for x_val, cube in enumerate(line, -1*x_size):
            if cube == ACTIVE:
                world.add(Cube(x_val, y_val, 0))
    print(len(world))

    for cycle in range(1,7):
        x_bound = ((-1*x_size) - cycle, x_size + 1 + cycle)
        y_bound = ((-1*y_size) - cycle, y_size + 1 + cycle)
        z_bound = (0 - cycle, 1 + cycle)
        world = traverse_world(world, x_bound, x_bound, z_bound)
        print("Cycle: {}, {} active".format(cycle, len(world)))

    print("Part 1: {}".format(len(world)))

def part2(data):
    world = set()
    if DEBUG:
        print(data)

    x_size = math.floor(len(data[0])/2)
    y_size = math.floor(len(data)/2)

    # Create world
    for y_val, line in enumerate(reversed(data),-1*y_size):
        for x_val, cube in enumerate(line, -1*x_size):
            if cube == ACTIVE:
                world.add(Cube(x_val, y_val, 0, 0))
    print(len(world))

    for cycle in range(1,7):
        x_bound = ((-1*x_size) - cycle, x_size + 1 + cycle)
        z_bound = (0 - cycle, 1 + cycle)
        w_bound = (0 - cycle, 1 + cycle)
        world = traverse_world2(world, x_bound, x_bound, z_bound, w_bound)
        print("Cycle: {}, {} active".format(cycle, len(world)))

    print("Part 2: {}".format(len(world)))

def main():
    print("Day {}".format(os.path.split(DIRPATH)[1]))

    with open(DATA) as file:
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
