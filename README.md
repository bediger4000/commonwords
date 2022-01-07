# Revisiting Knuth and McIlroy's word count programs

---
Read a file of text,
determine the n most frequently used words,
and print out a sorted list of those words along with their frequencies.

---

I found this problem statement on [this blog](https://franklinchen.com/blog/2011/12/08/revisiting-knuth-and-mcilroys-word-count-programs/)

PDF of original magazine exchange [here](https://www.cs.tufts.edu/~nr/cs257/archive/don-knuth/pearls-2.pdf)

I thought I'd do it in Go to see how hard it is.

I found a guy who uses this problem as a
[job interview question](https://benhoyt.com/writings/count-words/).
He offers quite a bit of what I call "interview analysis".
He's got a Go version to compare this code to.

## Results

* [McIlroy's shell script](mcilroy) version
* [My Go](wc.go) version, 68 lines, some blank, some boilerplate

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

This compares very favorably with the 9.5 hours spent
implementing a [C language](https://www.cs.upc.edu/~eipec/pdf/p583-van_wyk.pdf)
version published contemporaneously with Knuth's solution.

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

I'm going to both agree and disagree with the blogger, Franklin Chen.
I'm skeptical of literate programming for the same reasons he is.
But I don't think that "general purpose" programming languages
are a universal replacement for shell scripts.
It's far harder to get a C, Go or Haskell program correct
for some weird variant on word frequency,
than it is to get a shell script correct.
Knowing how to do the kind of text processing that McIlroy's
script illustrates is a valuable skill when dealing with more than
a page or two of information.
Modern sensibilities are distorted by Microsoft's vision of
what constitutes a good amount of information to process.
Word and Excel have lowered this "good amount" to about a page.
Simple, one-off text processing shell scripts have a lot of value.
