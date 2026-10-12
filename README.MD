# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Go 1.24+
Git (for repository cloning)

# Project Structure
- 'README.MD' - Project documentation
- 'grade-calculator/go.mod' - Go module file for the grade calculator
- 'grade-calculator/grade_calculator.go' - Main grade calculator code
- 'grade-calculator/grade_calculator_test.go' - Unit tests for the grade calculator

# Grading System
Assignments - 50%
Exams - 35%
Essays - 15%

90-100 A
80-89 B
70-79 C
60-69 D
0-59 F

A grade of C or higher is a passing grade.


Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```