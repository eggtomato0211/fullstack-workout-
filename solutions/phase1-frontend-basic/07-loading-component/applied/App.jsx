import Skeleton, { CardSkeleton } from './Skeleton';

function App() {
  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="space-y-8 text-center">
        <h1 className="text-2xl font-bold text-white">07 応用: Skeleton</h1>
        <div className="space-y-2">
          <p className="text-gray-400 text-sm">基本的なSkeleton</p>
          <div className="flex gap-4 items-end justify-center">
            <Skeleton width="80px" height="80px" rounded="full" />
            <Skeleton width="120px" height="20px" />
            <Skeleton width="200px" height="40px" rounded="lg" />
          </div>
        </div>
        <div className="space-y-2">
          <p className="text-gray-400 text-sm">CardSkeleton</p>
          <div className="flex gap-4 justify-center">
            <CardSkeleton />
            <CardSkeleton />
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
