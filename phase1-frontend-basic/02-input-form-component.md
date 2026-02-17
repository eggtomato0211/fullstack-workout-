# 02. Input/Formコンポーネント - text/email/password、バリデーション、エラー表示

## 🎯 このテーマで学ぶこと

- 制御コンポーネント（Controlled Component）パターンの理解
- `htmlFor` / `id` によるラベルとinputの紐付け（アクセシビリティ）
- input typeの使い分け（text / email / password）
- バリデーションロジックの設計とエラー表示
- フォーム送信時の `e.preventDefault()` によるデフォルト動作の制御

**なぜ重要か:** ほぼ全てのWebアプリケーションにフォームがあります。ログイン、会員登録、検索、設定変更、問い合わせ — ユーザーからのデータ入力はWebの根幹です。正しいバリデーションとエラー表示ができることは、フロントエンド開発者にとって必須スキルです。

## 📖 概念

Input/Formはユーザーからデータを受け取るための基本コンポーネントです。適切なtype指定、ラベルの紐付け、バリデーション、エラーメッセージの表示を組み合わせることで、使いやすく堅牢なフォームを構築できます。

**背景と設計意図:** Reactでは「制御コンポーネント」パターンが標準です。inputの値をstateで管理し、`onChange`で更新することで、Reactが常にフォームの状態を把握します。これにより、リアルタイムバリデーション、条件付きUI表示、送信前の値チェックなど、柔軟な制御が可能になります。HTMLのネイティブフォームバリデーションだけでは実現できない細かいUXを構築できるのがこのパターンの利点です。

**実務での活用場面:** ログインフォーム、会員登録フォーム、お問い合わせフォーム、管理画面のデータ入力フォームなど。特にECサイトの注文フォーム（住所、クレジットカード）では、リアルタイムバリデーションとわかりやすいエラー表示がコンバージョン率に直結します。

**よくある誤解:**

- ❌ 「inputにplaceholderがあればラベルは不要」→ ラベルはアクセシビリティに必須（スクリーンリーダー対応）
- ❌ 「バリデーションはsubmit時だけでいい」→ 入力中のリアルタイムフィードバックがUXを向上させる
- ❌ 「type="text"で全部まかなえる」→ type="email"やtype="password"はブラウザの補助機能を活用できる

## 💡 コード例

### 基本: テキスト入力

まず最もシンプルな形から始めます。Reactの「制御コンポーネント」パターンでは、inputの値をstateとして管理し、`onChange`イベントで同期します。こうすることで、Reactが常にフォームの現在の値を把握でき、表示や処理に自由に使えるようになります。ここではlabel、value、onChange、placeholderという最低限のpropsだけを受け取るInputコンポーネントを定義します。

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
      {/* ラベルを表示するが、この段階ではまだinputとの紐付けはしていない */}
      <label className="text-sm font-medium text-gray-700">
        {label}
      </label>
      {/* value と onChange で React が入力値を管理する（制御コンポーネント） */}
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
  // 親コンポーネントが状態を持ち、Inputに渡す（単一データソースの原則） */
  const [name, setName] = useState('');

  return (
    <div className="p-4 max-w-md">
      <Input
        label="名前"
        value={name}
        onChange={(e) => setName(e.target.value)}
        placeholder="山田太郎"
      />
      {/* stateをそのまま表示に使える — これが制御コンポーネントの利点 */}
      <p className="mt-2 text-sm text-gray-500">入力値: {name}</p>
    </div>
  );
}
```

> **💡 次のステップへ:** 基本ではtype固定・ラベル紐付けなしでした。次の応用では、`type` propで入力種別を切り替えられるようにし、`htmlFor`/`id`によるラベルとinputの正式な紐付けを追加します。これによりアクセシビリティが向上し、ラベルクリックでinputにフォーカスが移るようになります。

### 応用: type切替 + ラベル紐付け

基本ではtypeが`text`固定でしたが、実際のフォームではemail、passwordなど複数のtypeが必要です。typeをpropとして外部から受け取ることで、1つのInputコンポーネントをさまざまな入力欄に再利用できます。また、`htmlFor`と`id`を使ってlabelとinputをHTML仕様に沿って正しく紐付けることで、スクリーンリーダーがラベルを読み上げられるようになり、ラベルクリックでinputにフォーカスが移る動作も実現します。

```tsx
import { useState, type ChangeEvent } from 'react';

