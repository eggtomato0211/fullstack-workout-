# 05. Dropdownコンポーネント - メニュー表示、外側クリックで閉じる

## 🎯 このテーマで学ぶこと

- `useRef` によるDOM要素への参照取得
- `useRef` + `useEffect` + `contains()` による外側クリック検知パターン
- `position: absolute` によるメニューの位置制御
- キーボードナビゲーション（ArrowDown / ArrowUp / Enter / Escape）
- 選択状態の管理と親コンポーネントへのコールバック通知

**なぜ重要か:** 「外側クリックで閉じる」パターンはDropdownに限らず、ツールチップ、ポップオーバー、コンテキストメニューなど多くのUI要素で必要になります。ここで学ぶ `useRef` + `useEffect` + `contains()` の組み合わせは、React開発者が最も頻繁に使う「DOM連携パターン」の1つです。

## 📖 概念

Dropdownはトリガー要素のクリックでメニューを表示し、項目を選択させるコンポーネントです。外側クリックでメニューを閉じる仕組み（`useRef` + `useEffect`）はReact開発で頻出するパターンです。キーボード操作への対応まで含めると、アクセシブルなUIの実装力が身につきます。

**背景と設計意図:** HTMLネイティブの`<select>`要素はスタイリングの自由度が非常に低く、デザインの要件を満たせないことが多いです。そのため実務ではほぼ全てのプロジェクトでカスタムDropdownを実装します。ただし、自作する場合はネイティブ`<select>`が自動で提供してくれるキーボード操作やアクセシビリティを自分で実装する必要があり、それがこのテーマの難しさでもあります。

**実務での活用場面:** ヘッダーのユーザーメニュー、フィルタ・ソートの選択UI、設定画面のオプション選択、ナビゲーションのサブメニューなど。特にSaaSの管理画面ではテーブルの各行にアクションメニュー（編集・削除・複製等）をDropdownで配置するパターンが定番です。

**よくある誤解:**

- ❌ 「stateだけで開閉管理すればいい」→ 外側クリックで閉じるには`useRef`でDOM参照が必要
- ❌ 「onBlurで外側クリックを検知できる」→ onBlurはフォーカス管理に依存し、意図しない挙動になりやすい
- ❌ 「メニューはトグルボタンの直下に置けばいい」→ `position: absolute`で正確に位置を制御する必要がある

## 💡 コード例

### 基本: トグルで開閉

まずは最もシンプルな形から始めます。`useState`の真偽値1つだけでメニューの開閉を制御します。このステージの目的は、Dropdownの基本構造（トリガーボタン + メニューリスト）と、条件付きレンダリングによる表示切り替えの仕組みを理解することです。外側クリックやキーボード操作はまだ考えず、「開閉の骨格」に集中します。

```tsx
import { useState } from 'react';

type Props = {
  label: string;
  items: string[];
};

function Dropdown({ label, items }: Props) {
  const [isOpen, setIsOpen] = useState(false);

  return (
    // relative を指定することで、子要素の absolute 配置の基準点になる
    <div className="relative inline-block">
      <button
        onClick={() => setIsOpen(!isOpen)}
        className="bg-white border border-gray-300 rounded px-4 py-2 hover:bg-gray-50 flex items-center gap-2"
      >
        {label}
        {/* isOpen に応じて矢印を回転させ、開閉状態を視覚的にフィードバック */}
        <span className={`transition-transform ${isOpen ? 'rotate-180' : ''}`}>
          ▼
        </span>
      </button>

      {/* 条件付きレンダリング: isOpen が false のときはDOMに存在しない */}
      {isOpen && (
        <ul className="absolute top-full left-0 mt-1 w-48 bg-white border border-gray-200 rounded-lg shadow-lg py-1 z-10">
          {items.map((item) => (
            <li key={item}>
              <button
                className="w-full text-left px-4 py-2 hover:bg-gray-100 text-gray-700"
                onClick={() => {
                  alert(item);
                  // 項目選択後にメニューを閉じる（ユーザーの操作が完了したため）
                  setIsOpen(false);
                }}
              >
                {item}
              </button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

// 使用例
function App() {
  return (
    <div className="p-4">
      <Dropdown
        label="メニュー"
        items={['プロフィール', '設定', 'ログアウト']}
      />
    </div>
  );
}
```

> **💡 次のステップへ:** この基本形にはユーザビリティ上の大きな問題があります。メニューが開いた状態で「メニューの外側」をクリックしても閉じません。実際のUIではユーザーは「どこか別の場所をクリックすればメニューが閉じる」と期待するため、次のステージで `useRef` を使った外側クリック検知を追加します。

