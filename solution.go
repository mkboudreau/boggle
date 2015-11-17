package main

type Solution []*Coord

func (s *Solution) addCoord(c *Coord) bool {
	for _, self := range *s {
		if self.x == c.x && self.y == c.y {
			return false
		}
	}
	*s = append(*s, c)
	return true
}
func (s Solution) Len() int {
	return len(s)
}
func (s Solution) Last() *Coord {
	return s[s.Len()-1]
}
