/*
Author: Logan Munoz
**NOTE** - there are 3 separate packages, syntax - lexical - spFlags
		each has the code for those respective parts to make this 
		code here simpler, so feel free to open those up and take a look!
		they were all coded by myself as well

Future Plans: Look for a better way to output Scheme and Prolog code, and also a better syntax analysis

*/
package main
import (
	"fmt"
	"flag"
	"os"
	syntax "cpl/parser/syntax"
	lexical "cpl/parser/lexical"
	spFlags "cpl/parser/spFlags"
	"regexp"
)

func main(){
	var fileName string
	var sourceCode string
	//Checks to ensure that you have at least put in one argument
	if len(os.Args) < 2 {
		fmt.Println("Error: please include a filename, and a \"-s\" or \"-p\" flag if you wish")
		return
	}
//test
	//This checks to ensure that the first argument input is a valid filename - if so makes that the fileName var
//	filename_regex, _ := regexp.Compile("^[a-zA-Z0-9]+\\.[a-zA-Z0-]+")
//	if len(os.Args) >= 2  && !filename_regex.MatchString(os.Args[1]) {
//		fmt.Println("Error: first argument is not a valid filename")
//		fileName = ""
//		return
//	} else {
		fileName = os.Args[1]
//	}


	//Calls the ValidFile function which ensures the file exists and can be opened
	if !ValidFile(fileName) {
		return
	}

	//Runs the lexer and the parser and returns if they fail - also grabs the source code
	lexerParserValid, sourceCode := RunLexerParser(fileName)
	if !lexerParserValid {
		return
	} 

	readableCode_regex, _ := regexp.Compile("^([^.]*).*")
	readableSourceCode := readableCode_regex.FindString(sourceCode)

	//This block creates all the different flags, and ensures that they are only -s or -p, otherwise output an error
	//Also ensure there are not too many arguments
	myFlagSet := flag.NewFlagSet("",flag.ExitOnError)
	var schemePtr = myFlagSet.Bool("s",false,"Running Scheme?")
	var prologPtr = myFlagSet.Bool("p",false,"Running Prolog?")
	if len(os.Args) >= 3 {
		if os.Args[2] != "-s" && os.Args[2] != "-p" {
			fmt.Println("Error: unrecognized flag. Please use the following:\n\"-s\" to output scheme code\n\"-p\" to output prolog code")
			return
		}
		if len(os.Args) > 3 && os.Args[3] != "-s" && os.Args[3] != "-p" {
			fmt.Println("Error: unrecognized flag. Please use the following:\n\"-s\" to output scheme code\n\"-p\" to output prolog code")
			return
		}
		if len(os.Args) > 4 {
			fmt.Println("Error: too many arguments")
			return
		}
		myFlagSet.Parse(os.Args[2:])
	}

	//Outputs the correct things depending on the flag the user passes
	if *schemePtr {
		spFlags.SchemeFlag(readableSourceCode, fileName)
	}
	if *prologPtr {
		spFlags.PrologFlag(readableSourceCode, fileName)
	}

}

//Attempts to open the file and if it cannot then stops the program and lets the user know
//Returns: bool true if can open file, false if not
func ValidFile(fileName string) bool {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: \"" + fileName + "\" was unable to be opened")
		f.Close()
		return false
	}
	f.Close()
	return true
}

//Runs both the lexer and the parser, gathering the tokenList to send to the parser and the source code for the flags later
//Returns: bool true if both pass, false if not - and sourceCode if passed, nothing if not
func RunLexerParser(fileName string) (bool, string) {
	lexicalPass, sourceCode, tokenList := lexical.LexicalAnalysis(fileName)
	if lexicalPass &&  syntax.SyntaxAnalysis(tokenList)  {
		return true, sourceCode
	}
	return false, ""
}







