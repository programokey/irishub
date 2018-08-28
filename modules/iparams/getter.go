package iparams

import (

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Getter exposes methods related with only getting params
type Getter struct {
	k Keeper
}

// Get exposes get
func (k Getter) Get(ctx sdk.Context, key string, ptr interface{}) error {
	return k.k.get(ctx, key, ptr)
}

// GetRaw exposes getRaw
func (k Getter) GetRaw(ctx sdk.Context, key string) []byte {
	return k.k.getRaw(ctx, key)
}

// GetString is helper function for string params
func (k Getter) GetString(ctx sdk.Context, key string) (res string, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinary(bz, &res)
	return
}

// GetBool is helper function for bool params
func (k Getter) GetBool(ctx sdk.Context, key string) (res bool, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinary(bz, &res)
	return
}

// GetInt16 is helper function for int16 params
func (k Getter) GetInt16(ctx sdk.Context, key string) (res int16, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinary(bz, &res)
	return
}

// GetInt32 is helper function for int32 params
func (k Getter) GetInt32(ctx sdk.Context, key string) (res int32, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinary(bz, &res)
	return
}

// GetInt64 is helper function for int64 params
func (k Getter) GetInt64(ctx sdk.Context, key string) (res int64, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinary(bz, &res)
	return
}

// GetUint16 is helper function for uint16 params
func (k Getter) GetUint16(ctx sdk.Context, key string) (res uint16, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinary(bz, &res)
	return
}

// GetUint32 is helper function for uint32 params
func (k Getter) GetUint32(ctx sdk.Context, key string) (res uint32, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinary(bz, &res)
	return
}

// GetUint64 is helper function for uint64 params
func (k Getter) GetUint64(ctx sdk.Context, key string) (res uint64, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinary(bz, &res)
	return
}

// GetInt is helper function for sdk.Int params
func (k Getter) GetInt(ctx sdk.Context, key string) (res sdk.Int, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinary(bz, &res)
	return
}

// GetUint is helper function for sdk.Uint params
func (k Getter) GetUint(ctx sdk.Context, key string) (res sdk.Uint, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinary(bz, &res)
	return
}

// GetRat is helper function for rat params
func (k Getter) GetRat(ctx sdk.Context, key string) (res sdk.Rat, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinary(bz, &res)
	return
}

// GetStringWithDefault is helper function for string params with default value
func (k Getter) GetStringWithDefault(ctx sdk.Context, key string, def string) (res string) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinary(bz, &res)
	return
}

// GetBoolWithDefault is helper function for bool params with default value
func (k Getter) GetBoolWithDefault(ctx sdk.Context, key string, def bool) (res bool) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinary(bz, &res)
	return
}

// GetInt16WithDefault is helper function for int16 params with default value
func (k Getter) GetInt16WithDefault(ctx sdk.Context, key string, def int16) (res int16) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinary(bz, &res)
	return
}

// GetInt32WithDefault is helper function for int32 params with default value
func (k Getter) GetInt32WithDefault(ctx sdk.Context, key string, def int32) (res int32) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinary(bz, &res)
	return
}

// GetInt64WithDefault is helper function for int64 params with default value
func (k Getter) GetInt64WithDefault(ctx sdk.Context, key string, def int64) (res int64) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinary(bz, &res)
	return
}

// GetUint16WithDefault is helper function for uint16 params with default value
func (k Getter) GetUint16WithDefault(ctx sdk.Context, key string, def uint16) (res uint16) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinary(bz, &res)
	return
}

// GetUint32WithDefault is helper function for uint32 params with default value
func (k Getter) GetUint32WithDefault(ctx sdk.Context, key string, def uint32) (res uint32) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinary(bz, &res)
	return
}

// GetUint64WithDefault is helper function for uint64 params with default value
func (k Getter) GetUint64WithDefault(ctx sdk.Context, key string, def uint64) (res uint64) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinary(bz, &res)
	return
}

// GetIntWithDefault is helper function for sdk.Int params with default value
func (k Getter) GetIntWithDefault(ctx sdk.Context, key string, def sdk.Int) (res sdk.Int) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinary(bz, &res)
	return
}

// GetUintWithDefault is helper function for sdk.Uint params with default value
func (k Getter) GetUintWithDefault(ctx sdk.Context, key string, def sdk.Uint) (res sdk.Uint) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinary(bz, &res)
	return
}

// GetRatWithDefault is helper function for sdk.Rat params with default value
func (k Getter) GetRatWithDefault(ctx sdk.Context, key string, def sdk.Rat) (res sdk.Rat) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinary(bz, &res)
	return
}
