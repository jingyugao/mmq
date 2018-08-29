package mmq

func GetConsumeKey(name string) string {
	return "mmq::consume::" + name
}

func GetQueueKey(name string) string {
	return "mmq::queue::" + name
}

func GetQueueWorkingKey(name string) string {
	return "mmq::queue_work::" + name
}

func GetQueueSetKey() string {
	return "mmq::queueSet"
}
