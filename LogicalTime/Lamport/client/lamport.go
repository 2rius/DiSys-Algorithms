package main

func (c *Client) incrementClock() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.clock++
}

func (c *Client) LamportMax(x uint32) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.clock > x {
		c.clock++
	} else {
		c.clock = x + 1
	}
}