type Props = {
  label: string;
  type?: string;
  value: string;
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
  placeholder?: string;
  id?: string; // htmlFor/idによるラベル紐付けのために追加
};

function Input({ label, type = 'text', value, onChange, placeholder, id }: Props) {
  return (
    <div className="flex flex-col gap-1">
      {/* htmlFor で id を指定すると、ラベルクリックで対応するinputにフォーカスが移る */}
      <label
        htmlFor={id}
        className="text-sm font-medium text-gray-700"
      >
        {label}
      </label>
      {/* type を prop 化することで、同じコンポーネントを text/email/password 等に再利用できる */}
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
      {/* type="email" を指定すると、モバイルで@付きキーボードが表示されるなどブラウザ補助が効く */}
      <Input
        id="email"
        label="メールアドレス"
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        placeholder="example@mail.com"
      />
      {/* type="password" で入力値がマスクされ、ブラウザのパスワード管理機能とも連携する */}
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

> **💡 次のステップへ:** type切替とラベル紐付けができましたが、まだユーザーの入力ミスに対するフィードバックがありません。次の実践では、`error` propを追加してバリデーション結果を視覚的に伝え、`<form>`タグでの送信処理と`e.preventDefault()`によるデフォルト動作の制御を組み合わせます。

### 実践: バリデーション + エラーメッセージ表示

最後に、実際のログインフォームのようにバリデーションとエラー表示を組み込みます。`error` propの有無でボーダーの色を動的に切り替え、エラーメッセージを条件付きレンダリングすることで、ユーザーにどこを修正すべきか明確に伝えます。バリデーションロジックは送信ハンドラ内で一括実行し、エラー情報をRecordオブジェクトで管理することで、フィールドが増えても同じパターンで拡張できます。

```tsx
import { useState, type ChangeEvent, type FormEvent } from 'react';

type Props = {
  label: string;
  type?: string;
  value: string;
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
  placeholder?: string;
  id?: string;
  error?: string; // バリデーションエラーメッセージ（undefinedなら正常状態）
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
        // error の有無でボーダー色とフォーカスリングの色を切り替え、視覚的にエラー状態を伝える
        className={`border rounded px-3 py-2 focus:outline-none focus:ring-2 focus:border-transparent ${
          error
            ? 'border-red-500 focus:ring-red-500'
            : 'border-gray-300 focus:ring-blue-500'
        }`}
      />
      {/* エラーがある場合のみメッセージを表示（条件付きレンダリング） */}
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
  // Record<string, string> でフィールド名をキーにエラーメッセージを管理する
  const [errors, setErrors] = useState<Record<string, string>>({});

  const validate = () => {
    const newErrors: Record<string, string> = {};

    // 空チェック → 形式チェック の順で、より具体的なエラーメッセージを出し分ける
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
    // エラーが0件ならバリデーション通過
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = (e: FormEvent) => {
    // ブラウザのデフォルト送信（ページリロード）を防ぎ、JS側で処理を制御する
    e.preventDefault();
    if (validate()) {
      alert('送信成功！');
    }
  };

  return (
    // form タグ + onSubmit でEnterキー送信にも対応できる
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
      {/* type="submit" でフォームの onSubmit と連動する */}
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

## ✅ 重要ポイント

- [ ] `htmlFor`と`id`でラベルとinputを紐付け、アクセシビリティを確保する
- [ ] inputのtypeを適切に使い分ける（text/email/password）
- [ ] エラー状態はスタイル（赤枠）とメッセージの両方で伝える
- [ ] フォーム送信時は`e.preventDefault()`でデフォルト動作を防ぐ

**次のテーマ:** [03. Cardコンポーネント](./03-card-component.md)
