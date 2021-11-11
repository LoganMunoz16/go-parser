package lexical
import (
	"fmt"
	"os"
	"regexp"
)
//Runs the lexical analysis on the program, stopping if encountering an error
//Returns: bool true if successful false if not, string of sourceCode, and string of tokenList
func LexicalAnalysis(fileName string) (bool, string, string) {
	var sourceCodeBytes []byte
	var sourceCode string
	var sourceCodeCopy string
	var tokenList string
	tokenList = ""

	//Gathers the sourceCode from the file and creates a copy to return later
	sourceCodeBytes, _ = os.ReadFile(fileName)
	sourceCode = string(sourceCodeBytes)
	sourceCodeCopy = sourceCode
	
	//Regex block to check for all the different tokens
	id_regex, _ := regexp.Compile("^[a-z]+")
	num_regex, _ := regexp.Compile("^\\b[0-9]+")
	triangle_regex, _ := regexp.Compile("^\\b[t][r][i][a][n][g][l][e]\\b")
	square_regex, _ := regexp.Compile("^\\b[s][q][u][a][r][e]\\b")
	test_regex, _ := regexp.Compile("^\\b[t][e][s][t]\\b")
	point_regex, _ := regexp.Compile("^\\b[p][o][i][n][t]\\b")
	whitespace_regex, _ := regexp.Compile("^\\s")

	var currentPosition int
	//Runs through and tries to match each token
	//If it does, then it adds it to the tokenList along with the lexeme if applicable
	//To do so gets the length of the match, and discards that many characters from the front of the sourceCode
	//Prints a lexical error if there is something that is not recognized
	for true {
		currentPosition = 0
		if sourceCode[currentPosition] == ';' {
			tokenList += "SEMICOLON\n"
			sourceCode = sourceCode[1:]
		} else if sourceCode[currentPosition] == ',' {
			tokenList += "COMMA\n"
			sourceCode = sourceCode[1:]
		} else if sourceCode[currentPosition] == '.' {
			tokenList += "PERIOD\n"
			sourceCode = sourceCode[1:]
		} else if sourceCode[currentPosition] == '(' {
			tokenList += "LPAREN\n"
			sourceCode = sourceCode[1:]
		} else if sourceCode[currentPosition] == ')' {
			tokenList += "RPAREN\n"
			sourceCode = sourceCode[1:]
		} else if sourceCode[currentPosition] == '=' {
			tokenList += "ASSIGN\n"
			sourceCode = sourceCode[1:]
		} else if whitespace_regex.MatchString(sourceCode) {
			sourceCode = sourceCode[1:]
		} else if point_regex.MatchString(sourceCode) {
			matchLength := len(string(point_regex.FindString(sourceCode)))
			currentPosition += matchLength
			sourceCode = sourceCode[currentPosition:]
			tokenList += "POINT\n"
		} else if test_regex.MatchString(sourceCode) {
			matchLength := len(string(test_regex.FindString(sourceCode)))
			currentPosition += matchLength
			sourceCode = sourceCode[currentPosition:]
			tokenList += "TEST\n"
		} else if triangle_regex.MatchString(sourceCode) {
			matchLength := len(string(triangle_regex.FindString(sourceCode)))
			currentPosition += matchLength
			sourceCode = sourceCode[currentPosition:]
			tokenList += "TRIANGLE\n"
		} else if square_regex.MatchString(sourceCode) {
			matchLength := len(string(square_regex.FindString(sourceCode)))
			currentPosition += matchLength
			sourceCode = sourceCode[currentPosition:]
			tokenList += "SQUARE\n"
		} else if num_regex.MatchString(sourceCode) {
			matchLength := len(string(num_regex.FindString(sourceCode)))
			currentPosition += matchLength
			lexeme := sourceCode[:currentPosition]
			sourceCode = sourceCode[currentPosition:]
			tokenList = tokenList + "NUM " + lexeme + "\n"
		} else if id_regex.MatchString(sourceCode) {
			matchLength := len(string(id_regex.FindString(sourceCode)))
			currentPosition += matchLength
			lexeme := sourceCode[:currentPosition]
			sourceCode = sourceCode[currentPosition:]
			tokenList = tokenList + "ID  " + lexeme + "\n"
		} else {
			fmt.Println("Lexical error " + string(sourceCode[currentPosition]) + " not recognized")
			return false, "", ""
		}
		if len(sourceCode) == 0 {
			break
		}
	}
	return true, sourceCodeCopy, tokenList
}
