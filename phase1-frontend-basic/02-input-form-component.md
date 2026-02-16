# 02. Input/Formコンポーネント - text/email/password、バリデーション、エラー表示

## 📖 概念

Input/Formはユーザーからデータを受け取るための基本コンポーネントです。適切なtype指定、ラベルの紐付け、バリデーション、エラーメッセージの表示を組み合わせることで、使いやすく堅牢なフォームを構築できます。

**よくある誤解:**
- ❌ 「inputにplaceholderがあればラベルは不要」→ ラベルはアクセシビリティに必須（スクリーンリーダー対応）
- ❌ 「バリデーションはsubmit時だけでいい」→ 入力中のリアルタイムフィードバックがUXを向上させる
- ❌ 「type="text"で全部まかなえる」→ type="email"やtype="password"はブラウザの補助機能を活用できる

## 💡 コード例

### 基本: テキスト入力

```tsx
import { useState, type ChangeEvent } from 'react';

type Props = {
  label: string;
  value: string;
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
  placeholder?: string;
};

function Input({ label, value, onChange, placeholder }: Props) {
  return (
    <div className="flex flex-col gap-1">
      <label className="text-sm font-medium text-gray-700">
        {label}
      </label>
      <input
        type="text"
        value={value}
        onChange={onChange}
        placeholder={placeholder}
        className="border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
      />
    </div>
  );
}

// 使用例
function App() {
  const [name, setName] = useState('');

  return (
    <div className="p-4 max-w-md">
      <Input
        label="名前"
        value={name}
        onChange={(e) => setName(e.target.value)}
        placeholder="山田太郎"
      />
      <p className="mt-2 text-sm text-gray-500">入力値: {name}</p>
    </div>
  );
}
```

### 応用: type切替 + ラベル紐付け

```tsx
import { useState, type ChangeEvent } from 'react';

type Props = {
  label: string;
  type?: string;
  value: string;
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
  placeholder?: string;
  id?: string;
};

function Input({ label, type = 'text', value, onChange, placeholder, id }: Props) {
  return (
    <div className="flex flex-col gap-1">
      <label
        htmlFor={id}
        className="text-sm font-medium text-gray-700"
      >
        {label}
      </label>
      <input
        id={id}
        type={type}
        value={value}
        onChange={onChange}
        placeholder={placeholder}
        className="border border-gray-300 rounded px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
      />
    </div>
  );
}

// 使用例
function App() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  return (
    <div className="p-4 max-w-md space-y-4">
      <Input
        id="email"
        label="メールアドレス"
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        placeholder="example@mail.com"
      />
      <Input
        id="password"
        label="パスワード"
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        placeholder="8文字以上"
      />
    </div>
  );
}
```

### 実践: バリデーション + エラーメッセージ表示

```tsx
import { useState, type ChangeEvent, type FormEvent } from 'react';

type Props = {
  label: string;
  type?: string;
  value: string;
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
  placeholder?: string;
  id?: string;
  error?: string;
};

function Input({ label, type = 'text', value, onChange, placeholder, id, error }: Props) {
  return (
    <div className="flex flex-col gap-1">
      <label
        htmlFor={id}
        className="text-sm font-medium text-gray-700"
      >
        {label}
      </label>
      <input
        id={id}
        type={type}
        value={value}
        onChange={onChange}
        placeholder={placeholder}
        className={`border rounded px-3 py-2 focus:outline-none focus:ring-2 focus:border-transparent ${
          error
            ? 'border-red-500 focus:ring-red-500'
            : 'border-gray-300 focus:ring-blue-500'
        }`}
      />
      {error && (
        <p className="text-sm text-red-500">{error}</p>
      )}
    </div>
  );
}

// 使用例
function App() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [errors, setErrors] = useState<Record<string, string>>({});

  const validate = () => {
    const newErrors: Record<string, string> = {};

    if (!email) {
      newErrors.email = 'メールアドレスを入力してください';
    } else if (!/\S+@\S+\.\S+/.test(email)) {
      newErrors.email = '正しいメールアドレスを入力してください';
    }

    if (!password) {
      newErrors.password = 'パスワードを入力してください';
    } else if (password.length < 8) {
      newErrors.password = 'パスワードは8文字以上で入力してください';
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    if (validate()) {
      alert('送信成功！');
    }
  };

  return (
    <form onSubmit={handleSubmit} className="p-4 max-w-md space-y-4">
      <Input
        id="email"
        label="メールアドレス"
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        placeholder="example@mail.com"
        error={errors.email}
      />
      <Input
        id="password"
        label="パスワード"
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        placeholder="8文字以上"
        error={errors.password}
      />
      <button
        type="submit"
        className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700"
      >
        ログイン
      </button>
    </form>
  );
}
```

## 🎯 演習問題

### 基本: Inputコンポーネントの実装

ラベル付きのInputコンポーネントを作ってください。

```tsx
import type { ChangeEvent } from 'react';

type Props = {
  label: string;
  value: string;
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
  placeholder?: string;
};

function Input({ label, value, onChange, placeholder }: Props) {
  // ここにコードを書く
  // label と input を組み合わせる

  return (
    <div>
      {/* ラベルとinputを配置 */}
    </div>
  );
}
```

**期待される動作:**
- ラベルがinputの上に表示される
- テキストを入力すると親コンポーネントのstateが更新される
- focusすると青い枠線が表示される

---

### 応用: 複数type + エラー表示

type（`text`, `email`, `password`）を切り替えられ、エラーメッセージを表示できるInputコンポーネントを作ってください。

**要件:**
1. `type`, `label`, `value`, `onChange`, `error`, `id` をpropsで受け取る
2. デフォルト値: type=`text`
3. `htmlFor`と`id`でラベルとinputを紐付ける
4. `error`がある場合、ボーダーを赤くし、下にエラーメッセージを表示

**ヒント:**
```tsx
import type { ChangeEvent } from 'react';

type Props = {
  label: string;
  type?: string;
  value: string;
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
  error?: string;
  id?: string;
};

function Input({ label, type = 'text', value, onChange, error, id }: Props) {
  // error の有無でボーダーの色を切り替える
  // error がある場合は <p> でメッセージを表示
}
```

---

### 発展: フォーム全体 + 送信処理

Inputコンポーネントを使って、ログインフォーム（メールアドレス + パスワード）を作成してください。

**要件:**
1. メールアドレスとパスワードの2つのInputを配置
2. バリデーション: メールアドレスは空チェック+形式チェック、パスワードは空チェック+8文字以上
3. 送信ボタンクリック時にバリデーションを実行し、エラーがあればエラーメッセージを表示
4. バリデーション通過時は`alert('送信成功！')`を実行
5. `e.preventDefault()`でページリロードを防止

**完成イメージ:**
```
メールアドレス
[                    ]
正しいメールアドレスを入力してください  ← エラー時のみ表示

パスワード
[                    ]
パスワードは8文字以上で入力してください  ← エラー時のみ表示

[ログイン]
```

## ✅ 重要ポイント

- [ ] `htmlFor`と`id`でラベルとinputを紐付け、アクセシビリティを確保する
- [ ] inputのtypeを適切に使い分ける（text/email/password）
- [ ] エラー状態はスタイル（赤枠）とメッセージの両方で伝える
- [ ] フォーム送信時は`e.preventDefault()`でデフォルト動作を防ぐ

**次のテーマ:** [03. Cardコンポーネント](./03-card-component.md)
