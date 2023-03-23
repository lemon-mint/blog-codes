package main

type Solver struct {
	board [9][9]uint8

	x [9]uint16
	y [9]uint16
	g [9]uint16
}

func (s *Solver) init() {
	var i, j uint16
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			if s.board[i][j] != 0 {
				s.x[i] |= 1 << s.board[i][j]
				s.y[j] |= 1 << s.board[i][j]
				s.g[i/3*3+j/3] |= 1 << s.board[i][j]
			}
		}
	}
}

func (s *Solver) solve() bool {
	var i, j uint16
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			if s.board[i][j] == 0 {
				goto FOUND
			}
		}
	}
	return true
FOUND:
	for n := uint16(1); n <= 9; n++ {
		if s.x[i]&(1<<n) == 0 && s.y[j]&(1<<n) == 0 && s.g[i/3*3+j/3]&(1<<n) == 0 {
			s.x[i] |= 1 << n
			s.y[j] |= 1 << n
			s.g[i/3*3+j/3] |= 1 << n
			s.board[i][j] = uint8(n)
			if s.solve() {
				return true
			}
			s.x[i] &^= 1 << n
			s.y[j] &^= 1 << n
			s.g[i/3*3+j/3] &^= 1 << n
			s.board[i][j] = 0
		}
	}
	return false
}

func main() {
	var s Solver
	s.init()
	s.solve()
}
