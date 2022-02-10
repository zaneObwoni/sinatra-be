
VERSION=`git describe --tags`
XFLAGS="-X 'main.Version=${VERSION}'"
MONGODB_URL=mongodb+srv://sinatra:sinatra@cluster0.fdmuo.mongodb.net/myFirstDatabase?retryWrites=true&w=majority

build:
	@ go build -ldflags=${XFLAGS} -o "bin/sinatra"

dev:
	@ go run main.go \
		--log-dir="./" \
		--log-level="debug" \
		--skip-downloads

test:
	@ go test ./...

version:
	@ bin/sinatra --version