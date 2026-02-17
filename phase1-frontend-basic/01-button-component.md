# 01. Buttonコンポーネント - variant/size/state管理

## 🎯 このテーマで学ぶこと

- propsによるスタイル制御（variant / size）
- オブジェクトマッピングによるクラス管理
- disabled / loading の状態管理
- TypeScriptによる型安全なprops定義
- HTML属性とCSSの両面でアクセシビリティを確保する方法

**なぜ重要か:** Buttonは全てのUIライブラリの基本パターンです。ここで学ぶ「propsでバリエーションを制御し、オブジェクトでクラスを管理する」手法は、他の全コンポーネント（Input、Card、Alert等）にそのまま応用できます。実務では、デザインシステムの最初の構成要素としてButtonが設計されることがほとんどです。

## 📖 概念

Buttonは最も基本的なUIコンポーネントです。再利用可能なButtonを作ることで、アプリ全体で一貫したデザインを保てます。variant（見た目の種類）、size（サイズ）、state（状態）をpropsで制御するパターンを学びます。

**背景と設計意図:** アプリケーション内でボタンを毎回個別にスタイリングすると、見た目のバラつきが生まれ、修正時に全箇所を探す必要が出てきます。variant/sizeをpropsで定義し、1つのコンポーネントに集約することで「ここを直せば全部直る」状態を作れます。これはUIライブラリ（Material UI、shadcn/ui等）が採用している基本設計です。

**実務での活用場面:** ECサイトの「カートに追加」ボタン、管理画面の「保存」「削除」ボタン、フォームの「送信」ボタンなど、あらゆる画面で使われます。特にvariantによる視覚的な重要度の区別（primaryは主要アクション、dangerは破壊的操作）は、ユーザーが迷わず操作できるUIの基本です。

**よくある誤解:**

- ❌ 「classNameを直接書けばいい」→ variant/sizeごとにclassを切り替える設計が保守性を高める
- ❌ 「disabledはCSSだけでいい」→ `disabled`属性 + スタイルの両方が必要（アクセシビリティ）
- ❌ 「ボタンのスタイルは毎回書く」→ 共通コンポーネント化で一貫性を保つ

## 💡 コード例

### 基本: variant（見た目）の切り替え

まずは最もシンプルなパターンから始めます。ボタンの「見た目のバリエーション」をpropsで切り替えられるようにします。ポイントは、classNameを直接書くのではなく、**オブジェクトにvariantとクラスの対応表を持たせる**ことです。こうすることで、新しいvariantを追加するときもオブジェクトに1行追加するだけで済みます。

```tsx
import type { ReactNode } from 'react';

type Props = {
  variant?: 'primary' | 'secondary' | 'danger';
  children: ReactNode;
  onClick?: () => void;
};

function Button({ variant = 'primary', children, onClick }: Props) {
  // オブジェクトでvariantとクラスの対応を定義
  // → 新しいvariantの追加が1行で済み、if文の分岐が不要になる
  const variantClasses: Record<string, string> = {
    primary: 'bg-blue-600 text-white hover:bg-blue-700',
    secondary: 'bg-gray-200 text-gray-800 hover:bg-gray-300',
    danger: 'bg-red-600 text-white hover:bg-red-700',
  };

  return (
    <button
      // テンプレートリテラルで共通クラスとvariant固有クラスを結合
      className={`px-4 py-2 rounded font-medium ${variantClasses[variant]}`}
      onClick={onClick}
    >
      {children}
    </button>
  );
}

// 使用例
function App() {
  return (
    <div className="flex gap-4">
      <Button variant="primary">保存</Button>
      <Button variant="secondary">キャンセル</Button>
      <Button variant="danger">削除</Button>
    </div>
  );
}
```

> **💡 次のステップへ:** 基本ではvariant（色）だけを扱いましたが、実際のUIではサイズも統一したいケースがほとんどです。次の応用では、variantと同じ「オブジェクトマッピング」のパターンをsizeにも適用し、**複数の軸（見た目 × サイズ）を組み合わせる方法**を学びます。

### 応用: size（サイズ）の追加

実務のボタンは「小さいリンク風ボタン」「通常の操作ボタン」「目立たせたい大きなボタン」のようにサイズが必要です。variantと同じくオブジェクトマッピングで管理することで、組み合わせが増えてもコードが複雑にならないのがこのパターンの強みです。

```tsx
import type { ReactNode } from 'react';

type Props = {
  variant?: 'primary' | 'secondary' | 'danger';
  size?: 'sm' | 'md' | 'lg';
  children: ReactNode;
  onClick?: () => void;
};

function Button({ variant = 'primary', size = 'md', children, onClick }: Props) {
  // variant と size を別々のオブジェクトで管理
  // → 3 variant × 3 size = 9パターンだが、定義は6行で済む
  const variantClasses: Record<string, string> = {
    primary: 'bg-blue-600 text-white hover:bg-blue-700',
    secondary: 'bg-gray-200 text-gray-800 hover:bg-gray-300',
    danger: 'bg-red-600 text-white hover:bg-red-700',
  };

  // padding と font-size でサイズ感を制御
  const sizeClasses: Record<string, string> = {
    sm: 'px-3 py-1 text-sm',
    md: 'px-4 py-2 text-base',
    lg: 'px-6 py-3 text-lg',
  };

  return (
    <button
      // variant と size の両方のクラスを組み合わせる
      className={`rounded font-medium ${variantClasses[variant]} ${sizeClasses[size]}`}
      onClick={onClick}
    >
      {children}
    </button>
  );
}

// 使用例
function App() {
  return (
    <div className="flex items-center gap-4">
      <Button size="sm">小さいボタン</Button>
      <Button size="md">普通のボタン</Button>
      <Button size="lg">大きいボタン</Button>
    </div>
  );
}
```

