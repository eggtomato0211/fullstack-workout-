# 05. Dropdownコンポーネント - メニュー表示、外側クリックで閉じる

## 🎯 このテーマで学ぶこと

- `useRef` + `useEffect` + `contains()` による外側クリック検知
- `position: absolute` によるメニューの位置制御
- キーボードナビゲーション（ArrowDown / ArrowUp / Enter / Escape）

## 📖 なぜDropdownを自作する必要があるのか

HTMLには `<select>` 要素がありますが、スタイリングの自由度が非常に低いです。ブラウザごとに見た目が違い、デザイナーの意図通りのUIを実現できません。そのため実務ではほぼ必ずカスタムDropdownを実装します。

### こう書かないとどうなるか

「stateの真偽値だけでメニューを開閉すればいい」と思うかもしれませんが、大きな問題があります：

```tsx
// stateだけの場合、メニューの「外側」をクリックしても閉じない
const [isOpen, setIsOpen] = useState(false);

// ユーザーは「どこか別の場所をクリックすれば閉じる」と期待する
// → 閉じないと「壊れている」と感じる
```

外側クリックで閉じるには、Reactのイベントハンドラだけでは不十分です。外側のクリックはDropdownコンポーネントの外で発生するため、`document`レベルのイベントリスナーと`useRef`によるDOM参照を組み合わせる必要があります。

### このパターンが汎用的な理由

「外側クリックで閉じる」はDropdownだけでなく、ツールチップ、ポップオーバー、コンテキストメニューなど多くのUI要素で必要になります。ここで学ぶ `useRef` + `useEffect` + `contains()` の組み合わせは、React開発で最も頻繁に使う「DOM連携パターン」の1つです。

## 💡 コード例

### 基本: 外側クリックで閉じるDropdown

まずは「外側クリックで閉じる」という核心部分に集中します。`useRef`でDropdown全体のDOM要素を参照し、`contains()`で「クリックされた場所がDropdown内かどうか」を判定します。

```tsx
import { useState, useRef, useEffect } from 'react';

type Props = {
  label: string;
  items: string[];
};

function Dropdown({ label, items }: Props) {
  const [isOpen, setIsOpen] = useState(false);
  // useRefでDropdown全体のDOM要素を参照する
  // なぜuseRefが必要か → 「どこがクリックされたか」はDOM上の位置情報なので、
  // ReactのStateだけでは判定できない
  const dropdownRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (e: MouseEvent) => {
      // contains()で「クリック位置がDropdown内か」を判定
      // Dropdown外ならメニューを閉じる
      if (dropdownRef.current && !dropdownRef.current.contains(e.target as Node)) {
        setIsOpen(false);
      }
    };

    // メニューが開いているときだけリスナーを登録
    // なぜ常時ではなく条件付きか → 閉じているときの不要なイベント処理を避ける
    if (isOpen) {
      document.addEventListener('mousedown', handleClickOutside);
    }

    // クリーンアップ: これを忘れるとリスナーが積み重なりメモリリークする
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, [isOpen]);

  return (
    // refを最外層に付けることで、ボタンもメニューも含めた領域全体を参照
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

function App() {
  return (
    <div className="p-4">
      <Dropdown label="メニュー" items={['プロフィール', '設定', 'ログアウト']} />
    </div>
  );
}
```

### 実践: 選択値の管理 + キーボード操作

実務のDropdownは `value` + `onChange` で親コンポーネントが値を管理する「制御コンポーネント」として設計します。これはReactの `<input>` と同じパターンで、フォームライブラリとの統合が容易になります。

さらにキーボードナビゲーションを追加する理由は、マウスが使えないユーザーやスクリーンリーダー利用者への対応です。ネイティブの`<select>`が自動で提供してくれる操作性を、自作する場合は自分で実装する必要があります。

```tsx
import { useState, useRef, useEffect, type KeyboardEvent } from 'react';

type Props = {
  label: string;
  items: string[];
  value: string;
  onChange: (value: string) => void;
};

function Dropdown({ label, items, value, onChange }: Props) {
  const [isOpen, setIsOpen] = useState(false);
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
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, [isOpen]);

  const handleKeyDown = (e: KeyboardEvent) => {
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
        // 末尾→先頭に戻る循環ナビゲーション
        setHighlightedIndex((prev) => (prev < items.length - 1 ? prev + 1 : 0));
        break;
      case 'ArrowUp':
        e.preventDefault();
        setHighlightedIndex((prev) => (prev > 0 ? prev - 1 : items.length - 1));
        break;
      case 'Enter':
        e.preventDefault();
        if (highlightedIndex >= 0) {
          onChange(items[highlightedIndex]);
          setIsOpen(false);
        }
        break;
      case 'Escape':
        // Escapeは「操作のキャンセル」という普遍的なUI慣習
        setIsOpen(false);
        break;
    }
  };

  return (
    <div className="relative inline-block" ref={dropdownRef}>
      <button
        onClick={() => setIsOpen(!isOpen)}
        onKeyDown={handleKeyDown}
        className="bg-white border border-gray-300 rounded px-4 py-2 hover:bg-gray-50 flex items-center gap-2 min-w-[160px] justify-between"
      >
        <span className={value ? 'text-gray-900' : 'text-gray-400'}>
          {value || label}
        </span>
        <span className={`transition-transform ${isOpen ? 'rotate-180' : ''}`}>▼</span>
      </button>

      {isOpen && (
        <ul className="absolute top-full left-0 mt-1 w-full bg-white border border-gray-200 rounded-lg shadow-lg py-1 z-10">
          {items.map((item, index) => (
            <li key={item}>
              <button
                className={`w-full text-left px-4 py-2 text-gray-700 ${
                  index === highlightedIndex ? 'bg-blue-50 text-blue-700' : 'hover:bg-gray-100'
                } ${item === value ? 'font-medium' : ''}`}
                onClick={() => { onChange(item); setIsOpen(false); }}
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
      {selected && <p className="mt-4 text-gray-600">選択中: {selected}</p>}
    </div>
  );
}
```

## 🎯 演習問題

外側クリックで閉じる機能と、選択した値を表示する機能を持つDropdownを作ってください。

**要件:**

1. `useRef`でDropdown全体のDOM参照を取得
2. `useEffect`で`mousedown`イベントを監視し、外側クリック時に閉じる
3. イベントリスナーのクリーンアップを忘れない
4. 選択した値をトリガーボタンに表示
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
- [ ] メニューは`position: absolute`でトリガーの下に配置する
- [ ] キーボード操作（ArrowDown/ArrowUp/Enter/Escape）でアクセシビリティを確保する

**次のテーマ:** [06. Alertコンポーネント](./06-alert-component.md)
