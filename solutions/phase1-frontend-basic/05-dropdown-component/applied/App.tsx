import { useState } from 'react';
import Dropdown from './Dropdown';

function App() {
  const [selected, setSelected] = useState('');

  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="space-y-8 text-center">
        <h1 className="text-2xl font-bold text-white">05 応用: 外側クリック + 選択値管理</h1>
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
