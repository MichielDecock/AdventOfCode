import os
from dataclasses import dataclass

_TREE = '#'
_INCREASE_ROW = 1
_INCREASE_COL = 3

@dataclass
class Position:
    row: int
    col: int

    def slide(self, maxCols):
        self.row += _INCREASE_ROW
        self.col =  (self.col + _INCREASE_COL) % maxCols

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

def hits(pos, trees, limit):
    total = 0
    while pos.row < limit.row:
        pos.slide(limit.col)
        if pos in trees:
            total += 1

    return total


if __name__ == "__main__":
    [trees, limit] = parse('input')
    print(hits(Position(0,0), trees, limit))

