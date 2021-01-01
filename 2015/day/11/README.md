# [2015, Day 11: Corporate Policy](https://adventofcode.com/2015/day/11)

Implementing the rules as they are laid out works just fine, but this one was actually faster to solve in my head. The text hints at this, but when you combine all of the rules, you know that what you're likely doing is matching the pattern ending in something like `aabcc`, `bbcdd`, `ccdee`, etc.

However, the logic required to handle the edge cases for that "faster" method can quickly become unwieldy. When earlier parts of the string have one or two pairs, or a straight, the ending pattern could be different. And you could have to handle all of the permutations if you were working against adversarial input:

1. No rules satisfied by first part of the string
2. Straight, but no pairs in first part of the string.
3. One pair, no straight.
4. Two pair, no straight.
5. Straight and one pair.
6. Straight and two pair.

The logic may seem relatively clear at this point. However, what if the next match requires rolling over an earlier character? Enter recursion. Meanwhile, the "slow" version of the program can solve this problem a few thousand times a second. Yes, the "fast" version, can theoretically solve it 1000x faster, but it's just not necessary in this case.

All that said, I have learned that it is usually best to solve Part One close to the description, because you never know exactly what variation you will get for Part Two. In this case, it turns out the same approach would have worked well, but that's just luck of the draw.
