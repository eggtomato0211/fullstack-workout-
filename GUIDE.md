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

## 4. 環境一覧

| Phase | 環境ディレクトリ | 技術スタック | ポート |
|---|---|---|---|
| Phase 1 | `environments/frontend-basic/` | Vite + React + Tailwind CSS + TypeScript | 3000 |
| Phase 2 | `environments/frontend-advanced/` | Next.js (予定) | - |
| Phase 3-4 | `environments/backend/` | Go (予定) | - |
