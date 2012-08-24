#!/usr/bin/env python
# -*- coding: utf-8 -*-

def nabeatu(end, start=1):
    for x in xrange(start, end+1):
        if x % 3 == 0:
            print('Aho')
        elif '3' in str(x):
            print('Aho')
        else:
            print(x)

if __name__ == "__main__":
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument("endnum", metavar='ENDNUM', type=int, nargs=1)

    args = parser.parse_args()

    nabeatu(args.endnum[0])
