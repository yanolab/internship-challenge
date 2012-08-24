#!/usr/bin/env python
# -*- coding: utf-8 -*-

def fizzbuzz(end, start=1):
    for x in xrange(start, end+1):
        if x % 15 == 0:
            print('fizzbuzz')
        elif x % 3 == 0:
            print('fizz')
        elif x % 5 == 0:
            print('buzz')
        else:
            print(x)

if __name__ == "__main__":
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument("endnum", metavar='ENDNUM', type=int, nargs=1)

    args = parser.parse_args()

    fizzbuzz(args.endnum[0])
