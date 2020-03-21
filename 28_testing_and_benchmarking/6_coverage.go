package main

/*

go help testflag

...

	-cover
	    Enable coverage analysis.
	    Note that because coverage works by annotating the source
	    code before compilation, compilation and test failures with
	    coverage enabled may report line numbers that don't correspond
	    to the original sources.

...

`go test -cover`

If I run this in the 5_benchmarking directory...

"

PASS
coverage: 100.0% of statements
ok      _/Users/richardchernanko/Development/go-udemy/28_testing_and_benchmarking/5_benchmarking        0.006s


Also in the go help test...

"

	-coverprofile cover.out
	    Write a coverage profile to the file after all tests have passed.
	    Sets -cover.


"

Writing the coverage to a file...

go test -coverprofile richard

"

So, now, in the 5_benchmarking directory, I have a file called richard.txt

Its contents look like this:

mode: set
/Users/richardchernanko/Development/go-udemy/28_testing_and_benchmarking/5_benchmarking/main.go:9.29,11.2 1 1


Not massively useful at the moment. But if I run it with `go tool cover -html=richard`, it will open up in the browser. Pretty cool

For more documentation on go tools...

go tool cover

"

richardchernanko > ~ > $ go tool cover
Usage of 'go tool cover':
Given a coverage profile produced by 'go test':
	go test -coverprofile=c.out

Open a web browser displaying annotated source code:
	go tool cover -html=c.out

Write out an HTML file instead of launching a web browser:
	go tool cover -html=c.out -o coverage.html

Display coverage percentages to stdout for each function:
	go tool cover -func=c.out

Finally, to generate modified source code with coverage annotations
(what go test -cover does):
	go tool cover -mode=set -var=CoverageVariableName program.go

Flags:
  -V	print version and exit
  -func string
    	output coverage profile information for each function
  -html string
    	generate HTML representation of coverage profile
  -mode string
    	coverage mode: set, count, atomic
  -o string
    	file for output; default: stdout
  -var string
    	name of coverage variable to generate (default "GoCover")

  Only one of -html, -func, or -mode may be set.

"

*/
