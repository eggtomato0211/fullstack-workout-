import { useState, useRef, useEffect } from 'react';

function Dropdown({ label, items, value, onChange }) {
  const [isOpen, setIsOpen] = useState(false);
  const dropdownRef = useRef(null);

  useEffect(() => {
    const handleClickOutside = (e) => {
      if (dropdownRef.current && !dropdownRef.current.contains(e.target)) {
        setIsOpen(false);
      }
    };

    if (isOpen) {
      document.addEventListener('mousedown', handleClickOutside);
    }
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, [isOpen]);

  return (
    <div className="relative inline-block" ref={dropdownRef}>
      <button
        onClick={() => setIsOpen(!isOpen)}
        className="bg-white border border-gray-300 rounded px-4 py-2 hover:bg-gray-50 flex items-center gap-2 min-w-[160px] justify-between"
      >
        <span className={value ? 'text-gray-900' : 'text-gray-400'}>
          {value || label}
        </span>
        <span className={`transition-transform ${isOpen ? 'rotate-180' : ''}`}>▼</span>
      </button>

      {isOpen && (
        <ul className="absolute top-full left-0 mt-1 w-full bg-white border border-gray-200 rounded-lg shadow-lg py-1 z-10">
          {items.map((item) => (
            <li key={item}>
              <button
                className={`w-full text-left px-4 py-2 text-gray-700 ${
                  item === value ? 'bg-blue-50 font-medium' : 'hover:bg-gray-100'
                }`}
                onClick={() => {
                  onChange(item);
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
