package context

import (
	"context"
	"encoding/json"
)

type propagateKey struct{}

var PropagateCtxKey = propagateKey{}

func AddToPropagateCtx(ctx context.Context, key string, val interface{}) context.Context {
	propagate := ToMap(ctx)
	propagate[key] = val
	return context.WithValue(ctx, PropagateCtxKey, propagate)
}

func GetFromPropagateCtx(ctx context.Context, key string) interface{} {
	propagate := ToMap(ctx)
	if val, ok := propagate[key]; ok {
		return val
	}
	return nil
}

func ToMap(ctx context.Context) map[string]interface{} {
	if ctx == nil {
		return map[string]interface{}{}
	}
	p := ctx.Value(PropagateCtxKey)
	if p != nil {
		return p.(map[string]interface{})
	}
	return map[string]interface{}{}
}

func FromMap(val map[string]interface{}) context.Context {
	return context.WithValue(context.Background(), PropagateCtxKey, val)
}

func Encode(ctx context.Context) ([]byte, error) {
	m := ToMap(ctx)
	if len(m) > 0 {
		return json.Marshal(m)
	}
	return nil, nil
}

func Decode(m []byte) (context.Context, error) {
	if len(m) == 0 {
		return nil, nil
	}
	mp := make(map[string]interface{}, 0)
	err := json.Unmarshal(m, &mp)
	if err != nil {
		return nil, err
	}
	return FromMap(mp), nil
}
