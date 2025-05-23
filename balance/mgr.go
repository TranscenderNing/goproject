package balance

import "fmt"

type BalancerMgr struct {
	allBalance map[string]Balancer
}

var (
	mgr = BalancerMgr{
		allBalance: make(map[string]Balancer),
	}
)

func (p *BalancerMgr) registerBalancer(name string, b Balancer) {
	p.allBalance[name] = b
}

func RegisterBalancer(name string, b Balancer) {
	mgr.registerBalancer(name, b)
}


func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	
	fmt.Println(mgr.allBalance)
	balancer, ok := mgr.allBalance[name]
	if !ok {
		err = fmt.Errorf("not found %s", name)
		return 
	}
	inst, err = balancer.DoBalance(insts)
	return 
}
