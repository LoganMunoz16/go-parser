package spFlags
import (
	"fmt"
	"regexp"
	"unicode"
	"strings"
)

//This function creates a matrix with all of the IDs in the program and their associated values
//Returns: matrix described above
func CreateIDString(sourceCode string) [][]string {
	assignment_regex, _ := regexp.Compile("[a-zA-Z0-9_$]+\\s*[=].*[\\n]")
	idData_regex, _ := regexp.Compile("^.*[(]")
	allIDs := assignment_regex.FindAllString(sourceCode, -1)
	IDMatrix := make([][]string, len(allIDs))

	//Creating the rest of the slices for the matrix - each 3 long for ID, NUM, NUM
	for k := 0; k < len(allIDs); k++ {
		IDMatrix[k] = make([]string, 3)
	}
	var currentPosition int
	var currentCharacter rune

	//Going through the source code and gathering all of the IDs for the matrix
	for i := 0; i < len(allIDs); i++ {
		temp := allIDs[i]
		currentPosition = 0
		for true {
			currentPosition++
			currentCharacter = rune(temp[currentPosition])
			if currentCharacter == ' ' || currentCharacter == '=' {
				break
			}
		}
		IDMatrix[i][0] = temp[:currentPosition]
	}

	//This loops through the source code again and gather all the numbers associated with each ID
	for j := 0; j < len(allIDs); j++ {
		temp := allIDs[j]
		num1 := ""
		num2 := ""
		tempnum := ""
		count := 0
		currentPosition = 0
		matchLength := len(string(idData_regex.FindString(temp)))
		currentPosition += matchLength
		for true {
			currentCharacter = rune(temp[currentPosition])
			if currentCharacter == ')' {
				break
			} else if !unicode.IsNumber(currentCharacter) {
				currentPosition++
				continue
			} else if unicode.IsNumber(currentCharacter){
				for true {
					currentCharacter = rune(temp[currentPosition])
					if unicode.IsNumber(currentCharacter) {
						tempnum += string(currentCharacter)
						currentPosition++
					} else {
						break
					}
				}

				if count == 0 {
					num1 = tempnum
					count++
					tempnum = ""
				} else {
					num2 = tempnum
				}
			}
		}
		IDMatrix[j][1] = num1
		IDMatrix[j][2] = num2
	}
	return IDMatrix
}

//Takes in all the shapes being tested - all triangles or all squares -  and creates a matrix holding the ids needed for each shape
//Returns: matrix with ids for each shape
func AllShapeIDs (allShapes []string) [][]string {
	ShapeMatrix := make([][]string, len(allShapes))
	for k := 0; k < len(allShapes); k++ {
		ShapeMatrix[k] = make([]string, 10)
		allShapes[k] = strings.ReplaceAll(allShapes[k], " ", "")
	}
	var currentPosition int
	var currentCharacter rune
	var count int
	
	//Iterates through each test and attaches necessary IDs to it
	for i := 0; i < len(allShapes); i++ {
		count = 0
		temp := allShapes[i]
		currentPosition = 2
		for true {
			tempID := ""
			for true {
				currentPosition++
				currentCharacter = rune(temp[currentPosition])
				if currentCharacter == ')' {
					break
				} else if currentCharacter == ',' {
					break
				} else {
					tempID += string(currentCharacter)
				}
			}
			ShapeMatrix[i][count] = tempID
			count++
			if currentCharacter == ')' {
				break
			}
		}
	}
	return ShapeMatrix
}

