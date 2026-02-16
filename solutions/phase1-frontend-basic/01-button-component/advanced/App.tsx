import { useState } from 'react';
import Button from './Button';

function App() {
  const [isLoading, setIsLoading] = useState(false);

  const handleSave = () => {
    setIsLoading(true);
    setTimeout(() => setIsLoading(false), 2000);
  };

  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="space-y-8 text-center">
        <h1 className="text-2xl font-bold text-white">01 発展: ローディング状態</h1>

        {/* ローディングのデモ */}
        <div className="space-y-2">
          <p className="text-gray-400 text-sm">クリックすると2秒間ローディング</p>
          <div className="flex gap-4 justify-center">
            <Button variant="primary" isLoading={isLoading} onClick={handleSave}>
              保存
            </Button>
            <Button variant="secondary" disabled>
              無効なボタン
            </Button>
            <Button variant="danger" size="lg">
              削除
            </Button>
          </div>
        </div>

        {/* 全パターン一覧 */}
        <div className="space-y-2">
          <p className="text-gray-400 text-sm">サイズ一覧</p>
          <div className="flex gap-4 items-center justify-center">
            <Button size="sm">Small</Button>
            <Button size="md">Medium</Button>
            <Button size="lg">Large</Button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
