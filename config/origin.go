package config

import (
	"fmt"
	"github.com/duanhf2012/origin/v2/service"
)

type OriginModule struct {
	service.Module
}

func (m *OriginModule) OnInit() error {
	fmt.Printf("OriginModule OnInit.\n")
	return nil
}
