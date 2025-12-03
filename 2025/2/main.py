#! /opt/homebrew/bin/python3
from itertools import islice

if __name__ == "__main__":
    sum = 0
    with open("input.txt") as file:
        for ids in file.readline().split(','):
            limits = ids.split('-')
            for id in range(int(limits[0]), int(limits[1]) + 1):
                s = str(id)

                for div in range(2, max(2, len(s)) + 1):
                    if len(s) % div != 0:
                        continue

                    sub = ''
                    size = len(s) / div
                    for i in range(len(s)):
                        if i % size == 0:
                            sub += ',' + s[i]
                        else:
                            sub += s[i]
                    
                    parts = sub[1:].split(',')
                    if all(p == parts[0] for p in parts):
                        sum += id
                        break

    print(sum)