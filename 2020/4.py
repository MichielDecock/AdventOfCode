import os

_KEYS = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']
_OPTIONAL = 'cid'
_EYE_COLORS = ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']

def parse(filename: str):
    passports = []
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        keyValues = dict()
        for line in [line.strip('\n') for line in file]:
            if not line:
                passports.append(keyValues)
                keyValues = dict()
                continue

            keyValues |= dict([l for l in [l.split(':') for l in line.split(' ')] if l[0] != _OPTIONAL])

        if keyValues:
            passports.append(keyValues)

    return passports

def isComplete(passport):
    return len(passport) == len(_KEYS)

def checkBYR(value):
    return len(value) == 4 and 1920 <= int(value) and int(value) <= 2002

def checkIYR(value):
    return len(value) == 4 and 2010 <= int(value) and int(value) <= 2020

def checkEYR(value):
    return len(value) == 4 and 2020 <= int(value) and int(value) <= 2030

def checkHGT(value):
    if value.endswith('in'):
        return 59 <= int(value[:-2]) and int(value[:-2]) <= 76
    if value.endswith('cm'):
        return 150 <= int(value[:-2]) and int(value[:-2]) <= 193
    return False

def checkHCL(value):
    return len(value) == 7 and value[0] == '#' and all(c.isdigit() or c.islower() for c in value[1:])

def checkECL(value):
    return _EYE_COLORS.count(value) > 0

def checkPID(value):
    return len(value) == 9 and all(c.isdigit() for c in value)

def isValid(passport):
    return checkBYR(passport['byr']) and checkIYR(passport['iyr']) and checkEYR(passport['eyr']) and checkHGT(passport['hgt']) and checkHCL(passport['hcl']) and checkECL(passport['ecl']) and checkPID(passport['pid'])

if __name__ == "__main__":
    passports = parse('input')
    print(sum(1 for p in passports if isComplete(p) and isValid(p)))
