FROM node:20-alpine AS build

WORKDIR /app

COPY . .

RUN npm install
RUN npm run build

FROM node:20-alpine AS final

WORKDIR /root

COPY --from=build /app/.next ./.next
COPY --from=build /app/next.config.ts ./next.config.ts
COPY --from=build /app/node_modules ./node_modules
COPY --from=build /app/package.json ./package.json
COPY --from=build /app/package-lock.json ./package-lock.json

EXPOSE 4000

CMD ["npm", "start"]