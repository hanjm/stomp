package frame

import (
	. "gopkg.in/check.v1"
)

type EncodeSuite struct{}

var _ = Suite(&EncodeSuite{})

func (s *EncodeSuite) TestUnencodeValue(c *C) {
	val, err := unencodeValue([]byte(`Contains\r\nNewLine and \c colon and \\ backslash`))
	c.Check(err, IsNil)
	c.Check(val, Equals, "Contains\r\nNewLine and : colon and \\ backslash")
}
