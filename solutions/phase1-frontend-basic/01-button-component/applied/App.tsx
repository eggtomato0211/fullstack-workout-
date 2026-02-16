import Button from './Button';

function App() {
  return (
    <div className="min-h-screen bg-gray-900 flex items-center justify-center">
      <div className="space-y-8 text-center">
        <h1 className="text-2xl font-bold text-white">01 応用: variant + size + disabled</h1>

        {/* variant */}
        <div className="space-y-2">
          <p className="text-gray-400 text-sm">variant</p>
          <div className="flex gap-4 justify-center">
            <Button variant="primary">Primary</Button>
            <Button variant="secondary">Secondary</Button>
            <Button variant="danger">Danger</Button>
          </div>
        </div>

        {/* size */}
        <div className="space-y-2">
          <p className="text-gray-400 text-sm">size</p>
          <div className="flex gap-4 items-center justify-center">
            <Button size="sm">Small</Button>
            <Button size="md">Medium</Button>
            <Button size="lg">Large</Button>
          </div>
        </div>

        {/* disabled */}
        <div className="space-y-2">
          <p className="text-gray-400 text-sm">disabled</p>
          <div className="flex gap-4 justify-center">
            <Button variant="primary" disabled>Primary</Button>
            <Button variant="secondary" disabled>Secondary</Button>
            <Button variant="danger" disabled>Danger</Button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
