import { useState } from 'react';

type Props = {
  label: string;
  items: string[];
};

function Dropdown({ label, items }: Props) {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="relative inline-block">
      <button
        onClick={() => setIsOpen(!isOpen)}
        className="bg-white border border-gray-300 rounded px-4 py-2 hover:bg-gray-50 flex items-center gap-2"
      >
        {label}
        <span className={`transition-transform ${isOpen ? 'rotate-180' : ''}`}>▼</span>
      </button>

      {isOpen && (
        <ul className="absolute top-full left-0 mt-1 w-48 bg-white border border-gray-200 rounded-lg shadow-lg py-1 z-10">
          {items.map((item) => (
            <li key={item}>
              <button
                className="w-full text-left px-4 py-2 hover:bg-gray-100 text-gray-700"
                onClick={() => {
                  alert(item);
                  setIsOpen(false);
                }}
              >
                {item}
              </button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default Dropdown;
