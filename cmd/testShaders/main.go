package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/polis-interactive/slate-1/internal/types/shader"
	"log"
	"runtime"
	"time"
	"unsafe"
)

/*
	leaving this as a pseudo documentaiton; it doesn't work anymore after
	figuring out types/shader
*/

const programName = "Slate-1"
const shaderName = "slate-1"

const windowWidth = 800
const windowHeight = 600

func init() {
	// GLFW event handling must be run on the main OS thread
	runtime.LockOSThread()
}

func main() {

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to inifitialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "slate-1", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// Initialize Glow (go function bindings)
	if err := gl.Init(); err != nil {
		panic(err)
	}

	window.SetKeyCallback(keyCallback)

	err = programLoop(window)
	if err != nil {
		log.Fatal(err)
	}
}

func createFillRect() uint32 {

	vertices := []float32{
		// bottom left
		-1.0, -1.0, 0.0, // position
		0.0, 0.0, 0.0, // Color

		// bottom right
		1.0, -1.0, 0.0,
		0.0, 0.0, 0.0, // Color

		// top right
		1.0, 1.0, 0.0,
		0.0, 0.0, 0.0, // Color

		// top left
		-1.0, 1.0, 0.0,
		0.0, 0.0, 0.0, // Color
	}

	indices := []uint32{
		0, 1, 2, 3,
	}

	var VAO uint32
	gl.GenVertexArrays(1, &VAO)

	var VBO uint32
	gl.GenBuffers(1, &VBO)

	var EBO uint32;
	gl.GenBuffers(1, &EBO)

	// Bind the Vertex Array Object first, then bind and set vertex buffer(s) and attribute pointers()
	gl.BindVertexArray(VAO)

	// copy vertices data into VBO (it needs to be bound first)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	// copy indices into element buffer
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	// position
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	// unbind the VAO (safe practice so we don't accidentally (mis)configure it later)
	gl.BindVertexArray(0)

	return VAO
}

func programLoop(window *glfw.Window) error {

	fileName, err := shader.GetShaderQualifiedPath(shaderName, programName)
	if err != nil {
		log.Fatalln("Couldn't find shader")
		return err
	}
	log.Println(fileName)

	shaderProgram, err := shader.NewProgram(fileName)
	if err != nil {
		return err
	}
	defer shaderProgram.Delete()

	rect := createFillRect()

	start := time.Now()

	for !window.ShouldClose() {
		// poll events and call their registered callbacks
		glfw.PollEvents()

		// perform rendering
		gl.ClearColor(0.2, 0.2, 0.2, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// draw loop

		// draw triangle

		shaderProgram.Use()
		duration := time.Since(start)
		shaderProgram.SetUniform1f("time", float32(duration.Seconds()))
		shaderProgram.SetUniform2fv("resolution", []float32{windowWidth, windowHeight}, 1)
		gl.BindVertexArray(rect)
		gl.DrawElements(gl.TRIANGLE_FAN, 4, gl.UNSIGNED_INT, unsafe.Pointer(nil))
		gl.BindVertexArray(0)

		// end of draw loop

		// swap in the rendered buffer
		window.SwapBuffers()
	}

	return nil
}

func keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action,
	mods glfw.ModifierKey) {

	// When a user presses the escape key, we set the WindowShouldClose property to true,
	// which closes the application
	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}