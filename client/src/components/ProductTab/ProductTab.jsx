import { Button } from "antd"
import ProductsTable from "../Table/ProductsTable"

export const ProductTab = () => {
    const addNewProduct = () => {
        alert("Add new product")
    }

    return (
        <div className="product-container">
            <Button type="primary" style={{float: "left"}} onClick={addNewProduct}>New product</Button>
            <ProductsTable />
       </div>
    )
}