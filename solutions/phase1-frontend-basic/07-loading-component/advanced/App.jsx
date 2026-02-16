import { useState } from 'react';
import ProgressBar from './ProgressBar';

function App() {
  const [progress, setProgress] = useState(0);
  const [isLoading, setIsLoading] = useState(false);
  const [data, setData] = useState(null);

  const fetchData = () => {
    setIsLoading(true);
    setProgress(0);
    setData(null);

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
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="space-y-8 text-center w-96">
        <h1 className="text-2xl font-bold text-white">07 発展: Progress Bar</h1>
        <div className="bg-white rounded-lg shadow p-6 space-y-4">
          {isLoading && <ProgressBar value={progress} label="データ取得中" />}
          <button
            onClick={fetchData}
            disabled={isLoading}
            className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            {isLoading ? '読み込み中...' : 'データを取得'}
          </button>
          {data && <p className="text-green-600 font-medium">{data}</p>}
        </div>
      </div>
    </div>
  );
}

export default App;
