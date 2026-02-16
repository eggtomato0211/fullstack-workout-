import type { ReactNode } from 'react';

type Props = {
  children: ReactNode;
};

function Card({ children }: Props) {
  return (
    <div className="bg-white rounded-lg shadow-md p-6">
      {children}
    </div>
  );
}

export default Card;
