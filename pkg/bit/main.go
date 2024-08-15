package main

const (
	jcsGwRunning = 1
	rmsGwRunning = 2
)

type Heartbeat struct {
	running int
}

func main() {
	h := &Heartbeat{}
	h.setHeartBeatRunning(jcsGwRunning)
	if !h.IsRunning(jcsGwRunning) || h.IsRunning(rmsGwRunning) {
		panic("jcs not running or rms is running")
	}
	h.setHeartBeatRunning(rmsGwRunning)
	if !h.IsRunning(rmsGwRunning) || !h.IsRunning(jcsGwRunning) {
		panic("jcs is not running or rms is not running")
	}

	h.setHeartBeatStopped(rmsGwRunning)
	if h.IsRunning(rmsGwRunning) || !h.IsRunning(jcsGwRunning) {
		panic("jcs not running or rms is running")
	}
	h.setHeartBeatStopped(jcsGwRunning)
	if h.IsRunning(jcsGwRunning) || h.IsRunning(jcsGwRunning) {
		panic("jcs is running or rms is running")
	}
}

func (h *Heartbeat) IsRunning(running int) bool {
	return (h.running & running) == running
}

func (h *Heartbeat) setHeartBeatRunning(running int) {
	h.running |= running
	return
}

func (h *Heartbeat) setHeartBeatStopped(running int) {
	h.running ^= running
}
