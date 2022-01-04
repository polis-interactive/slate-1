
## Pi Install

- Install Go
- Install github.com/jgarff/rpi_ws281x
- Install github.com/go-gl/gl
    - go get -u github.com/go-gl/gl/v3.1/gles2
- Install github.com/go-gl/glfw
    - go get -u -tags=gles2 github.com/go-gl/glfw/v3.3/glfw
- Add go path to sudo
    - visudo, add secure_path="...:/usr/local/go/bin"
