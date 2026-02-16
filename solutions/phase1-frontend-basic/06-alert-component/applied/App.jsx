import { useState } from 'react';
import Alert from './Alert';

function App() {
  const [alerts, setAlerts] = useState([
    { id: 1, variant: 'success', message: '保存が完了しました。' },
    { id: 2, variant: 'warning', message: '入力内容を確認してください。' },
    { id: 3, variant: 'error', message: 'エラーが発生しました。' },
    { id: 4, variant: 'info', message: '新しいバージョンが利用可能です。' },
  ]);

  const removeAlert = (id) => {
    setAlerts((prev) => prev.filter((a) => a.id !== id));
  };

  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="space-y-6 w-full max-w-lg">
        <h1 className="text-2xl font-bold text-white text-center">06 応用: アイコン + 閉じるボタン</h1>
        <div className="space-y-4">
          {alerts.map((a) => (
            <Alert key={a.id} variant={a.variant} onClose={() => removeAlert(a.id)}>
              {a.message}
            </Alert>
          ))}
          {alerts.length === 0 && (
            <p className="text-gray-400 text-center">すべて閉じました</p>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
