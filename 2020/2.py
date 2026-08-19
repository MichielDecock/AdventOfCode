import os
from dataclasses import dataclass

@dataclass
class Data:
    password: str
    character: str
    minValue: int
    maxValue: int

def getInput(filename: str):
    lines = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        for line in file:
            [l1, password] = [l.strip() for l in line.split(':')]
            [minmax, char ] =[s.strip() for s in l1.split(' ')]
            [minValue, maxValue] = [int(el) for el in minmax.split('-')]
            lines.append(Data(password, char, minValue, maxValue))

    return lines

def isValid(data) -> bool:
    count = data.password.count(data.character)
    return data.minValue <= count and count <= data.maxValue


if __name__ == "__main__":
    data = getInput('input')
    total = sum(1 for d in data if isValid(d))
    print(total)
