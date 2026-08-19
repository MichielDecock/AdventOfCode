import os
from dataclasses import dataclass

@dataclass
class Data:
    password: str
    character: str
    pos1: int
    pos2: int

def getInput(filename: str):
    lines = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        for line in file:
            [l1, password] = [l.strip() for l in line.split(':')]
            [minmax, char ] =[s.strip() for s in l1.split(' ')]
            [pos1, pos2] = [int(el) for el in minmax.split('-')]
            lines.append(Data(password, char, pos1, pos2))

    return lines

def isValid(data) -> bool:
    count = sum(1 for el in [data.password[data.pos1 - 1], data.password[data.pos2 - 1]] if el == data.character)
    return count == 1


if __name__ == "__main__":
    data = getInput('input')
    total = sum(1 for d in data if isValid(d))
    print(total)
