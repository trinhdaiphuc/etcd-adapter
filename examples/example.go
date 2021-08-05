package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/trinhdaiphuc/etcd-adapter"
)

func main() {
	// Initialize a casbin etcd adapter and use it in a Casbin enforcer:
	// The adapter will use the ETCD and a named path with the key you give.
	// If not provided, the adapter will use the default value casbin_policy.
	a := etcdadapter.NewAdapter(
		etcdadapter.WithEndpoints([]string{"http://localhost:2379"}),
		etcdadapter.WithEtcdAuth("root", "password"),
		etcdadapter.WithKey("casbin_policy_test"),
	) // Your etcd endpoints and the path key.

	e, _ := casbin.NewEnforcer("rbac_model.conf", a)

	// Load the policy from ETCD.
	e.LoadPolicy()

	// Check the permission.
	result, err := e.Enforce("alice", "data1", "read")

	fmt.Printf("result %v, error %v\n", result, err)

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	e.SavePolicy()
}
