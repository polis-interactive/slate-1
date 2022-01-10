
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
- Install nymea
    - Add nymea repo, public key, insall
        - https://github.com/nymea/pi-gen/blob/berrylan-bullseye/stage2/04-berrylan/00-run-chroot.sh
        - https://github.com/nymea/nymea-networkmanager
    - Possible enable the service
    - Might need to add these hot fixes
        - https://unix.stackexchange.com/questions/382031/bluez-error-setting-privacy-on-raspbian
        - https://raspberrypi.stackexchange.com/questions/40839/sap-error-on-bluetooth-service-status
- Make Service
    - https://superuser.com/questions/544399/how-do-you-make-a-systemd-service-as-the-last-service-on-boot
    - Remove nymea from target graphical to target slate, resymlink