# 04. Modalコンポーネント - オーバーレイ、ESCキーで閉じる、背景スクロール防止

## 🎯 このテーマで学ぶこと

- `position: fixed` + `z-index` によるオーバーレイ実装
- `useEffect` によるイベントリスナーの登録とクリーンアップ
- ESCキー・外側クリック・背景スクロール防止

## 📖 なぜModalの実装を理解する必要があるのか

Modalは「useEffectの副作用管理」を実践的に学べるテーマです。イベントリスナーの登録・解除、DOMの直接操作など、React開発の中級スキルが一度に身につきます。

### こう書かないとどうなるか

「表示/非表示を切り替えるだけ」のモーダルは問題だらけです：

```tsx
// 最低限の実装だと...
{isOpen && <div className="modal">内容</div>}
// → ESCキーで閉じられない
// → 外側クリックで閉じられない
// → 背景がスクロールしてしまう
// → ユーザーは「閉じ方がわからない」
```

使いやすいモーダルには**3つの閉じ手段**が必要です：閉じるボタン、ESCキー、外側（オーバーレイ）クリック。これはWebのUX慣習であり、どれか1つでも欠けると「壊れている」と感じられます。

### なぜuseEffectのクリーンアップが重要か

ESCキーを検知するためにdocumentにイベントリスナーを登録しますが、モーダルが閉じた後もリスナーが残り続けると、メモリリークや意図しない動作の原因になります。useEffectのクリーンアップ関数でリスナーを解除することが必須です。

## 💡 コード例

### 基本: オーバーレイ + ESCキー + 外側クリック

モーダルの3つの閉じ手段を一気に実装します。`position: fixed`で画面全体を覆い、`e.stopPropagation()`でモーダル本体のクリックが外側クリックと判定されるのを防ぎます。

```tsx
import { useState, useEffect, type ReactNode } from 'react';

type Props = {
  isOpen: boolean;
  onClose: () => void;
  children: ReactNode;
};

function Modal({ isOpen, onClose, children }: Props) {
  // ESCキーの検知：Reactのイベントシステムでは捕捉できないため、
  // document に直接リスナーを登録する
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.key === 'Escape') onClose();
    };

    if (isOpen) {
      document.addEventListener('keydown', handleKeyDown);
    }

    // クリーンアップ：モーダルが閉じた/アンマウント時にリスナーを解除
    // これを忘れるとリスナーが蓄積してメモリリークになる
    return () => document.removeEventListener('keydown', handleKeyDown);
  }, [isOpen, onClose]);

  if (!isOpen) return null;

  return (
    // fixed + inset-0で画面全体を覆い、z-50で他の要素より前面に配置
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      {/* オーバーレイ：クリックでモーダルを閉じる */}
      <div className="fixed inset-0 bg-black bg-opacity-50" onClick={onClose} />
      {/* stopPropagation：モーダル内部のクリックがオーバーレイまで伝播するのを防ぐ
          これがないとモーダル内をクリックしても閉じてしまう */}
      <div
        className="relative bg-white rounded-lg shadow-xl p-6 max-w-md w-full mx-4"
        onClick={(e) => e.stopPropagation()}
      >
        <button onClick={onClose} className="absolute top-3 right-3 text-gray-400 hover:text-gray-600">
          ✕
        </button>
        {children}
      </div>
    </div>
  );
}

function App() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="p-4">
      <button onClick={() => setIsOpen(true)} className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">
        モーダルを開く
      </button>
      <Modal isOpen={isOpen} onClose={() => setIsOpen(false)}>
        <h2 className="text-lg font-bold mb-2">確認</h2>
        <p className="text-gray-600">ESCキーまたは外側クリックで閉じられます。</p>
      </Modal>
    </div>
  );
}
```

### 実践: 背景スクロール防止 + 確認ダイアログ

実務で最もよくある「削除確認モーダル」を実装します。背景スクロールを防止するために`body`の`overflow`を直接操作しますが、これはReactの宣言的なアプローチから外れるDOM直接操作です。useEffectのクリーンアップで確実に元に戻すことが重要です。

```tsx
import { useState, useEffect, type ReactNode } from 'react';

type ModalProps = {
  isOpen: boolean;
  onClose: () => void;
  children: ReactNode;
};

function Modal({ isOpen, onClose, children }: ModalProps) {
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.key === 'Escape') onClose();
    };
    if (isOpen) document.addEventListener('keydown', handleKeyDown);
    return () => document.removeEventListener('keydown', handleKeyDown);
  }, [isOpen, onClose]);

  // 背景スクロール防止：関心ごとを分離するためuseEffectを分ける
  useEffect(() => {
    if (isOpen) {
      document.body.style.overflow = 'hidden';
    }
    // クリーンアップ：空文字でブラウザのデフォルト挙動に戻す
    return () => { document.body.style.overflow = ''; };
  }, [isOpen]);

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      <div className="fixed inset-0 bg-black bg-opacity-50" onClick={onClose} />
      <div className="relative bg-white rounded-lg shadow-xl p-6 max-w-md w-full mx-4"
        onClick={(e) => e.stopPropagation()}>
        <button onClick={onClose} className="absolute top-3 right-3 text-gray-400 hover:text-gray-600">✕</button>
        {children}
      </div>
    </div>
  );
}

// 実務でよくある「削除確認ダイアログ」
function App() {
  const [isOpen, setIsOpen] = useState(false);

  const handleDelete = () => {
    alert('削除しました');
    setIsOpen(false);
  };

  return (
    <div className="p-4">
      <button onClick={() => setIsOpen(true)}
        className="bg-red-600 text-white px-4 py-2 rounded hover:bg-red-700">
        アイテムを削除
      </button>

      <Modal isOpen={isOpen} onClose={() => setIsOpen(false)}>
        <h2 className="text-lg font-bold mb-2">本当に削除しますか？</h2>
        <p className="text-gray-600 mb-4">この操作は取り消せません。</p>
        <div className="flex gap-3 justify-end">
          <button onClick={() => setIsOpen(false)}
            className="px-4 py-2 border border-gray-300 rounded hover:bg-gray-50">
            キャンセル
          </button>
          <button onClick={handleDelete}
            className="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700">
            削除する
          </button>
        </div>
      </Modal>

      {/* 背景スクロール確認用 */}
      <div className="mt-4 space-y-4">
        {Array.from({ length: 20 }, (_, i) => (
          <p key={i} className="text-gray-600">スクロール確認用テキスト {i + 1}</p>
        ))}
      </div>
    </div>
  );
}
```

## 🎯 演習問題

ESCキーと外側クリックで閉じられるModalを作ってください。

**要件:**

1. ESCキーを押すとモーダルが閉じる（`useEffect` + `addEventListener`）
2. オーバーレイをクリックするとモーダルが閉じる
3. モーダル本体をクリックしても閉じない（`e.stopPropagation()`）
4. クリーンアップでイベントリスナーを解除する

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
- [ ] ESCキーのリスナーはクリーンアップで解除する（メモリリーク防止）
- [ ] `e.stopPropagation()`でモーダル内クリックの伝播を防ぐ
- [ ] `document.body.style.overflow = 'hidden'`で背景スクロールを防止する

**次のテーマ:** [05. Dropdownコンポーネント](./05-dropdown-component.md)
