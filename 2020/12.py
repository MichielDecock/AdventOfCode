import os
from dataclasses import dataclass
from enum import Enum

class Dir(Enum):
    EAST = 0
    SOUTH = 1
    WEST = 2
    NORTH = 3

_DIR_MAP = {
    'E': Dir.EAST,
    'S': Dir.SOUTH,
    'W': Dir.WEST,
    'N': Dir.NORTH,
}

@dataclass
class Boat:
    x: int = 0
    y: int = 0
    dir: Dir = Dir.EAST


def parse(filename):
    lines = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        for line in [line.strip('\n') for line in file]:
            lines.append((line[0], int(line[1:])))
    return lines

def move(dir, value, boat):
    if dir == Dir.EAST:
        boat.x += value
    elif dir == Dir.SOUTH:
        boat.y -= value
    elif dir == Dir.WEST:
        boat.x -= value
    elif dir == Dir.NORTH:
        boat.y += value

def sail(instruction, boat):
    i, value = instruction
    
    if i == 'L':
        boat.dir = Dir((boat.dir.value - value // 90) % 4)
    elif i == 'R':
        boat.dir = Dir((boat.dir.value + value // 90) % 4)
    elif i == 'F':
        move(boat.dir, value, boat)
    else:
        move(_DIR_MAP[i], value, boat)

    return boat

if __name__ == "__main__":
    instructions = parse('input')

    boat = Boat()
    for i in instructions:
        boat = sail(i, boat)

    print(abs(boat.x) + abs(boat.y))
    
