def parse(file):
    with open(file) as f:
        map = {}
        lines = [l.strip() for l in f.readlines()]
        for line in lines:
            l = line.split('->')
            parent, weight = [s for s in l[0].split(' ') if s != '']
            weight = int(weight[1:-1])
            children = [l.strip() for l in l[1].split(',')] if len(l) > 1 else []
            map[parent] = (weight, children)
        return map
    
def findRoot(map):
    for parent in map.keys():
        stop = False
        for _, (_, children) in map.items():
            if parent in children:
                stop = True
                break
        
        if not stop:
            return parent

def totalWeight(item, map):
    weight, children = map[item]
    return weight + sum([totalWeight(child, map) for child in children])

def imbalancedItem(map):
    imbalancedItem = None
    toVisit = [findRoot(map)]

    while toVisit:
        item = toVisit.pop(0)
        weight, children = map[item]
        if not children:
            continue

        weights = [totalWeight(child, map) for child in children]
        if len(set(weights)) < 2:
            continue
        for child in children:
            toVisit.append(child)
        imbalancedItem = item
    
    return imbalancedItem

if __name__ == '__main__':
    map = parse('input')
    parent = imbalancedItem(map)
    _, children = map[parent]

    weights = [totalWeight(child, map) for child in children]
    outlier = [w for w in weights if weights.count(w) == 1][0]
    delta = max(weights) - min(weights)
    item = children[weights.index(outlier)]
    weight, _ = map[item]

    if outlier == max(weights):
        print(weight - delta)
    else:
        print(weight + delta)


