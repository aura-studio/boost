package regexp

import (
	"errors"
	"fmt"
	"sync"

	"github.com/aura-studio/boost/cast"
	"github.com/dlclark/regexp2"
)

var (
	ErrInvalidRemoteAddr = errors.New("invalid remote address")
	ErrInvalidIPType     = errors.New("invalid ip type")
)

type Regexp struct {
	sync.Map
}

func New() *Regexp {
	return &Regexp{}
}

func (r *Regexp) MatchString(pattern string, str string) (bool, error) {
	v, ok := r.Load(pattern)
	if !ok {
		v = regexp2.MustCompile(pattern, regexp2.RE2)
		r.Store(pattern, v)
	}

	re := v.(*regexp2.Regexp)
	return re.MatchString(str)
}

func (r *Regexp) ReplaceAllStringFunc(pattern string, str string, repl func(string) string) (string, error) {
	v, ok := r.Load(pattern)
	if !ok {
		v = regexp2.MustCompile(pattern, regexp2.RE2)
		r.Store(pattern, v)
	}

	re := v.(*regexp2.Regexp)
	return re.ReplaceFunc(str, func(m regexp2.Match) string {
		return repl(m.String())
	}, -1, -1)
}

var (
	reIPv4            = regexp2.MustCompile(`^(?P<ip>(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?))$`, regexp2.RE2)
	reIPv4Port        = regexp2.MustCompile(`^(?P<ip>(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)):(?P<port>([0-9]{1,5}))$`, regexp2.RE2)
	reIPv6            = regexp2.MustCompile(`^(?P<ip>([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]).){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]).){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$`, regexp2.RE2)
	reIPv6Port        = regexp2.MustCompile(`^(?P<ip>([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]).){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]).){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])):(?P<port>[0-9]{1,5})$`, regexp2.RE2)
	reIPv6Bracket     = regexp2.MustCompile(`^\[(?P<ip>([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]).){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]).){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))\]$`, regexp2.RE2)
	reIPv6BracektPort = regexp2.MustCompile(`^\[(?P<ip>([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]).){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]).){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))\]:(?P<port>[0-9]{1,5})$`, regexp2.RE2)
)

type IPType int

const (
	IPNil = iota
	IPv4
	IPv6
)

var ipTypeStringMap = map[IPType]string{
	IPNil: "nil",
	IPv4:  "ipv4",
	IPv6:  "ipv6",
}

func (t IPType) String() string {
	return ipTypeStringMap[t]
}

type RemoteAddr struct {
	IPType IPType
	IP     string
	Port   uint
}

func (ra *RemoteAddr) Parse(addr string) error {
	match, _ := reIPv4.FindStringMatch(addr)
	if match != nil {
		ra.IPType = IPv4
		ra.IP = match.GroupByName("ip").String()
		ra.Port = 0
		return nil
	}

	match, _ = reIPv4Port.FindStringMatch(addr)
	if match != nil {
		ra.IPType = IPv4
		ra.IP = match.GroupByName("ip").String()
		ra.Port = cast.ToUint(match.GroupByName("port").String())
		return nil
	}

	match, _ = reIPv6.FindStringMatch(addr)
	if match != nil {
		ra.IPType = IPv6
		ra.IP = match.GroupByName("ip").String()
		ra.Port = 0
		return nil
	}

	match, _ = reIPv6Port.FindStringMatch(addr)
	if match != nil {
		ra.IPType = IPv6
		ra.IP = match.GroupByName("ip").String()
		ra.Port = cast.ToUint(match.GroupByName("port").String())
		return nil
	}

	match, _ = reIPv6Bracket.FindStringMatch(addr)
	if match != nil {
		ra.IPType = IPv6
		ra.IP = match.GroupByName("ip").String()
		ra.Port = 0
		return nil
	}

	match, _ = reIPv6BracektPort.FindStringMatch(addr)
	if match != nil {
		ra.IPType = IPv6
		ra.IP = match.GroupByName("ip").String()
		ra.Port = cast.ToUint(match.GroupByName("port").String())
		return nil
	}

	return fmt.Errorf("%w: %s", ErrInvalidRemoteAddr, addr)
}

func (ra RemoteAddr) Format() (string, error) {
	switch ra.IPType {
	case IPv4:
		return fmt.Sprintf("%s:%d", ra.IP, ra.Port), nil
	case IPv6:
		return fmt.Sprintf("[%s]:%d", ra.IP, ra.Port), nil
	}

	return "", fmt.Errorf("%w: %s", ErrInvalidIPType, ra.IPType)
}

type remoteAddrAnalyzer struct{}

var RemoteAddrAnalyzer = remoteAddrAnalyzer{}

func (remoteAddrAnalyzer) Parse(addr string) (*RemoteAddr, error) {
	ra := &RemoteAddr{}
	if err := ra.Parse(addr); err != nil {
		return nil, err
	}
	return ra, nil
}

func (remoteAddrAnalyzer) Format(addr string) (string, error) {
	ra := &RemoteAddr{}
	if err := ra.Parse(addr); err != nil {
		return "", err
	}

	return ra.Format()
}
