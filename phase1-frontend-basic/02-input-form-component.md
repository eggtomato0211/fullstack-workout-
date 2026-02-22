# 02. Input/Formコンポーネント - text/email/password、バリデーション、エラー表示

## 🎯 このテーマで学ぶこと

- 制御コンポーネント（Controlled Component）パターン
- `htmlFor` / `id` によるラベルとinputの紐付け
- バリデーションとエラーメッセージの表示

## 📖 なぜ制御コンポーネントパターンを理解する必要があるのか

ほぼ全てのWebアプリケーションにフォームがあります。ログイン、会員登録、検索、設定変更 --- ユーザーからのデータ入力はWebの根幹です。

### こう書かないとどうなるか

Reactを使わずにHTMLのフォームをそのまま使うと、inputの値はDOMが管理します：

```tsx
// 非制御パターン：値はDOM内にあり、Reactは関知しない
<input type="text" />

// 値を取得するには毎回DOMを参照する必要がある
// → リアルタイムバリデーションや条件付きUI表示が難しい
```

制御コンポーネントでは、inputの値をReactのstateで管理します。これにより「今inputに何が入っているか」を常にReactが把握でき、リアルタイムバリデーション、条件付きUI表示、送信前チェックなど柔軟な制御が可能になります。

### ラベルとinputの紐付けが必要な理由

`placeholder`があればラベルは不要に見えますが、スクリーンリーダーはplaceholderを読み上げてくれないことがあります。`htmlFor`と`id`で紐付けることで、ラベルクリックでinputにフォーカスが移る操作性も得られます。

### input typeを正しく使う理由

`type="text"`で全部まかなえますが、`type="email"`にするとモバイルで@付きキーボードが表示され、`type="password"`にするとブラウザのパスワード管理機能が連携します。正しいtypeはブラウザの補助機能を引き出します。

## 💡 コード例

### 基本: type切替 + ラベル紐付きInput

制御コンポーネントの基本パターンと、アクセシビリティに必要なラベル紐付けを合わせて学びます。

```tsx
import { useState, type ChangeEvent } from 'react';

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
      {/* htmlForとidで紐付け → ラベルクリックでinputにフォーカスが移る */}
      <label htmlFor={id} className="text-sm font-medium text-gray-700">
        {label}
      </label>
      <input
        id={id}
        type={type}
        value={value}       // Reactのstateと同期（制御コンポーネント）
        onChange={onChange}  // ユーザーの入力をstateに反映
        placeholder={placeholder}
        // errorの有無でボーダー色を切り替え、視覚的にエラーを伝える
        className={`border rounded px-3 py-2 focus:outline-none focus:ring-2 focus:border-transparent ${
          error ? 'border-red-500 focus:ring-red-500' : 'border-gray-300 focus:ring-blue-500'
        }`}
      />
      {/* エラーがある場合のみメッセージを表示 */}
      {error && <p className="text-sm text-red-500">{error}</p>}
    </div>
  );
}

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

### 実践: バリデーション付きログインフォーム

バリデーションロジックを送信ハンドラ内で一括実行し、エラー情報を`Record`オブジェクトで管理します。フィールドが増えても同じパターンで拡張でき、`e.preventDefault()`でブラウザのデフォルト送信（ページリロード）を防ぎます。

```tsx
import { useState, type ChangeEvent, type FormEvent } from 'react';

type InputProps = {
  label: string;
  type?: string;
  value: string;
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
  placeholder?: string;
  id?: string;
  error?: string;
};

function Input({ label, type = 'text', value, onChange, placeholder, id, error }: InputProps) {
  return (
    <div className="flex flex-col gap-1">
      <label htmlFor={id} className="text-sm font-medium text-gray-700">{label}</label>
      <input
        id={id} type={type} value={value} onChange={onChange} placeholder={placeholder}
        className={`border rounded px-3 py-2 focus:outline-none focus:ring-2 focus:border-transparent ${
          error ? 'border-red-500 focus:ring-red-500' : 'border-gray-300 focus:ring-blue-500'
        }`}
      />
      {error && <p className="text-sm text-red-500">{error}</p>}
    </div>
  );
}

function App() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  // Record<string, string>でフィールド名をキーにエラーを管理
  const [errors, setErrors] = useState<Record<string, string>>({});

  const validate = () => {
    const newErrors: Record<string, string> = {};

    // 空チェック → 形式チェックの順で、具体的なメッセージを出し分ける
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
    // ブラウザのデフォルト送信（ページリロード）を防ぎ、JS側で制御
    e.preventDefault();
    if (validate()) {
      alert('送信成功！');
    }
  };

  return (
    <form onSubmit={handleSubmit} className="p-4 max-w-md space-y-4">
      <Input id="email" label="メールアドレス" type="email" value={email}
        onChange={(e) => setEmail(e.target.value)} placeholder="example@mail.com" error={errors.email} />
      <Input id="password" label="パスワード" type="password" value={password}
        onChange={(e) => setPassword(e.target.value)} placeholder="8文字以上" error={errors.password} />
      <button type="submit" className="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700">
        ログイン
      </button>
    </form>
  );
}
```

## 🎯 演習問題

type切替とエラー表示ができるInputコンポーネントを作ってください。

**要件:**

1. `type`, `label`, `value`, `onChange`, `error`, `id` をpropsで受け取る
2. `htmlFor`と`id`でラベルとinputを紐付ける
3. `error`がある場合、ボーダーを赤くし、下にエラーメッセージを表示

**ヒント:**

```tsx
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

## ✅ 重要ポイント

- [ ] `htmlFor`と`id`でラベルとinputを紐付け、アクセシビリティを確保する
- [ ] inputのtypeを適切に使い分ける（text/email/password）
- [ ] エラー状態はスタイル（赤枠）とメッセージの両方で伝える
- [ ] フォーム送信時は`e.preventDefault()`でデフォルト動作を防ぐ

**次のテーマ:** [03. Cardコンポーネント](./03-card-component.md)
