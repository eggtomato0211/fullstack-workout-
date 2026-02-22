# 07. Loading系コンポーネント - Spinner、Skeleton、Progressの実装

## 🎯 このテーマで学ぶこと

- Spinner・Skeleton・Progress Barの3種類の使い分け
- CSSアニメーション（`animate-spin` / `animate-pulse`）の活用
- `role="status"` と `sr-only` によるスクリーンリーダー対応

## 📖 なぜローディング表示を使い分ける必要があるのか

データの読み込み中に何も表示しないと、ユーザーは「アプリがフリーズした？」と感じてページを離脱してしまいます。ローディング表示は「今処理中です」というフィードバックで離脱を防ぎます。

### 3種類の使い分け

「全部Spinnerでいいのでは？」と思うかもしれませんが、場面によって最適な表示が違います：

- **Spinner**: 汎用的だが「いつ終わるか」が分からない。ボタン内やモーダルの中で使う
- **Skeleton**: 最終的なレイアウトの「形」を先に見せることで、体感速度を向上させる。Twitter/Xのタイムラインやプロフィールカードの読み込みで見かけるあれ
- **Progress Bar**: 進捗を数値で示し、ユーザーに「あとどれくらい」を伝える。ファイルアップロードなど

Skeletonを使うと、実際の読み込み時間は同じでも「速く感じる」という研究結果があります。表示するコンテンツの形が事前に分かっている場合は、Spinnerよりも良いUXを提供できます。

## 💡 コード例

### 基本: SpinnerとSkeleton

2つの最もよく使うローディングパターンを学びます。Spinnerは`animate-spin`で回転、Skeletonは`animate-pulse`で明滅します。

```tsx
// --- Spinner ---
type SpinnerProps = {
  size?: 'sm' | 'md' | 'lg';
};

function Spinner({ size = 'md' }: SpinnerProps) {
  const sizeClasses: Record<string, string> = {
    sm: 'h-4 w-4 border-2',  // ボタン内のインラインSpinner向け
    md: 'h-8 w-8 border-3',  // 一般的なローディング表示
    lg: 'h-12 w-12 border-4', // ページ全体のローディング
  };

  return (
    // role="status"でスクリーンリーダーに状態変化を伝える
    <div role="status">
      {/* border-t-transparentで上辺だけ透明にし、回転時に「切れ目」が見える */}
      <div className={`animate-spin rounded-full border-blue-600 border-t-transparent ${sizeClasses[size]}`} />
      {/* 視覚的には非表示、スクリーンリーダーには読み上げられる */}
      <span className="sr-only">読み込み中...</span>
    </div>
  );
}

// --- Skeleton ---
type SkeletonProps = {
  width: string;
  height: string;
  rounded?: 'none' | 'md' | 'full';
};

function Skeleton({ width, height, rounded = 'md' }: SkeletonProps) {
  const roundedClasses: Record<string, string> = {
    none: 'rounded-none',
    md: 'rounded',
    full: 'rounded-full', // アバター用の円形
  };

  return (
    // classNameではなくstyleでサイズ指定するのは、任意のpx/パーセント値に対応するため
    <div
      className={`bg-gray-200 animate-pulse ${roundedClasses[rounded]}`}
      style={{ width, height }}
    />
  );
}

// Skeletonを組み合わせて、実際のカードと同じレイアウトで表示する
// → データ到着後のレイアウトシフト（ガタッとずれる現象）を防ぐ
function CardSkeleton() {
  return (
    <div className="bg-white rounded-lg shadow p-6 max-w-sm">
      <div className="flex items-center gap-4 mb-4">
        <Skeleton width="48px" height="48px" rounded="full" />
        <div className="space-y-2">
          <Skeleton width="120px" height="16px" />
          <Skeleton width="80px" height="12px" />
        </div>
      </div>
      <div className="space-y-2">
        <Skeleton width="100%" height="14px" />
        <Skeleton width="100%" height="14px" />
        <Skeleton width="60%" height="14px" />
      </div>
    </div>
  );
}

function App() {
  return (
    <div className="p-4 space-y-8">
      <div className="flex items-center gap-6">
        <Spinner size="sm" />
        <Spinner size="md" />
        <Spinner size="lg" />
      </div>
      <CardSkeleton />
    </div>
  );
}
```

