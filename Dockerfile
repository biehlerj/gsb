FROM gnome-settings-backup-test:latest
WORKDIR /gsb
COPY . /gsb
RUN cd /gsb
RUN go mod download
RUN go build
