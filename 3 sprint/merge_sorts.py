def center(left, right):
    return int(((right - left) / 2 + left))


def merge(arr: list, left: int, mid: int, right: int) -> list:
    if right - left <= 1:
        return [arr[left]]

    larr = merge(arr, left, center(left, mid), mid)
    rarr = merge(arr, mid, center(mid, right), right)

    result = list()
    left, right = 0, 0
    while left < len(larr) and right < len(rarr):
        if larr[left] <= rarr[right]:
            result.append(larr[left])
            left += 1
        else:
            result.append(rarr[right])
            right += 1

    while left < len(larr):
        result.append(larr[left])
        left += 1

    while right < len(rarr):
        result.append(rarr[right])
        right += 1

    return result


def merge_sort(arr: list, left: int, right: int) -> None:
    c = center(left, right)
    res = merge(arr, left, c, right)
    arr.clear()
    arr.extend(res)


# def test():
#     # a = [1, 4, 9, 2, 10, 11]
#     # b = merge(a, 0, 3, 6)
#     # expected = [1, 2, 4, 9, 10, 11]
#     # assert b == expected
#     c = [1, 4, 2, 10, 1, 2]
#     merge_sort(c, 0, 6)
#     expected = [1, 1, 2, 2, 4, 10]
#     assert c == expected
#
#
# if __name__ == '__main__':
#     test()
