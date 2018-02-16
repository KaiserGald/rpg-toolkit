FROM golang:latest
ARG app_name
ARG src_path
ARG install_path
ENV GOBIN /go/bin
ENV BINARY_NAME $app_name
ENV INSTALLPATH $install_path
RUN mkdir /srv/$BINARY_NAME
RUN mkdir -p $src_path/unlicht-server
ADD . $src_path/unlicht-server
RUN cd $src_path/unlicht-server && make all
ENTRYPOINT $INSTALLPATH/$BINARY_NAME
EXPOSE 8080
EXPOSE 8081
