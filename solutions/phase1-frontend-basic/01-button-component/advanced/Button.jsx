function Button({
  variant = 'primary',
  size = 'md',
  disabled = false,
  isLoading = false,
  children,
  onClick,
}) {
  const variantClasses = {
    primary: 'bg-blue-600 text-white hover:bg-blue-700',
    secondary: 'bg-gray-200 text-gray-800 hover:bg-gray-300',
    danger: 'bg-red-600 text-white hover:bg-red-700',
  };

  const sizeClasses = {
    sm: 'px-3 py-1 text-sm',
    md: 'px-4 py-2 text-base',
    lg: 'px-6 py-3 text-lg',
  };

  const isDisabled = disabled || isLoading;

  return (
    <button
      className={`rounded font-medium transition-colors ${variantClasses[variant]} ${sizeClasses[size]} ${
        isDisabled ? 'opacity-50 cursor-not-allowed' : ''
      }`}
      onClick={onClick}
      disabled={isDisabled}
    >
      {isLoading ? (
        <span className="flex items-center gap-2">
          <span className="animate-spin h-4 w-4 border-2 border-white border-t-transparent rounded-full" />
          処理中...
        </span>
      ) : (
        children
      )}
    </button>
  );
}

export default Button;
