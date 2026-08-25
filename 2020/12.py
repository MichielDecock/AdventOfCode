import os
from dataclasses import dataclass
from enum import Enum
import math

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

@dataclass
class WayPoint:
    x: int = 10
    y: int = 1


def parse(filename):
    lines = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        for line in [line.strip('\n') for line in file]:
            lines.append((line[0], int(line[1:])))
    return lines

def move(dir, value, obj):
    if dir == Dir.EAST:
        obj.x += value
    elif dir == Dir.SOUTH:
        obj.y -= value
    elif dir == Dir.WEST:
        obj.x -= value
    elif dir == Dir.NORTH:
        obj.y += value

def rotate(instruction, boat, wayPoint):
    i, value = instruction
    d = [wayPoint.x, wayPoint.y]
    angle = value * math.pi / 180 * (1 if i == 'L' else -1)
    rot = [[math.cos(angle), -math.sin(angle)], [math.sin(angle), math.cos(angle)]]
    wayPoint.x = round(sum(rot[0][i] * d[i] for i in range(2)))
    wayPoint.y = round(sum(rot[1][i] * d[i] for i in range(2)))
        
def sail(instruction, boat, wayPoint):
    i, value = instruction
    
    if i == 'L' or i == 'R':
        rotate(instruction, boat, wayPoint)
    elif i == 'F':
        boat.x += value * wayPoint.x
        boat.y += value * wayPoint.y
    else:
        move(_DIR_MAP[i], value, wayPoint)

if __name__ == "__main__":
    instructions = parse('input')

    boat = Boat()
    wayPoint = WayPoint()
    for i in instructions:
        sail(i, boat, wayPoint)

    print((abs(boat.x)) + int(abs(boat.y)))
    
