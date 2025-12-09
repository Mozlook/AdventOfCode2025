package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type TileLocation struct {
	x int
	y int
}

type ScaledLocation struct {
	x, y int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func compressCoords(vals []int) ([]int, map[int]int) {
	if len(vals) == 0 {
		return nil, nil
	}

	sort.Ints(vals)

	u := vals[:0]
	for _, v := range vals {
		if len(u) == 0 || u[len(u)-1] != v {
			u = append(u, v)
		}
	}

	comp := make([]int, 0, len(u)*2)
	for i, v := range u {
		comp = append(comp, v)
		if i+1 < len(u) {
			next := u[i+1]
			if next > v+1 {
				comp = append(comp, v+1)
			}
		}
	}

	idx := make(map[int]int, len(comp))
	for i, v := range comp {
		idx[v] = i
	}

	return comp, idx
}

func markOutside(blocked [][]int) [][]bool {
	h := len(blocked)
	w := len(blocked[0])

	outside := make([][]bool, h)
	for y := range h {
		outside[y] = make([]bool, w)
	}

	var stack []ScaledLocation

	for y := range h {
		for x := range w {
			if (y == 0 || y == h-1 || x == 0 || x == w-1) && blocked[y][x] == 1 {
				outside[y][x] = true
				stack = append(stack, ScaledLocation{x: x, y: y})
			}
		}
	}

	directions := [][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	for len(stack) > 0 {
		n := len(stack)
		v := stack[n-1]
		stack = stack[:n-1]

		for _, d := range directions {
			nx := v.x + d[0]
			ny := v.y + d[1]

			if nx >= 0 && nx < w && ny >= 0 && ny < h {
				if blocked[ny][nx] == 1 && !outside[ny][nx] {
					outside[ny][nx] = true
					stack = append(stack, ScaledLocation{x: nx, y: ny})
				}
			}
		}
	}

	return outside
}

func buildPrefix(blocked [][]int) [][]int {
	h := len(blocked)
	w := len(blocked[0])

	pref := make([][]int, h+1)
	for y := 0; y <= h; y++ {
		pref[y] = make([]int, w+1)
	}

	for y := range h {
		for x := range w {
			pref[y+1][x+1] = blocked[y][x] +
				pref[y][x+1] +
				pref[y+1][x] -
				pref[y][x]
		}
	}

	return pref
}

func badInRect(pref [][]int, x1, y1, x2, y2 int) int {
	return pref[y2+1][x2+1] -
		pref[y1][x2+1] -
		pref[y2+1][x1] +
		pref[y1][x1]
}

func maxRectangleArea(tiles []TileLocation, xIndex, yIndex map[int]int, pref [][]int) int {
	maxArea := 0
	n := len(tiles)

	for i := range n {
		for j := i + 1; j < n; j++ {
			t1 := tiles[i]
			t2 := tiles[j]

			ix1 := xIndex[t1.x]
			iy1 := yIndex[t1.y]
			ix2 := xIndex[t2.x]
			iy2 := yIndex[t2.y]

			x1, x2 := ix1, ix2
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			y1, y2 := iy1, iy2
			if y1 > y2 {
				y1, y2 = y2, y1
			}

			bad := badInRect(pref, x1, y1, x2, y2)
			if bad == 0 {
				dx := abs(t1.x-t2.x) + 1
				dy := abs(t1.y-t2.y) + 1
				area := dx * dy
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

func main() {
	rawInput, err := os.Open("../input.txt")
	if err != nil {
		fmt.Printf("input error: %v\n", err)
		return
	}
	defer rawInput.Close()

	var tiles []TileLocation
	scanner := bufio.NewScanner(rawInput)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			fmt.Println("bad line:", line)
			continue
		}

		x, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			fmt.Println("parsing error x:", err)
			return
		}
		y, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			fmt.Println("parsing error y:", err)
			return
		}

		tiles = append(tiles, TileLocation{x: x, y: y})
	}

	if len(tiles) == 0 {
		fmt.Println("no tiles")
		return
	}

	minX, maxX := tiles[0].x, tiles[0].x
	minY, maxY := tiles[0].y, tiles[0].y

	for i := 1; i < len(tiles); i++ {
		if tiles[i].x < minX {
			minX = tiles[i].x
		}
		if tiles[i].x > maxX {
			maxX = tiles[i].x
		}
		if tiles[i].y < minY {
			minY = tiles[i].y
		}
		if tiles[i].y > maxY {
			maxY = tiles[i].y
		}
	}

	var xs, ys []int
	xs = make([]int, 0, len(tiles)+2)
	ys = make([]int, 0, len(tiles)+2)

	for _, t := range tiles {
		xs = append(xs, t.x)
		ys = append(ys, t.y)
	}
	xs = append(xs, minX-1, maxX+1)
	ys = append(ys, minY-1, maxY+1)

	compX, xIndex := compressCoords(xs)
	compY, yIndex := compressCoords(ys)

	w := len(compX)
	h := len(compY)

	blocked := make([][]int, h)
	for y := range h {
		blocked[y] = make([]int, w)
		for x := range w {
			blocked[y][x] = 1
		}
	}

	loop := make([]ScaledLocation, len(tiles))
	for i, tile := range tiles {
		ix := xIndex[tile.x]
		iy := yIndex[tile.y]
		loop[i] = ScaledLocation{x: ix, y: iy}
	}

	n := len(loop)
	for i := range n {
		a := loop[i]
		b := loop[(i+1)%n]

		if a.x != b.x && a.y != b.y {
			fmt.Println("bledne dane: skos")
		}

		if a.x == b.x {
			x := a.x
			y1, y2 := a.y, b.y
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for y := y1; y <= y2; y++ {
				blocked[y][x] = 0
			}
		} else if a.y == b.y {
			y := a.y
			x1, x2 := a.x, b.x
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for x := x1; x <= x2; x++ {
				blocked[y][x] = 0
			}
		}
	}

	outside := markOutside(blocked)

	for i := range h {
		for j := range w {
			if outside[i][j] {
				blocked[i][j] = 1
			} else {
				blocked[i][j] = 0
			}
		}
	}

	pref := buildPrefix(blocked)

	res := maxRectangleArea(tiles, xIndex, yIndex, pref)
	fmt.Println(res)
}
