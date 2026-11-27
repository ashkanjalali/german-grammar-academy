FROM node:22-alpine
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
# In the compose network, the API service is reachable by its service name.
RUN sed -i 's|http://localhost:8080|http://api:8080|' proxy.conf.json
EXPOSE 4200
CMD ["npm","start","--","--host","0.0.0.0"]
