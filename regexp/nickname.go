package regexp

import "github.com/dlclark/regexp2"

type NicknameWithBlank struct {
	name string
}

func (n *NicknameWithBlank) Parse(name string) error {
	re := regexp2.MustCompile(`[^A-Za-z0-9 ]`, regexp2.None)
	name, err := re.Replace(name, "", -1, -1)
	if err != nil {
		return err
	}

	n.name = name

	return nil
}

func (n *NicknameWithBlank) MustParse(name string) *NicknameWithBlank {
	if err := n.Parse(name); err != nil {
		panic(err)
	}

	return n
}

func (n *NicknameWithBlank) Format() (string, error) {
	return n.name, nil
}

func (n *NicknameWithBlank) MustFormat() string {
	s, err := n.Format()
	if err != nil {
		panic(err)
	}

	return s
}

type nicknameWithBlankAnalyzer struct{}

var NicknameWithBlankAnalyzer = nicknameWithBlankAnalyzer{}

func (nicknameWithBlankAnalyzer) Parse(name string) (*NicknameWithBlank, error) {
	n := &NicknameWithBlank{}
	if err := n.Parse(name); err != nil {
		return nil, err
	}
	return n, nil
}

func (nicknameWithBlankAnalyzer) MustParse(name string) *NicknameWithBlank {
	n := &NicknameWithBlank{}
	return n.MustParse(name)
}

func (nicknameWithBlankAnalyzer) Format(name string) (string, error) {
	n := &NicknameWithBlank{}
	if err := n.Parse(name); err != nil {
		return "", err
	}

	return n.Format()
}

func (nicknameWithBlankAnalyzer) MustFormat(name string) string {
	n := &NicknameWithBlank{}
	return n.MustFormat()
}

type NicknameWithoutBlank struct {
	name string
}

func (n *NicknameWithoutBlank) Parse(name string) error {
	re := regexp2.MustCompile(`[^A-Za-z0-9]`, regexp2.None)
	name, err := re.Replace(name, "", -1, -1)
	if err != nil {
		return err
	}

	n.name = name

	return nil
}

func (n *NicknameWithoutBlank) MustParse(name string) *NicknameWithoutBlank {
	if err := n.Parse(name); err != nil {
		panic(err)
	}

	return n
}

func (n *NicknameWithoutBlank) Format() (string, error) {
	return n.name, nil
}

func (n *NicknameWithoutBlank) MustFormat() string {
	s, err := n.Format()
	if err != nil {
		panic(err)
	}

	return s
}

type nicknameWithoutBlankAnalyzer struct{}

var NicknameWithoutBlankAnalyzer = nicknameWithoutBlankAnalyzer{}

func (nicknameWithoutBlankAnalyzer) Parse(name string) (*NicknameWithoutBlank, error) {
	n := &NicknameWithoutBlank{}
	if err := n.Parse(name); err != nil {
		return nil, err
	}
	return n, nil
}

func (nicknameWithoutBlankAnalyzer) MustParse(name string) *NicknameWithoutBlank {
	n := &NicknameWithoutBlank{}
	return n.MustParse(name)
}

func (nicknameWithoutBlankAnalyzer) Format(name string) (string, error) {
	n := &NicknameWithoutBlank{}
	if err := n.Parse(name); err != nil {
		return "", err
	}

	return n.Format()
}

func (nicknameWithoutBlankAnalyzer) MustFormat(name string) string {
	n := &NicknameWithoutBlank{}
	return n.MustFormat()
}
