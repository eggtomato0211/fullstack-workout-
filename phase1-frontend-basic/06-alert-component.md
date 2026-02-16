# 06. Alertコンポーネント - success/warning/error/infoの表示

## 🎯 このテーマで学ぶこと

- variant別の色・アイコンによる視覚的な状態区別
- `role="alert"` によるアクセシビリティ対応
- `useEffect` + `setTimeout` による自動消去とクリーンアップ
- 配列stateの追加・削除操作（リスト管理）
- 色だけに頼らない情報伝達（アイコン + テキストの併用）

**なぜ重要か:** ユーザーへのフィードバックはUXの根幹です。操作が成功したのか失敗したのか、すぐに視覚的に伝えられないアプリは使いづらいと感じられます。また、stateで通知リストを管理し、追加・削除を行うパターンは、ToDoリストやチャットなどリスト操作全般の基礎になります。

## 📖 概念

Alertはユーザーに状態や結果を伝えるための通知コンポーネントです。success（成功）、warning（警告）、error（エラー）、info（情報）の4種類を色とアイコンで視覚的に区別します。閉じるボタンや自動消去の機能を加えることで、邪魔にならない通知体験を実現できます。

**背景と設計意図:** Webアプリケーションでは、バックエンドとの通信結果（成功・エラー）やユーザーの入力ミスなど、様々な場面でフィードバックが必要です。それぞれの状態を色・アイコン・テキストの3つの手段で伝えることで、色覚に課題のあるユーザーにも情報が正しく届きます（WCAG準拠のアクセシビリティ）。自動消去の機能は、成功通知のように「確認だけでOK」な場合にユーザーの操作負担を減らします。

**実務での活用場面:** フォーム送信後の成功/エラー通知、システム障害の警告バナー、新機能のお知らせ、入力バリデーションの結果表示など。Toast通知（画面端に表示される一時通知）はAlertの発展形で、SaaS系アプリではほぼ必須の機能です。

**よくある誤解:**
- ❌ 「色だけ変えれば伝わる」→ アイコンとテキストも併用し、色覚に頼らない情報伝達が重要
- ❌ 「エラーも成功も同じ見た目でいい」→ variant別の色・アイコンで直感的に状態を区別させる
- ❌ 「表示したら消す必要はない」→ 成功通知などは自動消去するとUXが向上する

## 💡 コード例

### 基本: 4種類のAlert

```tsx
import type { ReactNode } from 'react';

type Variant = 'success' | 'warning' | 'error' | 'info';

type Props = {
  variant?: Variant;
  children: ReactNode;
};

function Alert({ variant = 'info', children }: Props) {
  const variantClasses: Record<Variant, string> = {
    success: 'bg-green-50 border-green-500 text-green-800',
    warning: 'bg-yellow-50 border-yellow-500 text-yellow-800',
    error: 'bg-red-50 border-red-500 text-red-800',
    info: 'bg-blue-50 border-blue-500 text-blue-800',
  };

  return (
    <div
      className={`border-l-4 px-4 py-3 rounded ${variantClasses[variant]}`}
      role="alert"
    >
      {children}
    </div>
  );
}

// 使用例
function App() {
  return (
    <div className="p-4 max-w-md space-y-4">
      <Alert variant="success">保存が完了しました。</Alert>
      <Alert variant="warning">入力内容を確認してください。</Alert>
      <Alert variant="error">エラーが発生しました。</Alert>
      <Alert variant="info">新しいバージョンが利用可能です。</Alert>
    </div>
  );
}
```

### 応用: アイコン付きAlert

```tsx
import type { ReactNode } from 'react';

type Variant = 'success' | 'warning' | 'error' | 'info';

type Props = {
  variant?: Variant;
  children: ReactNode;
};

function Alert({ variant = 'info', children }: Props) {
  const variantClasses: Record<Variant, string> = {
    success: 'bg-green-50 border-green-500 text-green-800',
    warning: 'bg-yellow-50 border-yellow-500 text-yellow-800',
    error: 'bg-red-50 border-red-500 text-red-800',
    info: 'bg-blue-50 border-blue-500 text-blue-800',
  };

  const icons: Record<Variant, string> = {
    success: '✓',
    warning: '⚠',
    error: '✕',
    info: 'ℹ',
  };

  const iconClasses: Record<Variant, string> = {
    success: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500',
    info: 'bg-blue-500',
  };

  return (
    <div
      className={`border-l-4 px-4 py-3 rounded flex items-center gap-3 ${variantClasses[variant]}`}
      role="alert"
    >
      <span
        className={`flex-shrink-0 w-6 h-6 rounded-full flex items-center justify-center text-white text-sm ${iconClasses[variant]}`}
      >
        {icons[variant]}
      </span>
      <span>{children}</span>
    </div>
  );
}

// 使用例
function App() {
  return (
    <div className="p-4 max-w-md space-y-4">
      <Alert variant="success">保存が完了しました。</Alert>
      <Alert variant="warning">入力内容を確認してください。</Alert>
      <Alert variant="error">エラーが発生しました。</Alert>
      <Alert variant="info">新しいバージョンが利用可能です。</Alert>
    </div>
  );
}
```

### 実践: 閉じるボタン + 自動消去

