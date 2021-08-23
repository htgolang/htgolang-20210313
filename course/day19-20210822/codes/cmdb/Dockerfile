FROM golang:1.16.6 as builder

WORKDIR /opt/cmdb/

ADD . /opt/cmdb/

CMD go env GO111MODULE=on \
    go env -w GOPROXY=https://goproxy.cn,direct \
    go build .

FROM centos:7

WORKDIR /opt/cmdb/

COPY --from=builder /opt/cmdb/cmdb /opt/cmdb/cmdb
COPY --from=builder /opt/cmdb/conf /opt/cmdb/conf
COPY --from=builder /opt/cmdb/views /opt/cmdb/views
COPY --from=builder /opt/cmdb/static /opt/cmdb/static

ENTRYPOINT [ "/opt/cmdb/cmdb" ]