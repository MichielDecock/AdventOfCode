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

if __name__ == '__main__':
    map = parse('input')
    res = findRoot(map)
    print(res)