```tsx
import { useState, useEffect, type ReactNode } from 'react';

type Variant = 'success' | 'warning' | 'error' | 'info';

type Props = {
  variant?: Variant;
  children: ReactNode;
  onClose?: () => void;
  autoClose?: number;
};

function Alert({ variant = 'info', children, onClose, autoClose = 0 }: Props) {
  const variantClasses: Record<Variant, string> = {
    success: 'bg-green-50 border-green-500 text-green-800',
    warning: 'bg-yellow-50 border-yellow-500 text-yellow-800',
    error: 'bg-red-50 border-red-500 text-red-800',
    info: 'bg-blue-50 border-blue-500 text-blue-800',
  };

  const icons: Record<Variant, string> = {
    success: '✓',
    warning: '⚠',
    error: '✕',
    info: 'ℹ',
  };

  const iconClasses: Record<Variant, string> = {
    success: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500',
    info: 'bg-blue-500',
  };

  // 自動消去
  useEffect(() => {
    if (autoClose > 0 && onClose) {
      const timer = setTimeout(onClose, autoClose);
      return () => clearTimeout(timer);
    }
  }, [autoClose, onClose]);

  return (
    <div
      className={`border-l-4 px-4 py-3 rounded flex items-center gap-3 ${variantClasses[variant]}`}
      role="alert"
    >
      <span
        className={`flex-shrink-0 w-6 h-6 rounded-full flex items-center justify-center text-white text-sm ${iconClasses[variant]}`}
      >
        {icons[variant]}
      </span>
      <span className="flex-1">{children}</span>
      {onClose && (
        <button
          onClick={onClose}
          className="flex-shrink-0 text-current opacity-50 hover:opacity-100"
        >
          ✕
        </button>
      )}
    </div>
  );
}

// 使用例
type AlertItem = {
  id: number;
  variant: Variant;
  message: string;
  autoClose: number;
};

function App() {
  const [alerts, setAlerts] = useState<AlertItem[]>([
    { id: 1, variant: 'success', message: '保存が完了しました。', autoClose: 3000 },
    { id: 2, variant: 'error', message: 'エラーが発生しました。', autoClose: 0 },
  ]);

  const removeAlert = (id: number) => {
    setAlerts((prev) => prev.filter((alert) => alert.id !== id));
  };

  const addAlert = () => {
    const newAlert: AlertItem = {
      id: Date.now(),
      variant: 'info',
      message: `新しい通知です（${new Date().toLocaleTimeString()}）`,
      autoClose: 5000,
    };
    setAlerts((prev) => [...prev, newAlert]);
  };

  return (
    <div className="p-4 max-w-md">
      <button
        onClick={addAlert}
        className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 mb-4"
      >
        通知を追加
      </button>

      <div className="space-y-3">
        {alerts.map((alert) => (
          <Alert
            key={alert.id}
            variant={alert.variant}
            onClose={() => removeAlert(alert.id)}
            autoClose={alert.autoClose}
          >
            {alert.message}
          </Alert>
        ))}
      </div>
    </div>
  );
}
```

## 🎯 演習問題

### 基本: 4種類のvariant

4種類のvariant（`success`, `warning`, `error`, `info`）に対応するAlertコンポーネントを作ってください。

```tsx
import type { ReactNode } from 'react';

type Variant = 'success' | 'warning' | 'error' | 'info';

type Props = {
  variant?: Variant;
  children: ReactNode;
};

function Alert({ variant = 'info', children }: Props) {
  // ここにコードを書く
  // variant に応じて背景色・ボーダー色・文字色を切り替える

  return (
    <div role="alert">
      {children}
    </div>
  );
}
```

**期待される動作:**
- `success`: 緑系の配色
- `warning`: 黄色系の配色
- `error`: 赤系の配色
- `info`: 青系の配色
- 左側にボーダー（`border-l-4`）が表示される

---

### 応用: アイコン + 閉じるボタン

基本のAlertにアイコンと閉じるボタンを追加してください。

**要件:**
1. 各variantに対応するアイコンを左側に表示（✓ / ⚠ / ✕ / ℹ）
2. アイコンは丸い背景付き（variant色）で白文字
3. `onClose`がpropsで渡された場合のみ、右端に閉じるボタン（✕）を表示
4. 閉じるボタンクリックで`onClose`コールバックを呼ぶ

**ヒント:**
```tsx
import type { ReactNode } from 'react';

type Variant = 'success' | 'warning' | 'error' | 'info';

type Props = {
  variant?: Variant;
  children: ReactNode;
  onClose?: () => void;
};

function Alert({ variant = 'info', children, onClose }: Props) {
  // icons オブジェクトで variant ごとのアイコンを定義
  // onClose が存在する場合のみ閉じるボタンを表示
}
```

---

### 発展: 自動消去 + Toast風通知

Alertを活用した通知システムを作成してください。

**要件:**
1. 「通知を追加」ボタンで新しいAlertを画面に追加
2. 各Alertは`autoClose`（ミリ秒）を指定でき、時間経過で自動消去する
3. 手動で閉じるボタンでも消去可能
4. 複数のAlertを同時に表示（縦に積む）
5. Alertの追加時にvariantをランダムに選択する

**完成イメージ:**
```
[通知を追加]

┌ ✓ 保存が完了しました。          ✕ ┐  ← 3秒後に自動消去
├ ⚠ 入力内容を確認してください。  ✕ ┤
└ ℹ 新しい通知です (12:34:56)    ✕ ┘  ← 5秒後に自動消去
```

## ✅ 重要ポイント

- [ ] variantごとに色・アイコンをオブジェクトで管理して一貫性を保つ
- [ ] `role="alert"`を付与してスクリーンリーダーに通知の意味を伝える
- [ ] 自動消去には`useEffect` + `setTimeout`を使い、クリーンアップで`clearTimeout`する
- [ ] 色だけでなくアイコンやテキストでも状態を伝え、アクセシビリティを確保する

**次のテーマ:** [07. Loading系コンポーネント](./07-loading-component.md)
