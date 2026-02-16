import ProductCard from './ProductCard';

function App() {
  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="space-y-8 text-center">
        <h1 className="text-2xl font-bold text-white">03 発展: 商品カード</h1>
        <div className="flex gap-6">
          <div className="w-72">
            <ProductCard
              name="ワイヤレスイヤホン"
              description="高音質Bluetooth対応。ノイズキャンセリング機能付き。"
              price="¥12,800"
            />
          </div>
          <div className="w-72">
            <ProductCard
              name="スマートウォッチ"
              description="健康管理と通知確認がこれ一つで。防水対応。"
              price="¥24,800"
              variant="outlined"
            />
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
