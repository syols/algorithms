import collections
import math


def solution():
    a = input()
    b = input()
    lena = len(a)
    lenb = len(b)

    if lena != lenb:
        print("NO")
        return

    oa = {}
    ob = {}
    for i in range(lena):
        va = a[i]
        vb = b[i]
        if va not in oa:
            oa[va] = i

        if vb not in ob:
            ob[vb] = i

    for i in range(lena):
        va = a[i]
        vb = b[i]
        if oa[va] != ob[vb]:
            print("NO")
            return
    print("YES")


if __name__ == '__main__':
    solution()