### 応用: 外側クリックで閉じる（useRef + useEffect）

基本形では `useState` だけで開閉を管理していましたが、「Dropdown領域の外側がクリックされた」ことを検知するには、ReactのState管理だけでは不十分です。なぜなら、外側のクリックはDropdownコンポーネントの外で発生するイベントであり、Reactのイベントハンドラでは捕捉できないからです。ここで `useRef` によるDOM参照と `document` レベルのイベントリスナーを組み合わせるパターンが必要になります。

```tsx
import { useState, useRef, useEffect } from 'react';

type Props = {
  label: string;
  items: string[];
};

function Dropdown({ label, items }: Props) {
  const [isOpen, setIsOpen] = useState(false);
  // useRef でDropdown全体のDOM要素を参照し、クリック位置の判定に使う
  const dropdownRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (e: MouseEvent) => {
      // contains() で「クリックされた要素がDropdown内かどうか」を判定
      // Dropdown外であれば閉じる
      if (dropdownRef.current && !dropdownRef.current.contains(e.target as Node)) {
        setIsOpen(false);
      }
    };

    // メニューが開いているときだけリスナーを登録し、不要なイベント処理を避ける
    if (isOpen) {
      document.addEventListener('mousedown', handleClickOutside);
    }

    // クリーンアップ: メニューが閉じたとき、またはアンマウント時にリスナーを解除
    // これを忘れるとメモリリークや意図しない動作の原因になる
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, [isOpen]);

  return (
    // ref を最外層の div に付けることで、ボタンもメニューも含めた領域全体を参照できる
    <div className="relative inline-block" ref={dropdownRef}>
      <button
        onClick={() => setIsOpen(!isOpen)}
        className="bg-white border border-gray-300 rounded px-4 py-2 hover:bg-gray-50 flex items-center gap-2"
      >
        {label}
        <span className={`transition-transform ${isOpen ? 'rotate-180' : ''}`}>
          ▼
        </span>
      </button>

      {isOpen && (
        <ul className="absolute top-full left-0 mt-1 w-48 bg-white border border-gray-200 rounded-lg shadow-lg py-1 z-10">
          {items.map((item) => (
            <li key={item}>
              <button
                className="w-full text-left px-4 py-2 hover:bg-gray-100 text-gray-700"
                onClick={() => {
                  alert(item);
                  setIsOpen(false);
                }}
              >
                {item}
              </button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

// 使用例
function App() {
  return (
    <div className="p-4">
      <Dropdown
        label="メニュー"
        items={['プロフィール', '設定', 'ログアウト']}
      />
      <p className="mt-4 text-gray-500">
        ドロップダウンの外側をクリックするとメニューが閉じます。
      </p>
    </div>
  );
}
```

> **💡 次のステップへ:** 外側クリックで閉じる機能が加わり、実用的なDropdownになりました。しかし実務ではさらに2つの機能が求められます。(1) 選択した値を親コンポーネントに伝えて状態を管理する機能、(2) キーボードだけで操作できるアクセシビリティ対応です。次のステージではこれらを実装し、プロダクション品質に近づけます。

### 実践: 選択値の管理 + キーボード操作

実務のDropdownは「何が選択されているか」を親コンポーネントが管理する制御コンポーネントとして設計します。`value` と `onChange` を props で受け取る設計は、React の `<input>` と同じパターンであり、フォームライブラリとの統合やバリデーションの実装が容易になります。さらにキーボードナビゲーションを追加することで、マウスが使えないユーザーやスクリーンリーダー利用者にも操作可能なUIになります。

