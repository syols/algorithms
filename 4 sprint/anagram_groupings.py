import math
from collections import OrderedDict


def solution():
    input()
    s = input().split()
    d = OrderedDict()
    for i, w in enumerate(s):
        sw = sorted(w)
        w = "".join(sw)
        if w in d:
            d[w].append(i)
        else:
            d[w] = [i]

    return [sorted(v) for k, v in d.items()]


if __name__ == '__main__':
    for i in solution():
        print(" ".join([str(intg) for intg in i]))
