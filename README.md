# Golang Parser
This is a project from my Concepts of Programming Langauges course, CSC 3310. October 2021.

We implemented a full parser in Golang to work with a simple programming language containing 12 tokens. This project was further extended to output code in both Scheme and Prolog with the goal of having programs in all three langauges work together. The parser was created to run in Linux, and the output was determined by the file name given as an argument and the "-s" or "-p" flags for indicating Scheme or Prolog output respectively. The language grammar and tokens used are included below.

This project was a great demonstration of how to implement a Lexer and a Parser, and the necessary logic behind each to check the given code. My implementation is operational, but in the future I want to optimize the code. Having studied recursive descent parsing in this course, my goal is to implement a top-down recursive descent parser to process this same example language but with greater efficiency and readability.

## Language Grammar

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

## Language Tokens

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
