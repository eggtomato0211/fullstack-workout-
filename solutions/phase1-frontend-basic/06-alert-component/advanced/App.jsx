import { useState } from 'react';
import Alert from './Alert';

function App() {
  const [alerts, setAlerts] = useState([]);

  const variants = ['success', 'warning', 'error', 'info'];

  const addAlert = () => {
    const variant = variants[Math.floor(Math.random() * variants.length)];
    const now = new Date().toLocaleTimeString('ja-JP');
    setAlerts((prev) => [
      ...prev,
      { id: Date.now(), variant, message: `新しい通知です（${now}）` },
    ]);
  };

  const removeAlert = (id) => {
    setAlerts((prev) => prev.filter((a) => a.id !== id));
  };

  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="space-y-6 w-full max-w-lg">
        <h1 className="text-2xl font-bold text-white text-center">06 発展: 自動消去 + Toast通知</h1>
        <div className="text-center">
          <button
            onClick={addAlert}
            className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
          >
            通知を追加
          </button>
        </div>
        <div className="space-y-3">
          {alerts.map((a) => (
            <Alert
              key={a.id}
              variant={a.variant}
              onClose={() => removeAlert(a.id)}
              autoClose={3000}
            >
              {a.message}
            </Alert>
          ))}
        </div>
      </div>
    </div>
  );
}

export default App;
