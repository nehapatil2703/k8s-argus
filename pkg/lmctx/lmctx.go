package lmctx

// LMContext will be used to pass metadata from one to other, for ex: log context.
type LMContext struct {
	kv map[string]interface{}
}

// Extract returns value against key.
func (lc *LMContext) Extract(key string) interface{} {
	return lc.kv[key]
}

// Set will store key val in map
func (lc *LMContext) Set(key string, val interface{}) {
	lc.kv[key] = val
}

// Copy copy
func (lc *LMContext) Copy() *LMContext {
	// copying variable instead of pointer to make deep copy
	l2 := *lc

	return &l2
}

// NewLMContext creates new context object
func NewLMContext() *LMContext {
	ctx := &LMContext{
		kv: make(map[string]interface{}),
	}

	return ctx
}
