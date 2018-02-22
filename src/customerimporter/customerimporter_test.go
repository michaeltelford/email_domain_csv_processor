package customerimporter

import "testing"

func TestGetEmailDomain(t *testing.T) {
	expected := `gmail.com`
	c := NewCustomer(`Joe`, `Bloggs`, `j.bloggs@gmail.com`, `male`, `192.168.0.1`)

	domain, err := c.GetEmailDomain()
	if err != nil {
		errorFail(t, err)
	}
	if domain != expected {
		fail(t, expected, domain)
	}
}

func TestGetEmailDomainFails(t *testing.T) {
	expected := ``
	c := NewCustomer(`Joe`, `Bloggs`, `j.bloggsgmail.com`, `male`, `192.168.0.1`)

	domain, err := c.GetEmailDomain()
	if err == nil {
		t.Error(`Expected an error`)
	}
	if domain != expected {
		fail(t, expected, domain)
	}
}

func TestImport(t *testing.T) {
	importer := NewCustomerImporter()
	stats, err := importer.Import(`test.csv`)

	if err != nil {
		errorFail(t, err)
	}
	if len(stats) != 2 {
		fail(t, 2, len(stats))
	}
	if stats[0].Domain != `cyberchimps.com` {
		fail(t, `cyberchimps.com`, stats[0].Domain)
	}
	if stats[1].Domain != `github.io` {
		fail(t, `github.io`, stats[1].Domain)
	}
}

func TestGetDomainStat(t *testing.T) {
	stats := []*DomainStatistic{}
	stats = append(stats, NewDomainStatistic(`gmail.com`, 2))

	stat, present := getDomainStat(stats, `gmail.com`)
	if stat.Domain != `gmail.com` {
		fail(t, `gmail.com`, stat.Domain)
	}
	if stat.Count != 2 {
		fail(t, 2, stat.Count)
	}
	if present != true {
		fail(t, true, present)
	}
}

func TestGetDomainStatNotFound(t *testing.T) {
	stats := []*DomainStatistic{}

	stat, present := getDomainStat(stats, `gmail.com`)
	if stat != nil {
		fail(t, nil, stat)
	}
	if present != false {
		fail(t, false, present)
	}
}

func TestSortAlphabeticallyByDomain(t *testing.T) {
	expected := `abc.com`
	stats := []*DomainStatistic{}
	stats = append(stats, NewDomainStatistic(`gmail.com`, 2))
	stats = append(stats, NewDomainStatistic(expected, 5))

	sorted := sortAlphabeticallyByDomain(stats)
	if sorted[0].Domain != expected {
		fail(t, expected, sorted[0].Domain)
	}
}

// Test helper funcs

func fail(t *testing.T, expected, actual interface{}) {
	t.Error(`Expected:`, expected, `Actual:`, actual)
}

func errorFail(t *testing.T, err error) {
	t.Error(`Unexpected error:`, err.Error())
}
