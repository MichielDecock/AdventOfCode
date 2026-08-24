from __future__ import annotations

import os
from dataclasses import dataclass


_LAST_ID = 0
_ID_MAP = dict()
_TARGET = 'shiny gold'

@dataclass
class Node:
    id: int
    children: list[tuple[Node, int]]

    def __eq__(self, other):
        if isinstance(other, int):
            return self.id == other
        return isinstance(other, Node) and self.id == other.id


def parse(filename):
    global _LAST_ID
    lines = dict()
    fullFilename = os.path.join(os.path.dirname(__file__), filename)
    with open(fullFilename) as file:
        for line in [line.strip('\n') for line in file]:
            line = line.replace('bags',' ')
            line = line.replace('bag',' ')
            line = line.replace('no other',' ')
            line = line.replace('.',' ')
            [key, _values] = [l.strip() for l in line.split('contain')]
            values = [v.strip() for v in _values.split(',') if v.strip()]
            if key not in _ID_MAP.keys():
                _ID_MAP[key] = _LAST_ID
                _LAST_ID += 1
            lines[key] = values

    return lines

def makeNodes(keys):
    return [Node(_ID_MAP[key], []) for key in keys]


def makeTree(lines):
    nodes = makeNodes(lines.keys())
    for key, values in lines.items():
        children = [(nodes[nodes.index(_ID_MAP[i[1]])], int(i[0])) for i in (v.split(' ', 1) for v in values)]
        nodes[nodes.index(_ID_MAP[key])].children = children

    return nodes

def totalBags(tree):
    node = tree[tree.index(_ID_MAP[_TARGET])]
    toVisit = [(node, 0)]

    total = 0

    while toVisit:
        node, localTotal = toVisit[0]
        total += localTotal
        toVisit.pop(0)

        for n, count in node.children:
            toVisit.append((n, max(localTotal, 1) * count))

    return total


if __name__ == "__main__":
    lines = parse('input')
    tree = makeTree(lines)
    print(totalBags(tree))