### 実践: Progress Barとローディングフローの統合

Progress Barは進捗を数値で表示します。`value`と`max`を受け取る設計にすることで、0-100%だけでなく「5/20件」のような任意のスケールにも対応できます。

```tsx
import { useState } from 'react';

type ProgressBarProps = {
  value: number;
  max?: number;
  label?: string;
};

function ProgressBar({ value, max = 100, label }: ProgressBarProps) {
  // 0〜100%にクランプ：外から不正な値が来てもUIが壊れない防御的処理
  const percentage = Math.min(Math.max((value / max) * 100, 0), 100);

  return (
    <div className="w-full">
      {label && (
        <div className="flex justify-between mb-1">
          <span className="text-sm font-medium text-gray-700">{label}</span>
          <span className="text-sm text-gray-500">{Math.round(percentage)}%</span>
        </div>
      )}
      <div className="w-full bg-gray-200 rounded-full h-3 overflow-hidden">
        {/* transition-allでwidthの変化を滑らかにアニメーション */}
        {/* これがないとバーがカクカク飛んでしまう */}
        <div
          className="bg-blue-600 h-full rounded-full transition-all duration-300"
          style={{ width: `${percentage}%` }}
        />
      </div>
    </div>
  );
}

function App() {
  const [progress, setProgress] = useState(0);
  const [isLoading, setIsLoading] = useState(false);

  const fetchData = () => {
    setIsLoading(true);
    setProgress(0);

    // 実務ではonUploadProgress等を使うが、ここではsetTimeoutでシミュレーション
    const steps = [20, 50, 75, 90, 100];
    steps.forEach((step, index) => {
      setTimeout(() => {
        setProgress(step);
        if (step === 100) {
          setTimeout(() => setIsLoading(false), 300);
        }
      }, (index + 1) * 500);
    });
  };

  return (
    <div className="p-4 max-w-md space-y-6">
      <button
        onClick={fetchData}
        disabled={isLoading}
        className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 disabled:opacity-50"
      >
        {isLoading ? '読み込み中...' : 'データを取得'}
      </button>
      {isLoading && <ProgressBar value={progress} label="データ取得中" />}
      {!isLoading && progress === 100 && (
        <p className="text-green-600 font-medium">データの読み込みが完了しました！</p>
      )}
    </div>
  );
}
```

## 🎯 演習問題

Skeletonコンポーネントと、それを組み合わせたCardSkeletonを作ってください。

**要件:**

1. `Skeleton`は`width`, `height`, `rounded`をpropsで受け取る
2. `animate-pulse`でローディング感を表現
3. `CardSkeleton`はアバター（丸型）+ テキスト行3行で構成
4. 実際のカードと同じレイアウトで配置する

**ヒント:**

```tsx
type Props = {
  width: string;
  height: string;
  rounded?: 'none' | 'md' | 'full';
};

function Skeleton({ width, height, rounded = 'md' }: Props) {
  // bg-gray-200 + animate-pulse でパルスアニメーション
  // style={{ width, height }} でサイズ指定
}
```

## ✅ 重要ポイント

- [ ] Spinner(`animate-spin`)、Skeleton(`animate-pulse`)、Progress Barを場面で使い分ける
- [ ] コンテンツの形が予測できる場合はSkeletonの方がUXが良い
- [ ] Progress Barは`transition`で滑らかに変化させる
- [ ] `role="status"`や`sr-only`でスクリーンリーダーに対応する

**次のテーマ:** [08. useStateの基本 - プリミティブ値](./08-usestate-primitive.md)
