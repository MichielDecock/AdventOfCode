#! /opt/homebrew/bin/python3

import math
from functools import reduce
from operator import mul

class Neighbor:
    def __init__(self, distance, index1, index2, h):
        self.distance = distance
        self.index1 = index1
        self.index2 = index2
        self.h = h
    
    def __lt__(self, other):
        return self.distance < other.distance
    
    def __eq__(self, other):
        return self.h == other.h
    
    def __str__(self):
        return f"{self.index1}\t{self.index2}\t{self.distance}"

    def __repr__(self):
        return str(self)
    

class Vector:
    def __init__(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z

    def distance(self, v):
        return math.sqrt(pow(self.x - v.x, 2) + pow(self.y - v.y, 2) + pow(self.z - v.z, 2))

def readFile(file):
    with open(file) as f:
        return [l.strip() for l in f.readlines()]
    
def getNodes(file):
    nodes = []
    for line in readFile(file):
        coords = [int(i) for i in line.split(',')]
        nodes.append(Vector(coords[0], coords[1], coords[2]))
    
    return nodes

def getNeighbors(nodes):
    neighbors = []
    for i, node in enumerate(nodes):
        for n in (ns := [Neighbor(node.distance(n), i, j, hash(frozenset({i, j}))) for j, n in enumerate(nodes) if not i == j]):
            neighbors.append(n)

    neighbors = {n.h: n for n in neighbors}.values()
    return sorted(neighbors)

def makeCircuit(circuits, neighbor):
    c1 = c2 = None

    for c in circuits:
        if neighbor.index1 in c:
            c1 = c
        if neighbor.index2 in c:
            c2 = c

    if c1 is None and c2 is None:
        circuits.append({neighbor.index1, neighbor.index2})
        return

    if c1 is not None and c2 is None:
        c1.add(neighbor.index2)
        return

    if c2 is not None and c1 is None:
        c2.add(neighbor.index1)
        return

    if c1 is not c2:
        c1 |= c2
        circuits.remove(c2)

if __name__ == "__main__":
    iterations = 10
    file = "test"

    real = True
    if real:
        iterations = 1000
        file = "input"

    neighbors = []

    nodes = getNodes(file)
    neighbors = getNeighbors(nodes)

    with open(f"out-{file}", "w") as file:
        for n in neighbors:
            file.write(str(n) + '\n')

    circuits = []

    for i in range(iterations):
        makeCircuit(circuits, neighbors.pop(0))

    circuits.sort(key = len, reverse=True)

    print("End result is", reduce(mul, [len(c) for c in circuits[:3]]))
