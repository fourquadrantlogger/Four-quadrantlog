FROM golang
COPY backend backend
ENV GOPROXY=https://proxy.golang.com.cn,direct
RUN cd backend &&\
    go build  -o fourquadrantlog app.go
CMD ./backend/fourquadrantlog