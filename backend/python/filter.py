numbers = [1, 2, 3, 4]


def is_even(num):
    return num % 2 == 0


even_numbers = filter(is_even, numbers)

even_numbers_list = list(even_numbers)

print(even_numbers_list)

odd_numbers = filter(lambda num: num % 2 == 1, numbers)

odd_numbers_list = list(odd_numbers)

print(odd_numbers_list)
