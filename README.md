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
| 08 | useStateの基本 - プリミティブ値 | カウンター、トグル、文字列の管理 |
| 09 | useStateの基本 - オブジェクト | オブジェクトstateの更新、スプレッド構文 |
| 10 | useStateの基本 - 配列 | 配列の追加・削除・更新 |
| 11 | prevStateパターン | 前の状態に基づく更新、クロージャ問題の回避 |
| 12 | Lazy初期化 | 関数を使った初期値の遅延評価 |
| 13 | 複数のstateの管理 | state分割 vs 統合の判断基準 |

### カテゴリC: Hooks基礎 - useEffect

| No | テーマ | 習得内容 |
|---|---|---|
| 14 | useEffectの基本 - マウント時の処理 | 初期データ取得、空の依存配列 |
| 15 | useEffectの基本 - 依存配列 | 依存配列の正しい指定、lintルールの理解 |
| 16 | useEffectによるAPI取得 | fetch実行、ローディング・エラー状態の管理 |
| 17 | useEffectのクリーンアップ - タイマー | setInterval/setTimeoutの解除 |
| 18 | useEffectのクリーンアップ - イベント | addEventListener/removeEventListener |
| 19 | useEffectのクリーンアップ - AbortController | fetch中断、メモリリーク防止 |
| 20 | useEffectの実行タイミング | レンダリングとの関係、useLayoutEffectとの違い |

### カテゴリD: Hooks基礎 - その他の基本Hook

| No | テーマ | 習得内容 |
|---|---|---|
| 21 | useRefの基本 - DOM参照 | input要素へのfocus、要素サイズの取得 |
| 22 | useRefの応用 - 値の保持 | 再レンダリングをトリガーしない値の保存 |
| 23 | useReducerの基本 | reducerパターン、複雑なstate管理 |
| 24 | useReducerの応用 | actionのtype定義、middleware的な処理 |

### カテゴリE: Hooks応用 - パフォーマンス最適化

| No | テーマ | 習得内容 |
|---|---|---|
| 25 | React.memoによるコンポーネントのメモ化 | 再レンダリングの最適化、propsの比較 |
| 26 | useMemoの基本 | 計算結果のメモ化、依存配列 |
| 27 | useMemoの応用 | フィルタリング・ソートの最適化 |
| 28 | useCallbackの基本 | 関数のメモ化、子コンポーネントへの影響 |
| 29 | useCallbackの応用 | イベントハンドラーの最適化 |
| 30 | 最適化の判断基準 | いつメモ化すべきか、過度な最適化の回避 |

### カテゴリF: Hooks応用 - Context API

| No | テーマ | 習得内容 |
|---|---|---|
| 31 | useContextの基本 | Context作成、Provider、Consumer |
| 32 | テーマ切り替え機能の実装 | Light/Darkモード、useContextの実践 |
| 33 | 認証状態の管理 | ユーザー情報の共有、ログイン/ログアウト |
| 34 | Contextの分割 | パフォーマンス改善、関心の分離 |
| 35 | Context + useReducerパターン | グローバルstate管理 |

### カテゴリG: Custom Hooks

| No | テーマ | 習得内容 |
|---|---|---|
| 36 | Custom Hooksの基本 | useToggleの実装、命名規則 |
| 37 | useLocalStorageの実装 | localStorage同期、型安全性 |
| 38 | useDebounceの実装 | 入力遅延、検索最適化 |
| 39 | useFetchの実装 | データ取得の共通化、再利用性 |
| 40 | useFormの実装 | フォーム状態管理、バリデーション |
| 41 | useMediaQueryの実装 | レスポンシブ判定、window.matchMedia |
| 42 | useIntersectionObserverの実装 | 無限スクロール、遅延読み込み |

### カテゴリH: フォーム処理

| No | テーマ | 習得内容 |
|---|---|---|
| 43 | 制御されたコンポーネント | value/onChangeによる制御 |
| 44 | 非制御コンポーネント | useRefによるフォーム値の取得 |
| 45 | フォームバリデーション (基本) | 入力チェック、エラー表示 |
| 46 | フォームバリデーション (応用) | リアルタイム検証、複数フィールドの関連検証 |
| 47 | フォーム送信処理 | 送信時の処理、preventDefault |

### カテゴリI: イベント処理

| No | テーマ | 習得内容 |
|---|---|---|
| 48 | onClick/onChangeイベント | 基本的なイベントハンドリング |
| 49 | onSubmit/onFocusイベント | フォーム関連イベント |
| 50 | onKeyDown/onKeyUpイベント | キーボード操作、Enter送信 |
| 51 | イベントバブリングと伝播 | stopPropagation、イベント委譲 |
| 52 | 合成イベント(SyntheticEvent) | Reactのイベントシステムの理解 |

