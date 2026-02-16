function Card({ variant = 'default', children }) {
  const variantClasses = {
    default: 'bg-white shadow-md',
    outlined: 'bg-white border-2 border-gray-200',
    elevated: 'bg-white shadow-xl',
  };

  return (
    <div className={`rounded-lg overflow-hidden ${variantClasses[variant]}`}>
      {children}
    </div>
  );
}

export function CardHeader({ children }) {
  return (
    <div className="px-6 py-4 border-b border-gray-200">
      {children}
    </div>
  );
}

export function CardBody({ children }) {
  return (
    <div className="px-6 py-4">
      {children}
    </div>
  );
}

export function CardFooter({ children }) {
  return (
    <div className="px-6 py-4 border-t border-gray-200 bg-gray-50">
      {children}
    </div>
  );
}

export default Card;
