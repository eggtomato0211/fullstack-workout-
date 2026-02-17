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

まずは最もシンプルな形から始めます。variantごとのスタイルを`Record`オブジェクトで一元管理することで、条件分岐（if/switch）を使わずにスタイルを切り替えられます。この「データ駆動」のアプローチは、variantが増えてもオブジェクトに1行追加するだけで済むため、保守性が高くなります。

```tsx
import type { ReactNode } from 'react';

type Variant = 'success' | 'warning' | 'error' | 'info';

type Props = {
  variant?: Variant;
  children: ReactNode;
};

function Alert({ variant = 'info', children }: Props) {
  // variantとクラス名の対応をオブジェクトで管理することで、
  // if/switchを使わず、新しいvariant追加時も1行で済む
  const variantClasses: Record<Variant, string> = {
    success: 'bg-green-50 border-green-500 text-green-800',
    warning: 'bg-yellow-50 border-yellow-500 text-yellow-800',
    error: 'bg-red-50 border-red-500 text-red-800',
    info: 'bg-blue-50 border-blue-500 text-blue-800',
  };

  return (
    <div
      className={`border-l-4 px-4 py-3 rounded ${variantClasses[variant]}`}
      // role="alert"でスクリーンリーダーに「これは通知である」と伝える
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

> **💡 次のステップへ:** 基本のAlertは色とテキストだけで状態を伝えていますが、色覚に課題のあるユーザーには色の違いが判別しにくい場合があります。次のステップでは、アイコンを追加して色に頼らない情報伝達を実現します。

### 応用: アイコン付きAlert

色だけでなくアイコンを併用することで、WCAG（Webアクセシビリティガイドライン）の「色だけに依存しない情報伝達」を満たします。アイコン用のスタイルも`Record`オブジェクトで管理し、variantClasses・icons・iconClassesの3つのオブジェクトを同じキーで揃えることで、不整合が起きにくい構造にしています。

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

  // 色覚に頼らず状態を伝えるためのアイコン
  // 実務ではreact-iconsやheroiconsなどのライブラリを使うことが多い
  const icons: Record<Variant, string> = {
    success: '✓',
    warning: '⚠',
    error: '✕',
    info: 'ℹ',
  };

  // アイコン背景色をvariantに合わせることで視覚的一貫性を保つ
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
      {/* flex-shrink-0でアイコンがテキスト量に応じて潰れるのを防ぐ */}
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

> **💡 次のステップへ:** 見た目は整いましたが、実際のアプリでは通知を「閉じる」「一定時間で消える」といったインタラクションが求められます。次のステップでは、`onClose`コールバックによる手動消去と、`useEffect` + `setTimeout`による自動消去を組み合わせた実践的な通知システムを作ります。

### 実践: 閉じるボタン + 自動消去

実際のアプリケーションでは、通知は動的に追加・削除されるものです。ここでは配列stateで複数の通知を管理し、手動（閉じるボタン）と自動（タイマー）の両方で削除できる仕組みを作ります。`useEffect`のクリーンアップ関数で`clearTimeout`することで、コンポーネントがアンマウントされたときにタイマーが残り続ける（メモリリーク）問題を防いでいます。

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

  // autoCloseが指定されている場合、指定ミリ秒後にonCloseを呼ぶ
  // クリーンアップでclearTimeoutすることで、コンポーネントが
  // アンマウントされた後にonCloseが呼ばれるのを防ぐ（メモリリーク対策）
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
      {/* flex-1でテキスト部分が残りスペースを埋め、閉じるボタンを右端に押しやる */}
      <span className="flex-1">{children}</span>
      {/* onCloseが渡された場合のみ閉じるボタンを表示（条件付きレンダリング） */}
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

  // filterで該当idを除外した新配列を作ることで、イミュータブルに削除する
  const removeAlert = (id: number) => {
    setAlerts((prev) => prev.filter((alert) => alert.id !== id));
  };

  // Date.now()をidに使うことで、簡易的にユニークなキーを生成する
  // 本番ではuuidなどの衝突しにくいID生成を推奨
  const addAlert = () => {
    const newAlert: AlertItem = {
      id: Date.now(),
      variant: 'info',
      message: `新しい通知です（${new Date().toLocaleTimeString()}）`,
      autoClose: 5000,
    };
    // スプレッド構文で既存配列を展開し、末尾に新要素を追加（イミュータブル更新）
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

## ✅ 重要ポイント

- [ ] variantごとに色・アイコンをオブジェクトで管理して一貫性を保つ
- [ ] `role="alert"`を付与してスクリーンリーダーに通知の意味を伝える
- [ ] 自動消去には`useEffect` + `setTimeout`を使い、クリーンアップで`clearTimeout`する
- [ ] 色だけでなくアイコンやテキストでも状態を伝え、アクセシビリティを確保する

**次のテーマ:** [07. Loading系コンポーネント](./07-loading-component.md)
