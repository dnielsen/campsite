FROM node:12-alpine

# Update packages.
RUN apk update

# Create the app directory.
WORKDIR /code

# Copy configs.
COPY . .

RUN yarn && \
  yarn build && \
  yarn cache clean

EXPOSE 3333

CMD ["yarn", "start"]
