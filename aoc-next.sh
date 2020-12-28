NEXT=$((`basename $(pwd)`+1))
mkdir "../$NEXT" && cd "../$NEXT" && ../../../aoc-init.sh && code -r "../$NEXT"
