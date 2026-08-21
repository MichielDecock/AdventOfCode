import os
from dataclasses import dataclass


@dataclass
class Seat:
    row: int
    col: int


def parse(filename: str):
    table = str.maketrans({
        'F': '0',
        'B': '1',
        'L': '0',
        'R': '1'
    })

    lines = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        for line in [line.strip('\n') for line in file]:
            line = line.translate(table)
            lines.append(line)

    return lines

def findSeat(line):
    return Seat(int(line[:7], 2),int(line[-3:], 2))


def seatID(seat):
    return seat.row * 8 + seat.col

if __name__ == "__main__":
    lines = parse('input')
    print(max([seatID(findSeat(line)) for line in lines]))
