package snsparser

import (
	"encoding/json"
	"strings"
)

type Parser struct {
	Entry *Entry
}

func NewSNSParser(jsonContent []byte) *Parser {
	entry := new(Entry)
	err := json.Unmarshal(jsonContent, entry)
	if err != nil {
		panic(err)
	}
	return &Parser{Entry: entry}
}

func (parser *Parser) IncludesMessage(message string) (bool, SNS) {
	contains := false
	sns := SNS{}
	for _, v := range parser.Entry.Records {
		contains = strings.Contains(v.SNS.Message, message)
		sns = v.SNS
		if contains {
			break
		}
	}
	return contains, sns
}
