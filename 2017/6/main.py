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

    print(banks)
    return len(history)

if __name__ == '__main__':
    print(reallocate([10, 9, 8, 7, 6, 5, 4, 3, 1, 1, 0, 15, 14, 13, 11, 12]))
