FROM node:20-alpine AS build
WORKDIR /app

ARG VITE_CLASSIFIER_API_URL=http://localhost:8080
ENV VITE_CLASSIFIER_API_URL=${VITE_CLASSIFIER_API_URL}

COPY package.json tsconfig.json vite.config.ts ./
COPY src ./src
COPY index.html ./

RUN npm install
RUN npm run build

FROM nginx:1.27-alpine
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80
