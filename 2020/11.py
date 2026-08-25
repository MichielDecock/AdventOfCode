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

def isDiagonal(refRow, refCol, row, col):
    return abs(refRow - row) == abs(refCol - col)

def addNeighbor(neighbors, candidates):
    if len(candidates) > 0:
        neighbors.append(candidates[0])
    
def getNeighbors(seats, rows, cols):
    for seat in seats:
        neighbors = []
        addNeighbor(neighbors, sorted([s for s in seats if seat.row > s.row and seat.col == s.col], key=lambda x: x.row, reverse= True)) #top
        addNeighbor(neighbors, sorted([s for s in seats if seat.row < s.row and seat.col == s.col], key=lambda x: x.row)) #bottom
        addNeighbor(neighbors, sorted([s for s in seats if seat.row == s.row and seat.col > s.col], key=lambda x: x.col, reverse= True)) #left
        addNeighbor(neighbors, sorted([s for s in seats if seat.row == s.row and seat.col < s.col], key=lambda x: x.col)) #right
        addNeighbor(neighbors, sorted([s for s in seats if isDiagonal(seat.row, seat.col, s.row, s.col) and s.row < seat.row and s.col < seat.col], key=lambda x: x.row, reverse=True)) #topLeft
        addNeighbor(neighbors, sorted([s for s in seats if isDiagonal(seat.row, seat.col, s.row, s.col) and s.row > seat.row and s.col < seat.col], key=lambda x: x.row)) #bottomLeft
        addNeighbor(neighbors, sorted([s for s in seats if isDiagonal(seat.row, seat.col, s.row, s.col) and s.row < seat.row and s.col > seat.col], key=lambda x: x.row, reverse= True)) #topRight
        addNeighbor(neighbors, sorted([s for s in seats if isDiagonal(seat.row, seat.col, s.row, s.col) and s.row > seat.row and s.col > seat.col], key=lambda x: x.row)) #bottomRight
        seat.neighbors = neighbors

def move(seats):
    mods = []

    for idx, seat in enumerate(seats):
        if not seat.occ and all(not n.occ for n in seat.neighbors):
            mods.append(idx)
        elif seat.occ and sum(1 for n in seat.neighbors if n.occ) >= 5:
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
