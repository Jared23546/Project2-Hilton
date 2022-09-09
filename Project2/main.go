package main

// import necessary classes
import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// this function is called to convert the text file to an array.
// this function returns the array and an error log
func createArray(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	x := 0
	for scanner.Scan() {
		x++
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// this function shuffles the array
func shuffleArr(lines []string) []string {
	rand.Seed(time.Now().UnixNano())
	for i := range lines {
		j := rand.Intn(i + 1)
		lines[i], lines[j] = lines[j], lines[i]
	}
	return lines
}

func main() {
	// convert the text file into an array
	lines, err := createArray("Names.txt")
	if err != nil {
		log.Fatalf("createArray: %s", err)
	}

	// find the number of students
	numPeople := len(lines)

	// shuffle the array
	lines = shuffleArr(lines)

	//get input from user
	var teamSize int
	fmt.Println("How many individuals per team? ")
	fmt.Scanln(&teamSize)

	numTeams := numPeople / teamSize
	remainder := numPeople % teamSize

	// this print is for visibility and is not necessary
	fmt.Println("\n \t", numPeople, "people and ", numTeams, " teams of ", teamSize, " with ", remainder, " people left over.")

	// create a map of teams ( key = (team number) and value = (list of team members))
	teams := make(map[int][]string)

	// loop through the array splitting it into teams
	var start int
	var end int
	for i := 0; i < numTeams; i++ {
		start = i * teamSize
		end = start + teamSize

		teams[i+1] = lines[start:end]
	}

	// take care of extra students if there are any
	if remainder != 0 {
		if remainder == (teamSize - 1) {
			i := len(teams) + 1
			teams[i] = lines[end : end+remainder]
		} else {
			for i := 0; i < remainder; i++ {
				teamToAddTo := i%len(teams) + 1
				teams[teamToAddTo] = append(teams[teamToAddTo], lines[end])
				end++
			}
		}
	}

	// print out the map of teams
	for key, value := range teams {
		fmt.Println("\nTeam ", key, value)
	}
}
