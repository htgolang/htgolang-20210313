if [[ "x" == "x${GOROOT}" ]]; then
    export GOROOT=/usr/local/go
    export GOPATH=${HOME}/go
    export PATH=${PATH}:${GOROOT}/bin:${GOPATH}/bin
fi
