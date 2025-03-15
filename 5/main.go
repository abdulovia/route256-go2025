package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Box struct {
	Name       string
	Area       int
	InnerBoxes map[string]interface{}
}

func parseMatrix(matrix []string, rows, cols int) map[string]interface{} {
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	boxes := make(map[string]interface{})

	var dfs func(x, y int) (int, int, []string)

	// Обход вложенных коробок
	dfs = func(x, y int) (int, int, []string) {
		// Логика обхода и подсчета
		return 0, 0, nil
	}

	// Основная логика парсинга матрицы
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] && matrix[i][j] == '+' {
				// Найдена новая коробка
				width, height, content := dfs(i, j)
				if len(content) > 0 {
					boxes[content[0]] = map[string]interface{}{}
				} else {
					boxes[content[0]] = width * height
				}
			}
		}
	}

	return boxes
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for inp := 0; inp < t; inp++ {
		var rows, cols int
		fmt.Fscan(in, &rows, &cols)
		matrix := make([]string, rows)
		for i := 0; i < rows; i++ {
			line, _ := in.ReadString('\n')
			matrix[i] = strings.TrimSpace(line)
		}

		result := make(map[string]interface{})
		if len(matrix) > 0 {
			result = parseMatrix(matrix, rows, cols)
		}
		jsonResult, _ := json.MarshalIndent(result, "", "  ")
		fmt.Fprintln(out, string(jsonResult))
	}
}
