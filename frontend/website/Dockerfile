FROM node:lts as builder
WORKDIR /app
COPY . .

ARG BACKEND_DOMAIN
ENV REACT_APP_BACKEND_DOMAIN ${BACKEND_DOMAIN}

RUN npm ci
RUN npm run build


FROM nginx:alpine
ENV NODE_ENV production
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 8000