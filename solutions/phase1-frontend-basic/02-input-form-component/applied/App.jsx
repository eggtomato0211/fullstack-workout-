import { useState } from 'react';
import Input from './Input';

function App() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="bg-white rounded-lg p-8 max-w-md w-full space-y-4">
        <h1 className="text-xl font-bold">02 応用: 複数type + エラー表示</h1>
        <Input
          id="email"
          label="メールアドレス"
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="example@mail.com"
        />
        <Input
          id="password"
          label="パスワード"
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="8文字以上"
        />
        <Input
          id="error-demo"
          label="エラー表示デモ"
          value=""
          onChange={() => {}}
          error="これはエラーメッセージです"
        />
      </div>
    </div>
  );
}

export default App;
