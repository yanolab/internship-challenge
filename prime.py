#! /usr/bin/env python
# -*- coding: utf-8 -*-

import math
from bisect import bisect
from itertools import dropwhile, imap, islice, count

def primes():
    """素数を2から順に返すジェネレータ"""
    yield 2
    _primes = [2]
    for i in count(3, 2):
      if all(imap(lambda x: i % x, islice(_primes, 0, bisect(_primes, math.sqrt(i))))):
          _primes.append(i)
          yield i

if __name__ == "__main__":
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument('position', metavar='POSITION', nargs=1, type=int)

    args = parser.parse_args()

    print next(dropwhile(lambda x: x[0] < args.position[0], enumerate(primes(), 1)))
