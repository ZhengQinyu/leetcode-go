package _202001

func solveSudoku(board [][]byte) {
	n := len(board)
	m := make([]bool, n*n*3)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if board[x][y] != '.' {
				m[n*x+int(board[x][y]-'1')] = true                 // 行
				m[81+n*y+int(board[x][y]-'1')] = true              // 列
				m[162+n*((x/3)*3+y/3)+int(board[x][y]-'1')] = true // 格
			}
		}
	}
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if board[x][y] == '.' {
				for z := 0; z < n; z++ {
					if m[n*x+z] || m[81+n*y+z] || m[162+n*((x/3)*3+y/3)+z] {
						continue
					} else {
						board[x][y] = byte(z + '1')
						m[n*x+z] = true
						m[81+n*y+z] = true
						m[162+n*((x/3)*3+y/3)+z] = true
						break
					}
				}
			}
		}
	}
}
