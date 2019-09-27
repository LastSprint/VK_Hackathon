FROM golang:latest 	

RUN git clone https://github.com/LastSprint/VK_Hackathon.git /app
WORKDIR /app/src/mocker

ENV GOPATH=/app/

RUN go build
RUN ./suncity