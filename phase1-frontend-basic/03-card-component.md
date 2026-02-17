# 03. Cardコンポーネント - header/body/footer構造、variant管理

## 🎯 このテーマで学ぶこと

- コンテンツの構造化（header / body / footer の3領域分割）
- 複数コンポーネントの組み合わせ（Card + CardHeader + CardBody + CardFooter）
- `children` を活用した柔軟なコンテンツ挿入
- variantパターンの応用（default / outlined / elevated）
- `overflow-hidden` による角丸と内部コンテンツの制御

**なぜ重要か:** Cardは情報を視覚的にグループ化する最も汎用的なレイアウトパターンです。1つの親コンポーネントと複数のサブコンポーネントを組み合わせる「Compound Components」の考え方は、複雑なUIを整理する上での基本設計です。

## 📖 概念

Cardはコンテンツをグループ化して表示するためのコンポーネントです。header（タイトル）、body（本文）、footer（アクション）の3つの領域に分割し、variant（見た目のバリエーション）で表示スタイルを切り替えます。構造的に情報を整理することで、ユーザーが内容を素早く把握できます。

**背景と設計意図:** 全ての情報をフラットに並べると、ユーザーはどこからどこまでが1つのまとまりなのか判断できません。Cardで囲むことで「これは1つの情報の塊です」と視覚的に伝えられます。さらにheader/body/footerに分離することで、タイトル・詳細・アクションの役割が明確になり、ユーザーは情報をスキャンしやすくなります。

**実務での活用場面:** ECサイトの商品一覧、ダッシュボードの統計パネル、SNSの投稿フィード、管理画面のデータ表示、プロフィールカードなど。Webアプリケーションのあらゆる一覧画面・詳細表示で使われる最頻出のレイアウトパターンです。

**よくある誤解:**

- ❌ 「Cardは見た目だけの問題」→ header/body/footerの構造化がコンテンツの意味を明確にする
- ❌ 「childrenに全部入れればいい」→ 領域を分けることで再利用性とレイアウトの一貫性が向上する
- ❌ 「variantはclassNameを直接渡せばいい」→ 定義済みvariantで統一的なデザインを維持する

## 💡 コード例

### 基本: シンプルなCard

まずは最もシンプルなCardから始めます。Cardの本質は「コンテンツを視覚的に1つのグループとして囲む」ことです。白背景・角丸・シャドウの3つを組み合わせることで、背景から浮き上がった「カード」の見た目を作り、ユーザーに「ここが1つの情報のまとまりである」と伝えます。`children`で中身を自由に差し込める設計にすることで、どんなコンテンツにも対応できる汎用的なコンポーネントになります。

```tsx
import type { ReactNode } from 'react';

type Props = {
  children: ReactNode;
};

// childrenで中身を受け取ることで、Card自体はレイアウトの責務だけを持つ
function Card({ children }: Props) {
  return (
    // shadow + rounded-lg で「浮いたカード」の視覚効果を作る
    <div className="bg-white rounded-lg shadow p-6">
      {children}
    </div>
  );
}

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

> **💡 次のステップへ:** シンプルなCardは便利ですが、タイトル・本文・アクションボタンが全て同じ領域に混在しています。次のステージでは、これらを header / body / footer に分離して、コンテンツの役割を構造的に明確にします。

### 応用: header/body/footer分離

Cardの中身をheader（タイトル領域）、body（メインコンテンツ）、footer（アクション領域）の3つのサブコンポーネントに分割します。こうすることで、各領域の役割が明確になり、ボーダーや背景色で視覚的な区切りも付けられます。親のCardに`overflow-hidden`を付けるのは、子要素のコンテンツが角丸からはみ出すのを防ぐためです。

```tsx
import type { ReactNode } from 'react';

type Props = {
  children: ReactNode;
};

function Card({ children }: Props) {
  return (
    // overflow-hidden: 子要素（画像など）が角丸の外にはみ出すのを防ぐ
    // p-6 を外して各セクションに余白を委譲し、ボーダーが端まで届くようにする
    <div className="bg-white rounded-lg shadow overflow-hidden">
      {children}
    </div>
  );
}

type SectionProps = {
  children: ReactNode;
};

// border-b でbodyとの境界線を作り、「ここがタイトル領域」と視覚的に伝える
function CardHeader({ children }: SectionProps) {
  return (
    <div className="px-6 py-4 border-b border-gray-200">
      {children}
    </div>
  );
}

// bodyは境界線なし。メインコンテンツが自然に広がる領域
function CardBody({ children }: SectionProps) {
  return (
    <div className="px-6 py-4">
      {children}
    </div>
  );
}

