def parse(file):
    with open(file) as f:
        return [int(l.strip()) for l in f.readlines()]
    
def escape(jumps):
    pos = 0
    back = len(jumps) - 1
    front = 0

    count = 0

    while front <= pos <= back:
        jump = jumps[pos]
        jumps[pos] += 1
        pos += jump
        count += 1
    
    return count


if __name__ == '__main__':
    jumps = parse('input')
    print(escape(jumps))
