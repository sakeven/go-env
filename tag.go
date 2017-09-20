package env

import (
	"strconv"
	"strings"
)

type structTag struct {
	name      string // name of tag
	omitempty bool   // use omit empty value
	skip      bool   // skip this tag
	dft       string // default value

	raw       string
	fieldName string
}

func newStructTag(fieldName string, rawStructTag string) *structTag {
	tag := &structTag{
		fieldName: fieldName,
		raw:       rawStructTag,
	}
	return tag.parseTag()
}

// withPrefix should be called after parseTag.
func (t *structTag) withPrefix(prefix string) *structTag {
	if len(prefix) > 0 {
		t.name = prefix + "_" + t.name
	}
	return t
}

func (t *structTag) defaultInt64() (int64, error) {
	defaultValue := int64(0)
	if t.omitempty != true {
		var err error
		defaultValue, err = strconv.ParseInt(t.dft, 10, 64)
		if err != nil {
			return 0, err
		}
	}

	return defaultValue, nil
}

func (t *structTag) defaultBool() (bool, error) {
	defaultValue := false
	if t.omitempty != true {
		var err error
		defaultValue, err = strconv.ParseBool(t.dft)
		if err != nil {
			return false, err
		}
	}
	return defaultValue, nil
}

func (t *structTag) defaultString() (string, error) {
	defaultValue := ""
	if t.omitempty != true {
		defaultValue = t.dft
	}
	return defaultValue, nil
}

func (t *structTag) parseTag() *structTag {
	list := strings.SplitN(t.raw, ",", 2)

	var options [2]string
	for i, op := range list {
		options[i] = fix(op)
	}

	//  tag name
	switch options[0] {
	case "-":
		t.skip = true
	case "":
		// use origin field name
		t.name = t.fieldName
	default:
		t.name = options[0]
	}

	// tag default value
	switch options[1] {
	case "":
		// use omit empty value
		t.omitempty = true
	default:
		t.dft = options[1]
	}

	return t
}

func fix(s string) string {
	s = strings.Trim(s, " ")
	s = strings.Trim(s, "\t")
	return s
}
