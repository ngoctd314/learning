symbols = "edbae"

numbers = [1, 2, 3, 4, 5, 6, 7, 8]


def is_even(num):
    return num % 2 == 0


even_numbers = filter(is_even, numbers)

event_number_lists = list(even_numbers)
print(event_number_lists)
