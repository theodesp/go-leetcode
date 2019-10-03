package array

/*
 Given a non-negative integer numRows,
 generate the first numRows of Pascal's triangle.
 */
/* Solution
 We notice that the i,j row is calculated by taking the sum of the upper left and
upper right cell. So create each cell by adding the above and above left cell.
If the cell is the first or last we put 1.
 */
func GeneratePascal(numRows int) [][]int {
	result := [][]int{}
	// For each row
	for i:= 0;i<numRows;i +=1 {
		// Every line has number of integers
		// equal to line number
		result = append(result, make([]int, i+1, i+1))
		for j:=0;j<=i;j+=1 {
			// First or last lines get an 1
			if j == 0 || j == i {
				result[i][j] = 1
			} else {
				// Other values are sum of values just
				// above and left of above
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
	}
	return result
}
