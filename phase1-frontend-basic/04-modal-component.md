# 04. Modalコンポーネント - オーバーレイ、ESCキーで閉じる、背景スクロール防止

## 📖 概念

Modalは画面の前面にオーバーレイ付きでコンテンツを表示するコンポーネントです。ユーザーの注意をモーダル内に集中させ、操作の確認やフォーム入力など、集中したインタラクションを提供します。ESCキーや外側クリックで閉じる操作、背景スクロールの防止など、UXに関わる実装が多く含まれます。

**よくある誤解:**
- ❌ 「表示/非表示の切り替えだけでいい」→ ESCキー・外側クリック・背景スクロール防止まで実装して初めて使いやすいモーダルになる
- ❌ 「オーバーレイはdivを重ねるだけ」→ `position: fixed`とz-indexで画面全体を覆う必要がある
- ❌ 「モーダルを閉じたらそれで終わり」→ イベントリスナーのクリーンアップを忘れるとメモリリークになる

## 💡 コード例

### 基本: オーバーレイ表示

```tsx
import { useState, type ReactNode } from 'react';

type Props = {
  isOpen: boolean;
  onClose: () => void;
  children: ReactNode;
};

function Modal({ isOpen, onClose, children }: Props) {
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      {/* オーバーレイ（半透明の背景） */}
      <div className="fixed inset-0 bg-black bg-opacity-50" />
      {/* モーダル本体 */}
      <div className="relative bg-white rounded-lg shadow-xl p-6 max-w-md w-full mx-4">
        <button
          onClick={onClose}
          className="absolute top-3 right-3 text-gray-400 hover:text-gray-600"
        >
          ✕
        </button>
        {children}
      </div>
    </div>
  );
}

// 使用例
function App() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="p-4">
      <button
        onClick={() => setIsOpen(true)}
        className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
      >
        モーダルを開く
      </button>

      <Modal isOpen={isOpen} onClose={() => setIsOpen(false)}>
        <h2 className="text-lg font-bold mb-2">タイトル</h2>
        <p className="text-gray-600">モーダルの内容です。</p>
      </Modal>
    </div>
  );
}
```

### 応用: ESCキー + 外側クリックで閉じる

```tsx
import { useState, useEffect, type ReactNode } from 'react';

type Props = {
  isOpen: boolean;
  onClose: () => void;
  children: ReactNode;
};

function Modal({ isOpen, onClose, children }: Props) {
  // ESCキーで閉じる
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.key === 'Escape') {
        onClose();
      }
    };

    if (isOpen) {
      document.addEventListener('keydown', handleKeyDown);
    }

    return () => {
      document.removeEventListener('keydown', handleKeyDown);
    };
  }, [isOpen, onClose]);

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      {/* オーバーレイ: クリックで閉じる */}
      <div
        className="fixed inset-0 bg-black bg-opacity-50"
        onClick={onClose}
      />
      {/* モーダル本体: クリックイベントの伝播を止める */}
      <div
        className="relative bg-white rounded-lg shadow-xl p-6 max-w-md w-full mx-4"
        onClick={(e) => e.stopPropagation()}
      >
        <button
          onClick={onClose}
          className="absolute top-3 right-3 text-gray-400 hover:text-gray-600"
        >
          ✕
        </button>
        {children}
      </div>
    </div>
  );
}

// 使用例
function App() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="p-4">
      <button
        onClick={() => setIsOpen(true)}
        className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
      >
        モーダルを開く
      </button>

      <Modal isOpen={isOpen} onClose={() => setIsOpen(false)}>
        <h2 className="text-lg font-bold mb-2">確認</h2>
        <p className="text-gray-600">
          ESCキーまたは外側クリックで閉じられます。
        </p>
      </Modal>
    </div>
  );
}
```

### 実践: スクロール防止 + アニメーション

