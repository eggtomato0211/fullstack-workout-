import Card from './Card';

function App() {
  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="space-y-8 text-center">
        <h1 className="text-2xl font-bold text-white">03 基本: シンプルなカード</h1>
        <div className="space-y-4">
          <Card>
            <h2 className="text-lg font-bold">カードタイトル</h2>
            <p className="text-gray-600 mt-2">
              これはシンプルなカードコンポーネントです。白背景、角丸、シャドウ、パディングがあります。
            </p>
          </Card>
          <Card>
            <h2 className="text-lg font-bold">もう一つのカード</h2>
            <p className="text-gray-600 mt-2">
              children propで自由にコンテンツを配置できます。
            </p>
          </Card>
        </div>
      </div>
    </div>
  );
}

export default App;
