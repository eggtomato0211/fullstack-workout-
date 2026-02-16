import Card, { CardBody, CardFooter } from './Card';

type Props = {
  name: string;
  description: string;
  price: string;
  variant?: 'default' | 'outlined';
};

function ProductCard({ name, description, price, variant = 'default' }: Props) {
  return (
    <Card variant={variant}>
      <div className="h-48 bg-gray-300 flex items-center justify-center">
        <span className="text-gray-500 text-sm">商品画像</span>
      </div>
      <CardBody>
        <h3 className="text-lg font-bold">{name}</h3>
        <p className="text-gray-600 mt-1">{description}</p>
      </CardBody>
      <CardFooter>
        <div className="flex items-center justify-between">
          <span className="text-xl font-bold">{price}</span>
          <button className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 transition-colors">
            カートに追加
          </button>
        </div>
      </CardFooter>
    </Card>
  );
}

export default ProductCard;
