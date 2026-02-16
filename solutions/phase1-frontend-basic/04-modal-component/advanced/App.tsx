import { useState } from 'react';
import Modal from './Modal';

function App() {
  const [isOpen, setIsOpen] = useState(false);

  const handleDelete = () => {
    alert('削除しました');
    setIsOpen(false);
  };

  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="text-center">
        <h1 className="text-2xl font-bold text-white mb-6">04 発展: 削除確認ダイアログ</h1>
        <button
          onClick={() => setIsOpen(true)}
          className="bg-red-600 text-white px-4 py-2 rounded hover:bg-red-700"
        >
          削除する
        </button>
        <Modal isOpen={isOpen} onClose={() => setIsOpen(false)}>
          <div className="text-center">
            <h2 className="text-lg font-bold mb-2">本当に削除しますか？</h2>
            <p className="text-gray-500 text-sm mb-6">この操作は取り消せません</p>
            <div className="flex gap-3 justify-center">
              <button
                onClick={() => setIsOpen(false)}
                className="px-4 py-2 bg-gray-200 text-gray-800 rounded hover:bg-gray-300"
              >
                キャンセル
              </button>
              <button
                onClick={handleDelete}
                className="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700"
              >
                削除する
              </button>
            </div>
          </div>
        </Modal>
      </div>
    </div>
  );
}

export default App;
