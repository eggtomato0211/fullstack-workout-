import { useState } from 'react';
import Dropdown from './Dropdown';

function App() {
  const [selected, setSelected] = useState('');

  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="space-y-8 text-center">
        <h1 className="text-2xl font-bold text-white">05 発展: キーボードナビゲーション</h1>
        <p className="text-gray-400 text-sm">矢印キーで移動、Enterで選択、Escで閉じる</p>
        <Dropdown
          label="フルーツを選択"
          items={['りんご', 'バナナ', 'オレンジ', 'ぶどう']}
          value={selected}
          onChange={setSelected}
        />
        {selected && <p className="text-gray-300">選択中: {selected}</p>}
      </div>
    </div>
  );
}

export default App;
