# 03. Cardコンポーネント - header/body/footer構造、variant管理

## 🎯 このテーマで学ぶこと

- header / body / footer の3領域分割
- `children` を活用した柔軟なコンテンツ挿入
- variantパターンの応用（default / outlined / elevated）

## 📖 なぜCardコンポーネントを理解する必要があるのか

Cardは情報を視覚的にグループ化する最も汎用的なレイアウトパターンです。ECサイトの商品一覧、SNSの投稿フィード、ダッシュボードの統計パネルなど、Webアプリケーションのあらゆる画面で使われます。

### こう書かないとどうなるか

情報をフラットに並べるとこうなります：

```tsx
// 構造なしでベタ書き
<h2>ユーザー情報</h2>
<p>山田太郎</p>
<p>yamada@example.com</p>
<button>編集する</button>

<h2>注文履歴</h2>
<p>商品A</p>
<button>詳細</button>
```

どこからどこまでが1つのまとまりなのか視覚的に判断できません。Cardで囲むことで「これは1つの情報の塊」と伝えられ、さらにheader/body/footerに分離することでタイトル・詳細・アクションの役割が明確になります。

### なぜサブコンポーネントに分けるのか

`children`に全部入れれば動きますが、header/body/footerに分けることでボーダーや背景色で視覚的な区切りを付けられ、レイアウトの一貫性が保たれます。

## 💡 コード例

### 基本: header/body/footer分離 + variant

Cardの構造（header/body/footer）とvariantによる見た目の切り替えを合わせて学びます。`overflow-hidden`は子要素（画像など）が角丸からはみ出すのを防ぐためです。

```tsx
import type { ReactNode } from 'react';

type CardProps = {
  variant?: 'default' | 'outlined' | 'elevated';
  children: ReactNode;
};

function Card({ variant = 'default', children }: CardProps) {
  // variantごとのクラスをオブジェクトで管理
  const variantClasses: Record<string, string> = {
    default: 'bg-white shadow',             // 標準的なカード
    outlined: 'bg-white border-2 border-gray-200',  // シャドウなし、軽い印象
    elevated: 'bg-white shadow-lg shadow-gray-200',  // 強いシャドウ、強調表示
  };

  return (
    // overflow-hidden: 子要素が角丸の外にはみ出すのを防ぐ
    <div className={`rounded-lg overflow-hidden ${variantClasses[variant]}`}>
      {children}
    </div>
  );
}

type SectionProps = { children: ReactNode };

// border-bでbodyとの境界線を作り、タイトル領域を明示する
function CardHeader({ children }: SectionProps) {
  return <div className="px-6 py-4 border-b border-gray-200">{children}</div>;
}

function CardBody({ children }: SectionProps) {
  return <div className="px-6 py-4">{children}</div>;
}

// bg-gray-50 + border-tでアクション領域を視覚的に分離
function CardFooter({ children }: SectionProps) {
  return <div className="px-6 py-4 border-t border-gray-200 bg-gray-50">{children}</div>;
}

function App() {
  return (
    <div className="p-4 max-w-md space-y-6">
      <Card variant="default">
        <CardHeader>
          <h2 className="font-bold">ユーザー情報</h2>
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

      <Card variant="outlined">
        <CardHeader><h2 className="font-bold">Outlined</h2></CardHeader>
        <CardBody><p className="text-gray-600">ボーダー付きカード</p></CardBody>
      </Card>

      <Card variant="elevated">
        <CardHeader><h2 className="font-bold">Elevated</h2></CardHeader>
        <CardBody><p className="text-gray-600">強いシャドウ付きカード</p></CardBody>
      </Card>
    </div>
  );
}
```

### 実践: 商品カード一覧

実務に近いパターンとして、APIから取得した商品データをCardで表示する例です。データの配列を`map`でレンダリングし、各Cardの構造を統一することで、デザインの一貫性が保たれます。

```tsx
import type { ReactNode } from 'react';

type CardProps = { variant?: 'default' | 'outlined' | 'elevated'; children: ReactNode };
type SectionProps = { children: ReactNode };

function Card({ variant = 'default', children }: CardProps) {
  const variantClasses: Record<string, string> = {
    default: 'bg-white shadow', outlined: 'bg-white border-2 border-gray-200', elevated: 'bg-white shadow-lg',
  };
  return <div className={`rounded-lg overflow-hidden ${variantClasses[variant]}`}>{children}</div>;
}
function CardHeader({ children }: SectionProps) { return <div className="px-6 py-4 border-b border-gray-200">{children}</div>; }
function CardBody({ children }: SectionProps) { return <div className="px-6 py-4">{children}</div>; }
function CardFooter({ children }: SectionProps) { return <div className="px-6 py-4 border-t border-gray-200 bg-gray-50">{children}</div>; }

// 実務では型定義を別ファイルに切り出すことが多い
type Product = {
  id: number;
  name: string;
  price: number;
  description: string;
};

const products: Product[] = [
  { id: 1, name: 'TypeScript入門', price: 3000, description: '型安全なJavaScript開発を学ぶ' },
  { id: 2, name: 'React実践ガイド', price: 3500, description: 'コンポーネント設計の基礎から応用まで' },
  { id: 3, name: 'Go言語の基礎', price: 2800, description: 'シンプルで高速なバックエンド開発' },
];

function App() {
  return (
    <div className="p-4 grid grid-cols-1 md:grid-cols-3 gap-6">
      {products.map((product) => (
        <Card key={product.id}>
          <CardHeader>
            <h2 className="font-bold">{product.name}</h2>
          </CardHeader>
          <CardBody>
            <p className="text-gray-600 text-sm">{product.description}</p>
            <p className="mt-2 text-lg font-bold text-blue-600">
              ¥{product.price.toLocaleString()}
            </p>
          </CardBody>
          <CardFooter>
            <button className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700">
              カートに追加
            </button>
          </CardFooter>
        </Card>
      ))}
    </div>
  );
}
```

## 🎯 演習問題

header/body/footer構造を持ち、variantを切り替えられるCardコンポーネント群を作ってください。

**要件:**

1. `Card`, `CardHeader`, `CardBody`, `CardFooter` の4コンポーネントを作成
2. `Card`は`variant`をpropsで受け取る（デフォルト: `default`）
3. `default`: 通常のシャドウ、`outlined`: ボーダー、`elevated`: 強いシャドウ

**ヒント:**

```tsx
type CardProps = {
  variant?: 'default' | 'outlined' | 'elevated';
  children: ReactNode;
};

function Card({ variant = 'default', children }: CardProps) {
  // variantClasses をオブジェクトで定義
}
```

## ✅ 重要ポイント

- [ ] header/body/footerの3領域分割でコンテンツの構造を明確にする
- [ ] variantはオブジェクトでクラスを管理し、統一的なデザインを維持する
- [ ] `overflow-hidden`で角丸と内部コンテンツの表示を制御する
- [ ] `children`で中身を柔軟に差し込める設計にする

**次のテーマ:** [04. Modalコンポーネント](./04-modal-component.md)
