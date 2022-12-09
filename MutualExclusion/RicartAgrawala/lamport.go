package main

func (p *Peer) incrementClock() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.clock++
}

func (p *Peer) LamportMax(x uint32) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.clock > x {
		p.clock++
	} else {
		p.clock = x + 1
	}
}
