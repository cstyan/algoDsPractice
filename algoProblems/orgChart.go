// input is something like
// [
//     'a,b,c,d',
//     'b,e,f',
// 	   'c',
//     'd,g,h'
// ]
//
// output should be
// a
// ....b
// ........e
// ........f
// ....c
// ....d
// ........g
// ........h

package main

import (
	"fmt"
	"strings"
)

// NOTE: solution makes the asumption that the array of manager/report strings
// is in hierarchical order, meaning if b, c, and d report to a, the string of
// a's reports should be before all the strings of it's reports that are managers
// in the array of manager/report strings, see example input above

func makeManagerMap(input []string) (map[string]string, string) {
	managerMap := make(map[string]string)
	// we need to keep track of which manager is at the top of the org.
	// since maps are unordered
	var startingManager string = ""
	for _, managerReportsString := range input {
		// this should split on just the first delimeter
		// 'a', 'b,c,d'
		split := strings.SplitN(managerReportsString, ",", 2)

		if startingManager == "" {
			startingManager = split[0]
		}

		// if there is ever a manager who doesn't have any direct reports
		if len(split) == 1 {
			managerMap[split[0]] = ""

		} else {
			managerMap[split[0]] = split[1]
		}
	}
	return managerMap, startingManager
}

func printOrgChart(input map[string]string, startingManager string) {
	numDelimeter := 0
	for {
		currentManager(input, startingManager, numDelimeter)
		// because we're calling currentManager recursively
		// we should be finished when it returns here
		break
	}
}

func currentManager(managerMap map[string]string, manager string, numDelimeter int) {
	printDelimeter(numDelimeter)
	fmt.Println(manager)

	reports := managerMap[manager]

	// split all of the reports so we can iterate over them
	split := strings.Split(reports, ",")
	// early exit case, if current manager has no direct reports for some reason
	if split[0] == "" {
		return
	}

	for _, report := range split {
		numDelimeter = numDelimeter + 1
		if _, ok := managerMap[report]; ok {
			currentManager(managerMap, report, numDelimeter)
		} else {
			printDelimeter(numDelimeter)
			fmt.Println(report)
		}
		numDelimeter = numDelimeter - 1
	}
}

func printDelimeter(numDelimeter int) {
	delimeter := "...."
	for i := 0; i < numDelimeter; i++ {
		fmt.Print(delimeter)
	}
}

func main() {
	managerReports := []string{
		"a,b,c,d",
		"b,e,f",
		"c",
		"d,g,h",
	}
	managers, startingManager := makeManagerMap(managerReports)
	printOrgChart(managers, startingManager)
}
