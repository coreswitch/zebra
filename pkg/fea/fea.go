package fea

import (
	"github.com/coreswitch/component"
)

var (
	Inits  []func() error
	VrfAdd func(vrfName string, index uint32) error
)

type FeaParam map[string]string

type FeaComponent struct {
	Params *FeaParam
}

func NewFeaComponent(param *FeaParam) *FeaComponent {
	return &FeaComponent{
		Params: param,
	}
}

func (fea *FeaComponent) Start() component.Component {
	//fea_linux.Ipv4ForwardingSet(true)
	//fmt.Println("ipv4 forwarding", fea_linux.Ipv4ForwardingGet())
	//fea_linux.Ipv6ForwardingSet(true)
	//fmt.Println("ipv6 forwarding", fea_linux.Ipv6ForwardingGet())
	return fea
}

func (fea *FeaComponent) Stop() component.Component {
	return fea
}
