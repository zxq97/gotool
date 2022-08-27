package constant

import "time"

const (
	TraceIDKey = "trace_id"
	SpanIDKey  = "span_id"

	EtcdLeaseTTL = 1000 * 60

	DefaultTimeout = time.Second
	DefaultTicker  = time.Minute

	EtcdScheme = "etcd"
)
