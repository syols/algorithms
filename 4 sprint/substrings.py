import math
from collections import OrderedDict


def solution():
    c = OrderedDict()
    s = input()
    max = 0
    for pos, r in enumerate(s):
        if r not in c:
            c[r] = pos
        else:
            value = c[r]
            c = {k: v for k, v in c.items() if v > value}
            length = len(c.items())
            if length > max:
                max = length
            c[r] = pos
        length = len(c.items())
        if length > max:
            max = length
    return max


if __name__ == '__main__':
    print(solution())
