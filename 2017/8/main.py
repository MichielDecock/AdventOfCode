def parse(file):
    instructions = []

    with open(file) as f:
        for line in [l.strip() for l in f.readlines()]:
            operation, condition = [l.strip() for l in line.split('if')]
            instructions.append((operation, condition))
            
    return instructions

def execute(operation, registers):
    register, operator, number = operation.split(' ')
    match operator:
        case 'inc':
            registers[register] += int(number)
        case 'dec':
            registers[register] -= int(number)

def check(condition, registers):
    register, operator, number = condition.split(' ')

    a = registers[register]
    b = int(number)

    match operator:
        case '==':
            return a == b
        case '!=':
            return a != b
        case '<':
            return a < b
        case '<=':
            return a <= b
        case '>':
            return a > b
        case '>=':
            return a >= b
    
    print("Another operator!" , operator)

def initRegisters(instructions):
    registers = {}

    for operation, _ in instructions:
        register, _, _ = operation.split(' ')
        registers[register] = 0
    
    return registers

def apply(instruction, registers):
    operation, condition = instruction
    if check(condition, registers):
        execute(operation, registers)

if __name__ == '__main__':
    instructions = parse('input')
    registers = initRegisters(instructions)
    for instruction in instructions:
        apply(instruction, registers)
    
    print(registers[max(registers, key=lambda k: registers[k])])
