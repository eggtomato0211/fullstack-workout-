# Web開発 実践問題集

## 📚 概要
React/Next.js (フロントエンド) と Go (バックエンド) を用いたWebアプリケーション開発の実践的な問題集です。
各テーマは「Why解説 → 実装例(2段階) → 演習問題」の構成です。

## 🗂️ ディレクトリ構造

```
web-dev-workbook/
├── README.md (本ファイル)
├── phase1-frontend-basic/     # フロント基礎 (React/Tailwind)
├── phase2-frontend-advanced/  # フロント応用 (Next.js)
├── phase3-backend-basic/      # バック基礎 (Go)
├── phase4-backend-advanced/   # バック応用 (設計/GORM)
└── phase5-devops/             # その他発展 (インフラ/CI/CD)
```

## 📖 学習の進め方

1. 各フェーズ内のテーマは番号順に進めることを推奨
2. 「なぜ必要か」を読み、コード例を理解した後、演習問題に取り組む
3. 実装したコードはGitHubにpushして進捗管理

---

## Phase 1: フロント基礎 (React/Tailwind)

### カテゴリA: 再利用可能なUIコンポーネント

| No | テーマ | 習得内容 |
|---|---|---|
| 01 | Buttonコンポーネント | variant/size/state管理の実装 |
| 02 | Input/Formコンポーネント | text/email/password、バリデーション、エラー表示 |
| 03 | Cardコンポーネント | header/body/footer構造、variant管理 |
| 04 | Modalコンポーネント | オーバーレイ、ESCキーで閉じる、背景スクロール防止 |
| 05 | Dropdownコンポーネント | メニュー表示、外側クリックで閉じる |
| 06 | Alertコンポーネント | success/warning/error/infoの表示 |
| 07 | Loading系コンポーネント | Spinner、Skeleton、Progressの実装 |

### カテゴリB: Hooks基礎 - useState

| No | テーマ | 習得内容 |
|---|---|---|
| 08 | useStateの基本 | プリミティブ値・オブジェクト・配列の管理、スプレッド構文 |
| 09 | prevStateパターンと応用 | 前の状態に基づく更新、Lazy初期化、state分割の判断 |

### カテゴリC: Hooks基礎 - useEffect

| No | テーマ | 習得内容 |
|---|---|---|
| 10 | useEffectの基本とAPI取得 | 依存配列、fetch実行、ローディング・エラー状態の管理 |
| 11 | useEffectのクリーンアップ | タイマー解除、イベントリスナー解除、AbortController |

### カテゴリD: Hooks基礎 - その他の基本Hook

| No | テーマ | 習得内容 |
|---|---|---|
| 12 | useRefの基本と応用 | DOM参照、再レンダリングをトリガーしない値の保存 |
| 13 | useReducerの基本と応用 | reducerパターン、複雑なstate管理、actionのtype定義 |

### カテゴリE: Hooks応用 - パフォーマンス最適化

| No | テーマ | 習得内容 |
|---|---|---|
| 14 | React.memo/useMemoによるメモ化 | コンポーネントと計算結果のメモ化、依存配列 |
| 15 | useCallbackと最適化の判断基準 | 関数のメモ化、いつメモ化すべきか |

### カテゴリF: Hooks応用 - Context API

| No | テーマ | 習得内容 |
|---|---|---|
| 16 | useContextの基本と実践 | Context作成、テーマ切り替え、認証状態の管理 |
| 17 | Context + useReducerパターン | Contextの分割、グローバルstate管理 |

### カテゴリG: Custom Hooks

| No | テーマ | 習得内容 |
|---|---|---|
| 18 | Custom Hooksの基本 | useToggle、useLocalStorageの実装、命名規則 |
| 19 | データ取得・入力系Custom Hooks | useDebounce、useFetch、useFormの実装 |

### カテゴリH: フォーム・イベント処理

| No | テーマ | 習得内容 |
|---|---|---|
| 20 | フォーム制御とバリデーション | 制御/非制御コンポーネント、入力チェック、エラー表示 |
| 21 | イベントハンドリング | onClick/onChange/onSubmit、バブリングと伝播、合成イベント |

### カテゴリI: コンポーネント設計

| No | テーマ | 習得内容 |
|---|---|---|
| 22 | コンポーネント設計パターン | children活用、Render Props、Compound Components、Container/Presenter |

### カテゴリJ: リスト・条件レンダリング

| No | テーマ | 習得内容 |
|---|---|---|
| 23 | リスト・条件レンダリング | map/filter、key属性、&&演算子、三項演算子、早期return |

**Phase 1 合計: 23テーマ**

---

## Phase 2: フロント応用 (Next.js)

### カテゴリA: レンダリング戦略

| No | テーマ | 習得内容 |
|---|---|---|
| 01 | Server/Client Componentsの使い分け | "use client"の適切な配置、パフォーマンスへの影響 |
| 02 | Streaming SSRとLoading/Error UI | Suspense、loading.tsx、error.tsxの活用 |

