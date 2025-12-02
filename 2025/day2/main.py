#! /opt/homebrew/bin/python3

from itertools import islice

if __name__ == "__main__":
    sum = 0
    with open("input.txt") as file:
        for ids in file.readline().split(','):
            limits = ids.split('-')
            for id in range(int(limits[0]), int(limits[1]) + 1):
                s = str(id)

                if len(s) % 2 != 0:
                    continue

                n1 = ''.join(islice(s, None, len(s)//2))
                n2 = ''.join(islice(s, len(s)//2, None))

                if n1 == n2:
                    sum += int(id)

    print(sum)