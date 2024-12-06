# another jolly season of advent of code!

## Running a problem:
```bash
go run cmd/main.go
```

## Adding a new problem to the current running context:
* Create a new folder named Day2147 and a respective solution.go + input if necessary
* Import the problem in main.go
  * EX: `import "github.com/tzdanows/advent-of-code/2024-advent/Day2147"`
* Add it to the days mapping
  * EX: 
```go
    "Day2147": {
        inputPath:  "2024-advent/Day2147/sampleInput.txt",
        solvePart1: Day2147.SolvePart1,
        solvePart2: Day2147.SolvePart2,
    },
```
* Update the current day
  * EX: `currDay := "Day2147"`
* Run the problem
```bash
go run cmd/main.go
```

# links:
- [🎅🎄🎁](https://adventofcode.com/)
- https://blog.jetbrains.com/kotlin/2024/11/advent-of-code-2024-in-kotlin/
- 