# 学習ガイド - 環境構築 & 進め方

## 1. 全体の流れ

```
1. GitHubの問題 (mdファイル) を読む
2. Docker環境を起動する
3. workspace/ 内でコードを実装する
4. ブラウザ (localhost:3000) で動作確認する
5. 完成したら solutions/ にコピーして保存する
6. 次の問題へ進む
```

## 2. Docker環境の使い方

### Phase 1 (React / Tailwind) の場合

```bash
# 環境ディレクトリに移動
cd environments/frontend-basic

# 初回起動（イメージのビルド + コンテナ起動）
docker compose up --build

# 2回目以降の起動
docker compose up

# バックグラウンドで起動する場合
docker compose up -d

# ログを確認（バックグラウンド起動時）
docker compose logs -f

# 停止
docker compose stop

# 停止 + コンテナ削除
docker compose down

# 完全リセット（node_modulesボリュームも削除）
docker compose down -v
```

**起動後のアクセス:** http://localhost:3000

### 構成図

```
ホストマシン                         Dockerコンテナ
┌──────────────────────┐           ┌──────────────────┐
│ environments/        │           │                  │
│   frontend-basic/    │           │  /app/           │
│     workspace/  ────────(マウント)────→  (Viteプロジェクト) │
│       src/           │           │    src/          │
│       package.json   │           │    package.json  │
│       ...            │           │    node_modules/ │
│                      │           │                  │
│ localhost:3000 ←──────────(ポート)─────── :5173 (Vite) │
└──────────────────────┘           └──────────────────┘
```

- `workspace/` をコンテナ内の `/app` にマウント
- ホストで `workspace/src/` を編集 → コンテナ側が自動検知 → HMRでブラウザに即反映
- `node_modules` はDockerボリューム管理（ホストには作られない）
- ポートは `3000 (ホスト) → 5173 (コンテナ内Vite)` にマッピング

### コンテナ内でコマンドを実行したい場合

```bash
# コンテナのシェルに入る
docker compose exec app sh

# 直接コマンドを実行
docker compose exec app npm install <パッケージ名>
```

### よくあるトラブル

| 症状 | 対処法 |
|---|---|
| `localhost:3000` にアクセスできない | `docker compose logs` でエラー確認 |
| パッケージが見つからない | `docker compose down -v && docker compose up --build` で再ビルド |
| HMRが効かない | ブラウザをリロード / `docker compose restart` |

---

## 3. 実装ファイルの管理

### ディレクトリ構成

```
web-dev-workout/
├── phase1-frontend-basic/          # 問題文 (md)
│   ├── 01-button-component.md
│   ├── 02-input-form-component.md
│   └── ...
│
├── environments/
│   └── frontend-basic/
│       └── workspace/              # 作業場所 (gitignore)
│           └── src/
│               ├── App.tsx         # ← ここを編集
│               └── components/     # ← コンポーネントを作成
│
└── solutions/                      # 解答保存 (git管理)
    └── phase1-frontend-basic/
        ├── 01-button-component/
        │   ├── App.tsx
        │   └── Button.tsx
        ├── 02-input-form-component/
        │   ├── App.tsx
        │   └── Input.tsx
        └── ...
```

### ワークフロー

#### 1. 問題を解く

`workspace/src/` 内で自由にファイルを作成・編集する。

```
workspace/src/
├── App.tsx              # エントリポイント（問題ごとに書き換え）
├── components/          # コンポーネントを切り出す場合
│   └── Button.tsx
└── index.css            # Tailwind読み込み（変更不要）
```

#### 2. 解答を保存する

問題が完成したら、実装したファイルを `solutions/` にコピーする。

```bash
# 例: 01-button-component の解答を保存
mkdir -p solutions/phase1-frontend-basic/01-button-component
cp environments/frontend-basic/workspace/src/App.tsx \
   solutions/phase1-frontend-basic/01-button-component/
cp environments/frontend-basic/workspace/src/components/Button.tsx \
   solutions/phase1-frontend-basic/01-button-component/
```

#### 3. 次の問題に進む

`workspace/src/` を初期状態に戻して次の問題を始める。

```bash
# App.jsxを初期状態に戻す（componentsディレクトリも削除）
rm -rf environments/frontend-basic/workspace/src/components
```

`App.tsx` は次の問題のコードで上書きすればOK。

#### 4. 解答をコミットする

