package proxy

// last second of current period, split by interval
func last(now int64, interval int64) int64 {
	reminder := now % interval
	return now - reminder + interval - 1
}
