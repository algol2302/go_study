# comparison go closures with python closures

def make_summer1():
    x = 0

    def inner():
        nonlocal x
        x +=1
        return x

    return inner

a = make_summer1()

print(a())
print(a())
print(a())


def make_summer2():
    data = []
    def summer(val):
        data.append(val)
        _sum = sum(data)
        return _sum
    return summer

b = make_summer2()

print(b(1))
print(b(1))
print(b(1))
