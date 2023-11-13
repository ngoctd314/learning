import math

numbers = [1, 2, 3, 4, 5]


def square(x):
    return x**2


squared_numbers = map(square, numbers)

squared_numbers_list = list(squared_numbers)

print(squared_numbers_list)

sqrt_numbers = map(lambda x: math.sqrt(x), numbers)

sqrt_numbers_list = list(sqrt_numbers)
print(sqrt_numbers_list)
