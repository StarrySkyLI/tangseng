package consts

const (
	KafkaAssignorRoundRobin = "roundrobin"
	KafkaAssignorSticky     = "sticky"
	KafkaAssignorRange      = "range"
)

const (
	KafkaCrawlTopic     = "kafka-crawl-topic"
	KafkaCSVLoaderTopic = "search-engine-csv-loader-topic"
	KafkaTrieTreeTopic  = "search-engine-trie-tree-topic"
)

const (
	KafkaCSVLoaderGroupId = "kafka-csv-loader-group-id"
)

const KafkaBatchProduceCount = 200
