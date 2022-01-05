
## Pi Install

- Install Go
    - https://www.jeremymorgan.com/tutorials/raspberry-pi/install-go-raspberry-pi/
    - Install armv6l
- Install github.com/jgarff/rpi_ws281x
- Install github.com/go-gl/gl
    - sudo go get -u github.com/go-gl/gl/v3.1/gles2
- Install github.com/go-gl/glfw
    - sudo go get -u -tags=gles2 github.com/go-gl/glfw/v3.3/glfw
- Add go path to sudo
    - visudo, add secure_path="...:/usr/local/go/bin"
- Build
    - go build ./cmd/runApplication/main.go 