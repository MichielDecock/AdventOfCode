def parse(file):
    with open(file) as f:
        return  [int(l) for l in f.readline().strip().split('\t')]

def reallocate(banks):
    history = []

    while banks not in history:
        history.append(banks[:])

        blocks = max(banks)
        index = banks.index(blocks)
        banks[index] = 0

        for i in range(blocks):
            banks[(index + 1 + i) % len(banks)] += 1

    return len(history)

if __name__ == '__main__':
    res = reallocate(parse('input'))
    print(res)
