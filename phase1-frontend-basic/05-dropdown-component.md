# 05. Dropdownコンポーネント - メニュー表示、外側クリックで閉じる

## 📖 概念

Dropdownはトリガー要素のクリックでメニューを表示し、項目を選択させるコンポーネントです。外側クリックでメニューを閉じる仕組み（`useRef` + `useEffect`）はReact開発で頻出するパターンです。キーボード操作への対応まで含めると、アクセシブルなUIの実装力が身につきます。

**よくある誤解:**
- ❌ 「stateだけで開閉管理すればいい」→ 外側クリックで閉じるには`useRef`でDOM参照が必要
- ❌ 「onBlurで外側クリックを検知できる」→ onBlurはフォーカス管理に依存し、意図しない挙動になりやすい
- ❌ 「メニューはトグルボタンの直下に置けばいい」→ `position: absolute`で正確に位置を制御する必要がある

## 💡 コード例

### 基本: トグルで開閉

```jsx
import { useState } from 'react';

function Dropdown({ label, items }) {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="relative inline-block">
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
          {items.map((item, index) => (
            <li key={index}>
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
    </div>
  );
}
```

### 応用: 外側クリックで閉じる（useRef + useEffect）

```jsx
import { useState, useRef, useEffect } from 'react';

function Dropdown({ label, items }) {
  const [isOpen, setIsOpen] = useState(false);
  const dropdownRef = useRef(null);

  // 外側クリックで閉じる
  useEffect(() => {
    const handleClickOutside = (e) => {
      if (dropdownRef.current && !dropdownRef.current.contains(e.target)) {
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

  return (
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
          {items.map((item, index) => (
            <li key={index}>
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

### 実践: 選択値の管理 + キーボード操作

```jsx
import { useState, useRef, useEffect } from 'react';

function Dropdown({ label, items, value, onChange }) {
  const [isOpen, setIsOpen] = useState(false);
  const [highlightedIndex, setHighlightedIndex] = useState(-1);
  const dropdownRef = useRef(null);

  // 外側クリックで閉じる
  useEffect(() => {
    const handleClickOutside = (e) => {
      if (dropdownRef.current && !dropdownRef.current.contains(e.target)) {
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

  // キーボード操作
  const handleKeyDown = (e) => {
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
        setHighlightedIndex((prev) =>
          prev < items.length - 1 ? prev + 1 : 0
        );
        break;
      case 'ArrowUp':
        e.preventDefault();
        setHighlightedIndex((prev) =>
          prev > 0 ? prev - 1 : items.length - 1
        );
        break;
      case 'Enter':
        e.preventDefault();
        if (highlightedIndex >= 0) {
          onChange(items[highlightedIndex]);
          setIsOpen(false);
        }
        break;
      case 'Escape':
        setIsOpen(false);
        break;
    }
  };

  const displayLabel = value || label;

  return (
    <div className="relative inline-block" ref={dropdownRef}>
      <button
        onClick={() => setIsOpen(!isOpen)}
        onKeyDown={handleKeyDown}
        className="bg-white border border-gray-300 rounded px-4 py-2 hover:bg-gray-50 flex items-center gap-2 min-w-[160px] justify-between"
      >
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
            <li key={index}>
              <button
                className={`w-full text-left px-4 py-2 text-gray-700 ${
                  index === highlightedIndex
                    ? 'bg-blue-50 text-blue-700'
                    : 'hover:bg-gray-100'
                } ${item === value ? 'font-medium' : ''}`}
                onClick={() => {
                  onChange(item);
                  setIsOpen(false);
                }}
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

### 基本: 開閉の切り替え

クリックでメニューを開閉するDropdownコンポーネントを作ってください。

```jsx
function Dropdown({ label, items }) {
  // ここにコードを書く
  // isOpen state でメニューの表示/非表示を切り替え
  // items を map でリスト表示

  return (
    <div>
      {/* トリガーボタンとメニューリストを配置 */}
    </div>
  );
}
```

**期待される動作:**
- ボタンクリックでメニューが開く
- もう一度クリックで閉じる
- メニュー項目をクリックするとメニューが閉じる

---

### 応用: 外側クリック + 選択値の管理

外側クリックで閉じる機能と、選択した値を表示する機能を追加してください。

**要件:**
1. `useRef`でDropdown全体のDOM参照を取得
2. `useEffect`で`mousedown`イベントを監視し、外側クリック時に閉じる
3. イベントリスナーのクリーンアップを忘れない
4. 選択した値を`value`として管理し、トリガーボタンに表示
5. `onChange`コールバックで親に選択値を通知

**ヒント:**
```jsx
const dropdownRef = useRef(null);

useEffect(() => {
  const handleClickOutside = (e) => {
    if (dropdownRef.current && !dropdownRef.current.contains(e.target)) {
      // 閉じる処理
    }
  };
  // mousedown イベントリスナーの登録と解除
}, [isOpen]);
```

---

### 発展: キーボードナビゲーション

Dropdownにキーボード操作を追加してください。

**要件:**
1. `ArrowDown` / `ArrowUp`でメニュー項目をハイライト移動
2. `Enter`でハイライト中の項目を選択
3. `Escape`でメニューを閉じる
4. ハイライト中の項目は背景色で強調表示
5. リストの端に達したら反対側にループする（最後→最初、最初→最後）

**完成イメージ:**
```
[フルーツを選択 ▼]
┌──────────┐
│ りんご     │  ← ↑↓で移動
│ バナナ  ◀─│  ← ハイライト表示
│ オレンジ   │
│ ぶどう     │
└──────────┘
```

## ✅ 重要ポイント

- [ ] 外側クリックの検知は`useRef` + `contains()`で実装する
- [ ] イベントリスナーはuseEffectのクリーンアップで必ず解除する
- [ ] メニューは`position: absolute`でトリガーの下に正確に配置する
- [ ] キーボード操作（ArrowDown/ArrowUp/Enter/Escape）でアクセシビリティを向上させる

**次のテーマ:** [06. Alertコンポーネント](./06-alert-component.md)
