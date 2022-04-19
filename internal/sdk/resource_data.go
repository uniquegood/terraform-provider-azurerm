package sdk

import "time"

type ResourceData interface {
	// Get returns a value from either the config/state depending on where this is called
	// in Create and Update functions this will return from the config
	// in Read, Exists and Import functions this will return from the state
	// NOTE: this should not be called from Delete functions.
	Get(key string) interface{}

	// GetChange returns the original and updated value, which can be useful in Update functions
	GetChange(key string) (original interface{}, updated interface{})

	// GetValue returns the value for this key, alongside a boolean determining whether
	// this field was set in the config
	GetValue(key string) (value interface{}, isSet bool)

	GetRawValue(key string) (value interface{}, isSet bool)

	HasChange(key string) bool

	HasChanges(keys ...string) bool

	Id() string

	IsNewResource() bool

	Set(key string, value interface{}) error

	SetConnInfo(v map[string]string)

	SetId(id string)

	Timeout(key string) time.Duration

	// TODO: add Get/Set helpers for each type

	//GetString(key string) * string

	// TODO: remove below here, just here to enable a seamless migration

	// Deprecated: use GetValue
	GetOk(key string) (interface{}, bool)

	// Deprecated: use GetRawValue
	GetOkExists(key string) (interface{}, bool)
}
