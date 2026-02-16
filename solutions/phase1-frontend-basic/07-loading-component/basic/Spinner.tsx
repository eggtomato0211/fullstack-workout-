type Props = {
  size?: 'sm' | 'md' | 'lg';
};

function Spinner({ size = 'md' }: Props) {
  const sizeClasses: Record<string, string> = {
    sm: 'h-4 w-4 border-2',
    md: 'h-8 w-8 border-3',
    lg: 'h-12 w-12 border-4',
  };

  return (
    <div role="status">
      <div
        className={`animate-spin rounded-full border-blue-600 border-t-transparent ${sizeClasses[size]}`}
      />
      <span className="sr-only">読み込み中...</span>
    </div>
  );
}

export default Spinner;
