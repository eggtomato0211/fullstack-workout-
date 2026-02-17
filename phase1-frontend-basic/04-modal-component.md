# 04. Modalコンポーネント - オーバーレイ、ESCキーで閉じる、背景スクロール防止

## 🎯 このテーマで学ぶこと

- `useEffect` によるイベントリスナーの登録とクリーンアップ
- `position: fixed` + `z-index` によるオーバーレイ実装
- ESCキー・外側クリックによる閉じる操作の実装
- `document.body.style.overflow` による背景スクロール防止
- `e.stopPropagation()` によるイベント伝播の制御

**なぜ重要か:** Modalは「useEffectの副作用管理」を実践的に学べるテーマです。イベントリスナーの登録・解除、DOMの直接操作（body overflow）、条件付きレンダリングなど、React開発の中級スキルを一度に身につけられます。実務ではほぼ全てのアプリに確認ダイアログやフォームモーダルが存在します。

## 📖 概念

Modalは画面の前面にオーバーレイ付きでコンテンツを表示するコンポーネントです。ユーザーの注意をモーダル内に集中させ、操作の確認やフォーム入力など、集中したインタラクションを提供します。ESCキーや外側クリックで閉じる操作、背景スクロールの防止など、UXに関わる実装が多く含まれます。

**背景と設計意図:** 通常のページ遷移ではなく、現在の画面の上にコンテンツを重ねる「モーダル」パターンは、ユーザーの文脈を維持したまま追加情報を表示したり、確認を求めたりするために使います。ただし、「閉じ方がわからないモーダル」はUXの大きな問題になるため、ESCキー・外側クリック・閉じるボタンの3つの閉じ手段を提供するのがベストプラクティスです。

**実務での活用場面:** 削除確認ダイアログ、ログインフォームのポップアップ、画像のプレビュー表示、利用規約の表示、設定パネルなど。特にSPAでは画面遷移を減らすためにモーダルが多用されます。

**よくある誤解:**

- ❌ 「表示/非表示の切り替えだけでいい」→ ESCキー・外側クリック・背景スクロール防止まで実装して初めて使いやすいモーダルになる
- ❌ 「オーバーレイはdivを重ねるだけ」→ `position: fixed`とz-indexで画面全体を覆う必要がある
- ❌ 「モーダルを閉じたらそれで終わり」→ イベントリスナーのクリーンアップを忘れるとメモリリークになる

## 💡 コード例

### 基本: オーバーレイ表示

まずはモーダルの最小構成から始めます。`isOpen`が`false`のときは`null`を返して何もレンダリングしないことで、条件付きレンダリングの基本パターンを押さえます。`position: fixed`と`z-index`を使ってオーバーレイを画面全体に重ねることが、モーダルの視覚的な土台になります。

```tsx
import { useState, type ReactNode } from 'react';

type Props = {
  isOpen: boolean;
  onClose: () => void;
  children: ReactNode;
};

function Modal({ isOpen, onClose, children }: Props) {
  // 閉じている状態では何もレンダリングしない（早期リターンパターン）
  if (!isOpen) return null;

  return (
    // fixed + inset-0 で画面全体を覆い、z-50 で他の要素より前面に配置する
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      {/* オーバーレイ: fixedで画面全体を覆い、半透明の黒背景でモーダルに注意を集中させる */}
      <div className="fixed inset-0 bg-black bg-opacity-50" />
      {/* モーダル本体: relativeでオーバーレイより前面に表示される */}
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

> **💡 次のステップへ:** 基本のモーダルは✕ボタンでしか閉じられません。実際のアプリでは、ユーザーがESCキーやオーバーレイクリックでも直感的に閉じられることが期待されます。次はこれらのインタラクションを`useEffect`で実装します。

### 応用: ESCキー + 外側クリックで閉じる

ここでは`useEffect`を使ってグローバルなキーボードイベントを監視します。Reactのイベントシステムではコンポーネント外のキー入力を拾えないため、`document`にイベントリスナーを直接登録する必要があります。オーバーレイクリックは`onClick`で実装し、モーダル本体には`stopPropagation`でクリックイベントの伝播を止めることで「外側だけ反応する」挙動を実現します。

```tsx
import { useState, useEffect, type ReactNode } from 'react';

