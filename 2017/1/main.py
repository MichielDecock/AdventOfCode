def parse(file):
    with open(file) as f:
        return f.readline()
    
def check(sequence):
    s = 0
    for index, digit in enumerate(sequence):
        nxt = sequence[(index + 1) % len(sequence)]
        if digit == nxt:
            s += int(digit)
    return s

if __name__ == "__main__":
    sequence = parse('input')
    result = check(sequence)
    print(result)