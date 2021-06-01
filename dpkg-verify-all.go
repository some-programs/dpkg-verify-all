package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

// output .
type output struct {
	Package string   // package name
	Output  []string // command output
	Error   bool     // error executing package
}

func (o output) String() string {
	return fmt.Sprintf("%v (err:%v): %v", o.Package, o.Error, o.Output)
}

func main() {
	var (
		nWorkers int
		verbose  int
	)
	flag.IntVar(&nWorkers, "workers", runtime.NumCPU()+1, "number of workers")
	flag.IntVar(&verbose, "verbose", 0, "verbose output 0=none, 5=max")
	flag.Parse()
	cmd := exec.Command("dpkg-query", "--show", "--showformat", "${binary:Package}\n")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(out), "\n")
	var packages []string
	for _, v := range lines {
		if v != "" {
			packages = append(packages, v)
		}
	}
	if verbose > 0 {
		log.Printf("Verifying %d packages using %d workers...", len(packages), nWorkers)
	}
	outputC := make(chan output, 100)
	var reporterWg sync.WaitGroup
	reporterWg.Add(1)
	go func() {
		var errors, outputs []output
		for e := range outputC {
			if verbose > 2 {
				log.Println(e)
			}
			if e.Error {
				errors = append(errors, e)
			} else {
				outputs = append(outputs, e)
			}
		}
		if len(outputs) > 0 {
			fmt.Printf("\n%d package verifications returned non empty output\n", len(outputs))
			for _, v := range outputs {
				for _, line := range v.Output {
					if strings.TrimSpace(line) != "" {
						fmt.Printf("%s: %s\n", v.Package, line)
					}
				}
			}
		}
		if len(errors) > 0 {
			fmt.Printf("\n dpkg --verify finished with non zero exit %d times:\n", len(errors))
			for _, v := range errors {
				for _, line := range v.Output {
					if strings.TrimSpace(line) != "" {
						fmt.Printf("%s: %s\n", v.Package, line)
					}
				}
			}
		}
		if len(outputs) > 0 || len(errors) > 0 {
			os.Exit(1)
		}
		reporterWg.Done()
	}()
	var workersWg sync.WaitGroup
	c := make(chan string, 0)
	for i := 0; i < nWorkers; i++ {
		workersWg.Add(1)
		go func() {
			for packageName := range c {
				cmd := exec.Command("dpkg", "--verify", packageName)
				out, err := cmd.CombinedOutput()
				if len(out) > 0 {
					outputC <- output{
						Package: packageName,
						Error:   err != nil,
						Output:  strings.Split(string(out), "\n"),
					}
				}
			}
			workersWg.Done()
		}()
	}
	n := 0
	for _, v := range packages {
		c <- v
		if verbose > 1 {
			n++
			if n%100 == 0 {
				log.Printf("Started verification of %d packagages...", n)
			}
		}
	}
	close(c)
	workersWg.Wait()
	close(outputC)
	reporterWg.Wait()
}
