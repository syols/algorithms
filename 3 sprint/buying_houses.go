def solution():
    k, n = list(map(int, input().split(' ')))
    hs = sorted(list(map(int, input().split(' '))))
    ok = 0
    for home in hs:
        if n - home < 0:
            return ok
        ok += 1
        n -= home
        continue
    return min(ok, k)


if __name__ == '__main__':
    print(solution())