# 06. Alertコンポーネント - success/warning/error/infoの表示

## 🎯 このテーマで学ぶこと

- variant別の色・アイコンによる状態の視覚的区別
- `role="alert"` によるアクセシビリティ対応
- `useEffect` + `setTimeout` による自動消去

## 📖 なぜAlertコンポーネントを理解する必要があるのか

ユーザーが操作した結果（成功したのか、失敗したのか）を即座にフィードバックすることは、UXの根幹です。フィードバックがないと「ボタン押したけど、どうなったの？」という不安を与えます。

### こう書かないとどうなるか

フィードバックを場当たり的に実装するとこうなります：

```tsx
// 場所ごとに別々のスタイルでフィードバックを実装
<p style={{ color: 'red' }}>エラーが発生しました</p>
<p style={{ color: 'green' }}>保存しました</p>
<div className="bg-yellow-100">注意してください</div>
```

見た目がバラバラで統一感がなく、新しい通知を追加するたびに毎回スタイルを考える必要があります。Alertコンポーネントに集約することで、`variant="success"` と書くだけで適切な色・アイコン・配置になります。

### 色だけに頼ってはいけない理由

色覚に課題のあるユーザーは、赤と緑の区別が難しい場合があります。色だけで成功/エラーを伝えると、一部のユーザーには区別がつきません。アイコン（✓ / ✕）とテキストを併用することで、色に頼らずに情報が伝わるようになります（WCAG準拠）。

## 💡 コード例

### 基本: アイコン付き4種類のAlert

variantごとのスタイル・アイコンを`Record`オブジェクトで管理します。if/switchを使わず、variantの追加が1行で済む設計です。

```tsx
import type { ReactNode } from 'react';

type Variant = 'success' | 'warning' | 'error' | 'info';

type Props = {
  variant?: Variant;
  children: ReactNode;
  onClose?: () => void;
};

function Alert({ variant = 'info', children, onClose }: Props) {
  const variantClasses: Record<Variant, string> = {
    success: 'bg-green-50 border-green-500 text-green-800',
    warning: 'bg-yellow-50 border-yellow-500 text-yellow-800',
    error: 'bg-red-50 border-red-500 text-red-800',
    info: 'bg-blue-50 border-blue-500 text-blue-800',
  };

  // 色だけでなくアイコンを併用することで、色覚に頼らない情報伝達を実現
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
      // role="alert"でスクリーンリーダーに「これは通知である」と伝える
      role="alert"
    >
      <span
        className={`flex-shrink-0 w-6 h-6 rounded-full flex items-center justify-center text-white text-sm ${iconClasses[variant]}`}
      >
        {icons[variant]}
      </span>
      <span className="flex-1">{children}</span>
      {/* onCloseが渡された場合のみ閉じるボタンを表示 */}
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

### 実践: 自動消去付き通知システム

実際のアプリでは通知が動的に追加・削除されます。配列stateで複数の通知を管理し、手動（閉じるボタン）と自動（タイマー）の両方で削除できるようにします。

なぜ`useEffect`のクリーンアップが重要か：タイマーを設定した後、ユーザーが手動で閉じた場合、タイマーだけが残り続けます。クリーンアップで`clearTimeout`しないと、既に消えた通知の削除処理が走り、エラーやメモリリークの原因になります。

```tsx
import { useState, useEffect, type ReactNode } from 'react';

type Variant = 'success' | 'warning' | 'error' | 'info';

type AlertProps = {
  variant?: Variant;
  children: ReactNode;
  onClose?: () => void;
  autoClose?: number; // ミリ秒。0なら自動消去しない
};

function Alert({ variant = 'info', children, onClose, autoClose = 0 }: AlertProps) {
  const variantClasses: Record<Variant, string> = {
    success: 'bg-green-50 border-green-500 text-green-800',
    warning: 'bg-yellow-50 border-yellow-500 text-yellow-800',
    error: 'bg-red-50 border-red-500 text-red-800',
    info: 'bg-blue-50 border-blue-500 text-blue-800',
  };
  const icons: Record<Variant, string> = { success: '✓', warning: '⚠', error: '✕', info: 'ℹ' };
  const iconClasses: Record<Variant, string> = {
    success: 'bg-green-500', warning: 'bg-yellow-500', error: 'bg-red-500', info: 'bg-blue-500',
  };

  // autoCloseが指定されている場合、指定ミリ秒後にonCloseを呼ぶ
  useEffect(() => {
    if (autoClose > 0 && onClose) {
      const timer = setTimeout(onClose, autoClose);
      // コンポーネントが消えた後にonCloseが呼ばれるのを防ぐ
      return () => clearTimeout(timer);
    }
  }, [autoClose, onClose]);

  return (
    <div className={`border-l-4 px-4 py-3 rounded flex items-center gap-3 ${variantClasses[variant]}`} role="alert">
      <span className={`flex-shrink-0 w-6 h-6 rounded-full flex items-center justify-center text-white text-sm ${iconClasses[variant]}`}>
        {icons[variant]}
      </span>
      <span className="flex-1">{children}</span>
      {onClose && (
        <button onClick={onClose} className="flex-shrink-0 text-current opacity-50 hover:opacity-100">✕</button>
      )}
    </div>
  );
}

type AlertItem = { id: number; variant: Variant; message: string; autoClose: number };

function App() {
  const [alerts, setAlerts] = useState<AlertItem[]>([
    { id: 1, variant: 'success', message: '保存が完了しました。', autoClose: 3000 },
    { id: 2, variant: 'error', message: 'エラーが発生しました。', autoClose: 0 },
  ]);

  // filterで該当idを除外することでイミュータブルに削除
  const removeAlert = (id: number) => {
    setAlerts((prev) => prev.filter((a) => a.id !== id));
  };

  const addAlert = () => {
    setAlerts((prev) => [...prev, {
      id: Date.now(),
      variant: 'info',
      message: `新しい通知です（${new Date().toLocaleTimeString()}）`,
      autoClose: 5000,
    }]);
  };

  return (
    <div className="p-4 max-w-md">
      <button onClick={addAlert} className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 mb-4">
        通知を追加
      </button>
      <div className="space-y-3">
        {alerts.map((alert) => (
          <Alert key={alert.id} variant={alert.variant} onClose={() => removeAlert(alert.id)} autoClose={alert.autoClose}>
            {alert.message}
          </Alert>
        ))}
      </div>
    </div>
  );
}
```

## 🎯 演習問題

アイコンと閉じるボタン付きのAlertコンポーネントを作ってください。

**要件:**

1. 各variantに対応するアイコンを左側に表示（✓ / ⚠ / ✕ / ℹ）
2. `onClose`がpropsで渡された場合のみ、右端に閉じるボタンを表示
3. `role="alert"`でアクセシビリティを確保

**ヒント:**

```tsx
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

## ✅ 重要ポイント

- [ ] variantごとの色・アイコンをオブジェクトで一元管理する
- [ ] `role="alert"`でスクリーンリーダーに通知の意味を伝える
- [ ] 自動消去には`useEffect` + `setTimeout` + クリーンアップを使う
- [ ] 色だけでなくアイコン・テキストでも状態を伝える（WCAG準拠）

**次のテーマ:** [07. Loading系コンポーネント](./07-loading-component.md)
