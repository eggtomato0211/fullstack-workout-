import Card, { CardHeader, CardBody, CardFooter } from './Card';

function App() {
  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="space-y-8 text-center">
        <h1 className="text-2xl font-bold text-white">03 応用: Header/Body/Footer + variant</h1>
        <div className="flex gap-6">
          <Card variant="default">
            <CardHeader><h2 className="font-bold">Default</h2></CardHeader>
            <CardBody><p className="text-gray-600">通常のシャドウ付きカード</p></CardBody>
            <CardFooter><p className="text-gray-500 text-sm">フッター</p></CardFooter>
          </Card>
          <Card variant="outlined">
            <CardHeader><h2 className="font-bold">Outlined</h2></CardHeader>
            <CardBody><p className="text-gray-600">ボーダー付きカード</p></CardBody>
            <CardFooter><p className="text-gray-500 text-sm">フッター</p></CardFooter>
          </Card>
          <Card variant="elevated">
            <CardHeader><h2 className="font-bold">Elevated</h2></CardHeader>
            <CardBody><p className="text-gray-600">強いシャドウ付きカード</p></CardBody>
            <CardFooter><p className="text-gray-500 text-sm">フッター</p></CardFooter>
          </Card>
        </div>
      </div>
    </div>
  );
}

export default App;
