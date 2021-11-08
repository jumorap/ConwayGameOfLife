package main

import (
	"math"
	"math/rand"
)
import "fmt"
import rl "github.com/gen2brain/raylib-go/raylib"

// widthHeight Define the world size as a square
const widthHeight int32 = 800

// cyclicStructure Define boolean values to generate an example of cyclic structures
var cyclicStructure = []bool{
	false, false, false, false, false, false, false, false, false, false, false,
	false, false, false, false, false, true, false, false, false, false, false,
	false, false, false, false, false, true, false, false, false, false, false,
	false, false, false, false, false, true, false, false, false, false, false,
	false, false, false, false, false, false, false, false, false, false, false,
	false, true, true, true, false, false, false, true, true, true, false,
	false, false, false, false, false, false, false, false, false, false, false,
	false, false, false, false, false, true, false, false, false, false, false,
	false, false, false, false, false, true, false, false, false, false, false,
	false, false, false, false, false, true, false, false, false, false, false,
	false, false, false, false, false, false, false, false, false, false, false,
}

// staticStructure Define boolean values to generate an example of static structures
var staticStructure = []bool{
	false, false, false, false, false, false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, true, false, false, false, false, false, false,
	false, false, false, false, false, true, false, true, false, false, false, false, false,
	false, false, false, false, false, true, false, true, false, false, false, false, false,
	false, false, false, false, false, false, true, false, false, false, false, false, false,
	false, false, true, true, false, false, false, false, false, true, true, false, false,
	false, true, false, false, true, false, false, false, true, false, false, true, false,
	false, false, true, true, false, false, false, false, false, true, true, false, false,
	false, false, false, false, false, false, true, false, false, false, false, false, false,
	false, false, false, false, false, true, false, true, false, false, false, false, false,
	false, false, false, false, false, true, false, true, false, false, false, false, false,
	false, false, false, false, false, false, true, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false, false, false, false, false, false,
}

// infiniteStructure Define boolean values to generate an example of infinite structures
var infiniteStructure = []bool{
	false, false, false, false, false, false, false, false, false, true, false,
	false, false, false, false, false, false, false, false, true, true, false,
	false, false, false, false, false, false, false, false, true, false, true,
	false, false, false, false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, true, false, false, false, false,
	false, false, false, false, false, true, true, false, false, false, false,
	false, false, false, false, false, true, false, true, false, false, false,
	false, false, false, false, false, false, false, false, false, false, false,
	false, false, false, true, false, false, false, false, false, false, false,
	false, false, true, true, false, false, false, false, false, false, false,
	false, false, true, false, true, false, false, false, false, false, false,
}


// GameOfLife saves the current state of the world
type GameOfLife struct {
	WorldWidth int32
	WorldHeight int32
	World [][]bool
	CurrentGeneration int
	Canvas rl.RenderTexture2D
}

func main()  {
	// Select game type to show
	election := menu()

	// Initialize the game and use a max of 10 frames per second by default
	rl.InitWindow(widthHeight, widthHeight, "Conway's Game of Life")
	rl.SetTargetFPS(10)

	gameOfLife := GameOfLife{}

	// Show the selected game
	switch election {
	case 0:
		gameOfLife.Init()
	case 1:
		gameOfLife.GenerateStructure(cyclicStructure)
	case 2:
		gameOfLife.GenerateStructure(staticStructure)
	case 3:
		gameOfLife.GenerateStructure(infiniteStructure)
	default:
		gameOfLife.Init()
    }

	// Main game loop drawing and updating generations
	for !rl.WindowShouldClose() {
		gameOfLife.Draw()
        gameOfLife.Update()
    }

	// Release resources
	rl.UnloadTexture(gameOfLife.Canvas.Texture)
	rl.CloseWindow()
}

// Init Define the initial state of the world
func (m *GameOfLife) Init() {
	m.Create(widthHeight / 10)
	// Cell vectors
	totalCells := m.WorldWidth * m.WorldHeight
    m.World = make([][]bool, 2)
	m.World[0] = make([]bool, totalCells)
	m.World[1] = make([]bool, totalCells)

    for i := int32(0); i < totalCells; i++ {
		if rand.Intn(2) == 0 {
			// Set the cell to alive
			m.World[0][i] = true
		} else {
			// Set the cell to dead
			m.World[0][i] = false
		}
		// No live cells at the beginning
        m.World[1][i] = false
    }
}

// GenerateStructure Set special games structures as initial state
func (m *GameOfLife) GenerateStructure(boolStructure []bool) {
	m.World = make([][]bool, 2)
	m.World[0] = boolStructure

	// Get World[0] Size and set it at Create
	lenWorld := len(m.World[0])
	xy := int32(math.Sqrt(float64(lenWorld)))
	m.Create(xy)

	// Fill the world with dead cells
	for i := 0; i <= lenWorld; i++ {
		m.World[1] = append(m.World[1], false)
	}
}

