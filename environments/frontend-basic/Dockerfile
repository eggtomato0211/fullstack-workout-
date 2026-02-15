FROM node:20-alpine

WORKDIR /app

# Install basic tools
RUN apk add --no-cache git

# Default command: install deps and start dev server
CMD ["sh", "-c", "if [ -f package.json ]; then npm install && npm run dev -- --host 0.0.0.0; else echo 'No project found. Create one with: npm create vite@latest . -- --template react' && tail -f /dev/null; fi"]
