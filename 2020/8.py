import os
from dataclasses import dataclass
from enum import Enum

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


if __name__ == "__main__":
    lines = parse('input')

    acc = 0
    pos = 0
    while True:
        instruction = lines[pos]
        if instruction.processed:
            print(acc)
            break

        instruction.processed = True

        if instruction.operation == Operation.JMP:
            pos += instruction.count
            continue

        if instruction.operation == Operation.ACC:
            acc += instruction.count

        pos += 1
            
