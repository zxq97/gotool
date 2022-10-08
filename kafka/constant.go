package kafka

import "time"

const (
	defaultDialTimeout     = 500 * time.Millisecond
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
	defaultConsumerTimeout = 5 * time.Second
)

const (
	TopicCommentPublish      = "comment_publish"
	TopicCommentCacheRebuild = "comment_cache_rebuild"
	TopicCommentOperator     = "comment_operator"

	TopicRelationFollow       = "relation_follow"
	TopicRelationCacheRebuild = "relation_cache_rebuild"
	TopicRelationOperator     = "relation_operator"

	TopicOpusOperator = "opus_operator"

	EventTypeCreate        = 1
	EventTypeReply         = 2
	EventTypeListMissed    = 3
	EventTypeSubListMissed = 4
	EventTypeLike          = 5
	EventTypeHate          = 6
	EventTypeDelete        = 7
)
