package syntax
import (
	"fmt"
	"regexp"
	"strings"
)

//Master function for running syntax analysis
//Returns: bool true if passed, false if not
//Gets the first token, which will indicate either an assignment statement or a test
//If so calls the respective validity check for that kind of statement - gets updated token list, and repeats
//If not returns a syntax error
func SyntaxAnalysis(tokenList string) bool {
	for len(tokenList) > 0 {
		nextToken, newTokenList := NextToken(tokenList)
		tokenList = newTokenList
		nextTokenClean, decoration := SeparateLexeme(nextToken)
		if nextTokenClean == "ID" {
			IDSyntaxValid, returnMessage := PointDef(tokenList)
			if !IDSyntaxValid {
				fmt.Println(returnMessage)
				return false
			} else if returnMessage == "done" {
				break
			}
			tokenList = returnMessage
		} else if nextTokenClean == "TEST" {
			TESTSyntaxValid, returnMessage := Test(tokenList)
			if !TESTSyntaxValid {
				fmt.Println(returnMessage)
				return false
			} else if returnMessage == "done" {
				break
			}
			tokenList = returnMessage
		} else {
			errorStatement := "Syntax error " + IdentifyToken(nextTokenClean, decoration)
			errorStatement += " found ID or test expected"
			fmt.Println(errorStatement)
			return false
		}
	}
	return true
}

//Gets the next token from the token list
//Returns: string of that token, and a string of the remaining token list
func NextToken(tokenList string) (string, string) {
	newLine_regex, _ := regexp.Compile("^.*[\\n]")
	if newLine_regex.MatchString(tokenList) {
		token := newLine_regex.FindString(tokenList)
		currentPosition := len(token)
		tokenList = tokenList[currentPosition:]
		return token, tokenList
	} else if len(tokenList) > 0 {
		token := tokenList
		tokenList = ""
		return token, tokenList
	} else {
		return "error", "error"
	}
}

//Essentially this is a translator for error messages
//Looks at the token and lexeme and returns a string for use in an error message
//Returns: string for error message
//For example, an ASSIGN would become a "=", so it looks better in the error message
//Or if the error has to do with an ID it will return the lexeme
func IdentifyToken(nextTokenClean string, decoration string) string {
	if nextTokenClean == "ID" {
		return decoration
	} else if nextTokenClean == "ASSIGN" {
		return "="
	} else if nextTokenClean == "SEMICOLON" {
		return ";"
	} else if nextTokenClean == "COMMA" {
		return ","
	} else if nextTokenClean == "PERIOD" {
		return "."
	} else if nextTokenClean == "LPAREN" {
		return "("
	} else if nextTokenClean == "RPAREN" {
		return ")"
	} else if nextTokenClean == "POINT" {
		return "point"
	} else if nextTokenClean == "NUM" {
		return decoration
	} else if nextTokenClean == "TRIANGLE" {
		return "triangle"
	} else if nextTokenClean == "SQUARE" {
		return "square"
	} else if nextTokenClean == "TEST"{
		return "test"
	} else {
		return "missing text"
	}
}

//Takes in a decorated token and separates it into the lexeme and the token itself
//Returns: string of just the token, string of the lexeme
func SeparateLexeme(tokenWithLexeme string ) (string, string) {
	token_regex, _ := regexp.Compile("^[A-Z]+")
	cleanToken := token_regex.FindString(tokenWithLexeme)
	decoration := strings.ReplaceAll(tokenWithLexeme[len(cleanToken):], " ", "")
	decoration = strings.ReplaceAll(decoration, "\n", "")
	return cleanToken, decoration
}

