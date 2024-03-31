# Use the official Hugo image as a parent image
FROM klakegg/hugo:latest as builder

# Copy the content of your site
COPY . /site
WORKDIR /site

# Build your site
RUN hugo

# Use Nginx to serve the site
FROM nginx:alpine
COPY --from=builder /site/public /usr/share/nginx/html