// Draw draws the current state of the world
func (m *GameOfLife) Draw() {
	// Clear the screen
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	// Initialize the texture and world is painted
	rl.BeginTextureMode(m.Canvas)
	rl.ClearBackground(rl.Black)

	// Go through the world and paint the cells
	for x := int32(0); x < m.WorldWidth; x++ {
		for y := int32(0); y < m.WorldHeight; y++ {
			cellIndexAt := y * m.WorldWidth + x
			if m.World[m.CurrentGeneration][cellIndexAt] {
				rl.DrawPixel(x, y, rl.NewColor(255, 255, 255, 255))
			}
		}
	}

	// Generate the texture and is painted by GPU
	rl.EndTextureMode()
	rl.DrawTexturePro(m.Canvas.Texture, rl.NewRectangle(0, 0, float32(m.Canvas.Texture.Width), float32(m.Canvas.Texture.Height)), rl.NewRectangle(0, 0, float32(widthHeight), float32(widthHeight)), rl.NewVector2(float32(0), float32(0)), 0, rl.RayWhite)
	rl.EndDrawing()
}

// Update updates the game state for generate the next generation
func (m *GameOfLife) Update() {
	// Generate a vector to save next generation
	nextGeneration := m.World[(m.CurrentGeneration + 1) % 2]

	// Loop through all cells
	for x := int32(0); x < m.WorldWidth; x++ {
		for y := int32(0); y < m.WorldHeight; y++ {
			cellIndex := y * m.WorldWidth + x
			liveNeighbors := m.GetLive(x, y)
            cellLive := m.World[m.CurrentGeneration][cellIndex]

			// Define the rules for the cells, to define the GAME OF LIFE
			// Rule #1: A cell is "dead" when have 3 neighbors alive around
			// Rule #2: A cell is "alive" with less of 2 of more than 3 neighbors alive
			if (cellLive && liveNeighbors == 2) || liveNeighbors == 3 {
                nextGeneration[cellIndex] = true
            } else {
                nextGeneration[cellIndex] = false
            }
		}
	}
	// Update the current generation
	m.CurrentGeneration = (m.CurrentGeneration + 1) % 2
}

// GetLive returns the number of live neighbors around a cell
func (m *GameOfLife) GetLive(x int32, y int32) int {
	var howManyLive = 0

	type position struct {
        x int32
        y int32
    }

	// Define the positions of the neighbors
	neighbors := []position{
        {x: -1, y: -1},
        {x: 0, y: -1},
        {x: 1, y: -1},
        {x: -1, y: 0},
        {x: 1, y: 0},
        {x: -1, y: 1},
        {x: 0, y: 1},
        {x: 1, y: 1},
    }

	// Loop through all neighbors
	// If the neighbor is alive, add 1 to the counter
	// If the neighbor is outside the world, ignore it
	for _, c := range neighbors {
		_x := mod(x + c.x, m.WorldWidth)
		_y := mod(y + c.y, m.WorldHeight)

		if m.World[m.CurrentGeneration][_y * m.WorldWidth + _x] {
            howManyLive++
        }
	}

	return howManyLive
}

// Create Define how many cells wide and high the world is. After is rendered
func (m *GameOfLife) Create(xy int32) {
	m.WorldWidth = xy
	m.WorldHeight = xy
	m.Canvas = rl.LoadRenderTexture(m.WorldWidth, m.WorldHeight)
	m.CurrentGeneration = 0
}

// mod returns the modulus of a and b
func mod(a int32, b int32) int32 {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

// hello World prints name's tool
func hello()  {
	fmt.Print("\n _______  _______  __   __  _______    _______  _______    ___      ___   _______  _______" +
		"\n|       ||   _   ||  |_|  ||       |  |       ||       |  |   |    |   | |       ||       |" +
		"\n|    ___||  |_|  ||       ||    ___|  |   _   ||    ___|  |   |    |   | |    ___||    ___|" +
		"\n|   | __ |       ||       ||   |___   |  | |  ||   |___   |   |    |   | |   |___ |   |___ " +
		"\n|   ||  ||       ||       ||    ___|  |  |_|  ||    ___|  |   |___ |   | |    ___||    ___|" +
		"\n|   |_| ||   _   || ||_|| ||   |___   |       ||   |      |       ||   | |   |    |   |___ " +
		"\n|_______||__| |__||_|   |_||_______|  |_______||___|      |_______||___| |___|    |_______|\n")
}

// menu shows options to the user and returns the option selected
func menu() int16 {
	hello()
	var election int16
	fmt.Println("\nWhich game of life do you want to visualize: ")
	fmt.Print(
		"[0] Random game (80 x 80)" +
		"\n[1] Cyclic game (11 x 11)" +
			"\n[2] Static game (13 x 13)" +
			"\n[3] Infinite game (11 x 11)" +
			"\n-> ")
	fmt.Scanln(&election)

	return election
}
