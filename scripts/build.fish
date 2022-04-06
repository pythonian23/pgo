echo Running from $(pwd)

function build-on -a os arch ver
    set file "-$ver-$os_$arch"
    if test "$os" = "windows"
        set file $file.exe
    else if test "$os" = "js"
        set file $file.asm
    end
    GOOS=$os GOARCH=$arch go build -o=./build/pgo$file ./cmd/pgo
    echo "build/pgo$file"
    GOOS=$os GOARCH=$arch go build -o=./build/pgo-discord$file ./cmd/pgo-discord
    echo "build/pgo-discord$file"
end

gofmt -w .
goimports -w .

set ver $argv[1]
set platforms "linux amd64
linux arm
linux arm64
android arm
android arm64
windows amd64
js wasm"

for platform in $(string split \n $platforms)
    build-on $(string split " " $platform) $ver
end
