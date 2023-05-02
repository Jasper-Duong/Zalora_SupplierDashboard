import './App.css';
import ProductsTable from './components/Table/ProductsTable'
import SuppliersTable from './components/Table/SuppliersTable'

function App() {
  return (
    <div className="App">
      <ProductsTable></ProductsTable>
      <SuppliersTable></SuppliersTable>
    </div>
  );
}

export default App;
