インターンシップ課題
====================

次のQ1-Q6の設問に答えなさい。
ソースコードには適宜コメントを書き、動作の説明を行えるようにしなさい。

    制限時間：3時間
    使用言語：自由

Q1. FizzBuzz
------------

 コマンドラインから任意の数字xを受け取り、1からxまで順番に数を数え、その数が3で割り切れるなら"Fizz"、5で割り切れるなら"Buzz"、両方で割り切れるなら"FizzBuzz"と表示するプログラムを作成せよ。

    > fizzbuzz.exe 100
    1
    2
    Fizz
    4
    Buzz
    Fizz
    7
    8
    Fizz
    Buzz
    11
    Fizz
    13
    14
    FizzBuzz
    16
    .
    .
    .

Q2. 世界のナベアツ問題
----------------------

 コマンドラインから任意の数字xを受け取り、1からxまでの３の倍数と３のつく数字のときだけAhoと表示するプログラムを作成せよ。

    > nabeatu.exe 100
    1
    2
    Aho
    4
    5
    Aho
    7
    8
    Aho
    10
    .
    .
    .
    29
    Aho
    Aho
    Aho
    .
    .
    .

Q3. スライド関数
----------------

 引数に文字列とウインドウサイズをとり、1文字ずつずらしたウインドウの配列を戻り値として返す関数を作成しなさい。また、その関数を用いてコマンドラインで与えられたウインドウサイズで文字列を分割して出力するプログラムを作成せよ。

    > slidestr.exe 2 ABCDEFG
    AB
    BC
    CD
    DE
    EF
    FG

    > slidestr.exe 3 ABCDEFG
    ABC
    BCD
    CDE
    DEF
    EFG

Q5. 文字列処理
--------------

 次のURL([https://raw.github.com/yanolab/internship-challenge/master/digits.txt](https://raw.github.com/yanolab/internship-challenge/master/digits.txt))で与えられた数字から5つの連続する数字を取り出して その積を計算する。そのような積の中で最大のものの値はいくらかもとめるプログラムを作成せよ。ただし、ファイルはプログラムから直接ロードしても、ローカルにダウンロードしてから利用しても構わない。

    # file -> 123456789
    > digits.exe file
    15120

    > digits.exe https://raw.github.com/yanolab/internship-challenge/master/digits.txt
    31752

Q6. 素数
--------

 素数を小さい方から6つ並べると 2, 3, 5, 7, 11, 13 であり、6番目の素数は 13 である。コマンドラインから数xを受け取り、x番目の素数を求めるプログラムを作成せよ。

    > prime.exe 10001
    104743
