echo Running from $(pwd)
echo Version:
read ver

function build-on -a os arch ver
    set file "-$ver-$os""_$arch"
    if test "$os" = "windows"
        set file $file.exe
    else if test "$os" = "js"
        set file $file.wasm
    end
    GOOS=$os GOARCH=$arch go build -o=./bin/pgo$file ./cmd/pgo && echo SUCCESS
    echo "build/pgo$file"
end

gofmt -w .
goimports -w .

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
