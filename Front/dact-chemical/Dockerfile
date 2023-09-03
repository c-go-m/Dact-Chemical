FROM node:18-alpine as build-step

RUN mkdir -p /app
WORKDIR /app
COPY . /app

RUN npm install
RUN npm run build --prod

FROM nginx:1.23.1-alpine 
COPY --from=build-step /app/dist/dact-chemical /usr/share/nginx/html
