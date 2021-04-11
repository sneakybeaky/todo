package todo

import (
	"bufio"
	"fmt"
	"os"
)

/*TransactionLogger records events such that a todo list can be recreated
We’ll assume that your transaction log will be read only when your service is restarted or otherwise needs to recover its state, and that it’ll be read from top to bottom, sequentially replaying each event. It follows that your transaction log will consist of an ordered list of mutating events. For speed and simplicity, a transaction log is also generally append-only, so when a record is deleted from your key-value store, for example, a “delete” is recorded in the log.
*/
type TransactionLogger interface {
	WriteAdd(todo Todo)
	ReadEvents() (<-chan Event, <-chan error)
	Err() <-chan error
}

type Event struct {
	Sequence  uint64    // A unique record ID
	EventType EventType // The action taken
	Payload   string    // The data from this event
}

type EventType byte

const (
	_                  = iota // iota == 0; ignore the zero value
	EventAdd EventType = iota // iota == 1
)

// FileTransactionLogger is a naive file based transaction logger.
//First, for simplicity, the log will be written in plain text; a binary, compressed format might be more time and space efficient, but we can always optimize later. Second, each entry will be written on its own line; this will make it much easier to read the data later.
// - Sequence number : A unique record ID, in monotonically increasing order.
// - Event type : descriptor of the type of action taken; can be ADD or DELETE.
// - Payload : the data from the event. For an ADD this would be the TODO
type FileTransactionLogger struct {
	events       chan<- Event // Write-only channel for sending events
	errors       <-chan error // Read-only channel for receiving errors
	lastSequence uint64       // The last used event sequence number
	file         *os.File     // The location of the transaction log
}

func (l FileTransactionLogger) WriteAdd(todo Todo) {
	l.events <- Event{EventType: EventAdd, Payload: todo.Title}
}

func (l FileTransactionLogger) Err() <-chan error {
	return l.errors
}

func NewFileTransactionLogger(filename string) (*FileTransactionLogger, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return nil, fmt.Errorf("cannot open transaction log file: %w", err)
	}

	return &FileTransactionLogger{file: file}, nil
}

func (l FileTransactionLogger) Run() {
	events := make(chan Event, 16) // Make an events channel
	l.events = events

	errors := make(chan error, 1) // Make an errors channel
	l.errors = errors

	go func() {
		for e := range events { // Retrieve the next Event

			l.lastSequence++ // Increment sequence number

			_, err := fmt.Fprintf( // Write the event to the log
				l.file,
				"%d\t%d\t%s\n",
				l.lastSequence, e.EventType, e.Payload)

			if err != nil {
				errors <- err
				return
			}
		}
	}()
}

func (l FileTransactionLogger) ReadEvents() (<-chan Event, <-chan error) {
	scanner := bufio.NewScanner(l.file) // Create a Scanner for l.file
	outEvent := make(chan Event)        // An unbuffered Event channel
	outError := make(chan error, 1)     // A buffered error channel

	go func() {
		var e Event

		defer close(outEvent) // Close the channels when the
		defer close(outError) // goroutine ends

		for scanner.Scan() {
			line := scanner.Text()

			if _, err := fmt.Sscanf(line, "%d\t%d\t%s", &e.Sequence, &e.EventType, &e.Payload); err != nil {
				outError <- fmt.Errorf("input parse error: %w", err)
				return
			}

			// Sanity check! Are the sequence numbers in increasing order?
			if l.lastSequence >= e.Sequence {
				outError <- fmt.Errorf("transaction numbers out of sequence")
				return
			}

			l.lastSequence = e.Sequence // Update last used sequence #

			outEvent <- e // Send the event along
		}

		if err := scanner.Err(); err != nil {
			outError <- fmt.Errorf("transaction log read failure: %w", err)
			return
		}
	}()

	return outEvent, outError
}
