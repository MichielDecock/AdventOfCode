import os
from dataclasses import dataclass


@dataclass
class Seat:
    row: int
    col: int


def parse(filename: str, selectRow: str = ''):
    table = str.maketrans({
        'F': '0',
        'B': '1',
        'L': '0',
        'R': '1'
    })

    lines = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        for line in [line.strip('\n') for line in file if not selectRow or line[:7] == selectRow]:
            line = line.translate(table)
            lines.append(line)

    return lines

def row(line):
    return int(line[:7], 2)

def column(line):
    return int(line[-3:], 2)

def rows(lines):
    rows = sorted([row(line) for line in lines])
    return [el for el in rows if el != min(rows) and el != max(rows)]

def columns(lines):
    return sorted([column(line) for line in lines])


def findRow(rows):
    for row in rows:
        if rows.count(row) != 8:
            return row
    return None

def rowToString(row):
    table = str.maketrans({
            '0': 'F',
            '1': 'B'
        })

    return str(format(row, 'b').translate(table))

def seatID(seat):
    return seat.row * 8 + seat.col

if __name__ == "__main__":
    lines = parse('input')
    candidateRow = findRow(rows(lines))
    candidates = columns(parse('input', rowToString(candidateRow)))
    for col in range(8):
        if col not in candidates:
            assert(candidateRow)
            print(seatID(Seat(candidateRow, col)))