> **💡 次のステップへ:** ここまでは「見た目」の制御でしたが、実際のアプリではボタンの**状態管理**も重要です。たとえば、フォーム送信中にボタンを連打されると二重送信が起きます。次の実践では、`disabled`と`isLoading`という「状態」をpropsに追加し、**見た目と動作の両方を制御する方法**を学びます。

### 実践: disabled状態とローディング状態の管理

実務で最もよく実装するパターンです。API呼び出し中にボタンを無効化してスピナーを表示することで、「処理中であること」をユーザーに伝え、二重送信を防ぎます。ここで大切なのは、CSSで見た目を変えるだけでなく、**HTML属性の`disabled`も必ず設定する**ことです（スクリーンリーダー等のアクセシビリティ対応）。

```tsx
import { useState } from 'react';
import type { ReactNode } from 'react';

type Props = {
  variant?: 'primary' | 'secondary' | 'danger';
  size?: 'sm' | 'md' | 'lg';
  disabled?: boolean;
  isLoading?: boolean;
  children: ReactNode;
  onClick?: () => void;
};

function Button({
  variant = 'primary',
  size = 'md',
  disabled = false,
  isLoading = false,
  children,
  onClick,
}: Props) {
  const variantClasses: Record<string, string> = {
    primary: 'bg-blue-600 text-white hover:bg-blue-700',
    secondary: 'bg-gray-200 text-gray-800 hover:bg-gray-300',
    danger: 'bg-red-600 text-white hover:bg-red-700',
  };

  const sizeClasses: Record<string, string> = {
    sm: 'px-3 py-1 text-sm',
    md: 'px-4 py-2 text-base',
    lg: 'px-6 py-3 text-lg',
  };

  // ローディング中も操作不可にするため、どちらかがtrueならdisabled扱い
  const isDisabled = disabled || isLoading;

  return (
    <button
      className={`rounded font-medium transition-colors ${variantClasses[variant]} ${sizeClasses[size]} ${
        isDisabled ? 'opacity-50 cursor-not-allowed' : '' // 見た目で「押せない」ことを示す
      }`}
      onClick={onClick}
      disabled={isDisabled} // HTML属性でもクリックを無効化（アクセシビリティ対応）
    >
      {isLoading ? (
        // ローディング中はスピナー + テキストに差し替え
        <span className="flex items-center gap-2">
          <span className="animate-spin h-4 w-4 border-2 border-white border-t-transparent rounded-full"></span>
          処理中...
        </span>
      ) : (
        children
      )}
    </button>
  );
}

// 使用例: ボタンクリックで2秒間ローディング状態にするデモ
function App() {
  const [isLoading, setIsLoading] = useState(false);

  const handleSave = () => {
    setIsLoading(true);
    // 2秒後にローディング解除（実務ではAPI呼び出しの完了後に解除する）
    setTimeout(() => setIsLoading(false), 2000);
  };

  return (
    <div className="flex gap-4">
      <Button variant="primary" isLoading={isLoading} onClick={handleSave}>
        保存
      </Button>
      <Button variant="secondary" disabled>
        無効なボタン
      </Button>
    </div>
  );
}
```

## 🎯 演習問題

variant（`primary`, `secondary`, `danger`）、size（`sm`, `md`, `lg`）、`disabled`を全て対応するButtonを作ってください。

**要件:**

1. variant, size, disabledをpropsで受け取る
2. デフォルト値: variant=`primary`, size=`md`, disabled=`false`
3. disabled時は`opacity-50`と`cursor-not-allowed`を適用し、`disabled`属性を付与
4. Tailwind CSSでスタイリング

**ヒント:**

```tsx
import type { ReactNode } from 'react';

type Props = {
  variant?: 'primary' | 'secondary' | 'danger';
  size?: 'sm' | 'md' | 'lg';
  disabled?: boolean;
  children: ReactNode;
  onClick?: () => void;
};

function Button({ variant = 'primary', size = 'md', disabled = false, children, onClick }: Props) {
  // variantClasses, sizeClasses をオブジェクトで定義
  // disabled の場合のスタイルも追加
}
```

**期待される動作:**

- `primary`: 青背景・白文字、`secondary`: グレー背景・黒文字、`danger`: 赤背景・白文字
- `sm`/`md`/`lg`でパディングとフォントサイズが変わる
- `disabled`時は半透明になりクリック不可

## ✅ 重要ポイント

- [ ] variantやsizeはオブジェクトでクラスを管理すると拡張しやすい
- [ ] `disabled`はCSS（見た目）とHTML属性（操作無効化）の両方を設定する
- [ ] propsにデフォルト値を設定し、省略可能にする
- [ ] ローディング中はユーザーの二重クリックを防止する

**次のテーマ:** [02. Input/Formコンポーネント](./02-input-form-component.md)
