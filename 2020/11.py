from __future__ import annotations
import os
from dataclasses import dataclass

@dataclass
class Seat:
    row: int
    col: int
    neighbors: list[Seat]
    occ: bool = False

def parse(filename):
    seats = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        lines = [line.strip('\n') for line in file]
        rows = len(lines)
        cols = len(lines[0])
        for row, line in enumerate(lines):
            for col, c in enumerate(line):
                if c == '.':
                    continue
                seats.append(Seat(row, col, []))
    return (seats, rows, cols)
    
def getNeighbors(seats, rows, cols):
    for seat in seats:
        neighbors = []
        for nRow in range(max(0, seat.row - 1), min(rows, seat.row + 2)):
            for nCol in range(max(0, seat.col - 1), min(cols, seat.col + 2)):
                if nRow == seat.row and nCol == seat.col:
                    continue

                match = next((seat for seat in seats if seat.row == nRow and seat.col == nCol), None)
                if match is not None:
                    neighbors.append(match)

        seat.neighbors = neighbors

def move(seats):
    mods = []

    for idx, seat in enumerate(seats):
        if not seat.occ and all(not n.occ for n in seat.neighbors):
            mods.append(idx)
        elif seat.occ and sum(1 for n in seat.neighbors if n.occ) >= 4:
            mods.append(idx)

    for m in mods:
        seats[m].occ = not seats[m].occ

    return len(mods) > 0

if __name__ == "__main__":
    seats, rows, cols = parse('input')
    getNeighbors(seats, rows, cols)

    while move(seats):
        pass

    print(sum(1 for s in seats if s.occ))
