package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

const YEAR = 2025

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dir, err := os.Executable()
	check(err)
	os.Chdir(filepath.Dir(dir))
	dayPtr := flag.CommandLine.Int("d", 0, "Day")
	partPtr := flag.CommandLine.Int("p", 0, "Part")
	answerPtr := flag.CommandLine.Int64("a", 0, "Answer")
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "No args provided.")
		os.Exit(1)
	}
	_ = flag.CommandLine.Parse(os.Args[2:])
	if *dayPtr == 0 {
		*dayPtr = dayNow()
	}
	switch os.Args[1] {
	case "fetch":
		session := readSession()
		input := fetch(YEAR, *dayPtr, session)
		write(*dayPtr, input)
	case "tmpl":
		tmpl(*dayPtr)
	case "run":
		run(*dayPtr)
	case "test":
		if *partPtr == 0 {
			test(*dayPtr, 1)
			test(*dayPtr, 2)
			return
		}
		test(*dayPtr, *partPtr)
	case "bench":
		if *partPtr == 0 {
			bench(*dayPtr, 1)
			bench(*dayPtr, 2)
			return
		}
		bench(*dayPtr, *partPtr)
	case "upload":
		session := readSession()
		upload(*dayPtr, *partPtr, *answerPtr, session)
	default:
		fmt.Fprintln(os.Stderr, "Not IMPLEMENTED")
		os.Exit(1)
	}
}

func dayNow() int {
	now := time.Now()
	_, _, date := now.Date()
	return date
}

func readSession() string {
	dat, err := os.ReadFile(".session")
	check(err)
	return string(dat)
}

func fetch(year int, day int, session string) []byte {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	check(err)
	req.AddCookie(&http.Cookie{Name: "session", Value: session})
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	check(err)
	return body
}

func write(day int, input []byte) {
	fileDir := fmt.Sprintf("./input/day%d/", day)
	err := os.MkdirAll(fileDir, 0755)
	check(err)
	err = os.WriteFile(string(fileDir+"input"), bytes.TrimSuffix(input, []byte("\n")), 0644)
	check(err)
}

func tmpl(day int) {
	fileDir := fmt.Sprintf("./src/day%d/", day)
	err := os.MkdirAll(fileDir, 0755)
	check(err)
	test_tmpl, err := template.ParseFiles("tmpl/test.tmpl")
	check(err)
	main_tmpl, err := template.ParseFiles("tmpl/sol.tmpl")
	check(err)
	test_file, err := os.Create(fileDir + "main_test.go")
	check(err)
	defer test_file.Close()
	main_file, err := os.Create(fileDir + "main.go")
	check(err)
	defer main_file.Close()
	test_tmpl.Execute(test_file, struct{ Day int }{day})
	main_tmpl.Execute(main_file, struct{ Day int }{day})
}

func run(day int) {
	cmd := exec.Command("go", "run", fmt.Sprintf("./src/day%d/main.go", day))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func test(day int, part int) {
	cmd := exec.Command("go", "test", fmt.Sprintf("-run=TestPart%d", part), "-v", fmt.Sprintf("./src/day%d", day))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func bench(day int, part int) {
	cmd := exec.Command("go", "test", fmt.Sprintf("-bench=BenchmarkPart%d", part), fmt.Sprintf("./src/day%d", day))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func upload(day int, part int, answer int64, session string) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", YEAR, day)
	data := fmt.Sprintf("level=%d&answer=%d", part, answer)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	check(err)
	req.AddCookie(&http.Cookie{Name: "session", Value: session})
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	check(err)
	if strings.Contains(string(body), "That's the right answer") {
		fmt.Println("Correct!")
	} else if strings.Contains(string(body), "That's not the right answer") {
		fmt.Println("Wrong answer")
	} else if strings.Contains(string(body), "You gave an answer too recently") {
		fmt.Println("Too recent - wait before submitting again.")
	} else {
		fmt.Println("Wrong Day :)")
	}
}