```tsx
import { useState, useRef, useEffect, type KeyboardEvent } from 'react';

type Props = {
  label: string;
  items: string[];
  // value + onChange の組み合わせで「制御コンポーネント」パターンを実現
  // 親が状態を持ち、子は表示と操作のみを担当する
  value: string;
  onChange: (value: string) => void;
};

function Dropdown({ label, items, value, onChange }: Props) {
  const [isOpen, setIsOpen] = useState(false);
  // キーボードで現在どの項目がハイライトされているかを追跡
  const [highlightedIndex, setHighlightedIndex] = useState(-1);
  const dropdownRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (e: MouseEvent) => {
      if (dropdownRef.current && !dropdownRef.current.contains(e.target as Node)) {
        setIsOpen(false);
      }
    };

    if (isOpen) {
      document.addEventListener('mousedown', handleClickOutside);
    }

    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, [isOpen]);

  const handleKeyDown = (e: KeyboardEvent) => {
    // メニューが閉じている状態でのキー操作: Enter/Space/ArrowDown でメニューを開く
    if (!isOpen) {
      if (e.key === 'Enter' || e.key === ' ' || e.key === 'ArrowDown') {
        e.preventDefault();
        setIsOpen(true);
        setHighlightedIndex(0);
      }
      return;
    }

    switch (e.key) {
      case 'ArrowDown':
        e.preventDefault();
        // 末尾に達したら先頭に戻る（循環ナビゲーション）
        setHighlightedIndex((prev) =>
          prev < items.length - 1 ? prev + 1 : 0
        );
        break;
      case 'ArrowUp':
        e.preventDefault();
        // 先頭に達したら末尾に戻る（循環ナビゲーション）
        setHighlightedIndex((prev) =>
          prev > 0 ? prev - 1 : items.length - 1
        );
        break;
      case 'Enter':
        e.preventDefault();
        if (highlightedIndex >= 0) {
          // onChange で親に選択値を通知（制御コンポーネントパターン）
          onChange(items[highlightedIndex]);
          setIsOpen(false);
        }
        break;
      case 'Escape':
        // Escape は「操作のキャンセル」という普遍的なUIの慣習に対応
        setIsOpen(false);
        break;
    }
  };

  // 値が未選択のときは label をプレースホルダーとして表示
  const displayLabel = value || label;

  return (
    <div className="relative inline-block" ref={dropdownRef}>
      <button
        onClick={() => setIsOpen(!isOpen)}
        onKeyDown={handleKeyDown}
        className="bg-white border border-gray-300 rounded px-4 py-2 hover:bg-gray-50 flex items-center gap-2 min-w-[160px] justify-between"
      >
        {/* 選択済みかどうかでテキスト色を変え、プレースホルダーと区別する */}
        <span className={value ? 'text-gray-900' : 'text-gray-400'}>
          {displayLabel}
        </span>
        <span className={`transition-transform ${isOpen ? 'rotate-180' : ''}`}>
          ▼
        </span>
      </button>

      {isOpen && (
        <ul className="absolute top-full left-0 mt-1 w-full bg-white border border-gray-200 rounded-lg shadow-lg py-1 z-10">
          {items.map((item, index) => (
            <li key={item}>
              <button
                className={`w-full text-left px-4 py-2 text-gray-700 ${
                  // キーボードハイライトとマウスホバーの視覚フィードバックを分離
                  index === highlightedIndex
                    ? 'bg-blue-50 text-blue-700'
                    : 'hover:bg-gray-100'
                } ${item === value ? 'font-medium' : ''}`}
                onClick={() => {
                  onChange(item);
                  setIsOpen(false);
                }}
                // マウスホバーでもハイライトを更新し、キーボードとマウスの併用に対応
                onMouseEnter={() => setHighlightedIndex(index)}
              >
                {item}
              </button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

// 使用例
function App() {
  const [selected, setSelected] = useState('');

  return (
    <div className="p-4">
      <Dropdown
        label="フルーツを選択"
        items={['りんご', 'バナナ', 'オレンジ', 'ぶどう']}
        value={selected}
        onChange={setSelected}
      />
      {selected && (
        <p className="mt-4 text-gray-600">選択中: {selected}</p>
      )}
    </div>
  );
}
```

## 🎯 演習問題

外側クリックで閉じる機能と、選択した値を表示する機能を追加してください。

**要件:**

1. `useRef`でDropdown全体のDOM参照を取得
2. `useEffect`で`mousedown`イベントを監視し、外側クリック時に閉じる
3. イベントリスナーのクリーンアップを忘れない
4. 選択した値を`value`として管理し、トリガーボタンに表示
5. `onChange`コールバックで親に選択値を通知

**ヒント:**

```tsx
const dropdownRef = useRef<HTMLDivElement>(null);

useEffect(() => {
  const handleClickOutside = (e: MouseEvent) => {
    if (dropdownRef.current && !dropdownRef.current.contains(e.target as Node)) {
      // 閉じる処理
    }
  };
  // mousedown イベントリスナーの登録と解除
}, [isOpen]);
```

## ✅ 重要ポイント

- [ ] 外側クリックの検知は`useRef` + `contains()`で実装する
- [ ] イベントリスナーはuseEffectのクリーンアップで必ず解除する
- [ ] メニューは`position: absolute`でトリガーの下に正確に配置する
- [ ] キーボード操作（ArrowDown/ArrowUp/Enter/Escape）でアクセシビリティを向上させる

**次のテーマ:** [06. Alertコンポーネント](./06-alert-component.md)
