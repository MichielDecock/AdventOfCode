#! /opt/homebrew/bin/python3
from itertools import islice
import copy

if __name__ == "__main__":
    sum = 0
    with open("input") as file:
        matrix = [list(line.rstrip("\n")) for line in file]
        
        for i, el in enumerate(matrix):
            for j, el2 in enumerate(el):
                if el2 != '@':
                    continue

                hits = 0
                for k in range(max(i - 1, 0), min(i + 1, len(matrix) - 1) + 1):
                    for l in range(max(j - 1, 0), min(j + 1, len(el) - 1) + 1):
                        if k == i and l == j:
                            continue

                        if matrix[k][l] == '@':
                            hits += 1
                
                if hits < 4:
                    sum += 1

    print(sum)