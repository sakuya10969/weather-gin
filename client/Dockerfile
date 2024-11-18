# 開発環境用のビルドステージ
FROM node:20-alpine AS development

WORKDIR /app

# パッケージインストール
COPY package.json package-lock.json ./
RUN npm install

# ソースコードをコピー
COPY . ./

# 開発用ホットリロードをサポート
CMD ["npm", "start"]

EXPOSE 3000

# 本番環境用のビルドステージ
FROM node:20-alpine AS build

WORKDIR /app

# パッケージインストールとビルド
COPY package.json package-lock.json ./
RUN npm install
COPY . ./
RUN npm run build

# 本番環境の軽量なNginxを利用
FROM nginx:alpine AS production

WORKDIR /usr/share/nginx/html

# ビルドされたファイルをコピー
COPY --from=build /app/build .

# カスタムNginx設定を追加（必要に応じて）
# COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]