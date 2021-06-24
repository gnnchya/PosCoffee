package kafka

import "fmt"

func makeHostKafka(host []string) (hosts []string) {
	hosts = make([]string, len(host))
	for i, v := range host {
		hosts[i] = fmt.Sprintf("%s", v)
	}
	return hosts
}
