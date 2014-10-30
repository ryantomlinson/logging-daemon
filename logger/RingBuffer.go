package logger

// RingBuffer ...
type RingBuffer struct {
	inputChannel  <-chan string
	outputChannel chan string
}

// Run : start the ringbuffer
func (r *RingBuffer) Run() {
	for v := range r.inputChannel {
		select {
		case r.outputChannel <- v:
		default:
			<-r.outputChannel
			r.outputChannel <- v
		}
	}
	close(r.outputChannel)
}

// NewRingBuffer : initialises a new RingBuffer from the given channels
func NewRingBuffer(inputChannel <-chan string, outputChannel chan string) *RingBuffer {
	return &RingBuffer{inputChannel, outputChannel}
}
