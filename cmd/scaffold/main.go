package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		usage(nil)
	}

	problemNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage(err)
	}

	dirName := fmt.Sprintf("%02d", problemNumber)

	log.Println("Creating package directory...")
	err = os.Mkdir(dirName, 0744)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Creating files...")
	{
		mainT := template.Must(template.New("main").Parse(mainTemplate))
		f, err := os.Create(path.Join(dirName, "main.go"))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer f.Close()
		err = mainT.Execute(f, dirName)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	log.Println("Creating test file...")
	testT := template.Must(template.New("test").Parse(testTemplate))
	{
		f, err := os.Create(path.Join(dirName, "main_test.go"))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer f.Close()
		err = testT.Execute(f, dirName)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	log.Println("Creating empty input file...")
	{
		f, err := os.Create(path.Join(dirName, "input"))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer f.Close()
	}

	log.Println("Creating empty puzzle file...")
	{
		f, err := os.Create(path.Join(dirName, "puzzle.txt"))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer f.Close()
	}

	log.Println("All done!")
}

func usage(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n\n", err.Error())
	}
	fmt.Fprintf(os.Stderr, "usage: %s <problem number:int>\n", os.Args[0])
	os.Exit(1)
}

const testTemplate = `
package main

import (
	"testing"

	h "github.com/jmhobbs/advent-of-code-2018/helpers"
)

func TestTodo(t *testing.T) {
	h.Assert(t, false, "false should be true")
}
`

const mainTemplate = `
package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error reading input:", err)
	}

	fmt.Println(todo())
}

func todo() string {
	return "replace-me"
}
`
