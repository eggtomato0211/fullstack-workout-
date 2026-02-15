FROM node:20-alpine

WORKDIR /app

# Install basic tools
RUN apk add --no-cache git

# Default command: install deps and start dev server
CMD ["sh", "-c", "if [ -f package.json ]; then npm install && npm run dev; else echo 'No project found. Create one with: npx create-next-app@latest .' && tail -f /dev/null; fi"]
