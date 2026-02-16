import type { ReactNode } from 'react';

type CardProps = {
  variant?: 'default' | 'outlined';
  children: ReactNode;
};

function Card({ variant = 'default', children }: CardProps) {
  const variantClasses: Record<string, string> = {
    default: 'bg-white shadow-md',
    outlined: 'bg-white border-2 border-gray-200',
  };

  return (
    <div className={`rounded-lg overflow-hidden ${variantClasses[variant]}`}>
      {children}
    </div>
  );
}

type SectionProps = {
  children: ReactNode;
};

export function CardBody({ children }: SectionProps) {
  return <div className="px-6 py-4">{children}</div>;
}

export function CardFooter({ children }: SectionProps) {
  return (
    <div className="px-6 py-4 border-t border-gray-200 bg-gray-50">
      {children}
    </div>
  );
}

export default Card;
