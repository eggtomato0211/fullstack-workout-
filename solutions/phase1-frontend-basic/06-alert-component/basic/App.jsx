import Alert from './Alert';

function App() {
  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="space-y-6 w-full max-w-lg">
        <h1 className="text-2xl font-bold text-white text-center">06 基本: 4つのvariant</h1>
        <div className="space-y-4">
          <Alert variant="success">保存が完了しました。</Alert>
          <Alert variant="warning">入力内容を確認してください。</Alert>
          <Alert variant="error">エラーが発生しました。</Alert>
          <Alert variant="info">新しいバージョンが利用可能です。</Alert>
        </div>
      </div>
    </div>
  );
}

export default App;
