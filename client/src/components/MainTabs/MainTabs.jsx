import { Tabs } from 'antd';
import { ProductTab } from '../ProductTab/ProductTab';
import { SupplierTab } from '../SupplierTab/SupplierTab';

const items = [
    {
        key: "1",
        label: "Products",
        children: <ProductTab />
    },
    {
        key: "2",
        label: "Suppliers",
        children: <SupplierTab />
    }
]

export const MainTabs = () => {
    return (
        <div className="main-tabs" style={{marginLeft: "2%", marginRight: "2%"}}>
            <Tabs 
                defaultActiveKey="1"
                items={items}>
            </Tabs>
        </div>
    )
}