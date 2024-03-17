# Lesson Notes - 01-quiz

## flag package

Package flag implements command-line flag parsing. You can use --help to get help output for your binary.

```bash
./quiz --help

# Usage of bin/quiz:
#  -csv string
#        a csv file in the format of question,answer (default "questions.csv")
```

## encoding/csv

allows reading csv files

```go
file, err := os.Open(*fileName)
// if err ...
r := csv.NewReader(file)
lines, err := r.ReadAll()
```

## trimming string spaces with strings package

```go
strings.TrimSpace(" lionel  messi  ")
// outputs "lionel messi"
```

## os package to open a file and exit from program

Open a file

```go
fileName := flag.String("csv", "questions.csv", "a csv file in the format of question,answer")
flag.Parse()
// adding * in front of the fileName because it is the address of the fileName parameter. os.Open wants the actual value so we are dereferencing(?) with the star
file, err := os.Open(*fileName)
```

Exit from the program with an error

```
os.Exit(1)
```

## creating a timer - not a ticker

```go
timer := time.NewTimer(time.Duration(*timeLimit * int(time.Second)))
```

## select case

select blocks the code. it is like a switch case where there are special cases and default case.

```go
select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if answer == p.a {
				correctCount++
			}
		}
```

this is a typical go select block where we have two different cases listens to different channels. If we get sth from timer channel we are breaking the loop (in go loops can have labels), if we get answer from the answerChannel we evaluate the answer and increase the counter.
