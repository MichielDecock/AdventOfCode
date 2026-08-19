import os

_TARGET = 2020

def getNumbers(filename: str) -> list[int]:
    lines = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        for line in file:
            lines.append(int(line))
    return lines

def findPair(numbers: list[int]) -> list[int] | None:
    for x in numbers:
        for y in numbers:
            if x == y:
                continue

            if x + y == _TARGET:
                return [x, y]
    return []


if __name__ == "__main__":
    numbers = getNumbers('input')
    if pair := findPair(numbers):
        print(pair[0] * pair[1])
