package models

import whoisparser "github.com/likexian/whois-parser"

type WhoISDTO struct {
	Error          string
	Domain         whoisparser.Domain
	Registrar      whoisparser.Contact
	Registrant     whoisparser.Contact
	Administrative whoisparser.Contact
	Technical      whoisparser.Contact
	Billing        whoisparser.Contact
}
