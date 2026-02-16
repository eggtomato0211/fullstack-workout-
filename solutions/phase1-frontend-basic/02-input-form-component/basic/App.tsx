import { useState } from 'react';
import Input from './Input';

function App() {
  const [name, setName] = useState('');

  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="bg-white rounded-lg p-8 max-w-md w-full">
        <h1 className="text-xl font-bold mb-4">02 基本: Input</h1>
        <Input
          label="名前"
          value={name}
          onChange={(e) => setName(e.target.value)}
          placeholder="山田太郎"
        />
        <p className="mt-4 text-sm text-gray-500">入力値: {name}</p>
      </div>
    </div>
  );
}

export default App;