type Props = {
  isOpen: boolean;
  onClose: () => void;
  children: ReactNode;
};

function Modal({ isOpen, onClose, children }: Props) {
  // グローバルなキーボードイベントはReactのJSXでは捕捉できないため、
  // useEffect + document.addEventListener で直接DOMに登録する
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.key === 'Escape') {
        onClose();
      }
    };

    // モーダルが開いているときだけリスナーを登録する（不要な監視を避ける）
    if (isOpen) {
      document.addEventListener('keydown', handleKeyDown);
    }

    // クリーンアップ: モーダルが閉じたとき・アンマウント時にリスナーを解除する
    // これを忘れるとリスナーが蓄積してメモリリークになる
    return () => {
      document.removeEventListener('keydown', handleKeyDown);
    };
  }, [isOpen, onClose]);

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      {/* オーバーレイ: onClickを設定し、背景クリックでモーダルを閉じる */}
      <div
        className="fixed inset-0 bg-black bg-opacity-50"
        onClick={onClose}
      />
      {/* stopPropagation: モーダル内部のクリックがオーバーレイまで伝播するのを防ぐ
          これがないとモーダル内をクリックしても閉じてしまう */}
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

> **💡 次のステップへ:** ESCキーと外側クリックで閉じる操作が実装できました。しかし、モーダルが開いている間も背景がスクロールできてしまうと、ユーザーの注意がモーダルの外に逸れてしまいます。次は`document.body.style.overflow`による背景スクロール防止とアニメーションを追加して、本格的なモーダルに仕上げます。

### 実践: スクロール防止 + アニメーション

本格的なモーダルでは、表示中に背景がスクロールしないよう`body`の`overflow`を直接操作します。これはReactの宣言的なアプローチから外れるDOM直接操作ですが、CSSの制約上この方法が最も確実です。アニメーションを加えることで、モーダルの出現が唐突にならず、ユーザーに「新しいレイヤーが重なった」ことを視覚的に伝えられます。

```tsx
import { useState, useEffect, type ReactNode } from 'react';

type Props = {
  isOpen: boolean;
  onClose: () => void;
  children: ReactNode;
};

function Modal({ isOpen, onClose, children }: Props) {
  // ESCキーで閉じる（応用ステージと同じパターン）
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

  // 背景スクロール防止: bodyのoverflowを直接操作する
  // useEffectを分けることで、関心ごとが明確に分離される
  useEffect(() => {
    if (isOpen) {
      // hidden にすることでページ全体のスクロールを無効化する
      document.body.style.overflow = 'hidden';
    }

    // クリーンアップ: 元のスクロール状態に戻す
    // 空文字を代入するとブラウザのデフォルト挙動に戻る
    return () => {
      document.body.style.overflow = '';
    };
  }, [isOpen]);

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      {/* フェードインで「背景が暗くなる」過程を見せ、レイヤーの出現を伝える */}
      <div
        className="fixed inset-0 bg-black bg-opacity-50 animate-fade-in"
        onClick={onClose}
      />
      {/* スケールアニメーションで「手前に飛び出す」印象を与える */}
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

## ✅ 重要ポイント

- [ ] `position: fixed`と`z-index`でオーバーレイを画面全体に表示する
- [ ] ESCキーのイベントリスナーはクリーンアップ関数で解除する（メモリリーク防止）
- [ ] `e.stopPropagation()`でモーダル本体のクリックが外側クリックと判定されるのを防ぐ
- [ ] モーダル表示中は`document.body.style.overflow = 'hidden'`で背景スクロールを防止する

**次のテーマ:** [05. Dropdownコンポーネント](./05-dropdown-component.md)
