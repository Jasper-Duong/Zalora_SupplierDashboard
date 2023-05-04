import { Button } from "antd"
import SuppliersTable from "../Table/SuppliersTable"

export const SupplierTab = () => {
    const addNewSupplier = () => {
        alert("Add new supplier")
    }

    return (
        <div className="supplier-container">
            <Button type="primary" style={{float: "left"}} onClick={addNewSupplier}>New supplier</Button>
            <SuppliersTable />
       </div>
    )
}
