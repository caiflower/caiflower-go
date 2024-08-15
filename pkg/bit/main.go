package main

const (
	jRunning = 1
	rRunning = 2
)

type Heartbeat struct {
	running int
}

func main() {
	h := &Heartbeat{}
	h.setHeartBeatRunning(jRunning)
	if !h.IsRunning(jRunning) || h.IsRunning(rRunning) {
		panic("jcs not running or rms is running")
	}
	h.setHeartBeatRunning(rRunning)
	if !h.IsRunning(rRunning) || !h.IsRunning(jRunning) {
		panic("jcs is not running or rms is not running")
	}

	h.setHeartBeatStopped(rRunning)
	if h.IsRunning(rRunning) || !h.IsRunning(jRunning) {
		panic("jcs not running or rms is running")
	}
	h.setHeartBeatStopped(jRunning)
	if h.IsRunning(jRunning) || h.IsRunning(jRunning) {
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