### カテゴリJ: コンポーネント設計

| No | テーマ | 習得内容 |
|---|---|---|
| 53 | Propsの型定義 | TypeScript、PropTypesの活用 |
| 54 | childrenの活用 | コンポーネント合成、スロットパターン |
| 55 | Render Propsパターン | 柔軟なコンポーネント設計 |
| 56 | Compound Componentsパターン | 関連コンポーネントのグループ化 |
| 57 | Container/Presenterパターン | ロジックとUIの分離 |
| 58 | Atomic Designの基礎 | Atoms/Molecules/Organismsの概念 |

### カテゴリK: リスト・条件レンダリング

| No | テーマ | 習得内容 |
|---|---|---|
| 59 | mapによるリストレンダリング | key属性の重要性 |
| 60 | filterとmapの組み合わせ | 条件付きリスト表示 |
| 61 | 条件レンダリング (&&演算子) | 真偽値による表示切り替え |
| 62 | 条件レンダリング (三項演算子) | if-else的な表示切り替え |
| 63 | 早期return | 条件による早期終了 |

**Phase 1 合計: 63テーマ**

---

## Phase 2: フロント応用 (Next.js)

### カテゴリA: レンダリング戦略

| No | テーマ | 習得内容 |
|---|---|---|
| 01 | Server ComponentsとClient Componentsの使い分け | "use client"の適切な配置、パフォーマンスへの影響 |
| 02 | Streaming SSRの実装 | Suspenseによる段階的レンダリング |
| 03 | Loading UIとエラーハンドリング | loading.tsx, error.tsxの活用 |

### カテゴリB: データフェッチ

| No | テーマ | 習得内容 |
|---|---|---|
| 04 | Server Actionsによるフォーム処理 | formDataの取得、バリデーション、revalidate |
| 05 | fetchのキャッシュ制御 | force-cache, no-store, revalidateの使い分け |
| 06 | Parallel Data Fetchingとシーケンシャル取得 | Promise.allによる並列化、waterfallの回避 |

### カテゴリC: ルーティングとナビゲーション

| No | テーマ | 習得内容 |
|---|---|---|
| 07 | Dynamic Routesの実装 | [id], [...slug]の活用 |
| 08 | Route Groupsとレイアウト | (admin), (auth)によるレイアウト分離 |
| 09 | Middlewareによる認証ガード | リダイレクト、条件分岐の実装 |

### カテゴリD: パフォーマンス最適化

| No | テーマ | 習得内容 |
|---|---|---|
| 10 | next/imageによる画像最適化 | priority, sizes, fill属性の活用 |
| 11 | next/fontによるフォント最適化 | variable fontの使用、FOUT/FOITの防止 |
| 12 | Code Splittingとdynamic import | React.lazyとnext/dynamicの使い分け |

### カテゴリE: SEO・メタデータ

| No | テーマ | 習得内容 |
|---|---|---|
| 13 | 動的メタデータの生成 | generateMetadata関数の実装 |
| 14 | OGP画像の動的生成 | ImageResponseを用いたOG画像作成 |

### カテゴリF: 設計・セキュリティ

| No | テーマ | 習得内容 |
|---|---|---|
| 15 | Feature-based設計 | 機能単位のディレクトリ構成 |
| 16 | NextAuth.jsによる認証実装 | Google/GitHub OAuth、JWTセッション |
| 17 | XSS対策 | dangerouslySetInnerHTMLの安全な使用、サニタイズ |
| 18 | CSRF対策 | CSRFトークン、SameSite Cookie属性 |

**Phase 2 合計: 18テーマ**

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
| 13 | Goroutineの基本 | go キーワード、並行実行の仕組み |
| 14 | sync.WaitGroupの使用 | Goroutineの完了待機 |
| 15 | Channelの基本 | チャネル作成、送受信、閉じる |
| 16 | Buffered Channel | バッファ付きチャネル、容量の指定 |
| 17 | select文の活用 | 複数チャネルの操作、タイムアウト |
| 18 | sync.Mutexによる排他制御 | 競合状態の回避、Lock/Unlock |
| 19 | sync.RWMutexの活用 | 読み込みと書き込みの分離 |
| 20 | Contextの基本 | context.Context、キャンセル、タイムアウト |
| 21 | Contextの応用 | 値の伝搬、リクエストスコープの管理 |

### カテゴリC: パッケージとモジュール

| No | テーマ | 習得内容 |
|---|---|---|
| 22 | パッケージの基本 | package宣言、import、可視性 |
| 23 | Go Modulesの基礎 | go.mod/go.sum、go get/mod tidy |
| 24 | 依存関係の管理 | バージョン指定、vendor、replace |
| 25 | 内部パッケージの設計 | ディレクトリ構成、循環依存の回避 |