```tsx
import { useState, useEffect, type ReactNode } from 'react';

type Props = {
  isOpen: boolean;
  onClose: () => void;
  children: ReactNode;
};

function Modal({ isOpen, onClose, children }: Props) {
  // ESCキーで閉じる
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.key === 'Escape') {
        onClose();
      }
    };

    if (isOpen) {
      document.addEventListener('keydown', handleKeyDown);
    }

    return () => {
      document.removeEventListener('keydown', handleKeyDown);
    };
  }, [isOpen, onClose]);

  // 背景スクロール防止
  useEffect(() => {
    if (isOpen) {
      document.body.style.overflow = 'hidden';
    }

    return () => {
      document.body.style.overflow = '';
    };
  }, [isOpen]);

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      {/* オーバーレイ: フェードインアニメーション */}
      <div
        className="fixed inset-0 bg-black bg-opacity-50 animate-fade-in"
        onClick={onClose}
      />
      {/* モーダル本体: スケールアニメーション */}
      <div
        className="relative bg-white rounded-lg shadow-xl p-6 max-w-md w-full mx-4 animate-scale-in"
        onClick={(e) => e.stopPropagation()}
      >
        <button
          onClick={onClose}
          className="absolute top-3 right-3 text-gray-400 hover:text-gray-600"
        >
          ✕
        </button>
        {children}
      </div>
    </div>
  );
}

// 使用例
function App() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="p-4">
      <button
        onClick={() => setIsOpen(true)}
        className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
      >
        モーダルを開く
      </button>

      {/* 背景スクロール確認用の長いコンテンツ */}
      <div className="mt-4 space-y-4">
        {Array.from({ length: 20 }, (_, i) => (
          <p key={i} className="text-gray-600">
            スクロール確認用テキスト {i + 1}
          </p>
        ))}
      </div>

      <Modal isOpen={isOpen} onClose={() => setIsOpen(false)}>
        <h2 className="text-lg font-bold mb-2">モーダル</h2>
        <p className="text-gray-600">
          背景のスクロールが防止されていることを確認してください。
        </p>
      </Modal>
    </div>
  );
}
```

> **補足:** `animate-fade-in`と`animate-scale-in`はTailwind CSSのカスタムアニメーションです。`tailwind.config.js`に以下を追加してください:
> ```js
> // tailwind.config.js
> module.exports = {
>   theme: {
>     extend: {
>       keyframes: {
>         'fade-in': { '0%': { opacity: '0' }, '100%': { opacity: '1' } },
>         'scale-in': { '0%': { opacity: '0', transform: 'scale(0.95)' }, '100%': { opacity: '1', transform: 'scale(1)' } },
>       },
>       animation: {
>         'fade-in': 'fade-in 0.2s ease-out',
>         'scale-in': 'scale-in 0.2s ease-out',
>       },
>     },
>   },
> };
> ```

## 🎯 演習問題

### 基本: 表示/非表示の切り替え

`isOpen`で表示/非表示を切り替えるModalコンポーネントを作ってください。

```tsx
import type { ReactNode } from 'react';

type Props = {
  isOpen: boolean;
  onClose: () => void;
  children: ReactNode;
};

function Modal({ isOpen, onClose, children }: Props) {
  // ここにコードを書く
  // isOpen が false なら何も表示しない
  // オーバーレイ + モーダル本体 + 閉じるボタンを実装

  return null;
}
```

**期待される動作:**
- ボタンクリックでモーダルが開く
- 半透明の黒い背景（オーバーレイ）が表示される
- ✕ボタンでモーダルが閉じる

---

### 応用: ESCキー + 外側クリック

基本のModalに、ESCキーと外側クリックで閉じる機能を追加してください。

**要件:**
1. ESCキーを押すとモーダルが閉じる（`useEffect` + `addEventListener`）
2. オーバーレイ部分をクリックするとモーダルが閉じる
3. モーダル本体をクリックしても閉じない（`e.stopPropagation()`）
4. モーダルが閉じた後にイベントリスナーをクリーンアップする

**ヒント:**
```tsx
useEffect(() => {
  const handleKeyDown = (e: KeyboardEvent) => {
    // ESCキーの判定
  };
  // addEventListener と return でクリーンアップ
}, [isOpen, onClose]);
```

---

### 発展: スクロール防止 + 確認ダイアログ

Modalを使って「削除確認ダイアログ」を作成してください。

**要件:**
1. 背景スクロール防止（`document.body.style.overflow`）
2. 「本当に削除しますか？」のメッセージ表示
3. 「キャンセル」と「削除する」の2つのボタンを配置
4. 「削除する」ボタンは赤い色で目立たせる
5. 「削除する」クリックで`alert('削除しました')`を実行してモーダルを閉じる
6. ESCキーとオーバーレイクリックでも閉じる

**完成イメージ:**
```
┌──────────────────────┐
│                    ✕ │
│  本当に削除しますか？   │
│                      │
│  この操作は取り消せません │
│                      │
│  [キャンセル] [削除する] │
└──────────────────────┘
```

## ✅ 重要ポイント

- [ ] `position: fixed`と`z-index`でオーバーレイを画面全体に表示する
- [ ] ESCキーのイベントリスナーはクリーンアップ関数で解除する（メモリリーク防止）
- [ ] `e.stopPropagation()`でモーダル本体のクリックが外側クリックと判定されるのを防ぐ
- [ ] モーダル表示中は`document.body.style.overflow = 'hidden'`で背景スクロールを防止する

**次のテーマ:** [05. Dropdownコンポーネント](./05-dropdown-component.md)
