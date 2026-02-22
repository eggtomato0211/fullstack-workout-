# 01. Buttonコンポーネント - variant/size/state管理

## 🎯 このテーマで学ぶこと

- propsによるスタイル制御（variant / size）
- オブジェクトマッピングによるクラス管理
- disabled / loading の状態管理

## 📖 なぜButtonコンポーネントを理解する必要があるのか

Buttonは全てのUIライブラリの最も基本的なパターンです。ここで学ぶ「propsでバリエーションを制御し、オブジェクトでクラスを管理する」手法は、Input、Card、Alert等すべてのコンポーネントにそのまま応用できます。

### こう書かないとどうなるか

ボタンを毎回個別にスタイリングするとこうなります：

```tsx
// 各画面で場当たり的にスタイルを書く
<button className="bg-blue-600 text-white px-4 py-2 rounded">保存</button>
<button className="bg-blue-500 text-white px-3 py-1 rounded">送信</button>  {/* 微妙に色が違う */}
<button className="bg-red-600 text-white px-4 py-2">削除</button>           {/* roundedを忘れた */}
```

色のバラつき、角丸の有無など見た目が統一されず、修正時に全箇所を探す必要があります。variantとsizeをpropsで定義し、1つのコンポーネントに集約すれば「ここを直せば全部直る」状態を作れます。

### disabledはCSSだけでは不十分

見た目を半透明にしただけでは、スクリーンリーダーはボタンが押せると判断してしまいます。HTML属性の`disabled`も必ず設定することで、アクセシビリティを確保します。

## 💡 コード例

### 基本: variant + sizeの制御

variantとsizeを別々のオブジェクトで管理することで、3 variant x 3 size = 9パターンでも定義は6行で済みます。if文の分岐が不要で、新しいvariantの追加も1行です。

```tsx
import type { ReactNode } from 'react';

type Props = {
  variant?: 'primary' | 'secondary' | 'danger';
  size?: 'sm' | 'md' | 'lg';
  children: ReactNode;
  onClick?: () => void;
};

function Button({ variant = 'primary', size = 'md', children, onClick }: Props) {
  // なぜオブジェクトで管理するか → if/switchが不要になり、variantの追加が1行で済む
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

  return (
    <button
      className={`rounded font-medium ${variantClasses[variant]} ${sizeClasses[size]}`}
      onClick={onClick}
    >
      {children}
    </button>
  );
}

function App() {
  return (
    <div className="flex items-center gap-4">
      <Button variant="primary" size="sm">小さいボタン</Button>
      <Button variant="secondary">キャンセル</Button>
      <Button variant="danger" size="lg">削除</Button>
    </div>
  );
}
```

### 実践: disabled状態とローディング状態

API呼び出し中にボタンを無効化してスピナーを表示することで、二重送信を防ぎます。CSSで見た目を変えるだけでなく、**HTML属性の`disabled`も必ず設定する**ことがポイントです（スクリーンリーダー対応）。

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

function Button({ variant = 'primary', size = 'md', disabled = false, isLoading = false, children, onClick }: Props) {
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

  // ローディング中もクリック不可にする
  const isDisabled = disabled || isLoading;

  return (
    <button
      className={`rounded font-medium transition-colors ${variantClasses[variant]} ${sizeClasses[size]} ${
        isDisabled ? 'opacity-50 cursor-not-allowed' : ''
      }`}
      onClick={onClick}
      // CSSだけでなくHTML属性でもクリックを無効化（アクセシビリティ対応）
      disabled={isDisabled}
    >
      {isLoading ? (
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

function App() {
  const [isLoading, setIsLoading] = useState(false);

  const handleSave = () => {
    setIsLoading(true);
    // 実務ではAPI呼び出しの完了後に解除する
    setTimeout(() => setIsLoading(false), 2000);
  };

  return (
    <div className="flex gap-4">
      <Button variant="primary" isLoading={isLoading} onClick={handleSave}>保存</Button>
      <Button variant="secondary" disabled>無効なボタン</Button>
    </div>
  );
}
```

## 🎯 演習問題

variant、size、disabledを全て対応するButtonを作ってください。

**要件:**

1. variant(`primary`, `secondary`, `danger`)、size(`sm`, `md`, `lg`)、disabledをpropsで受け取る
2. デフォルト値: variant=`primary`, size=`md`, disabled=`false`
3. disabled時は`opacity-50`と`cursor-not-allowed`を適用し、`disabled`属性を付与

**ヒント:**

```tsx
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

## ✅ 重要ポイント

- [ ] variantやsizeはオブジェクトでクラスを管理すると拡張しやすい
- [ ] `disabled`はCSS（見た目）とHTML属性（操作無効化）の両方を設定する
- [ ] propsにデフォルト値を設定し、省略可能にする
- [ ] ローディング中はユーザーの二重クリックを防止する

**次のテーマ:** [02. Input/Formコンポーネント](./02-input-form-component.md)
