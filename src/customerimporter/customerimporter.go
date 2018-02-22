package customerimporter

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	CSVColumnSeparator   = `,`
	EmailDomainSeparator = `@`
)

// Interfaces

type Customer interface {
	GetEmailDomain() (string, error)
}

type CustomerImporter interface {
	Import(string) ([]*DomainStatistic, error)
}

// Structs and methods

type customer struct {
	Forename  string
	Surname   string
	Email     string
	Gender    string
	IPAddress string
}

// Return everything after @ (the email domain) or an error
func (c *customer) GetEmailDomain() (string, error) {
	splitEmail := strings.Split(c.Email, EmailDomainSeparator)
	if len(splitEmail) < 2 {
		return ``, errors.New(fmt.Sprintf(`Invalid email address '%s', missing %s symbol`, c.Email, EmailDomainSeparator))
	}
	return splitEmail[1], nil
}

type customerImporter struct{}

// Imports a files contents efficiently by using bufio.Scan() which processes
// the file line by line. This effectively means the same amount of system
// resources are needed for any size of file.
//
// Returns a map of email domains with number of occurrences or an error.
func (*customerImporter) Import(filepath string) ([]*DomainStatistic, error) {
	// Open the file for loading into the buffer
	file, err := os.Open(filepath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf(`Error opening file: `, filepath))
	}

	domainStats := make([]*DomainStatistic, 0)
	scanner := bufio.NewScanner(file)

	// Iterate over and process each line, loop will end at EOF
	for scanner.Scan() {
		line := scanner.Text()
		if line == `` {
			// Encountered blank line, skipping...
			continue
		}
		if columns := strings.Split(line, CSVColumnSeparator); len(columns) < 5 {
			// Encountered invalid line, skipping...
			continue
		} else {
			customer := NewCustomer(columns[0], columns[1], columns[2], columns[3], columns[4])
			// Continue to the next line/customer if there is an invalid email
			if domain, err := customer.GetEmailDomain(); err == nil {
				// Process the email domain into domainStats
				if domainStat, present := getDomainStat(domainStats, domain); present {
					domainStat.Count++
				} else {
					domainStats = append(domainStats, NewDomainStatistic(domain, 1))
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, errors.New(fmt.Sprintf(`Error reading file: `, filepath))
	}

	return sortAlphabeticallyByDomain(domainStats), nil
}

type DomainStatistic struct {
	Domain string
	Count  int
}

// Factory funcs

func NewCustomer(forename, surname, email, gender, ip_address string) Customer {
	return &customer{
		Forename:  forename,
		Surname:   surname,
		Email:     email,
		Gender:    gender,
		IPAddress: ip_address,
	}
}

func NewCustomerImporter() CustomerImporter {
	return &customerImporter{}
}

func NewDomainStatistic(domain string, count int) *DomainStatistic {
	return &DomainStatistic{
		Domain: domain,
		Count:  count,
	}
}

// Helper funcs

func getDomainStat(domainStats []*DomainStatistic, domain string) (*DomainStatistic, bool) {
	for _, stat := range domainStats {
		if stat.Domain == domain {
			return stat, true
		}
	}
	return nil, false
}

func sortAlphabeticallyByDomain(stats []*DomainStatistic) []*DomainStatistic {
	// Extract and sort the domains
	domains := make([]string, len(stats))
	for _, stat := range stats {
		domains = append(domains, stat.Domain)
	}
	sort.Strings(domains)

	// Rebuild the slice ordered by the email domains
	sorted := make([]*DomainStatistic, 0)
	for _, domain := range domains {
		if domainStat, present := getDomainStat(stats, domain); present {
			sorted = append(sorted, domainStat)
		}
	}

	return sorted
}
