import os
from dataclasses import dataclass

_TREE = '#'

@dataclass
class Position:
    row: int
    col: int

    def slide(self, maxCols, slope):
        self.row += slope.row
        self.col =  (self.col + slope.col) % maxCols

_SLOPES = [Position(1,1), Position(1,3), Position(1,5), Position(1,7), Position(2,1), ]

def parse(filename: str):
    lines = []
    maxRows = 0
    maxCols = 0
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        for row, line in enumerate([line.strip('\n') for line in file]):
            maxRows += 1
            maxCols = len(line)
            for col, el in enumerate(line):
                if el != _TREE:
                    continue
                lines.append(Position(row,col))

    return [lines, Position(maxRows, maxCols)]

def hits(pos, trees, limit, slope):
    total = 0
    while pos.row < limit.row:
        pos.slide(limit.col, slope)
        if pos in trees:
            total += 1

    return total


if __name__ == "__main__":
    [trees, limit] = parse('input')
    total = 1
    for slope in _SLOPES:
        total *= hits(Position(0,0), trees, limit, slope)
    print(total)

