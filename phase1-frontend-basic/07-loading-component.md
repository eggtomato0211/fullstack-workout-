# 07. Loading系コンポーネント - Spinner、Skeleton、Progressの実装

## 🎯 このテーマで学ぶこと

- CSSアニメーション（`animate-spin` / `animate-pulse`）の活用
- Spinner・Skeleton・Progress Barの3種類の使い分け
- `style`プロパティによる動的なサイズ・幅の制御
- `setTimeout` による非同期処理のシミュレーション
- `role="status"` と `sr-only` によるスクリーンリーダー対応

**なぜ重要か:** 現代のWebアプリケーションはAPIからデータを取得して表示するケースが大半です。データの読み込み中にユーザーに適切なフィードバックを返せるかどうかで、アプリの体感品質は大きく変わります。Spinnerしか知らない開発者が多い中、SkeletonやProgress Barを場面に応じて使い分けられると、一段上のUXを提供できます。

## 📖 概念

Loading系コンポーネントは、データ取得や処理の待ち時間にユーザーへフィードバックを返すためのUIです。Spinner（回転インジケーター）、Skeleton（コンテンツの骨格表示）、Progress Bar（進捗表示）の3種類を使い分けることで、待ち時間のUXを大きく改善できます。

**背景と設計意図:** 読み込み中に何も表示しないと、ユーザーはアプリがフリーズしたと感じてページを離脱してしまいます。ローディング表示は「今処理しています」というフィードバックを与えることで離脱を防ぎます。Spinnerは汎用的ですが「いつ終わるか不明」な印象を与えます。Skeletonは最終的なレイアウトを先に見せることで体感速度を向上させ、Progress Barは進捗を数値で示すことで安心感を与えます。

**実務での活用場面:** API通信中のSpinner、一覧ページ読み込み時のSkeleton（Twitter/Xのタイムライン、YouTubeのサムネイル一覧）、ファイルアップロード時のProgress Bar、ダッシュボードの初期読み込みなど。特にSkeletonはNext.jsの`loading.tsx`やReactのSuspenseと組み合わせて使うパターンが普及しています。

**よくある誤解:**

- ❌ 「Spinnerだけあれば十分」→ コンテンツの形が予測できる場合はSkeletonの方がUXが良い（体感速度の向上）
- ❌ 「ローディング中は何も表示しなくていい」→ フィードバックがないとユーザーはフリーズと誤解する
- ❌ 「Progress Barの値は適当でいい」→ 実際の進捗に基づかない嘘のプログレスはユーザーの信頼を損なう

## 💡 コード例

### 基本: Spinner

Spinnerはローディング表示の最も基本的なパターンです。CSSの`animate-spin`だけで回転アニメーションを実現できるため、実装コストが低く汎用的に使えます。`size`プロパティを用意しておくことで、ボタン内の小さなインジケーターからページ全体のローディングまで、1つのコンポーネントで対応できるようにしています。

```tsx
type Props = {
  size?: 'sm' | 'md' | 'lg';
};

function Spinner({ size = 'md' }: Props) {
  // サイズごとにTailwindクラスをマッピング。borderの太さもサイズに比例させることで、
  // どのサイズでも視覚的なバランスが保たれる
  const sizeClasses: Record<string, string> = {
    sm: 'h-4 w-4 border-2',
    md: 'h-8 w-8 border-3',
    lg: 'h-12 w-12 border-4',
  };

  return (
    // role="status"を付与することで、スクリーンリーダーがこの領域を
    // ライブリージョンとして認識し、状態変化を読み上げてくれる
    <div role="status">
      {/* border-t-transparent で上辺だけ透明にし、回転時に「切れ目」が見える仕組み */}
      <div
        className={`animate-spin rounded-full border-blue-600 border-t-transparent ${sizeClasses[size]}`}
      />
      {/* 視覚的には非表示だが、スクリーンリーダーには「読み込み中」と伝わる */}
      <span className="sr-only">読み込み中...</span>
    </div>
  );
}

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

> **💡 次のステップへ:** Spinnerは「何かが読み込まれている」ことは伝わりますが、最終的にどんなコンテンツが表示されるかは分かりません。次のSkeletonでは、コンテンツの「形」を先に見せることで、ユーザーの体感待ち時間を短縮するアプローチを学びます。

### 応用: Skeleton

Skeletonはコンテンツが表示される前にその「骨格」をグレーのブロックで描画するパターンです。ユーザーに最終的なレイアウトを予測させることで、Spinnerよりも体感速度が向上します。汎用的な`Skeleton`コンポーネントを1つ作り、`width`・`height`・`rounded`を外から指定できるようにすることで、カードやリストなどあらゆるレイアウトに再利用できます。

```tsx
type SkeletonProps = {
  width: string;
  height: string;
  rounded?: 'none' | 'sm' | 'md' | 'lg' | 'full';
};