//This is called when a point assignment statement is seen
//Cascading ifs which match the correct order of such a statement
//If there are any issues it will create an error message for the specific issue and abort
//Returns: bool true if passed, false if not - string of error message
func PointDef(tokenList string) (bool, string) {
	nextToken, newTokenList := NextToken(tokenList)
	tokenList = newTokenList
	nextTokenClean, decoration := SeparateLexeme(nextToken)
	if nextTokenClean == "ASSIGN" {
		nextToken, newTokenList = NextToken(tokenList)
		tokenList = newTokenList
		nextTokenClean, decoration = SeparateLexeme(nextToken)

		if nextTokenClean == "POINT" {
			nextToken, newTokenList = NextToken(tokenList)
			tokenList = newTokenList
			nextTokenClean, decoration = SeparateLexeme(nextToken)

			if nextTokenClean == "LPAREN" {
				nextToken, newTokenList = NextToken(tokenList)
				tokenList = newTokenList
				nextTokenClean, decoration = SeparateLexeme(nextToken)

				if nextTokenClean == "NUM" {
					nextToken, newTokenList = NextToken(tokenList)
					tokenList = newTokenList
					nextTokenClean, decoration = SeparateLexeme(nextToken)
					
					if nextTokenClean == "COMMA" {
						nextToken, newTokenList = NextToken(tokenList)
						tokenList = newTokenList
						nextTokenClean, decoration = SeparateLexeme(nextToken)

						if nextTokenClean == "NUM" {
							nextToken, newTokenList = NextToken(tokenList)
							tokenList = newTokenList
							nextTokenClean, decoration = SeparateLexeme(nextToken)

							if nextTokenClean == "RPAREN" {
								nextToken, newTokenList = NextToken(tokenList)
								tokenList = newTokenList
								nextTokenClean, decoration = SeparateLexeme(nextToken)
								if nextTokenClean == "SEMICOLON" && len(tokenList) > 0 {
									return true, tokenList
								} else if nextTokenClean == "SEMICOLON" && len(tokenList) == 0 {
									errorStatement := "Syntax error ; found . expected"
									return false, errorStatement
								} else if nextTokenClean == "PERIOD" {
									return true, "done"
								} else {
									errorStatement := "Syntax error " + IdentifyToken(nextTokenClean, decoration)
									errorStatement += " found ; or . expected"
									return false, errorStatement
								}
							} else {
								errorStatement := "Syntax error " + IdentifyToken(nextTokenClean, decoration)
								errorStatement += " found ) expected"
								return false, errorStatement
							}
						} else {
							errorStatement := "Syntax error " + IdentifyToken(nextTokenClean, decoration)
							errorStatement += " found number expected"
							return false, errorStatement
						}
					} else {
						errorStatement := "Syntax error " +IdentifyToken(nextTokenClean, decoration)
						errorStatement += " found , expected"
						return false, errorStatement
					}
				} else {
					errorStatement := "Syntax error " + IdentifyToken(nextTokenClean, decoration)
					errorStatement += " found number expected"
					return false, errorStatement
				}
			} else {
				errorStatement := "Syntax error " + IdentifyToken(nextTokenClean, decoration)
				errorStatement += " found ( expected"
				return false, errorStatement
			}
		} else {
			errorStatement := "Syntax error " + IdentifyToken(nextTokenClean, decoration)
			errorStatement += " found point expected"
			return false, errorStatement
		}

	} else {
		errorStatement := "Syntax error " + IdentifyToken(nextTokenClean, decoration)
		errorStatement += " found = expected"
		return false, errorStatement
	}
}

//Validity test for a test statement when found
//Just like the previous function, is a series of cascading ifs to match the correct syntax of a test statement
//If there is an issue it will abort and send an error message
//Returns: bool true if passed, false if not - string of error message
func Test(tokenList string) (bool, string) {
	nextToken, newTokenList := NextToken(tokenList)
	tokenList = newTokenList
	nextTokenClean, decoration := SeparateLexeme(nextToken)
	if nextTokenClean == "LPAREN" {
		nextToken, newTokenList = NextToken(tokenList)
		tokenList = newTokenList
		nextTokenClean, decoration = SeparateLexeme(nextToken)

		if nextTokenClean == "TRIANGLE" || nextTokenClean == "SQUARE" {
			nextToken, newTokenList = NextToken(tokenList)
			tokenList = newTokenList
			nextTokenClean, decoration = SeparateLexeme(nextToken)

			if nextTokenClean == "COMMA" {
				nextToken, newTokenList = NextToken(tokenList)
				tokenList = newTokenList
				nextTokenClean, decoration = SeparateLexeme(nextToken)
				
				for true {
					if nextTokenClean == "ID" {
						nextToken, newTokenList = NextToken(tokenList)
						tokenList = newTokenList
						nextTokenClean, decoration = SeparateLexeme(nextToken)

						if nextTokenClean == "COMMA" {
							nextToken, newTokenList = NextToken(tokenList)
							tokenList = newTokenList
							nextTokenClean, decoration = SeparateLexeme(nextToken)
						} else if nextTokenClean == "RPAREN" {
						    	nextToken, newTokenList = NextToken(tokenList)
                        				tokenList = newTokenList
                        				nextTokenClean, decoration = SeparateLexeme(nextToken)
						    	break
						}else {
                                			errorStatement := "Syntax error " +IdentifyToken(nextTokenClean, decoration)
                                			errorStatement += " found , or ) expected"
                                			return false, errorStatement
                                			}
                    			} else {
						errorStatement := "Syntax error " + IdentifyToken(nextTokenClean, decoration)
						errorStatement += " found ID expected"
						return false, errorStatement
					}
				}
                        		if nextTokenClean == "SEMICOLON" && len(tokenList) > 0 {
                            			return true, tokenList
                        		} else if nextTokenClean == "SEMICOLON" && len(tokenList) == 0 {
                            			errorStatement := "Syntax error ; found . expected"
                            			return false, errorStatement
                        		} else if nextTokenClean == "PERIOD" {
                            			return true, "done"
                        		} else {
                            			errorStatement := "Syntax error " + IdentifyToken(nextTokenClean, decoration)
                            			errorStatement += " found ; or . expected"
                            			return false, errorStatement
                        		}
			} else {
				errorStatement := "Syntax error " + IdentifyToken(nextTokenClean, decoration)
				errorStatement += " found , expected"
				return false, errorStatement
			}
		} else {
			errorStatement := "Syntax error " + IdentifyToken(nextTokenClean, decoration)
			errorStatement += " found options triangle or square expected"
			return false, errorStatement
		}

	} else {
		errorStatement := "Syntax error " + IdentifyToken(nextTokenClean, decoration)
		errorStatement += " found ( expected"
		return false, errorStatement
	}
}

