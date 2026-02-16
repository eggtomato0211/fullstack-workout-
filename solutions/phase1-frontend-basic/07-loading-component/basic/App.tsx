import Spinner from './Spinner';

function App() {
  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="space-y-8 text-center">
        <h1 className="text-2xl font-bold text-white">07 基本: Spinner</h1>
        <div className="flex gap-8 items-center justify-center">
          <div className="space-y-2">
            <Spinner size="sm" />
            <p className="text-gray-400 text-xs">sm</p>
          </div>
          <div className="space-y-2">
            <Spinner size="md" />
            <p className="text-gray-400 text-xs">md</p>
          </div>
          <div className="space-y-2">
            <Spinner size="lg" />
            <p className="text-gray-400 text-xs">lg</p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
