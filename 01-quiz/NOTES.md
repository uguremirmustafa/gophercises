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
