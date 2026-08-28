import os
from dataclasses import dataclass

_TARGET = 2020

@dataclass
class Number():
    turn: int
    count: int = 1

    def _update(self, turn):
        self.turn = turn
        self.count += 1

def parse(filename):
    lines = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        for line in [line.strip('\n') for line in file]:
            lines.extend([int(x) for x in line.split(',')])

    return lines

def initDict(numbers):
    d = {}
    for i, n in enumerate(numbers):
        d[n] = Number(i + 1)
    return d

if __name__ == "__main__":
    numbers = parse('input')
    d = initDict(numbers[:-1])
    lenN = len(numbers)

    number = numbers[lenN - 1]
    for turn in range(lenN, _TARGET + 1):
        if number in d:
            diff = turn - d[number].turn
            d[number]._update(turn)
            number = diff
        else:
            d[number] = Number(turn)
            number = 0

    print(next(key for key, value in d.items() if value.turn == _TARGET))
