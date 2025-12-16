def parse(file):
    with open(file) as f:
        return [[''.join(sorted(word)) for word in l.strip().split(' ')] for l in f.readlines()]
    
def check(words):
    return len(words) == len(set(words))

if __name__ == "__main__":
    phrases = parse('input')
    print(sum([check(p) for p in phrases]))