### カテゴリB: データフェッチ

| No | テーマ | 習得内容 |
|---|---|---|
| 03 | Server Actionsによるフォーム処理 | formDataの取得、バリデーション、revalidate |
| 04 | fetchのキャッシュとParallel Data Fetching | cache制御、Promise.allによる並列化 |

### カテゴリC: ルーティングとナビゲーション

| No | テーマ | 習得内容 |
|---|---|---|
| 05 | Dynamic RoutesとRoute Groups | [id]/[...slug]、レイアウト分離 |
| 06 | Middlewareによる認証ガード | リダイレクト、条件分岐の実装 |

### カテゴリD: パフォーマンス最適化

| No | テーマ | 習得内容 |
|---|---|---|
| 07 | next/image・next/fontによる最適化 | 画像最適化、variable font、FOUT/FOIT防止 |
| 08 | Code Splittingとdynamic import | React.lazyとnext/dynamicの使い分け |

### カテゴリE: SEO・メタデータ

| No | テーマ | 習得内容 |
|---|---|---|
| 09 | メタデータとOGP画像の動的生成 | generateMetadata、ImageResponse |

### カテゴリF: 設計・セキュリティ

| No | テーマ | 習得内容 |
|---|---|---|
| 10 | NextAuth.jsによる認証実装 | Google/GitHub OAuth、JWTセッション |
| 11 | XSS/CSRF対策 | サニタイズ、CSRFトークン、SameSite Cookie |

**Phase 2 合計: 11テーマ**

---

## Phase 3: バック基礎 (Go)

### カテゴリA: Goらしさの基礎

| No | テーマ | 習得内容 |
|---|---|---|
| 01 | 構造体と埋め込み | struct定義、初期化、埋め込みによる擬似継承 |
| 02 | ポインタの基本 | 参照と逆参照、値渡しとポインタ渡しの使い分け |
| 03 | メソッドとレシーバー | 値レシーバーとポインタレシーバーの使い分け |
| 04 | Interfaceの基本 | interface定義、暗黙的な実装 |
| 05 | Interfaceの応用 | 型アサーション、型スイッチ、空インターフェース |
| 06 | エラーハンドリングの基本 | error型、errorsパッケージ |
| 07 | カスタムエラーの実装 | errors.New、fmt.Errorf、カスタムエラー型 |
| 08 | errors.Is/Asの活用 | エラーラッピング、エラーチェーン |
| 09 | defer文の活用 | リソース解放、実行順序、エラーハンドリングとの組み合わせ |
| 10 | panic/recoverの理解 | 例外的状況の処理、ミドルウェアでの活用 |
| 11 | encoding/jsonと構造体タグ | JSON Marshal/Unmarshal、構造体タグの活用 |
| 12 | io.Reader/io.Writerの理解 | Goのインターフェースパターンの代表例、ストリーム処理 |

### カテゴリB: 並行処理

| No | テーマ | 習得内容 |
|---|---|---|
| 13 | Goroutineとsync.WaitGroup | 並行実行の仕組み、完了待機 |
| 14 | Channelの基本 | チャネル送受信、Buffered Channel、閉じる |
| 15 | select文の活用 | 複数チャネルの操作、タイムアウト |
| 16 | sync.Mutex/RWMutexによる排他制御 | 競合状態の回避、読み書き分離 |
| 17 | Contextの基本と応用 | キャンセル、タイムアウト、値の伝搬 |

### カテゴリC: パッケージとモジュール

| No | テーマ | 習得内容 |
|---|---|---|
| 18 | Go Modulesとパッケージ管理 | go.mod/go.sum、import、可視性、循環依存の回避 |

### カテゴリD: HTTP APIの基礎

| No | テーマ | 習得内容 |
|---|---|---|
| 19 | net/httpとHandlerFuncの基本 | サーバー起動、GETリクエスト処理 |
| 20 | リクエストの解析とボディ読み取り | パラメータ、ヘッダー、JSON decoding、バリデーション |
| 21 | レスポンスとエラーレスポンスの設計 | JSON encoding、ステータスコード、統一エラー構造 |

### カテゴリE: ルーティング

| No | テーマ | 習得内容 |
|---|---|---|
| 22 | ServeMuxとgorilla/mux | 基本ルーティング、パスパラメータ、メソッドルーティング |
| 23 | RESTful API設計 | リソース指向、HTTPメソッドの使い分け、バージョニング |

### カテゴリF: ミドルウェア

| No | テーマ | 習得内容 |
|---|---|---|
| 24 | ミドルウェアの基本と認証 | http.Handlerのラップ、ロギング、JWT検証 |
| 25 | CORS/エラーハンドリングミドルウェア | クロスオリジン許可、panic recovery、レート制限 |

### カテゴリG: バリデーション

| No | テーマ | 習得内容 |
|---|---|---|
| 26 | リクエストバリデーション | 構造体タグ、go-playground/validator、カスタム実装 |

### カテゴリH: データベース基礎

