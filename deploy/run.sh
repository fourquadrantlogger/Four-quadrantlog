APISERVER_PORT=10008 ./backend/fourquadrantlog &
cd vite-proj && npm run dev &
nginx -g 'daemon off;'