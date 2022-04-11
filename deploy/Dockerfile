FROM golang as BACKEND
WORKDIR /opt/project
COPY backend backend
ENV GOPROXY=https://proxy.golang.com.cn,direct
RUN pwd && cd backend &&\
    go build  -o fourquadrantlog app.go
FROM node:17-alpine
WORKDIR /opt/project
COPY --from=BACKEND /opt/project/backend/fourquadrantlog backend/fourquadrantlog
COPY vite-proj vite-proj
RUN pwd && cd vite-proj &&\
    npm install
COPY vue-web vue-web
RUN cd vue-web &&\
    npm install
