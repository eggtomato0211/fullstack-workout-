import { useState } from 'react';
import Modal from './Modal';

function App() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="text-center">
        <h1 className="text-2xl font-bold text-white mb-6">04 基本: モーダルの開閉</h1>
        <button
          onClick={() => setIsOpen(true)}
          className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
        >
          モーダルを開く
        </button>
        <Modal isOpen={isOpen} onClose={() => setIsOpen(false)}>
          <h2 className="text-lg font-bold mb-2">タイトル</h2>
          <p className="text-gray-600">モーダルの内容です。</p>
        </Modal>
      </div>
    </div>
  );
}

export default App;
