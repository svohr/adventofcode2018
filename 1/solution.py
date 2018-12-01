'''
Advent of Code 2018 - Day 1
'''
from __future__ import print_function
import sys
import itertools

def read_frequency_delta(delta_str):
    sign = delta_str[0]
    value = int(delta_str[1:])
    return value * (1 if sign == '+' else -1)


def apply_frequency_changes(initial_freq, delta_in):
    frequency = initial_freq
    for delta_str in delta_in:
        frequency += read_frequency_delta(delta_str)
    return frequency


def find_first_repeat(initial_freq, delta_in):
    frequency = initial_freq
    prev_freqs = set([frequency])
    for delta_str in itertools.cycle(delta_in):
        frequency += read_frequency_delta(delta_str)
        if frequency in prev_freqs:
            return frequency
        prev_freqs.add(frequency)
    return None  # cannot reach here.


def test_apply_frequency_changes_simple():
    assert apply_frequency_changes(0, ['+10']) == 10
    assert apply_frequency_changes(0, ['-10']) == -10
    assert apply_frequency_changes(0, ['+1', '-2', '+3', '-4']) == -2


def test_find_first_repeat_simple():
    assert find_first_repeat(0, ['+1', '+2', '-3']) == 0
    assert find_first_repeat(0, ['+5', '-4']) == 5


def main():
    with open('input.txt', 'r') as delta_in:
        delta_input = delta_in.readlines()
    initial_freq = 0
    current_freq = apply_frequency_changes(initial_freq, delta_input)
    print('Updated Frequency: {}'.format(current_freq))

    first_repeat = find_first_repeat(initial_freq, delta_input)
    print('First Repeated Frequency: {}'.format(first_repeat))
    return 0


if __name__ == '__main__':
    sys.exit(main())
