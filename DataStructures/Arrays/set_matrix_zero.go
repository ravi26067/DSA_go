package main

// Time : O(m*n) , Space: O(m+n)
//https://leetcode.com/problems/set-matrix-zeroes/description/

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	colmj := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				if j != 0 {
					matrix[0][j] = 0
				} else {
					colmj = 0
				}
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] != 0 {
				if matrix[i][0] == 0 || matrix[0][j] == 0 {
					matrix[i][j] = 0
				}
			}
		}
	}

	if matrix[0][0] == 0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
	if colmj == 0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

}

/*
// Time : O(m*n) , Space: O(1)

func setZeroes(matrix [][]int)  {
    m:=len(matrix)
    n:=len(matrix[0])
    rows:=make([]bool,m)
    colm:=make([]bool,n)

    for i:=range(m){
        for j:=range(n){
            if matrix[i][j]==0{
                rows[i]=true
                colm[j]=true
            }
        }
    }

    for i:=range(m){
        for j:=range(n){
            if rows[i] || colm[j]{
                matrix[i][j]=0
            }
        }
    }
}
*/
