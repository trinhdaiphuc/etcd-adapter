package etcdadapter

type AdapterOption func(*Adapter)

func WithEndpoints(etcdEndpoints []string) AdapterOption {
	return func(a *Adapter) {
		a.etcdEndpoints = etcdEndpoints
	}
}

func WithEtcdAuth(username, password string) AdapterOption {
	return func(a *Adapter) {
		a.etcdUsername = username
		a.etcdPassword = password
	}
}

func WithKey(key string) AdapterOption {
	return func(a *Adapter) {
		a.key = key
	}
}
