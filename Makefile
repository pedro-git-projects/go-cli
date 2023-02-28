OBJ_NAME = 
LDFLAGS = 
PKG_NAME =
ARCH =
OS = 
greeter:
	$(eval PKG_NAME+= "greeter")
	$(eval ARCH += "amd64")
	$(eval OS += "linux")
	$(eval LDFLAGS += "-w -s")
	$(eval OBJ_NAME += greeter)
	cd ./cmd/$(PKG_NAME); GOOS=$(OS) GOARCH=$(ARCH) go build -v -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../../bin 
doc:
	cd ./cmd/; godoc -http=:6060
