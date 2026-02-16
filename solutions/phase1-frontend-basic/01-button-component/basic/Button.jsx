function Button({ variant = 'primary', children, onClick }) {
  const variantClasses = {
    primary: 'bg-blue-600 text-white hover:bg-blue-700',
    secondary: 'bg-gray-200 text-gray-800 hover:bg-gray-300',
    outline: 'bg-white text-blue-600 border border-blue-600 hover:bg-blue-50',
  };

  return (
    <button
      className={`px-4 py-2 rounded font-medium transition-colors ${variantClasses[variant]}`}
      onClick={onClick}
    >
      {children}
    </button>
  );
}

export default Button;