```bash
git add solutions/phase1-frontend-basic/01-button-component/
git commit -m "01-button-component の解答を追加"
```

---

## 4. IDEの補完を有効にする

workspace内のファイルでReactやTypeScriptの型補完を効かせるには、ホスト側にも`node_modules`が必要です。

```bash
# workspaceディレクトリに移動してインストール
cd environments/frontend-basic/workspace
npm install
```

これにより、VS Code等のIDEが以下を提供できるようになります:

- React (`useState`, `useEffect`等) の型補完・ホバードキュメント
- Tailwind CSSのクラス名補完（拡張機能が必要）
- TypeScriptの型チェック・エラー表示
- `import`文のパス補完

> **補足:** Docker環境の`node_modules`はnamed volumeで管理されているため、ホスト側の`node_modules`はIDEの補完専用です。コンテナ内の実行には影響しません。

---

## 5. Phase 3-4 (Go) の場合

### 起動方法

```bash
# 環境ディレクトリに移動
cd environments/backend

# 初回起動（イメージのビルド + コンテナ起動）
docker compose up --build

# 2回目以降の起動
docker compose up

# バックグラウンドで起動する場合
docker compose up -d

# ログを確認（バックグラウンド起動時）
docker compose logs -f

# 停止
docker compose stop

# 停止 + コンテナ削除
docker compose down
```

**カテゴリH以降（DB必要時）:**

```bash
# DBも一緒に起動する場合（profileを指定）
docker compose --profile db up --build
```

### 構成図

```
ホストマシン                         Dockerコンテナ
┌──────────────────────┐           ┌──────────────────┐
│ environments/        │           │                  │
│   backend/           │           │  /app/           │
│     workspace/  ────────(マウント)────→  (Goプロジェクト)  │
│       main.go        │           │    main.go       │
│       go.mod         │           │    go.mod        │
│       .air.toml      │           │    tmp/ (ビルド)  │
│                      │           │                  │
│ ターミナル出力 ←──────────(ログ)──────── Air (hot reload) │
└──────────────────────┘           └──────────────────┘
```

- `workspace/` をコンテナ内の `/app` にマウント
- ホストで `workspace/main.go` を編集 → Air が自動検知 → 再ビルド＆実行
- `go run .` の代わりに Air が自動でホットリロードしてくれる

### ワークフロー（Go版）

#### 1. 問題を解く

`workspace/main.go` を編集する。カテゴリAの問題は全て単一ファイルで完結。

```
workspace/
├── main.go     # ← ここを編集（問題ごとに書き換え）
├── go.mod      # モジュール定義（変更不要）
└── .air.toml   # Air設定（変更不要）
```

#### 2. 実行して確認

Air が自動で再ビルド＆実行するので、ファイルを保存するだけでOK。
手動で実行したい場合:

```bash
# コンテナ内でGoコマンドを直接実行
docker compose exec app go run .

# コンテナのシェルに入る
docker compose exec app sh
```

#### 3. 解答を保存する

```bash
# 例: 01-struct-and-embedding の基本レベル解答を保存
mkdir -p solutions/phase3-backend-basic/01-struct-and-embedding/basic
cp environments/backend/workspace/main.go \
   solutions/phase3-backend-basic/01-struct-and-embedding/basic/main.go
```

#### 4. 次の問題に進む

`workspace/main.go` を次の問題のコードで上書きすればOK。

### コンテナ内でコマンドを実行したい場合

```bash
# コンテナのシェルに入る
docker compose exec app sh

# 直接コマンドを実行
docker compose exec app go run .
docker compose exec app go test ./...
docker compose exec app go fmt ./...
```

### よくあるトラブル

| 症状 | 対処法 |
|---|---|
| `air: command not found` | `docker compose down && docker compose up --build` で再ビルド |
| ビルドエラーが出る | `docker compose logs` でエラー内容を確認 |
| ホットリロードが効かない | `docker compose restart` |
| パッケージが見つからない | コンテナ内で `go mod tidy` を実行 |

---

## 6. 環境一覧

| Phase | 環境ディレクトリ | 技術スタック | ポート |
|---|---|---|---|
| Phase 1 | `environments/frontend-basic/` | Vite + React + Tailwind CSS + TypeScript | 3000 |
| Phase 2 | `environments/frontend-advanced/` | Next.js (予定) | - |
| Phase 3-4 | `environments/backend/` | Go 1.21 + Air (hot reload) + PostgreSQL | 8080 |
