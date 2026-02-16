# 03. Cardコンポーネント - header/body/footer構造、variant管理

## 📖 概念

Cardはコンテンツをグループ化して表示するためのコンポーネントです。header（タイトル）、body（本文）、footer（アクション）の3つの領域に分割し、variant（見た目のバリエーション）で表示スタイルを切り替えます。構造的に情報を整理することで、ユーザーが内容を素早く把握できます。

**よくある誤解:**
- ❌ 「Cardは見た目だけの問題」→ header/body/footerの構造化がコンテンツの意味を明確にする
- ❌ 「childrenに全部入れればいい」→ 領域を分けることで再利用性とレイアウトの一貫性が向上する
- ❌ 「variantはclassNameを直接渡せばいい」→ 定義済みvariantで統一的なデザインを維持する

## 💡 コード例

### 基本: シンプルなCard

```jsx
function Card({ children }) {
  return (
    <div className="bg-white rounded-lg shadow p-6">
      {children}
    </div>
  );
}

// 使用例
function App() {
  return (
    <div className="p-4 max-w-sm">
      <Card>
        <h2 className="text-lg font-bold">カードタイトル</h2>
        <p className="mt-2 text-gray-600">
          これはカードの本文です。コンテンツをまとめて表示します。
        </p>
      </Card>
    </div>
  );
}
```

### 応用: header/body/footer分離

```jsx
function Card({ children }) {
  return (
    <div className="bg-white rounded-lg shadow overflow-hidden">
      {children}
    </div>
  );
}

function CardHeader({ children }) {
  return (
    <div className="px-6 py-4 border-b border-gray-200">
      {children}
    </div>
  );
}

function CardBody({ children }) {
  return (
    <div className="px-6 py-4">
      {children}
    </div>
  );
}

function CardFooter({ children }) {
  return (
    <div className="px-6 py-4 border-t border-gray-200 bg-gray-50">
      {children}
    </div>
  );
}

// 使用例
function App() {
  return (
    <div className="p-4 max-w-sm">
      <Card>
        <CardHeader>
          <h2 className="text-lg font-bold">ユーザー情報</h2>
        </CardHeader>
        <CardBody>
          <p className="text-gray-600">山田太郎</p>
          <p className="text-sm text-gray-400">yamada@example.com</p>
        </CardBody>
        <CardFooter>
          <button className="text-blue-600 hover:text-blue-800 text-sm font-medium">
            編集する
          </button>
        </CardFooter>
      </Card>
    </div>
  );
}
```

### 実践: variant（default/outlined/elevated）

```jsx
function Card({ variant = 'default', children }) {
  const variantClasses = {
    default: 'bg-white shadow',
    outlined: 'bg-white border-2 border-gray-200',
    elevated: 'bg-white shadow-lg shadow-gray-200',
  };

  return (
    <div className={`rounded-lg overflow-hidden ${variantClasses[variant]}`}>
      {children}
    </div>
  );
}

function CardHeader({ children }) {
  return (
    <div className="px-6 py-4 border-b border-gray-200">
      {children}
    </div>
  );
}

function CardBody({ children }) {
  return (
    <div className="px-6 py-4">
      {children}
    </div>
  );
}

function CardFooter({ children }) {
  return (
    <div className="px-6 py-4 border-t border-gray-200 bg-gray-50">
      {children}
    </div>
  );
}

// 使用例
function App() {
  return (
    <div className="p-4 max-w-md space-y-6">
      <Card variant="default">
        <CardHeader>
          <h2 className="font-bold">Default</h2>
        </CardHeader>
        <CardBody>
          <p className="text-gray-600">通常のシャドウ付きカード</p>
        </CardBody>
      </Card>

      <Card variant="outlined">
        <CardHeader>
          <h2 className="font-bold">Outlined</h2>
        </CardHeader>
        <CardBody>
          <p className="text-gray-600">ボーダー付きカード</p>
        </CardBody>
      </Card>

      <Card variant="elevated">
        <CardHeader>
          <h2 className="font-bold">Elevated</h2>
        </CardHeader>
        <CardBody>
          <p className="text-gray-600">強いシャドウ付きカード</p>
        </CardBody>
      </Card>
    </div>
  );
}
```

## 🎯 演習問題

### 基本: シンプルなCardの実装

`children`を受け取り、白背景・角丸・シャドウで囲むCardコンポーネントを作ってください。

```jsx
function Card({ children }) {
  // ここにコードを書く
  // 白背景、角丸、シャドウを適用する

  return (
    <div>
      {children}
    </div>
  );
}
```

**期待される動作:**
- 白背景で角丸のカードが表示される
- シャドウがついて浮いて見える
- 内側にパディングがある

---

### 応用: header/body/footer + variant

header/body/footer構造を持ち、variant（`default`, `outlined`, `elevated`）を切り替えられるCardコンポーネント群を作ってください。

**要件:**
1. `Card`, `CardHeader`, `CardBody`, `CardFooter` の4つのコンポーネントを作成
2. `Card`は`variant`をpropsで受け取る（デフォルト: `default`）
3. `default`: 通常のシャドウ、`outlined`: ボーダー表示、`elevated`: 強いシャドウ
4. `CardHeader`と`CardFooter`にはボーダー（上下の区切り線）を付ける

**ヒント:**
```jsx
function Card({ variant = 'default', children }) {
  // variantClasses をオブジェクトで定義
  // variant に応じてクラスを切り替える
}
```

---

### 発展: 画像付きカード + アクションボタン

商品カードを作成してください。

**要件:**
1. カード上部に画像エリア（`<img>`タグ、高さ固定）
2. body部分に商品名（太字）と説明文
3. footer部分に価格表示と「カートに追加」ボタン
4. variant（`default`, `outlined`）を切り替え可能
5. 使用例としてAppコンポーネントで2つの商品カードを横に並べる

**完成イメージ:**
```
┌──────────────────┐
│   [商品画像]      │
├──────────────────┤
│ 商品名            │
│ 説明テキスト...    │
├──────────────────┤
│ ¥1,980   [カートに追加] │
└──────────────────┘
```

## ✅ 重要ポイント

- [ ] header/body/footerの3領域に分割し、コンテンツの構造を明確にする
- [ ] variantはオブジェクトでクラスを管理し、統一的なデザインを維持する
- [ ] `overflow-hidden`で角丸と内部コンテンツの表示を正しく制御する
- [ ] 各サブコンポーネント（CardHeader等）は`children`で柔軟にコンテンツを受け取る

**次のテーマ:** [04. Modalコンポーネント](./04-modal-component.md)
