// Command gotraceanalyzer turns golang tracebacks into useful summaries.
//
// See README.md for instructions.
package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"hash"
	"io"
	"os"
	"regexp"
	"sort"
)

func main() {
	var input io.Reader
	if len(os.Args) == 2 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fatal(err)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	var (
		parser = NewParser(input)
		index  = map[string]*UniqueGoroutine{}
		list   = UniqueGoroutines{}
		total  = 0
	)
	for {
		routine := &Goroutine{}
		if err := parser.ReadRoutine(routine); err == io.EOF {
			break
		} else if err != nil {
			fatal(err)
		}
		total++
		hash := routine.Hash()
		if unique, ok := index[hash]; !ok {
			unique := &UniqueGoroutine{routine, 1}
			index[hash] = unique
			list = append(list, unique)
		} else {
			unique.Count++
		}
	}
	sort.Sort(list)
	fmt.Printf("Showing %d unique goroutines out of %d total:\n\n", len(list), total)
	for _, unique := range list {
		fmt.Printf("%d goroutines [%s]:\n%s\n", unique.Count, unique.State(), unique.Trace())
	}
}

func fatal(err interface{}) {
	fmt.Printf("error: %s\n", err)
	os.Exit(1)
}

func NewParser(r io.Reader) *Parser {
	return &Parser{scanner: bufio.NewScanner(r), first: true}
}

type Parser struct {
	scanner *bufio.Scanner
	first   bool
	buf     Goroutine
	err     error
}

var (
	headerRegexp = regexp.MustCompile("goroutine \\d+ \\[(.+)\\]")
	fileRegexp   = regexp.MustCompile("^\t(.+:\\d+)")
)

func (p *Parser) ReadRoutine(r *Goroutine) error {
	if p.err != nil {
		return p.err
	}
	if !p.first {
		*r = p.buf
	}
	for p.scanner.Scan() {
		line := p.scanner.Text()
		if line == "" {
			continue
		}

		if m := headerRegexp.FindStringSubmatch(line); len(m) == 2 {
			p.buf = Goroutine{state: m[1], hash: sha1.New()}
			_, p.err = p.buf.hash.Write([]byte(m[1]))
			if p.err != nil {
				return p.err
			}

			if p.first {
				*r = p.buf
				p.first = false
			} else {
				return nil
			}
		} else if m := fileRegexp.FindStringSubmatch(line); len(m) == 2 {
			_, p.err = r.hash.Write([]byte(m[1]))
			if p.err != nil {
				return p.err
			}
		}
		r.trace += line + "\n"
	}
	p.err = p.scanner.Err()
	if p.err != nil {
		return p.err
	}
	p.err = io.EOF
	return nil
}

type Goroutine struct {
	state string
	trace string
	hash  hash.Hash
}

func (g *Goroutine) State() string {
	return g.state
}

func (g *Goroutine) Trace() string {
	return g.trace
}

func (g *Goroutine) Hash() string {
	return fmt.Sprintf("%x", g.hash.Sum(nil))
}

type UniqueGoroutine struct {
	*Goroutine
	Count int
}

type UniqueGoroutines []*UniqueGoroutine

func (u UniqueGoroutines) Len() int {
	return len(u)
}

func (u UniqueGoroutines) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u UniqueGoroutines) Less(i, j int) bool {
	return u[i].Count > u[j].Count
}