//Generates the output for the scheme flag
//Takes the sourceCode and makes the required matrices to like the IDs with each test and each ID with its numbers
func SchemeFlag(sourceCode string, fileName string) {
	test_regex, _ := regexp.Compile("[t][e][s][t].*[)]")
	testSquare_regex, _ := regexp.Compile("[s][q][u][a][r][e]")
	testTriangle_regex, _ := regexp.Compile("[t][r][i][a][n][g][l][e]")
	squareCodeLine_regex, _ := regexp.Compile("[r][e].*[)]")
	triangleCodeLine_regex, _ := regexp.Compile("[l][e].*[)]")

	allTests := test_regex.FindAllString(sourceCode, -1)
	allTriangles := triangleCodeLine_regex.FindAllString(sourceCode, -1)
	allSquares := squareCodeLine_regex.FindAllString(sourceCode, -1)
	
	//Creating all the required matrices
	IDMatrix := CreateIDString(sourceCode)
	allTriangleIDs := AllShapeIDs(allTriangles)
	allSquareIDs := AllShapeIDs(allSquares)

	var schemeOutput string
	schemeOutput = "; processing input file " + fileName
	schemeOutput += "\n; Lexical and Syntax analysis passed\n; Generating Scheme Code\n"
	
	s_count := 0
	t_count := 0
	var relevantTest []string

	if len(allTests) == 0 {
		schemeOutput += "Error: no tests in given code\n"
	}

	//Generates the output based on the information in the matrices
	for i := 0; i < len(allTests); i++ {
		if testSquare_regex.MatchString(allTests[i]) {
			schemeOutput += "(process-square "
			relevantTest = allSquareIDs[s_count]
			s_count++
		} else if testTriangle_regex.MatchString(allTests[i]) {
			schemeOutput += "(process-triangle "
			relevantTest = allTriangleIDs[t_count]
			t_count++
		}
		
		
			for index, ID := range relevantTest {
				if relevantTest[index] == "" {
					break
				}
				noIDCount := 0
				schemeOutput += "(make-point "
				for j := 0; j < len(IDMatrix); j++ {
					if IDMatrix[j][0] == ID {
						schemeOutput += IDMatrix[j][1]
						schemeOutput += " "
						schemeOutput += IDMatrix[j][2]
						schemeOutput += ") "
						break
					}
					noIDCount++
				}
				if noIDCount == len(IDMatrix) {
					schemeOutput = schemeOutput[:(len(schemeOutput)-12)]
					schemeOutput += "(Error ID "
					schemeOutput += ID
					schemeOutput += " not defined) "
				}
			}
			schemeOutput = schemeOutput[:(len(schemeOutput)-1)]
			schemeOutput += ")\n"

	}

	schemeOutput = schemeOutput[:(len(schemeOutput)-1)]
	fmt.Println(schemeOutput)
}