### カテゴリD: HTTP APIの基礎

| No | テーマ | 習得内容 |
|---|---|---|
| 26 | net/httpの基本 | http.Server、http.ListenAndServe |
| 27 | http.HandlerFuncの実装 | GETリクエストの処理 |
| 28 | リクエストの解析 | クエリパラメータ、パスパラメータ、ヘッダー |
| 29 | リクエストボディの読み取り | JSON decoding、バリデーション |
| 30 | レスポンスの作成 | JSON encoding、ステータスコード |
| 31 | エラーレスポンスの設計 | 統一されたエラー構造 |

### カテゴリE: ルーティング

| No | テーマ | 習得内容 |
|---|---|---|
| 32 | http.ServeMuxの使用 | 基本的なルーティング |
| 33 | gorilla/muxの導入 | パスパラメータ、メソッドルーティング |
| 34 | RESTful APIのルーティング設計 | リソース指向、HTTPメソッドの使い分け |
| 35 | APIバージョニング | /v1, /v2 のルーティング |

### カテゴリF: ミドルウェア

| No | テーマ | 習得内容 |
|---|---|---|
| 36 | ミドルウェアの基本 | http.Handlerのラップ |
| 37 | ロギングミドルウェア | リクエスト/レスポンスのログ出力 |
| 38 | 認証ミドルウェア | JWTトークン検証、Authorizationヘッダー |
| 39 | CORSミドルウェア | クロスオリジンリクエストの許可 |
| 40 | エラーハンドリングミドルウェア | panicのrecovery、統一エラーレスポンス |
| 41 | レート制限ミドルウェア | リクエスト数制限、IPベース制限 |

### カテゴリG: バリデーション

| No | テーマ | 習得内容 |
|---|---|---|
| 42 | 構造体タグによるバリデーション | validate:"required,email"等 |
| 43 | go-playground/validatorの使用 | カスタムバリデーション |
| 44 | リクエストバリデーション関数の実装 | 再利用可能なバリデーター |

### カテゴリH: データベース基礎

| No | テーマ | 習得内容 |
|---|---|---|
| 45 | database/sqlの基本 | DB接続、sql.Open |
| 46 | Prepared Statementの使用 | SQLインジェクション対策 |
| 47 | Query/QueryRowの使用 | SELECT文の実行、結果の取得 |
| 48 | Exec/ExecContextの使用 | INSERT/UPDATE/DELETE文の実行 |
| 49 | トランザクションの基本 | Begin/Commit/Rollback |
| 50 | NULL値の扱い | sql.NullString/NullInt64等 |
| 51 | 接続プールの設定 | MaxOpenConns/MaxIdleConns |

### カテゴリI: SQLクエリ実践

| No | テーマ | 習得内容 |
|---|---|---|
| 52 | 基本的なCRUD操作 - CREATE | INSERT文、LastInsertId |
| 53 | 基本的なCRUD操作 - READ | SELECT文、WHERE条件 |
| 54 | 基本的なCRUD操作 - UPDATE | UPDATE文、RowsAffected |
| 55 | 基本的なCRUD操作 - DELETE | DELETE文、論理削除 vs 物理削除 |
| 56 | JOINクエリの実装 | INNER/LEFT JOIN、複数テーブルの結合 |
| 57 | ページネーションの実装 | LIMIT/OFFSET、cursor-based pagination |
| 58 | 検索機能の実装 | LIKE、全文検索 |
| 59 | ソート機能の実装 | ORDER BY、動的ソート |

### カテゴリJ: テスト基礎

| No | テーマ | 習得内容 |
|---|---|---|
| 60 | testing パッケージの基本 | テスト関数、t.Error/Fatal |
| 61 | Table Driven Testの実装 | 複数テストケースの効率的な実行 |
| 62 | サブテストの活用 | t.Run、テストの構造化 |
| 63 | テストヘルパー関数 | setup/teardown、共通処理 |
| 64 | モックの基本 | interfaceを用いたモック |
| 65 | HTTPハンドラーのテスト | httptest.ResponseRecorder |
| 66 | DBのテスト | テストDB、トランザクションロールバック |

### カテゴリK: コード品質

| No | テーマ | 習得内容 |
|---|---|---|
| 67 | gofmtによるフォーマット | コードスタイルの統一 |
| 68 | golintの使用 | コーディング規約のチェック |
| 69 | go vetの活用 | 潜在的なバグの検出 |
| 70 | golangci-lintの導入 | 包括的な静的解析 |
| 71 | Effective Goの理解 | Goのイディオム、ベストプラクティス |

### カテゴリL: 環境設定・設定管理

