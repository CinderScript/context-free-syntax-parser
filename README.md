# Go Programming Assignment: 4Point Grammar Lexical and Syntax Analysis


## Motivation
Go is a programming language designed at Google. From the Go [website](https://golang.org/ref/spec):
> Go is a general-purpose language designed with systems programming in mind. It is strongly typed and garbage-collected and has explicit support for concurrent programming. Programs are constructed from packages, whose properties allow efficient management of dependencies.
> The grammar is compact and simple to parse, allowing for easy analysis by automatic tools such as integrated development environments.

This project consists in the development of the front end of a compiler. By programming the Lexical Analyzer (Scanner) and Syntax Analyzer (Parser) for the 4Point grammar you will gain further understanding of the lexical analysis and the production of tokens needed for the Syntax Analyzer (Parser), and how to consume those tokens by the Parser to verify that the syntax is correct.


## Description
Write a program in Go that takes a program written in 4Point, and outputs:
1. If the code has lexical or syntax errors, the error that was found. Use panic version of error handling (once an error is found report the error and stop the process).
1. If the code is OK, depending on a command line flag the program will produce:
   1.	If the flag is `-s` the program will output function calls in Scheme that is going to be called by a program in Scheme that will calculate properties of those three points.
   1. If the flag is `-p` the program will output a series of queries about those three points.

The program should run like this:
```
prompt>go run . input.txt -s
; Processing Input File input.txt
; Lexical and Syntax analysis passed
; Generating Scheme Code
(process-triangle (make-point 2 3) (make-point 1 4) (make-point 3 4))
prompt>
```

## Grammar

```
START      --> STMT_LIST
STMT_LIST  --> STMT. |
               STMT; STMT_LIST
STMT       --> POINT_DEF |
               TEST
POINT_DEF  --> ID = point(NUM, NUM)
TEST       --> test(OPTION, POINT_LIST)
ID         --> LETTER+
NUM        --> DIGIT+
OPTION     --> triangle |
               square
POINT_LIST --> ID |
               ID, POINT_LIST
LETTER     --> a | b | c | d | e | f | g | ... | z
DIGIT      --> 0 | 1 | 2 | 3 | 4 | 5 | 6 | ... | 9

```

The tokens of this grammar are:

Token | Lexeme
------ | ------
`POINT` | `point`
`ID` | `identifier`
`NUM` | `234`
`SEMICOLON` | `;`
`COMMA` | `,`
`PERIOD` | `.`
`LPAREN` | `(`
`RPAREN` | `)`
`ASSIGN` | `=`
`TRIANGLE` | `triangle`
`SQUARE` | `square`
`TEST` | `test`

Given the following program written in this language:
```
a = point(2, 3);
b = point(1, 1);
c = point(1, 3);
d = point(0, 0);
test(square, a, b, c, d);
test(triangle, a, b, c).
```
The tokens that it would generate are:
```
ID  a
ASSIGN
POINT
LPAREN
NUM 2
COMMA
NUM 3
RPAREN
SEMICOLON
ID  b
ASSIGN
POINT
LPAREN
NUM 1
COMMA
NUM 1
RPAREN
SEMICOLON
ID  c
ASSIGN
POINT
LPAREN
NUM 1
COMMA
NUM 3
RPAREN
SEMICOLON
ID  d
ASSIGN
POINT
LPAREN
NUM 0
COMMA
NUM 0
RPAREN
SEMICOLON
TEST
LPAREN
SQUARE
COMMA
ID a
COMMA
ID b
COMMA
ID c
COMMA
ID d
RPAREN
SEMICOLON
TEST
LPAREN
TRIANGLE
COMMA
ID a
COMMA
ID b
COMMA
ID c
RPAREN
PERIOD
```

Notice that the ID and NUM tokens have their lexeme associated. Also notice that in the language the elements do not need to be separated by space.

---
**NOTE**


The purpose of this assignment is _not_ to just output the list of tokens of a given programs. One of the steps of your program can be to produce this list, but the actual goal is to do the following phase of syntax analysis and producing the code on Scheme or Prolog. Furthermore, remember to remove that temporary outout, your program will be graded by a script, and the list of tokens is *not* part of the expected output.

---


## How to run the program

The following examples assume that `input.txt` contains the following code:
```
a = point(2, 3);
b = point(1, 1);
c = point(1, 3);
d = point(0, 0);
test(square, a, b, c, d);
test(triangle, a, b, c).
```

### Scheme Output
To generate scheme output you will add the `-s` flag at the end of the command:
```
prompt> go run . input.txt -s
; processing input file input.txt
; Lexical and Syntax analysis passed
; Generating Scheme Code
(process-square (make-point 2 3) (make-point 1 1) (make-point 1 3) (make-point 0 0))
(process-triangle (make-point 2 3) (make-point 1 1) (make-point 1 3))
```

### Prolog Output
To generate scheme output you will add the `-p` flag at the end of the command:
```
prompt> go run .  input.txt -p
/* processing input file input.txt
   Lexical and Syntax analysis passed
   Generating Prolog Code */

 /* Processing test(square, a, b, c, d) */
 query(square(point2d(2, 3), point2d(1, 1), point2d(1, 3), point2d(0, 0)))

 /* Professing test(triangle, a, b, c) */
 query(line(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(triangle(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(vertical(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(horizontal(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(equilateral(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(isosceles(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(right(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(scalene(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(acute(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(obtuse(point2d(2,3), point2d(1,1), point2d(1, 3))).
 
 /* Query Processing */
 writeln(T) :- write(T), nl.
 main:- forall(query(Q), Q-> (writeln(‘yes’)) ; (writeln(‘no’))),
       halt.

```
### Some code with errors

The following example assumes that `another.txt` contains the following code:
```
one=pont(1,1);
two=point(2,2);
three=point(3,3).
```
Independently of the flag (`-p` or `-s`).
```
prompt> go run . another.txt -p
Syntax error pont found point expected
```

The following example assumes that `thelast.txt` contains the following code:
```
misa=point[1,2];
jarjar=point(0,0);
binks=point(@,@).
```
Independently of the flag (`-p` or `-s`).
```
prompt> go run . thelast.txt -p
Lexical error [ not recognized
```

Later we will redirect the output to Scheme and Prolog programs respectively.

## Assignment Requirements
- Good programming practices
  - Indentation
  - Meaningful identifier naming
  - Consistent variable naming convention
- This activity is strictly individual

## Delivery
You will use this repository and commit and push to it. Remember to push your last version before the deadline.
What files should be in your repository:
- `parser.go` Source code in Go for your lexical and syntax analysis
- `go.mod` Modules file
- `test0.cpl`, `test1.cpl`, `test2.cpl`, `test3.cpl`, `test4.cpl`, `test5.cpl`, `test6.cpl`,the test files provided for you to test. Notice that `test3.cpl`  and `test4.cpl` have some issues, additionally `test5.cpl` has a lexical error and `test6.cpl` has a syntax error.

## Assessment and Grading
Assessment will consider the following factors in the grading of this assignment:
-	Good programming practices
-	Your program will be tested with the test programs that were provided and some others that will be made to test, some with lexical errors, some with syntax errors, some without any errors.
-	Adherence to instructions
-	Correct function of the program
-	No runtime errors
-	Late deliveries will have a zero mark
-	Plagiarism will have a double zero mark (in addition to losing 10% of your final grade, the person that plagiarizes will lose an additional 10% of their final grade), besides there will be a report filed in the students’ academic record.

## Extra Challenge

Create an additional file with a decorated (Tokens with Lexemes on the Leaves) parse tree. The file should be named like the input file but with the extension `.pt` (stands for parse tree), for instance if the input file is `test1.cpl` the parse tree should be in file `test1.pt`.

## Academic Integrity
This programming assignment is to be done on an **individual** basis. At the same time, it is understood that learning from your peers is valid and you are encouraged to talk among yourselves about programming in general and current assignments in particular.  Keep in mind, however, that each individual student must do the work in order to learn.  Hence, the following guidelines are established:
- Feel free to discuss any and all programming assignments but do not allow other students to look at or copy your code. Do not give any student an electronic or printed copy of any program you write for this class.
- Gaining the ability to properly analyze common programming errors is an important experience. Do not deprive a fellow student of his/her opportunity to practice problem solving: control the urge to show them what to do by writing the code for them.
- If you’ve given the assignment a fair effort and still need help, see the instructor or a lab assistant.
- **If there is any evidence that a program or other written assignment was copied from another student, neither student will receive any credit for it. This rule will be enforced.**
- Protect yourself: Handle throw-away program listings carefully, keep your repository private.

Refer to the ECS Department Policy on Academic Integrity that is included in the class syllabus.