//Just like the scheme option, takes in source code and generates needed matrices
//Then uses those ID and test matrices to match up all the IDS and NUMs for the needed output
func PrologFlag(sourceCode string, fileName string) {
	var prologOutput string

	test_regex, _ := regexp.Compile("[t][e][s][t].*[)]")
	testSquare_regex, _ := regexp.Compile("[s][q][u][a][r][e]")
	testTriangle_regex, _ := regexp.Compile("[t][r][i][a][n][g][l][e]")
	squareCodeLine_regex, _ := regexp.Compile("[r][e].*[)]")
	triangleCodeLine_regex, _ := regexp.Compile("[l][e].*[)]")

	allTests := test_regex.FindAllString(sourceCode, -1)
	allTriangles := triangleCodeLine_regex.FindAllString(sourceCode, -1)
	allSquares := squareCodeLine_regex.FindAllString(sourceCode, -1)
	
	//Creating all the required matrices
	IDMatrix := CreateIDString(sourceCode)
	allTriangleIDs := AllShapeIDs(allTriangles)
	allSquareIDs := AllShapeIDs(allSquares)
	
	prologOutput = "/* processing input file " + fileName
	prologOutput += "\n   Lexical and Syntax analysis passed\n   Generating Prolog Code */\n\n"
	
	s_count := 0
	t_count := 0
	var relevantTest []string

	//The large list of different triangle types to output
	triangleTypeList := [10]string{"line", "triangle", "vertical", "horizontal", "equilateral", "isosceles", "right", "scalene", "acute", "obtuse"}

	if len(allTests) == 0 {
		prologOutput += "Error: no tests in given code\n\n"
	}

	//This is almost the same as the previous scheme loop, but has different parts for square or triangle due to the large number
	//	of triangle types that need to be output for the prolog part
	for i := 0; i < len(allTests); i++ {
		if testSquare_regex.MatchString(allTests[i]) {
			prologOutput += " /* Processing test(square, "
			relevantTest = allSquareIDs[s_count]
			for m := 0; m < len(relevantTest); m++ {
				if relevantTest[m] != "" {
					prologOutput += relevantTest[m]
					prologOutput += ", "
				}
			}
			prologOutput = prologOutput[:(len(prologOutput)-2)]
			prologOutput += ") */\n query(square("
			s_count++
			for index, ID := range relevantTest {
				if relevantTest[index] == "" {
					break
				}
				prologOutput += "point2d("
				noIDCount := 0
				for j := 0; j < len(IDMatrix); j++ {
					if IDMatrix[j][0] == ID {
						prologOutput += IDMatrix[j][1]
						prologOutput += ", "
						prologOutput += IDMatrix[j][2]
						prologOutput += "), "
						break
					}
					noIDCount++
				}
				if noIDCount == len(IDMatrix) {
					prologOutput = prologOutput[:(len(prologOutput)-8)]
					prologOutput += "(Error ID "
					prologOutput += ID
					prologOutput += " not defined)  "
				}
			}
			prologOutput = prologOutput[:(len(prologOutput)-2)]
			prologOutput += ")).\n\n"

		//This is where the two differ - we have extra steps in here to get the triangle data to look right with the 10 different types
		} else if testTriangle_regex.MatchString(allTests[i]) {
			prologOutput += " /* Processing test(triangle, "
			relevantTest = allTriangleIDs[t_count]
			for m := 0; m < len(relevantTest); m++ {
				if relevantTest[m] != "" {
					prologOutput += relevantTest[m]
					prologOutput += ", "
				}
			}
			prologOutput = prologOutput[:(len(prologOutput)-2)]
			prologOutput += ") */\n"
			for k := 0; k < 10; k++ {
				prologOutput += " query("
				prologOutput += triangleTypeList[k]
				prologOutput += "("
		
				for index, ID := range relevantTest {
					//If there is no more tests, adding a space in to match formatting where needed
					if relevantTest[index] == "" {
						prologOutput = prologOutput[:(len(prologOutput) - 4)] + " " + prologOutput[(len(prologOutput) - 4):]
						break
					}
					prologOutput += "point2d("
					noIDCount := 0
					for j := 0; j < len(IDMatrix); j++ {
						if IDMatrix[j][0] == ID {
							prologOutput += IDMatrix[j][1]
							if index == len(relevantTest) - 1 {
								prologOutput += ", "
							} else {
								prologOutput += ","
							}
							prologOutput += IDMatrix[j][2]
							prologOutput += "), "
							break
						}
						noIDCount++
					}
					if noIDCount == len(IDMatrix) {
					prologOutput = prologOutput[:(len(prologOutput)-8)]
					prologOutput += "(Error ID "
					prologOutput += ID
					prologOutput += " not defined)  "
					}
				}
				prologOutput = prologOutput[:(len(prologOutput)-2)]
				if k == 9 {
					prologOutput += ")).\n\n"
				} else {
					prologOutput += ")).\n"
				}
			}
			t_count++
		}
	}

	prologOutput = prologOutput[:(len(prologOutput)-1)]
	prologOutput += "\n\n /* Query Processing */\n writeln(T) :- write(T), nl.\n"
	prologOutput += " main:- forall(query(Q), Q-> (writeln('yes')) ; (writeln('no'))),\n      halt."
	fmt.Println(prologOutput)
}