// border-t + bg-gray-50 でアクション領域を視覚的に分離し、操作可能な場所だと示す
function CardFooter({ children }: SectionProps) {
  return (
    <div className="px-6 py-4 border-t border-gray-200 bg-gray-50">
      {children}
    </div>
  );
}

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

> **💡 次のステップへ:** 構造は整いましたが、現状では全てのCardが同じ見た目です。実際のアプリケーションでは、文脈に応じてCardの外観を切り替えたい場面があります（例: 通常表示、軽い囲み、強調表示）。次のステージでは variant props を導入して、見た目のバリエーションを型安全に管理する方法を学びます。

### 実践: variant（default/outlined/elevated）

Buttonコンポーネントと同様に、Cardにも`variant` propsを導入します。見た目の切り替えロジックをオブジェクトマッピングで管理することで、新しいvariantの追加が容易になり、コンポーネント利用側は`variant="outlined"`と指定するだけで統一されたデザインを適用できます。TypeScriptのユニオン型で許可されるvariantを制限しているため、タイポや未定義のvariantを使おうとするとコンパイル時にエラーになります。

```tsx
import type { ReactNode } from 'react';

type CardProps = {
  // ユニオン型でvariantを制限し、未定義の値を渡せないようにする
  variant?: 'default' | 'outlined' | 'elevated';
  children: ReactNode;
};

function Card({ variant = 'default', children }: CardProps) {
  // variantごとのクラスをオブジェクトで一元管理
  // 新しいvariantを追加する際もここに1行足すだけで済む
  const variantClasses: Record<string, string> = {
    default: 'bg-white shadow',
    outlined: 'bg-white border-2 border-gray-200',  // シャドウの代わりにボーダーで囲む
    elevated: 'bg-white shadow-lg shadow-gray-200',  // より強いシャドウで浮遊感を強調
  };

  return (
    <div className={`rounded-lg overflow-hidden ${variantClasses[variant]}`}>
      {children}
    </div>
  );
}

type SectionProps = {
  children: ReactNode;
};

function CardHeader({ children }: SectionProps) {
  return (
    <div className="px-6 py-4 border-b border-gray-200">
      {children}
    </div>
  );
}

function CardBody({ children }: SectionProps) {
  return (
    <div className="px-6 py-4">
      {children}
    </div>
  );
}

function CardFooter({ children }: SectionProps) {
  return (
    <div className="px-6 py-4 border-t border-gray-200 bg-gray-50">
      {children}
    </div>
  );
}

function App() {
  return (
    <div className="p-4 max-w-md space-y-6">
      {/* default: 一般的な情報表示に使う標準スタイル */}
      <Card variant="default">
        <CardHeader>
          <h2 className="font-bold">Default</h2>
        </CardHeader>
        <CardBody>
          <p className="text-gray-600">通常のシャドウ付きカード</p>
        </CardBody>
      </Card>

      {/* outlined: シャドウを使わず軽い印象にしたい場面向け */}
      <Card variant="outlined">
        <CardHeader>
          <h2 className="font-bold">Outlined</h2>
        </CardHeader>
        <CardBody>
          <p className="text-gray-600">ボーダー付きカード</p>
        </CardBody>
      </Card>

      {/* elevated: 注目させたい重要な情報やCTA向け */}
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

header/body/footer構造を持ち、variant（`default`, `outlined`, `elevated`）を切り替えられるCardコンポーネント群を作ってください。

**要件:**

1. `Card`, `CardHeader`, `CardBody`, `CardFooter` の4つのコンポーネントを作成
2. `Card`は`variant`をpropsで受け取る（デフォルト: `default`）
3. `default`: 通常のシャドウ、`outlined`: ボーダー表示、`elevated`: 強いシャドウ
4. `CardHeader`と`CardFooter`にはボーダー（上下の区切り線）を付ける

**ヒント:**

```tsx
import type { ReactNode } from 'react';

type CardProps = {
  variant?: 'default' | 'outlined' | 'elevated';
  children: ReactNode;
};

function Card({ variant = 'default', children }: CardProps) {
  // variantClasses をオブジェクトで定義
  // variant に応じてクラスを切り替える
}
```

## ✅ 重要ポイント

- [ ] header/body/footerの3領域に分割し、コンテンツの構造を明確にする
- [ ] variantはオブジェクトでクラスを管理し、統一的なデザインを維持する
- [ ] `overflow-hidden`で角丸と内部コンテンツの表示を正しく制御する
- [ ] 各サブコンポーネント（CardHeader等）は`children`で柔軟にコンテンツを受け取る

**次のテーマ:** [04. Modalコンポーネント](./04-modal-component.md)
