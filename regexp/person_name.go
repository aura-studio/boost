package regexp

import (
	"strings"

	"github.com/dlclark/regexp2"
)

type PersonName struct {
	names []string
}

func (pn *PersonName) Parse(name string) error {
	re := regexp2.MustCompile(`[^A-Za-z ]`, regexp2.None)
	name, err := re.Replace(name, "", -1, -1)
	if err != nil {
		return err
	}

	parts := strings.Split(name, " ")

	var names []string
	for _, part := range parts {
		if part == "" {
			continue
		}

		names = append(names, strings.ToTitle(part))
	}

	pn.names = names

	return nil
}

func (pn *PersonName) MustParse(name string) *PersonName {
	if err := pn.Parse(name); err != nil {
		panic(err)
	}

	return pn
}

func (pn *PersonName) Format() (string, error) {
	return strings.Join(pn.names, " "), nil
}

func (pn *PersonName) MustFormat() string {
	s, err := pn.Format()
	if err != nil {
		panic(err)
	}

	return s
}

type personNameAnalyzer struct{}

var PersonNameAnalyzer = personNameAnalyzer{}

func (personNameAnalyzer) Parse(name string) (*PersonName, error) {
	personName := &PersonName{}
	if err := personName.Parse(name); err != nil {
		return nil, err
	}

	return personName, nil
}

func (personNameAnalyzer) MustParse(name string) *PersonName {
	personName := &PersonName{}
	return personName.MustParse(name)
}

func (personNameAnalyzer) Format(name string) (string, error) {
	personName := &PersonName{}
	if err := personName.Parse(name); err != nil {
		return "", err
	}

	return personName.Format()
}

func (personNameAnalyzer) MustFormat(name string) string {
	personName := &PersonName{}
	return personName.MustParse(name).MustFormat()
}
