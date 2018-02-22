package main

import (
	"fmt"
	"log"

	"github.com/michaeltelford/email_domain_csv_processor/src/customerimporter"
)

func main() {
	importer := customerimporter.NewCustomerImporter()
	domainStats, err := importer.Import(`customers.csv`)
	if err != nil {
		log.Fatal(err.Error())
	}

	logDomainStats(domainStats)
}

func logDomainStats(domainStats []*customerimporter.DomainStatistic) {
	log.Println(fmt.Sprintf(`There are %d unique email domains with the following number of customers associated for each:`, len(domainStats)))

	for _, stat := range domainStats {
		log.Println(stat.Domain, stat.Count)
	}
}
