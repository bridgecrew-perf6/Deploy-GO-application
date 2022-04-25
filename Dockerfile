FROM gcr.io/masec-docker-images/linux-sap-nwrfcsdk:latest

# Define workdir
WORKDIR /app/src
RUN yum update -y
RUN yum install wget tar git gcc -y
RUN wget https://golang.org/dl/go1.17.3.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.17.3.linux-amd64.tar.gz
RUN yum remove wget tar -y
ENV GOROOT=/usr/local/go
ENV PATH=$PATH:/usr/local/go/bin
ENV GOBIN=/usr/local/go/bin
RUN mkdir -p /root/go/bin /root/go/src /root/go/pkg
ENV GOPATH=/root/go
ENV PATH=$PATH:/root/go/bin
# Install required applications
RUN go version
RUN go get github.com/cosmtrek/air
COPY .air.toml /root/
# Copy the Source Code
COPY ./src ./ 
# COPY ./.env ../
# Install the dependencies
RUN go get
# Run the program
CMD [ "air", "-c", "/root/.air.toml" ]