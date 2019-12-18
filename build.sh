build_package() {
    # static builds http://blog.wrouesnel.com/articles/Totally%20static%20Go%20builds/
    CGO_ENABLED=0 GOOS=$1 GOARCH=$2 go build -o terraform-provider-subregcz_$1_$2 -a -ldflags '-extldflags "-static"' .
}

build_package linux amd64
build_package darwin amd64
