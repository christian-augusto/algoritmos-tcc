import time

from random import randint


def generate_random_int_array(size: int, min: int, max: int):
    array = []

    for _ in range(0, size):
        array.append(randint(min, max))

    return array


def copy_array(array: list):
    new_array = []

    for e in array:
        new_array.append(e)

    return new_array


def current_milli_time():
    return round(time.time() * 1000)
