#!/usr/bin/env python
# -*- coding: utf-8 -*-

def sliding(lst, size, step=1):
    return (lst[idx:idx+size] for idx in xrange(0, len(lst)-size+1, step))

if __name__ == "__main__":
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument('-s', '--size', type=int)
    parser.add_argument('strings', metavar='STRINGS', nargs=1)

    args = parser.parse_args()

    for window in sliding(args.strings[0], args.size):
        print("".join(window))
