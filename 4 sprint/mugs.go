import math
from collections import OrderedDict


def solution():
    a = int(input())
    od = OrderedDict()
    for i in range(a):
        od[input()] = None
    return od


if __name__ == '__main__':
    for key, value in solution().items():
        print(key)
