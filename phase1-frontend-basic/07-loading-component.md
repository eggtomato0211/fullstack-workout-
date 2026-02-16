# 07. Loading系コンポーネント - Spinner、Skeleton、Progressの実装

## 📖 概念

Loading系コンポーネントは、データ取得や処理の待ち時間にユーザーへフィードバックを返すためのUIです。Spinner（回転インジケーター）、Skeleton（コンテンツの骨格表示）、Progress Bar（進捗表示）の3種類を使い分けることで、待ち時間のUXを大きく改善できます。

**よくある誤解:**
- ❌ 「Spinnerだけあれば十分」→ コンテンツの形が予測できる場合はSkeletonの方がUXが良い（体感速度の向上）
- ❌ 「ローディング中は何も表示しなくていい」→ フィードバックがないとユーザーはフリーズと誤解する
- ❌ 「Progress Barの値は適当でいい」→ 実際の進捗に基づかない嘘のプログレスはユーザーの信頼を損なう

## 💡 コード例

### 基本: Spinner

```tsx
type Props = {
  size?: 'sm' | 'md' | 'lg';
};

function Spinner({ size = 'md' }: Props) {
  const sizeClasses: Record<string, string> = {
    sm: 'h-4 w-4 border-2',
    md: 'h-8 w-8 border-3',
    lg: 'h-12 w-12 border-4',
  };

  return (
    <div role="status">
      <div
        className={`animate-spin rounded-full border-blue-600 border-t-transparent ${sizeClasses[size]}`}
      />
      <span className="sr-only">読み込み中...</span>
    </div>
  );
}

// 使用例
function App() {
  return (
    <div className="p-4 flex items-center gap-6">
      <Spinner size="sm" />
      <Spinner size="md" />
      <Spinner size="lg" />
    </div>
  );
}
```

### 応用: Skeleton

```tsx
type SkeletonProps = {
  width: string;
  height: string;
  rounded?: 'none' | 'sm' | 'md' | 'lg' | 'full';
};

function Skeleton({ width, height, rounded = 'md' }: SkeletonProps) {
  const roundedClasses: Record<string, string> = {
    none: 'rounded-none',
    sm: 'rounded-sm',
    md: 'rounded',
    lg: 'rounded-lg',
    full: 'rounded-full',
  };

  return (
    <div
      className={`bg-gray-200 animate-pulse ${roundedClasses[rounded]}`}
      style={{ width, height }}
    />
  );
}

function CardSkeleton() {
  return (
    <div className="bg-white rounded-lg shadow p-6 max-w-sm">
      {/* アバター */}
      <div className="flex items-center gap-4 mb-4">
        <Skeleton width="48px" height="48px" rounded="full" />
        <div className="space-y-2">
          <Skeleton width="120px" height="16px" />
          <Skeleton width="80px" height="12px" />
        </div>
      </div>
      {/* テキスト行 */}
      <div className="space-y-2">
        <Skeleton width="100%" height="14px" />
        <Skeleton width="100%" height="14px" />
        <Skeleton width="60%" height="14px" />
      </div>
    </div>
  );
}

// 使用例
function App() {
  return (
    <div className="p-4 space-y-6">
      <h2 className="text-lg font-bold">Skeleton表示</h2>
      <CardSkeleton />
    </div>
  );
}
```

### 実践: Progress Bar

```tsx
import { useState } from 'react';

type ProgressBarProps = {
  value: number;
  max?: number;
  label?: string;
};

function ProgressBar({ value, max = 100, label }: ProgressBarProps) {
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
        <div
          className="bg-blue-600 h-full rounded-full transition-all duration-300"
          style={{ width: `${percentage}%` }}
        />
      </div>
    </div>
  );
}

// 使用例: データ取得と連動
function App() {
  const [progress, setProgress] = useState(0);
  const [isLoading, setIsLoading] = useState(false);
  const [data, setData] = useState<string | null>(null);

  const fetchData = () => {
    setIsLoading(true);
    setProgress(0);
    setData(null);

    // 段階的に進捗を更新するシミュレーション
    const steps = [20, 50, 75, 90, 100];
    steps.forEach((step, index) => {
      setTimeout(() => {
        setProgress(step);
        if (step === 100) {
          setTimeout(() => {
            setData('データの読み込みが完了しました！');
            setIsLoading(false);
          }, 300);
        }
      }, (index + 1) * 500);
    });
  };

  return (
    <div className="p-4 max-w-md space-y-6">
      <button
        onClick={fetchData}
        disabled={isLoading}
        className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        {isLoading ? '読み込み中...' : 'データを取得'}
      </button>

      {isLoading && (
        <ProgressBar value={progress} label="データ取得中" />
      )}

      {data && (
        <p className="text-green-600 font-medium">{data}</p>
      )}
    </div>
  );
}
```

## 🎯 演習問題

### 基本: Spinnerの実装

3サイズ（`sm`, `md`, `lg`）に対応するSpinnerコンポーネントを作ってください。

```tsx
type Props = {
  size?: 'sm' | 'md' | 'lg';
};

function Spinner({ size = 'md' }: Props) {
  // ここにコードを書く
  // size に応じてサイズを切り替える
  // CSSアニメーション（animate-spin）で回転させる

  return (
    <div role="status">
      {/* スピナー要素 */}
    </div>
  );
}
```

**期待される動作:**
- `sm`: 小さいスピナー（16px）
- `md`: 標準スピナー（32px）
- `lg`: 大きいスピナー（48px）
- 回転アニメーションが常に動いている
- `role="status"`でアクセシビリティに配慮

---

### 応用: Skeletonの実装

幅・高さ・角丸を指定できるSkeletonコンポーネントと、それを組み合わせたCardSkeletonを作ってください。

**要件:**
1. `Skeleton`は`width`, `height`, `rounded`をpropsで受け取る
2. パルスアニメーション（`animate-pulse`）でローディング感を表現
3. `CardSkeleton`はアバター（丸型）+ テキスト行3行で構成
4. 実際のカードコンテンツと同じレイアウトでSkeletonを配置

**ヒント:**
```tsx
type Props = {
  width: string;
  height: string;
  rounded?: 'none' | 'sm' | 'md' | 'lg' | 'full';
};

function Skeleton({ width, height, rounded = 'md' }: Props) {
  // bg-gray-200 + animate-pulse でパルスアニメーション
  // style={{ width, height }} でサイズ指定
}
```

---

### 発展: Progress Bar + データ取得連動

Progress Barコンポーネントを作成し、データ取得のシミュレーションと連動させてください。

**要件:**
1. `value`（現在値）と`max`（最大値、デフォルト100）をpropsで受け取る
2. ラベルと現在のパーセンテージを表示
3. バーの幅を`value/max`のパーセンテージで制御
4. ボタンクリックでデータ取得をシミュレーション（`setTimeout`で段階的に進捗更新）
5. 100%到達時に完了メッセージを表示

**完成イメージ:**
```
[データを取得]

データ取得中                         75%
[██████████████████░░░░░░]

↓ 完了後

データの読み込みが完了しました！
```

## ✅ 重要ポイント

- [ ] Spinnerは`animate-spin`、Skeletonは`animate-pulse`でアニメーションを実装する
- [ ] コンテンツの形が予測できる場合はSpinnerよりSkeletonがUXに優れる
- [ ] Progress Barは`transition`で値の変化を滑らかに見せる
- [ ] `role="status"`や`sr-only`テキストでスクリーンリーダーに対応する

**次のテーマ:** [08. useStateの基本 - プリミティブ値](./08-usestate-primitive.md)
