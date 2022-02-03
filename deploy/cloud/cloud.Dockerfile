# debian lts with node lts
FROM node:14.15.4-buster

WORKDIR /usr/src/app

RUN npm install --global --unsafe-perm sequelize-cli

COPY backend-cloud/package*.json ./backend-cloud/

RUN npm install --prefix ./backend-cloud

COPY ./backend-cloud ./backend-cloud

EXPOSE 3030

ENV NODE_ENV=test

CMD ["npm", "run", "--prefix", "./backend-cloud", "start"]
