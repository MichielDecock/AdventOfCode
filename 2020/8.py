import os
from dataclasses import dataclass
from enum import Enum
import copy

class Operation(Enum):
    ACC = 1,
    JMP = 2,
    NOP = 3

@dataclass
class Instruction:
    operation: Operation
    count: int
    processed = False

def toOperation(name):
    if name == 'acc':
        return Operation.ACC
    if name == 'jmp':
        return Operation.JMP

    return Operation.NOP

def parse(filename):
    lines = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        for line in [line.strip('\n') for line in file]:
            operation, count = line.split(' ')
            lines.append(Instruction(toOperation(operation), int(count)))

    return lines

def reachEnd(lines, pos, acc):
    numLines = len(lines)
    while True:
        if pos >= numLines:
            return True, acc

        instruction = lines[pos]
        if instruction.processed:
            return False

        instruction.processed = True

        if instruction.operation == Operation.JMP:
            pos += instruction.count
            continue

        if instruction.operation == Operation.ACC:
            acc += instruction.count

        pos += 1


if __name__ == "__main__":
    lines = parse('input')

    acc = 0
    pos = 0
    while True:
        instruction = lines[pos]

        if instruction.operation == Operation.ACC:
            instruction.processed = True
            acc += instruction.count
            pos += 1
            continue

        if out := reachEnd(copy.deepcopy(lines), pos, acc):
            print(out)
            break

        copyLines = copy.deepcopy(lines)
        if copyLines[pos].operation == Operation.JMP:
            copyLines[pos].operation = Operation.NOP
        elif copyLines[pos].operation == Operation.NOP:
            copyLines[pos].operation = Operation.JMP

        if out := reachEnd(copyLines, pos, acc):
            print(out)
            break

        if lines[pos].operation == Operation.JMP:
            pos += copyLines[pos].count
        elif lines[pos].operation == Operation.NOP:
            pos += 1

        instruction.processed = True
            
