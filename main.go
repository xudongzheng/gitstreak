package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Define the daily character.
const char = "■"

// Define the date format for parsing.
const timeFormat = "2006-01-02"

// Store the author information if the graph is to be limited to one author.
var author = flag.String("author", "", "limit statistics to a specific author")

// Map each day to the number of commits.
var counter = make(map[int]int)

// Determine the earliest commit date, which helps calculate the longest streak.
var firstDate = time.Now()

func token(t time.Time) int {
	return t.Year()*10000 + int(t.Month())*100 + t.Day()
}

func handlePath(path string) error {
	// Initiate git log command to fetch commit timestamps. Set current working
	// directory to src and pipe output to buffer. TODO consider --no-merges.
	cmd := exec.Command("git", "log", "--pretty=format:%aI %ae")
	cmd.Dir = path

	// Obtain pipe for stdout and run the command. Defer wait, which will close
	// the stdout pipe when the process exits.
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}
	defer cmd.Wait()

	// Initialize buffered reader for stdout.
	buf := bufio.NewReader(stdout)

	var authors []string
	if *author != "" {
		authors = strings.Split(*author, ",")
		for i, a := range authors {
			authors[i] = strings.Trim(a, " ")
		}
	}

	// Read commit timestamps line by line. Parse in ISO8601 format, which is
	// virtually identical to RFC3339.
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF && line == "" {
			break
		} else if (err != nil && err != io.EOF) || line == "" {
			return err
		}
		line = strings.TrimSuffix(line, "\n")
		pieces := strings.SplitN(line, " ", 2)
		if *author != "" {
			found := false
			for _, a := range authors {
				if pieces[1] == a {
					found = true
					continue
				}
			}
			if !found {
				continue
			}
		}
		parse, err := time.Parse(timeFormat, pieces[0][:10])
		if err != nil {
			return err
		}
		counter[token(parse)]++
		if parse.Before(firstDate) {
			firstDate = parse
		}
	}

	return nil
}

func main() {
	// Parse command line flags.
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatal("please specify one (or multiple) repositories")
	}

	// Determine the graph range. The graph should start at the beginning of a
	// week on a Sunday.
	endDate := time.Now()
	startDate := endDate.AddDate(-1, 0, 0)
	startDate = startDate.AddDate(0, 0, int(time.Sunday-startDate.Weekday()))

	// Count commits for each repository. Use arguments as repository list.
	for _, value := range flag.Args() {
		if err := handlePath(value); err != nil {
			log.Fatal(err)
		}
	}

	// Define the colors for the graph. The first is for day with no commits,
	// with the remaining four being each responsible for a "quartile".
	legend := []int{255, 6, 2, 11, 9}

	// Calculate the max maximum streak and current streak. This accounts for
	// the life of the repository.
	streakLiveActive := true
	var streakMax, streakLive, streakLoop int
	var streakStart time.Time
	for t := endDate; !t.Before(firstDate); t = t.AddDate(0, 0, -1) {
		// If a streak is ongoing, no additional handling is necessary.
		count := counter[token(t)]
		if count > 0 {
			streakLoop++
			continue
		}

		// Set the current streak if it ends on either endDate or the day
		// before, which is the same way Github would calculate it.
		if streakLiveActive && t != endDate {
			streakLive = streakLoop
			streakLiveActive = false
		}

		// Set the current streak if it is the longest.
		if streakLoop > streakMax {
			streakMax = streakLoop
			streakStart = t.AddDate(0, 0, 1)
		}

		streakLoop = 0
	}

	// Calculate the total number of commits and the highest number of commits
	// in a day. The latter is used to determine the color scale for the graph.
	// Unlike the previous loop, this only accounts for the past year.
	var countTotal, countMax int
	for t := endDate; !t.Before(startDate); t = t.AddDate(0, 0, -1) {
		count := counter[token(t)]
		countTotal += count
		if count > countMax {
			countMax = count
		}
	}

	// Make sure that data exists.
	if countTotal == 0 {
		log.Fatal("unable to find commit data")
	}

	// Calculate color for each day of the year. The upper limit for the graph
	// is 53 weeks and a day, hence the capacity.The quartiles are simply the
	// day with the most commits divided into 4 pieces. To calculate the color,
	// multiply by 4, divide by the countMax, and use the ceiling integer trick.
	colors := make([]int, 0, 53*7+1)
	for t := endDate; !t.Before(startDate); t = t.AddDate(0, 0, -1) {
		count := counter[token(t)]
		colors = append(colors, legend[(count*4+countMax-1)/countMax])
	}

	// Determine the number of weeks (columns) to display. Divide the number of
	// days by 7 and take the ceiling. Use this nice integer division trick.
	weeks := (len(colors) + 6) / 7

	// Initialize 2D array to store chart data and fill it with data. This is
	// very straightforward because it always starts on a Sunday. We must
	// "invert" the key because counter goes from present to past, whereas chart
	// goes from past to present.
	formatter := func(value int) string {
		return "\x1b[38;5;" + strconv.Itoa(value) + "m" + char + " \x1b[0m"
	}
	chart := make([][7]string, weeks)
	for key, value := range colors {
		inv := len(colors) - 1 - key
		chart[inv/7][inv%7] = formatter(value)
	}
	fmt.Println()

	// Print the months of the year.
	fmt.Print(strings.Repeat(" ", 11))
	for t := startDate; t.Before(endDate); t = t.AddDate(0, 0, 7) {
		if t.Day() <= 7 {
			fmt.Print(t.Month().String()[:3] + " ")
		} else if t.Day() > 14 {
			fmt.Print("  ")
		}
	}
	fmt.Println()

	// Print each day of the week.
	for i := time.Sunday; i <= time.Saturday; i++ {
		fmt.Printf("%10s ", i.String())
		for j := 0; j < len(chart); j++ {
			if len(chart[j][i]) > 0 {
				fmt.Print(chart[j][i])
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	// Print legend and statistics.
	streakStartStr := streakStart.Format("2006-01-02")
	legendStr := "Total Commits: " + strconv.Itoa(countTotal) + " | "
	legendStr += "Current Streak: " + strconv.Itoa(streakLive) + " days | "
	legendStr += "Longest Streak: " + strconv.Itoa(streakMax) + " days (from " + streakStartStr + ") | "
	legendStr += "Less "
	for _, value := range legend {
		legendStr += "\x1b[38;5;" + strconv.Itoa(value) + "m" + char + " \x1b[0m"
	}
	legendStr += "More"
	spaces := weeks*2 + 10 - len(legendStr) + 78
	fmt.Println(strings.Repeat(" ", spaces) + legendStr)
	fmt.Println()
}
