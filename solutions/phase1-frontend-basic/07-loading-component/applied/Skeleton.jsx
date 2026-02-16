function Skeleton({ width, height, rounded = 'md' }) {
  const roundedClasses = {
    none: 'rounded-none',
    sm: 'rounded-sm',
    md: 'rounded',
    lg: 'rounded-lg',
    full: 'rounded-full',
  };

  return (
    <div
      className={`bg-gray-200 animate-pulse ${roundedClasses[rounded]}`}
      style={{ width, height }}
    />
  );
}

export function CardSkeleton() {
  return (
    <div className="bg-white rounded-lg shadow p-6 max-w-sm">
      <div className="flex items-center gap-4 mb-4">
        <Skeleton width="48px" height="48px" rounded="full" />
        <div className="space-y-2">
          <Skeleton width="120px" height="16px" />
          <Skeleton width="80px" height="12px" />
        </div>
      </div>
      <div className="space-y-2">
        <Skeleton width="100%" height="14px" />
        <Skeleton width="100%" height="14px" />
        <Skeleton width="60%" height="14px" />
      </div>
    </div>
  );
}

export default Skeleton;
