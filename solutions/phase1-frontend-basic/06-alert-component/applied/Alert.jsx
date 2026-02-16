function Alert({ variant = 'info', children, onClose }) {
  const variantClasses = {
    success: 'bg-green-50 border-green-500 text-green-800',
    warning: 'bg-yellow-50 border-yellow-500 text-yellow-800',
    error: 'bg-red-50 border-red-500 text-red-800',
    info: 'bg-blue-50 border-blue-500 text-blue-800',
  };

  const icons = { success: '✓', warning: '⚠', error: '✕', info: 'ℹ' };

  const iconClasses = {
    success: 'bg-green-500',
    warning: 'bg-yellow-500',
    error: 'bg-red-500',
    info: 'bg-blue-500',
  };

  return (
    <div
      className={`border-l-4 px-4 py-3 rounded flex items-center gap-3 ${variantClasses[variant]}`}
      role="alert"
    >
      <span
        className={`flex-shrink-0 w-6 h-6 rounded-full flex items-center justify-center text-white text-sm ${iconClasses[variant]}`}
      >
        {icons[variant]}
      </span>
      <span className="flex-1">{children}</span>
      {onClose && (
        <button
          onClick={onClose}
          className="flex-shrink-0 text-current opacity-50 hover:opacity-100"
        >
          ✕
        </button>
      )}
    </div>
  );
}

export default Alert;
