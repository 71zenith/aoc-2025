package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"path/filepath"
)

const YEAR = 2024
const TEMP_MAIN = "tmpl/sol.tmpl"
const TEMP_TEST = "tmpl/test.tmpl"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dir, err := os.Executable()
	check(err)
	os.Chdir(filepath.Dir(dir))
	fetchFlag := flag.NewFlagSet("fetch", flag.ExitOnError)
	dayFetch := fetchFlag.Int("d", 1, "Day")

	tmplFlag := flag.NewFlagSet("tmpl", flag.ExitOnError)
	dayTmpl := tmplFlag.Int("d", 1, "Day")

	runFlag := flag.NewFlagSet("run", flag.ExitOnError)
	dayRun := runFlag.Int("d", 1, "Day")

	testFlag := flag.NewFlagSet("test", flag.ExitOnError)
	dayTest := testFlag.Int("d", 1, "Day")
	partTest := testFlag.Int("p", 1, "Part")

	benchFlag := flag.NewFlagSet("bench", flag.ExitOnError)
	dayBench := benchFlag.Int("d", 1, "Day")
	partBench := benchFlag.Int("p", 1, "Part")

	uploadFlag := flag.NewFlagSet("upload", flag.ExitOnError)
	dayUpload := uploadFlag.Int("d", 1, "Day")
	partUpload := uploadFlag.Int("p", 1, "Part")
	answerUpload := uploadFlag.Int64("a", 1, "Answer")

	switch os.Args[1] {
	case "fetch":
		session, err := readSession()
		check(err)
		input := fetch(YEAR, *dayFetch, session)
		write(*dayFetch, input)
	case "tmpl":
		tmpl(*dayTmpl)
	case "run":
		run(*dayRun)
	case "test":
		test(*dayTest, *partTest)
	case "bench":
		bench(*dayBench, *partBench)
	case "upload":
		session, err := readSession()
		check(err)
		upload(*dayUpload, *partUpload, *answerUpload, session)
	}
}

func readSession() (string, error) {
    dat, err := os.ReadFile(".session")
	return string(dat), err
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
	fileDir := fmt.Sprintf("input/day%d/", day)
	err := os.MkdirAll(fileDir, 0755)
	check(err)
	err = os.WriteFile(string(fileDir + "input"), input, 0644)
	check(err)
}

func tmpl(day int) {
	fileDir := fmt.Sprintf("src/day%d/", day)
	err := os.MkdirAll(fileDir, 0755)
	check(err)
	test_tmpl, err := template.ParseFiles(TEMP_TEST)
	check(err)
	main_tmpl, err := template.ParseFiles(TEMP_MAIN)
	check(err)
	test_file, err := os.Create(fileDir + "test.go")
	check(err)
	defer test_file.Close()
	main_file, err := os.Create(fileDir + "main.go")
	check(err)
	defer main_file.Close()
	test_tmpl.Execute(test_file, struct{Day int} {day});
	main_tmpl.Execute(main_file, struct{Day int} {day});
}

func run(day int) {
    cmd := exec.Command("go", "run", fmt.Sprintf("src/day%d/main.go", day))
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}

func test(day int, part int) {
    cmd := exec.Command("go", "test", fmt.Sprintf("src/day%d", day), "-v" , "-run", fmt.Sprintf("TestPart%d", part))
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}

func bench(day int, part int) {
    cmd := exec.Command("go", "test", "-bench=", fmt.Sprintf("BenchmarkPart%d", part), fmt.Sprintf("src/day%d", day))
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
        fmt.Println("Too recent - wait before submitting again")
    }
}
