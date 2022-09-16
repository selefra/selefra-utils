package channel_util

func SendSliceToChannel[T any](slice []T, channel chan<- any) {
	if len(slice) == 0 {
		return
	}
	for _, x := range slice {
		channel <- x
	}
}
