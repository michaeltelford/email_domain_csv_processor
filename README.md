
# Customer Importer

Package `customerimporter` reads from the given `customers.csv` file and returns a sorted (data structure of your choice) of email domains along with the number of customers with e-mail addresses for each domain. Any errors should be logged (or handled). Performance matters (this is only ~3k lines, but *could* be 1m lines or run on a small machine).

## Usage

Place the `customerimporter.go` file in your `$GOPATH` and import it from somewhere else e.g. a `main.go` file.

There is a `main.go` in this directory as an example for convenience but the import path at the top will obviously need changed to run on a different machine.

### Output

```sh
05:13:00 ~/Development/Go/src/github.com/teamwork.com/teamworkgotests
$ go run main.go
2018/02/03 17:13:05 There are 500 unique email domains with the following number of customers associated for each:
2018/02/03 17:13:05 123-reg.co.uk 8
2018/02/03 17:13:05 163.com 6
2018/02/03 17:13:05 1688.com 3
2018/02/03 17:13:05 1und1.de 5
2018/02/03 17:13:05 360.cn 6
2018/02/03 17:13:05 4shared.com 5
2018/02/03 17:13:05 51.la 4
2018/02/03 17:13:05 a8.net 6
2018/02/03 17:13:05 abc.net.au 7
2018/02/03 17:13:05 about.com 5
2018/02/03 17:13:05 about.me 2
...
```

## Design

- A simple `map[string]int` type would not be adequate as a data structure because it's unordered. Instead I chose a slice of `domainStatistic`'s (key value pairs) for the ability to sort alphabetically by email domain. A different method of sorting can be easily implemented if needed.
- The Golang `bufio` package was used to scan the CSV file as performance was a concern given the requirement for supporting large file sizes. By scanning one line of the file at a time, minimal resources are required.
- The code resides in one file for convenience but would probably benefit from a refactor to split some of the functionality out into different files for maintainability purposes e.g. have separate `customer.go` and `customerimporter.go` files.
- The structure of the code is standard Golang best practices. For example, use of interfaces and factory functions where possible produces more maintainable decoupled code. This has a knock on effect of being easier to test as well as facilitating the easy swapping out of one implementation for another in the future, if required.
- All errors are handled or returned from each function/method. The `main.go` file uses Golang's `log` package to print to `STDOUT` to create a CLI UI. The `Import()` method tries to be as robust as possible. For example, it skips any lines which may be blank or invalid making for good defensive programming.
- The code is tested using Golang's built in `testing` package for simplicity but on a larger project I would suggest a BDD tool such as `goconvey`.
