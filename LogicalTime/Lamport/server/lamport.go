package main

func (s *Server) incrementClock() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clock++
}

func (s *Server) LamportMax(x uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.clock > x {
		s.clock++
	} else {
		s.clock = x + 1
	}
}
