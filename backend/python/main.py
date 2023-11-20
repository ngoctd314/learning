a = []

a.append(1)
print(a)
a[len(a) :] = [2]
print(a)

a.extend(range(3, 5))
print(a)
a[len(a) :] = [5, 6]
print(a)

a.insert(0, -1)
print(a)

a.remove(-1)
print(a)
