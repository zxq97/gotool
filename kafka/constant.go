package kafka

import "time"

const (
	defaultDialTimeout     = 500 * time.Millisecond
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
	defaultConsumerTimeout = 5 * time.Second
)
