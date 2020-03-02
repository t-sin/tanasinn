# tanasinn

> Don't think. Feel and you'll be tanasinn.

```
　　　　　　　　　　_,. -‐''"∴∵``' ‐ .､._
　　　　　　　　　 ,.‐'´∴∵∴∵∴∵∴∵ `‐.､
　　　　　　　 .／∴∵∴∵∴∵∴∵∴∵∴∵＼
　　　　　　 ,i´∴∵∴∵∴∵∴∵∴∵∴∵∴∵ヽ
.　　 　 　 /∴∵∴∵∴∵∴∵∴∵∴∵∴∵∴∵i、
　　　　　,i∴∵∴∵∴∵∴∵∴∵∴∵∴∵∴∵∴l
　　　 　 ∴∵∴∵∴∵∴(･)∴∴.(･)∴∵∴∵∴∵
.　　　　 ∴∵∴∵∴∵∴ ／ ○＼ ∴∵∴∵∴∵
.　 　 　 ∴∵∴∵∴∵∴/三　　三∴∵∴∵∴∵ 　tanasinn
.　　　　 ∴∵∴∵∴∵∴　＿_＿_ ∴∵∴∵∴∵l
　　　　　l∴∵∴∵∴∵∴　　===　.∴∵∴∵∴ .,l
.　　　　　ﾞi ∴∵∴∵∴∵ ＼＿＿／∴∵∴∵∴,i
　　　　　　ヽ∴ ∵∴∵∴∵∴∵∴∵∴∵∴∵/
　　　　 　 　 ＼∴∵∴∵∴∵∴∵∴∵∴∵／
　　　　　　　　　`‐.､∴∵∴∵∴∵∴∵∴,.‐'´
　　　　　　　　　　　 `:‐.､. _∴∵∴∵.-‐''"
```

# Installation

```sh
go get github.com/t-sin/tanasinn
```

# Usage

```sh
# from stdin
$ echo "Don't think. Feel and you'll be tanasinn." | ./tanasinn
･･on't thin... Feel an..∴yo∵'ll be∴∵t･:na∴inn.

# from file
$ echo "Don't think. Feel and you'll be tanasinn." > test.txt
$ ./tanasinn test.txt
Don't thin... Feel∴･and you'll be ∴anasi..n.

# change *tanasinnize threashold*
$ echo "Don't think. Feel and you'll be tanasinn." | ./tanasinn -t 0.9
Don't thin... (･)eel and..you'll b) ta･･asinn)
$ echo "Don't think. Feel and you'll be tanasinn." | ./tanasinn -t 0.1
∴･..∴)∵･･..∴∵..(･∴∵∴∵..･･)e.. (･)∴∵∴∵･:････))･･l∴(･)e･:∴∵)n(･)s∴∵･･.

# *tanasinnize threashold sequence
$ echo "Don't think. Feel and you'll be tanasinn." | ./tanasinn -t 1,0.7,0.3
Don't think. Feel and..(･)ou∴∴l･:be∵t∴∵(････:･:･････:
```

# Author

- tanasinn <shinichi.tanaka45@gmail.com>

# License

This program is licensed under the MIT license.
