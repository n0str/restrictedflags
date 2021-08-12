package restrictedflags

import (
	"fmt"
)

type RestrictedFlag struct {
	AllowedValues []string `json:"allowedValues"`
	Value string `json:"value"`
}

func New(allowedValues []string) RestrictedFlag {
	return RestrictedFlag{AllowedValues:allowedValues}
}

func (rf *RestrictedFlag) String() string {
	return rf.Value
}

func (rf *RestrictedFlag) Set(s string) error {
	for _, value := range rf.AllowedValues {
		if s == value {
			rf.Value = s
			return nil
		}
	}
	return fmt.Errorf("value `%s` is not in the allowed set", s)
}