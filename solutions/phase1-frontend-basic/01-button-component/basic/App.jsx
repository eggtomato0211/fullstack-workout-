import Button from './Button';

function App() {
  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="space-y-6 text-center">
        <h1 className="text-2xl font-bold text-white">01 基本: variant</h1>
        <div className="flex gap-4">
          <Button variant="primary">保存</Button>
          <Button variant="secondary">キャンセル</Button>
          <Button variant="outline">詳細</Button>
        </div>
      </div>
    </div>
  );
}

export default App;
