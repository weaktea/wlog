package wlog

import (
	"fmt"
	"testing"
)

func TestLoadXmlConfig(t *testing.T) {
	LoadXmlConfig()
	fmt.Printf("%+v", config)
}