| No | テーマ | 習得内容 |
|---|---|---|
| 27 | database/sqlの基本 | DB接続、Prepared Statement、接続プール設定 |
| 28 | CRUD操作とトランザクション | Query/Exec、Begin/Commit/Rollback、NULL値の扱い |

### カテゴリI: SQLクエリ実践

| No | テーマ | 習得内容 |
|---|---|---|
| 29 | JOINとページネーション | テーブル結合、LIMIT/OFFSET、cursor-based pagination |
| 30 | 検索・ソート機能の実装 | LIKE、全文検索、ORDER BY、動的ソート |

### カテゴリJ: テスト基礎

| No | テーマ | 習得内容 |
|---|---|---|
| 31 | testingパッケージとTable Driven Test | テスト関数、サブテスト、テストヘルパー |
| 32 | モックとHTTP/DBテスト | interfaceモック、httptest、テストDB |

### カテゴリK: コード品質

| No | テーマ | 習得内容 |
|---|---|---|
| 33 | 静的解析ツールとGoイディオム | gofmt、go vet、golangci-lint、Effective Go |

### カテゴリL: 環境設定・設定管理

| No | テーマ | 習得内容 |
|---|---|---|
| 34 | 環境変数と設定管理 | os.Getenv、.env、設定構造体、環境別設定 |

**Phase 3 合計: 34テーマ**

---

## Phase 4: バック応用 (設計/GORM)

### カテゴリA: クリーンアーキテクチャ

| No | テーマ | 習得内容 |
|---|---|---|
| 01 | レイヤー分離の基本 | Domain/Usecase/Interface/Infrastructureの役割 |
| 02 | 各レイヤーの実装 | エンティティ、ビジネスルール、HTTPハンドラー、Repository |

### カテゴリB: 依存性の注入(DI)

| No | テーマ | 習得内容 |
|---|---|---|
| 03 | DIの基本とモック作成 | コンストラクタインジェクション、テスト用モック、wire |

### カテゴリC: GORM基礎

| No | テーマ | 習得内容 |
|---|---|---|
| 04 | GORMの基本操作 | Create/Find/Update/Delete、Association定義 |
| 05 | マイグレーション管理 | AutoMigrate、migration file |

### カテゴリD: GORM最適化

| No | テーマ | 習得内容 |
|---|---|---|
| 06 | N+1問題とLoading戦略 | Preload/Joinsの使い分け、パフォーマンス |
| 07 | クエリ最適化とトランザクション | Select/Pluck、ネストトランザクション、SavePoint |

### カテゴリE: アンチパターンの回避

| No | テーマ | 習得内容 |
|---|---|---|
| 08 | よくあるアンチパターン | ドメインロジック漏洩、グローバル変数乱用、不適切なエラーハンドリング |

**Phase 4 合計: 8テーマ**

---

## Phase 5: その他発展 (インフラ/CI/CD)

### カテゴリA: Docker

| No | テーマ | 習得内容 |
|---|---|---|
| 01 | Dockerfileとマルチステージビルド | FROM/COPY/RUN/CMD、本番イメージの軽量化 |
| 02 | docker-composeによる開発環境 | DB/アプリの連携、ボリューム管理 |

### カテゴリB: CI/CD

| No | テーマ | 習得内容 |
|---|---|---|
| 03 | GitHub Actionsの基本 | workflow作成、トリガー設定 |
| 04 | 自動テスト・Lint・デプロイ | CI統合、ESLint/golangci-lint、デプロイフロー |

### カテゴリC: デプロイ・運用

| No | テーマ | 習得内容 |
|---|---|---|
| 05 | クラウドデプロイ | AWS App Runner、Cloud Runへのデプロイ |
| 06 | 構造化ログとヘルスチェック | slog、Liveness/Readiness probe |

### カテゴリD: 通信・連携

| No | テーマ | 習得内容 |
|---|---|---|
| 07 | gRPCの基本 | Protocol Buffers、Unary/Streaming通信パターン |
| 08 | OpenAPI・WebSocket | スキーマ駆動開発、リアルタイム通信 |

**Phase 5 合計: 8テーマ**

---

## 📊 全体サマリー

| フェーズ | テーマ数 |
|---|---|
| Phase 1: フロント基礎 | 23 |
| Phase 2: フロント応用 | 11 |
| Phase 3: バック基礎 | 34 |
| Phase 4: バック応用 | 8 |
| Phase 5: その他発展 | 8 |
| **合計** | **84テーマ** |

---

## 🎯 次のステップ

1. ✅ 全体構成の確認 (このファイル)
2. 📝 取り組みたいフェーズを選択
3. 🚀 各テーマのMarkdownファイルを作成開始

**推奨学習順序:**
- 初心者: Phase 1 → Phase 3 → Phase 2 → Phase 4 → Phase 5
- フロント重点: Phase 1 → Phase 2 → Phase 3 → Phase 4 → Phase 5
- バック重点: Phase 3 → Phase 1 → Phase 4 → Phase 2 → Phase 5

どのフェーズから始めますか?
