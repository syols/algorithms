import math



def solution():
    a = int(input())
    m = int(input())
    s = input()

    hs = 0
    for v in s:
        hs = ((hs * a) + ord(v)) % m
    return hs % m


if __name__ == '__main__':
    print(solution())



