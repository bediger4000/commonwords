# Revisiting Knuth and McIlroy's word count programs

---
Read a file of text,
determine the n most frequently used words,
and print out a sorted list of those words along with their frequencies.
---

I found this problem statement on [this blog](https://franklinchen.com/blog/2011/12/08/revisiting-knuth-and-mcilroys-word-count-programs/)

I thought I'd do it in Go to see how hard it is.

## Results

* [McIlroy's shell script] version
* [My Go] version, 68 lines, some blank, some boilerplate

Test inputs:

1. Small count [word frequency](test1.in)
2. Upper/lower case [word frequency](test2.in)

Even a few simple test cases revealed bugs in my program,
so it's probably not as good as Knuth's literate,
almost guaranteed-correct version.

## Analysis

The Go program imports 7 packages from the standard library,
and defines 2 new types.
It makes heavy use of standard library for line-oriented I/O,
and to break a line into words.
One of the new types arises because of the way the "top N" requirement
interacts with the `sort` package's use of Go interfaces.
The actual unique-word counting gets done using a Go `map` type.

It took far less than 3 hours to do this task.
The git logs show amost exactly 3 hours between first and last
check in of [wc.go](wc.go).
I did eat a meal in there,
and have a longish discussion with a family member.

The real problem is deciding what's the correct output
for something like this very `README.md` file.

```sh
1499 % ./mcilroy 5 < README.md
     13 the
     10 a
      9 and
      9 go
      9 of
1500 % ./wc 5 < README.md
./wc 5 < README.md
     13 the
     10 a
      9 of
      8 and
      6 go
```

McIlroy's shell script turns all punctuation marks into newlines,
exposing "-and-", which appears in a URL, as the word "and".
My program does not truck with punctuation, so it misses that "and".
Should the string "revisiting-knuth-and-mcilroys-word-count-programs"
count as 7 "words", or just one?

This small difference illustrates a major problem in software engineering.
What's obvious to the requirements writer may not be obvious to the programmer.
Programming entails attention to "minor" details like that,
and "correct" vs "incorrect" programs hinge on those minor details.
