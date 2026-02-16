import type { ReactNode } from 'react';

type Variant = 'success' | 'warning' | 'error' | 'info';

type Props = {
  variant?: Variant;
  children: ReactNode;
};

function Alert({ variant = 'info', children }: Props) {
  const variantClasses: Record<Variant, string> = {
    success: 'bg-green-50 border-green-500 text-green-800',
    warning: 'bg-yellow-50 border-yellow-500 text-yellow-800',
    error: 'bg-red-50 border-red-500 text-red-800',
    info: 'bg-blue-50 border-blue-500 text-blue-800',
  };

  return (
    <div
      className={`border-l-4 px-4 py-3 rounded ${variantClasses[variant]}`}
      role="alert"
    >
      {children}
    </div>
  );
}

export default Alert;
