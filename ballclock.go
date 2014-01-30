package main

import "fmt"
import "flag"

const MinutesPerDay int = 1440
const TrackOneSize int = 5
const TrackTwoSize int = 12
const TrackThreeSize int = 12

type ClockData struct {
	track1      []int
	track2      []int
	track3      []int
	queue       []int
	queueLength int
}

func ElevateBall(data *ClockData) bool {
	var ball = data.queue[0]
	data.queue = data.queue[1:]
	data.track1 = append(data.track1, ball)

	if len(data.track1) == TrackOneSize {
		// Move `ball` to track2
		// Return first 4 balls to queue in reverse order.
		for i := TrackOneSize - 2; i > -1; i-- {
			data.queue = append(data.queue, data.track1[i])
		}
		data.track2 = append(data.track2, data.track1[4])
		data.track1 = make([]int, 0)
	}

	if len(data.track2) == TrackTwoSize {
		// Move `ball` to track3
		// Return first 11 balls to queue in reverse order.
		for i := TrackTwoSize - 2; i > -1; i-- {
			data.queue = append(data.queue, data.track2[i])
		}
		data.track3 = append(data.track3, data.track2[11])
		data.track2 = make([]int, 0)
	}

	if len(data.track3) == TrackThreeSize {
		// Hour track is full.
		// Return 12 balls to queue in reverse order
		for i := TrackThreeSize - 1; i > -1; i-- {
			data.queue = append(data.queue, data.track3[i])
		}
		data.track3 = make([]int, 0)
	}
	return true
}

func CheckQueue(data *ClockData) bool {
	if len(data.queue) != data.queueLength {
		return false
	}
	for i := 0; i < data.queueLength; i++ {
		if data.queue[i] != i+1 {
			return false
		}
	}
	return true
}

func main() {
	var queueLength = flag.Int("queuelength", 27, "27 <= queueLength <= 127")
	flag.Parse()
	fmt.Printf("Queue length is %d\n", *queueLength)

	// initialize queue
	var queue = make([]int, *queueLength)
	for i := 0; i < *queueLength; i++ {
		queue[i] = i + 1
	}

	data := ClockData{
		track1:      make([]int, 0),
		track2:      make([]int, 0),
		track3:      make([]int, 0),
		queue:       queue,
		queueLength: *queueLength,
	}

	var minutes int
	for minutes = 1; ElevateBall(&data); minutes++ {
		if CheckQueue(&data) {
			break
		}
	}

	fmt.Printf("The cycle completed in %d days.\n", minutes/MinutesPerDay)
}
