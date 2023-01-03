FROM node:19

WORKDIR /usr/src/app
COPY package.json yarn.lock .
RUN yarn install

COPY *.js .
CMD ["yarn", "start"]
