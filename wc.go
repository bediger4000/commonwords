package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type wcount struct {
	word  string
	count int
}

type swc []*wcount

func main() {
	fin := os.Stdin

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 2 {
		filename := os.Args[2]
		fin, err = os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
	}

	scanner := bufio.NewScanner(fin)

	wordcount := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			wordcount[word]++
		}
	}

	var freq swc

	for word, count := range wordcount {
		freq = append(freq, &wcount{word: word, count: count})
	}

	sort.Sort(freq)

	for i := 0; i < n; i++ {
		fmt.Printf("%s\t%d\n", freq[i].word, freq[i].count)
	}
}

func (wc swc) Len() int           { return len(wc) }
func (wc swc) Less(i, j int) bool { return wc[i].count < wc[j].count }
func (wc swc) Swap(i, j int)      { wc[i], wc[j] = wc[j], wc[i] }
