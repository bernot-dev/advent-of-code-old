# Solutions to Advent of Code problems

I started learning Go a little while ago, and decided to participate in [Advent of Code](https://adventofcode.com/) 2020 to improve my skills. Now, I'm starting from the beginning and sharing my solutions. In general, I am optimizing code for readability and testability over speed of development, performance, and cleverness. The goal is to learn good habits in Go, rather than make the best use of every microsecond of processing time or go code golfing.

Priorities
1. Code is readable
2. Tests cover important logic
3. [Idomatic Go](https://golang.org/doc/effective_go.html)

[...]

10. Performant code, beyond the extent necessary for the program to run in a reasonable amount of time (less than 10 seconds).

[...]

50. Competition-style code, written as quickly as possible.

[...]

100. [Code golf](https://en.wikipedia.org/wiki/Code_golf).

[...]

1000. [Cleverness](https://twitter.com/eliotpeper/status/1276899756464238593).

## Structure

Folder structure follows URL path structure of Advent of Code. For instance, the solution to 2015, Day 3 will be at [adventofcode.com/2015/day/3](https://adventofcode.com/2015/day/3) and `./2015/day/3`.

A starting template that I use for each challenge lives at `./template`. A growing collection of helper functions that perform simple tasks live at `./helpers`.

### Main function
Regardless of the problem, the `main` function is kept slim. Generally, just this:

```golang
func main() {
	p1 := PartOne(ReadInput("input.txt"))
	fmt.Println("Part 1 Solution:", p1)

	p2 := PartTwo(ReadInput("input.txt"))
	fmt.Println("Part 2 Solution:", p2)
}
```

### main.go
The `ReadInput` function reads the data into a usable format. Depending on the chalenge, this may involve a simple helper function that returns a string, or some fairly heavy lifting building and populating structs. Corresponding tests live in `main_test.go`.

### partone.go
Code necessary to solve Part One of the challenge. Corresponding tests live in `partone_test.go`.

### parttwo.go
Code necessary to solve Part Two of the challenge. Ideally, this reuses functions/stucts/etc. from Part One, when applicable. Corresponding tests live in `parttwo_test.go`.

### sample.txt
Contains test puzzle input with known solution.

### input.txt
Not included in repo. The creator of Advent of Code does not permit sharing puzzle input, so those files have been excluded from this repo. However, solutions should work with any puzzle input. You can [sign up for your own Advent of Code account](https://adventofcode.com/2020/auth/login) to get your own puzzle input.


## Code Reviews Welcome!
I would love to hear your feedback! Please [open an issue](../../issues/new) to share thoughts/comments.