function Skeleton({ width, height, rounded = 'md' }: SkeletonProps) {
  // 角丸のバリエーションをマッピング。アバター用にfull（円形）も用意しておく
  const roundedClasses: Record<string, string> = {
    none: 'rounded-none',
    sm: 'rounded-sm',
    md: 'rounded',
    lg: 'rounded-lg',
    full: 'rounded-full',
  };

  return (
    // animate-pulseで明滅させることで「まだ読み込み中」であることを視覚的に伝える。
    // classNameではなくstyleでサイズを指定するのは、任意のピクセル値やパーセント値に対応するため
    <div
      className={`bg-gray-200 animate-pulse ${roundedClasses[rounded]}`}
      style={{ width, height }}
    />
  );
}

// 実際のカードUIと同じレイアウト構造でSkeletonを配置する。
// こうすることで、データ到着後にレイアウトがガタッとずれる（レイアウトシフト）を防げる
function CardSkeleton() {
  return (
    <div className="bg-white rounded-lg shadow p-6 max-w-sm">
      {/* アバター + ユーザー名エリア：実際のカードと同じflex配置にする */}
      <div className="flex items-center gap-4 mb-4">
        <Skeleton width="48px" height="48px" rounded="full" />
        <div className="space-y-2">
          <Skeleton width="120px" height="16px" />
          <Skeleton width="80px" height="12px" />
        </div>
      </div>
      {/* 本文テキストエリア：最後の行だけ短くして、自然なテキストの終端を表現 */}
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
    <div className="p-4 space-y-6">
      <h2 className="text-lg font-bold">Skeleton表示</h2>
      <CardSkeleton />
    </div>
  );
}
```

> **💡 次のステップへ:** Spinnerは「処理中」、Skeletonは「コンテンツの形を予告」しますが、どちらも「あとどれくらいかかるか」は伝えられません。次のProgress Barでは、進捗を数値で可視化してユーザーに安心感を与える方法を学びます。

### 実践: Progress Bar

Progress Barは処理の進捗を数値（パーセンテージ）で表示するパターンです。ファイルアップロードやデータ取得のように「全体の何%が完了したか」が分かる処理で使います。`value`と`max`を受け取る設計にしておくことで、0-100だけでなく「5/20件処理済み」のような任意のスケールにも対応できます。ここでは`setTimeout`で段階的に進捗を更新するシミュレーションも組み合わせ、実際のデータ取得フローに近い体験を実装します。

```tsx
import { useState } from 'react';

type ProgressBarProps = {
  value: number;
  max?: number;
  label?: string;
};

function ProgressBar({ value, max = 100, label }: ProgressBarProps) {
  // Math.min/Maxで0〜100%の範囲にクランプする。
  // 外から不正な値（負数や超過値）が渡されてもUIが壊れないようにする防御的な処理
  const percentage = Math.min(Math.max((value / max) * 100, 0), 100);

  return (
    <div className="w-full">
      {label && (
        <div className="flex justify-between mb-1">
          <span className="text-sm font-medium text-gray-700">{label}</span>
          <span className="text-sm text-gray-500">{Math.round(percentage)}%</span>
        </div>
      )}
      {/* overflow-hiddenで内側のバーが角丸からはみ出すのを防ぐ */}
      <div className="w-full bg-gray-200 rounded-full h-3 overflow-hidden">
        {/* transition-allにより、widthが変わるたびに滑らかにアニメーションする。
            これがないとバーがカクカク飛んでしまい、体験が悪くなる */}
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
  const [data, setData] = useState<string | null>(null);

  const fetchData = () => {
    setIsLoading(true);
    setProgress(0);
    setData(null);

    // 実際のAPI通信では進捗イベント（onUploadProgressなど）を使うが、
    // ここではsetTimeoutで段階的に進捗を更新するシミュレーションを行う
    const steps = [20, 50, 75, 90, 100];
    steps.forEach((step, index) => {
      setTimeout(() => {
        setProgress(step);
        if (step === 100) {
          // 100%到達後に少し間を置いてから完了表示にする。
          // すぐ消すとユーザーが「100%」を認識できないため
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

## ✅ 重要ポイント

- [ ] Spinnerは`animate-spin`、Skeletonは`animate-pulse`でアニメーションを実装する
- [ ] コンテンツの形が予測できる場合はSpinnerよりSkeletonがUXに優れる
- [ ] Progress Barは`transition`で値の変化を滑らかに見せる
- [ ] `role="status"`や`sr-only`テキストでスクリーンリーダーに対応する

**次のテーマ:** [08. useStateの基本 - プリミティブ値](./08-usestate-primitive.md)