| No | テーマ | 習得内容 |
|---|---|---|
| 72 | 環境変数の読み取り | os.Getenv、デフォルト値 |
| 73 | .envファイルの使用 | godotenvによる環境変数管理 |
| 74 | 設定構造体の設計 | アプリケーション設定の一元管理 |
| 75 | 環境別設定の切り替え | dev/staging/production |

**Phase 3 合計: 75テーマ**

---

## Phase 4: バック応用 (設計/GORM)

### カテゴリA: クリーンアーキテクチャ

| No | テーマ | 習得内容 |
|---|---|---|
| 01 | レイヤー分離の基本 | Domain/Usecase/Interface/Infrastructureの役割 |
| 02 | Domainレイヤーの実装 | エンティティ、ビジネスルールの定義 |
| 03 | Usecaseレイヤーの実装 | アプリケーションロジック、orchestration |
| 04 | Interfaceレイヤーの実装 | HTTPハンドラー、リクエスト/レスポンス変換 |
| 05 | Infrastructureレイヤーの実装 | Repository、外部API連携 |

### カテゴリB: 依存性の注入(DI)

| No | テーマ | 習得内容 |
|---|---|---|
| 06 | DIの基本概念 | コンストラクタインジェクション |
| 07 | インターフェースを用いたモック作成 | テスト用のモック実装 |
| 08 | DIコンテナの実装 | wire、google/wireの活用 |

### カテゴリC: GORM基礎

| No | テーマ | 習得内容 |
|---|---|---|
| 09 | GORMの基本操作 | Create/Find/Update/Delete |
| 10 | Association定義 | has one/has many/belongs to/many to many |
| 11 | マイグレーション管理 | AutoMigrate、migration file |

### カテゴリD: GORM最適化

| No | テーマ | 習得内容 |
|---|---|---|
| 12 | N+1問題の理解と解決 | Preload/Joinsの使い分け |
| 13 | Eager LoadingとLazy Loading | パフォーマンスへの影響 |
| 14 | クエリの最適化 | Select、Pluckによる不要なデータ取得の削減 |
| 15 | トランザクション制御 | ネストしたトランザクション、SavePoint |

### カテゴリE: アンチパターンの回避

| No | テーマ | 習得内容 |
|---|---|---|
| 16 | ドメインロジックの漏洩防止 | Repositoryへのビジネスロジック混入を防ぐ |
| 17 | グローバル変数の乱用回避 | DB接続のグローバル化の問題点 |
| 18 | 不適切なエラーハンドリングの修正 | panic/recover の誤用、エラーの握り潰し |
| 19 | 巨大な構造体の分割 | Single Responsibility Principleの実践 |

**Phase 4 合計: 19テーマ**

---

## Phase 5: その他発展 (インフラ/CI/CD)

### カテゴリA: Docker

| No | テーマ | 習得内容 |
|---|---|---|
| 01 | Dockerfileの基本 | FROM/COPY/RUN/CMD/ENTRYPOINTの理解 |
| 02 | マルチステージビルド | 本番用イメージの軽量化 |
| 03 | docker-composeによる開発環境 | DB/アプリの連携、ボリューム管理 |

### カテゴリB: CI/CD

| No | テーマ | 習得内容 |
|---|---|---|
| 04 | GitHub Actionsの基本 | workflow作成、トリガー設定 |
| 05 | 自動テストの実装 | unit test/integration testの自動実行 |
| 06 | Linterの自動実行 | ESLint/golangci-lintのCI統合 |
| 07 | 自動デプロイフローの構築 | staging/productionへのデプロイ |

### カテゴリC: デプロイ・運用

| No | テーマ | 習得内容 |
|---|---|---|
| 08 | AWS App Runnerへのデプロイ | コンテナベースのデプロイ |
| 09 | Cloud Runへのデプロイ | GCPでのサーバーレス実行 |
| 10 | 構造化ログ(slog)の実装 | JSON形式のログ、可観測性の向上 |
| 11 | ヘルスチェックエンドポイント | Liveness/Readiness probeの実装 |

### カテゴリD: 通信・連携

| No | テーマ | 習得内容 |
|---|---|---|
| 12 | gRPCの基本 | Protocol Buffersの定義、サーバー/クライアント実装 |
| 13 | gRPC StreamingとUnary | 4つの通信パターンの使い分け |
| 14 | OpenAPI (Swagger) によるスキーマ駆動開発 | spec定義、コード生成 |
| 15 | WebSocketによるリアルタイム通信 | チャット機能の実装 |

**Phase 5 合計: 15テーマ**

---

## 📊 全体サマリー

| フェーズ | テーマ数 |
|---|---|
| Phase 1: フロント基礎 | 63 |
| Phase 2: フロント応用 | 18 |
| Phase 3: バック基礎 | 75 |
| Phase 4: バック応用 | 19 |
| Phase 5: その他発展 | 15 |
| **合計** | **190テーマ** |

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
