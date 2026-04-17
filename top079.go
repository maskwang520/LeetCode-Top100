package main

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {

			}
		}
	}
}

func iterate(board [][]byte, word string, list []byte) bool {

	for i := 0; i < len(word); i++ {
		list = append(list, word[i])

	}
}
