PINK = 0
YELLOW = 1
CRIMSON = 2


def solution():
    n = int(input())
    v = input()
    if n == 0:
        return []

    hs = list(map(int, v.split(' ')))
    pink, yellow, crimson = 0, 0, 0

    for h in hs:
        if h == PINK:
            pink += 1
        if h == YELLOW:
            yellow += 1
        if h == CRIMSON:
            crimson += 1

    return [PINK] * pink + [YELLOW] * yellow + [CRIMSON] * crimson


if __name__ == '__main__':
    print(' '.join(list(map(str, solution()))))