import Dropdown from './Dropdown';

function App() {
  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="space-y-8 text-center">
        <h1 className="text-2xl font-bold text-white">05 基本: ドロップダウン</h1>
        <Dropdown label="メニュー" items={['プロフィール', '設定', 'ログアウト']} />
      </div>
    </div>
  );
}

export default App;
