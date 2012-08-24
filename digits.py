#!/usr/bin/env python
# -*- coding: utf-8 -*-

if __name__ == "__main__":
    import sys
    import argparse
    from operator import mul

    import slide

    parser = argparse.ArgumentParser()
    parser.add_argument('file', type=argparse.FileType('r'), default=sys.stdin, nargs='?')
    parser.add_argument('-s', '--splitsize', nargs=1, type=int)

    args = parser.parse_args()

    size = args.splitsize[0]
    indigits = args.file.read().strip()

    print(max(reduce(mul, map(int, x)) for x in slide.sliding(indigits, size) if len(x) == size))
