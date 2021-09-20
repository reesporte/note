# note
a simple command line tool for simple note taking

## installation
`go get -u github.com/reesporte/note && go install github.com/reesporte/note@latest`

## usage
set `$NOTES_HOME` to the directory you wish to store your notes in, then `note [subcommand [args]]`.
`note` will automatically setup the directory if it does not exist.

subcommands:   
     * new : add the arguments to new to the notes for today   
     * today : print today's notes   
     * <YYYY.MM.DD> : print the notes from the given date   
 
as a bare command, note will read from stdin and add to today's note.

## examples
```
$ note new this is a new note 
YYYY.MM.DD
this is a new note

$ note new "this is another note"
YYYY.MM.DD
this is a new note
this is another note

$ echo "hey there\!" | note
YYYY.MM.DD
this is a new note
this is another note
hey there!

$ note today 
YYYY.MM.DD
this is a new note
this is another note
hey there!

```
