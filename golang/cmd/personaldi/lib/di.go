package di

type DI struct {
	deps map[string]Injectable
	lazy map[string]Dependency
}

type Injectable interface{}

type Dependency struct {
	Key   string
	Setup func() Injectable
	Lazy  bool
}

func NewDIContainer() *DI {
	return &DI{
		deps: make(map[string]Injectable),
		lazy: make(map[string]Dependency),
	}
}

func (r *DI) Provide(dep Dependency) {
	if !dep.Lazy {
		r.deps[dep.Key] = dep.Setup()
		return
	}
	r.lazy[dep.Key] = dep
}

func (r *DI) Inject(key string) Injectable {
	dep, ok := r.deps[key]
	if ok {
		return dep
	}
	lazy, ok := r.lazy[key]
	if !ok {
		panic("failed to load lazy dependency. Key: " + key)
	}
	return lazy.Setup()
}
