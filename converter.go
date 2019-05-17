package xml2json

import (
	"bytes"
	"io"
	"strings"
)

// Convert converts the given XML document to JSON
func Convert(r io.Reader, ps ...plugin) (string, error) {
	// Decode XML document
	root := &Node{}
	err := NewDecoder(r, ps...).Decode(root)
	if err != nil {
		return "", err
	}

	// Then encode it in JSON
	buf := new(bytes.Buffer)
	e := NewEncoder(buf, ps...)
	err = e.Encode(root)
	if err != nil {
		return "", err
	}

	 jsonStr := strings.Trim(buf.String(), `"`+string([]byte{10, 13, 32}))

	return jsonStr, nil
}
