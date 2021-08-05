etcd-adapter
====

[![Build Status](https://travis-ci.org/sebastianliu/etcd-adapter.svg?branch=master)](https://travis-ci.org/sebastianliu/etcd-adapter)
[![Coverage Status](https://coveralls.io/repos/github/sebastianliu/etcd-adapter/badge.svg)](https://coveralls.io/github/sebastianliu/etcd-adapter)
[![Godoc](https://godoc.org/github.com/sebastianliu/etcd-adapter?status.svg)](https://godoc.org/github.com/sebastianliu/etcd-adapter)

ETCD adapter is the policy storage adapter for [Casbin](https://github.com/casbin/casbin). With this library, Casbin can load policy from ETCD and save policy to it. ETCD adapter support the __Auto-Save__ feature for Casbin policy. This means it can support adding a single policy rule to the storage, or removing a single policy rule from the storage.

## Installation
```bash
go get github.com/sebastianliu/etcd-adapter
```

## Sample Example
```go
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

	e.EnableAutoSave(true)

	// Check the permission.
	result, err := e.Enforce("alice", "data1", "read")

	fmt.Printf("result %v, error %v\n", result, err)

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	e.SavePolicy()
}
```
