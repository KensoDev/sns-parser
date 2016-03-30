package snsparser

import (
	. "gopkg.in/check.v1"
	"io/ioutil"
	"testing"
)

func TestParser(t *testing.T) { TestingT(t) }

type ParserSuite struct{}

var _ = Suite(&ParserSuite{})

func (s *ParserSuite) TestParsingJson(c *C) {
	content, _ := ioutil.ReadFile("fixtures/sns-message.json")
	parser := NewSNSParser(content)
	entry := parser.Entry
	c.Assert(len(entry.Records), Equals, 1)
	c.Assert(entry.Records[0].EventSource, Equals, "aws:sns")
	c.Assert(entry.Records[0].SNS.Subject, Equals, "My First Message")
}

func (s *ParserSuite) TestContainsMessage(c *C) {
	content, _ := ioutil.ReadFile("fixtures/sns-message.json")
	parser := NewSNSParser(content)
	includes, _ := parser.IncludesMessage("Hello")
	c.Assert(includes, Equals, true)
}
