package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

var USAGE = `USAGE: note [subcommand [args]]
    subcommands: 
        * new : add the args to the notes for today
        * today : print today's notes
        * ls : print the dates of previous notes
        * <YYYY.MM.DD> : print the notes from the given date
    
As a bare command, note will read from stdin and add to today's note.
`

// FatalPrintln is like log.Fatal but with fmt.Println instead
func FatalPrintln(err ...interface{}) {
	fmt.Println(err...)
	os.Exit(1)
}

// WriteNote creates or opens an existing file and writes the new notes to file
func WriteNote(date string, notes string) {
	note, err := os.OpenFile(path.Join(os.Getenv("NOTES_HOME"), date), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		FatalPrintln("ERROR: could note create note: ", err)
	}
	defer note.Close()
	_, err = note.WriteString(notes + "\n")
	if err != nil {
		FatalPrintln("ERROR: could note create note: ", err)
	}
}

// ReadNote reads an existing file and prints it to stdout
func ReadNote(date string) {
	notes, err := ioutil.ReadFile(path.Join(os.Getenv("NOTES_HOME"), date))
	if err != nil {
		FatalPrintln("ERROR: could note read note: ", err)
	}
	fmt.Println(date)
	fmt.Println(string(notes))
}

// ListNotes prints a listing of notes
func ListNotes() {
	f, err := os.Open(os.Getenv("NOTES_HOME"))
	if err != nil {
		FatalPrintln("ERROR: could not get notes listing: ", err)
	}
	defer f.Close()
	list, err := f.Readdirnames(-1)
	if err != nil {
		FatalPrintln("ERROR: could not get notes listing: ", err)
	}

	fmt.Println("NOTES:")
	for _, name := range list {
		fmt.Println(name)
	}
}

func main() {
	subcommand := os.Args
	notesHomeDir := os.Getenv("NOTES_HOME")
	if notesHomeDir == "" {
		FatalPrintln("ERROR: $NOTES_HOME environment variable not set, please set this variable before continuing")
	}
	if err := os.MkdirAll(notesHomeDir, fs.ModeDir|fs.ModePerm); err != nil {
		FatalPrintln("ERROR: could not start note with error: ", err)
	}

	date := time.Now().Format("2006.01.02")

	if len(subcommand) > 1 {
		switch subcommand[1] {
		case "new":
			if len(subcommand) > 2 {
				WriteNote(date, strings.Join(subcommand[2:], " "))
				ReadNote(date)
			} else {
				fmt.Println(USAGE)
				return
			}
		case "today":
			ReadNote(date)
		case "ls":
			ListNotes()
		default:
			if _, err := time.Parse("2006.01.02", subcommand[1]); err != nil {
				fmt.Println(USAGE)
				return
			}
			ReadNote(subcommand[1])
		}
		return
	}

	s := bufio.NewScanner(os.Stdin)
	notes := ""
	for s.Scan() {
		notes += s.Text() + "\n"
	}

	WriteNote(date, notes)
	ReadNote(date)

}
