import { Button, Space } from 'antd'
import { EditFilled, DeleteFilled } from '@ant-design/icons'
import ExtendedTable from './ExtendedTable'

const ProductsTable = () => { 
    const columns = [
        {
            title: 'Name',
            dataIndex: 'name',
            key: 'name',
            sorter: {
                multiple: 2
            }
        },
        {
            title: 'Brand',
            dataIndex: 'brand',
            key: 'brand',
            filters: [      
                { text: 'Brand 1', value: 'Brand 1', },     
                { text: 'Brand 2', value: 'Brand 2', }, 
            ],
            filterSearch: true,
        },
        {
            title: 'SKU',
            dataIndex: 'SKU',
            key: 'SKU',
        },
        {
            title: 'Size',
            dataIndex: 'size',
            key: 'size',
            sorter: {
                multiple: 3,
            }
        },
        {
            title: 'Color',
            dataIndex: 'color',
            key: 'color',
            filters: [      
                { text: 'Đen,', value: 'Đen,', },     
                { text: 'Hồng', value: 'Hồng', }, 
            ],
            filterSearch: true,
        },
        /*{
            title: 'Supplier',
            dataIndex: 'supplier',
            key: 'supplier',
            filters: [      
                { text: 'Supplier 1', value: 'Supplier 1', },     
                { text: 'Supplier 2', value: 'Supplier 2', }, 
            ],
            filterSearch: true,
            render: (suppliers) => (
                <>
                    {suppliers.map((supplier) => (
                        <p key={supplier}>
                            {supplier.toUpperCase()}
                        </p>
                    ))}
                </>
            )
        },
        {
            title: 'Status',
            dataIndex: 'status',
            key: 'status',
            filters: [      
                { text: 'Active', value: true, },     
                { text: 'Not active', value: false, }, 
            ],
            render: (status) => (
                status ? "Active" : "Not active"
            ),
        },*/
        {
            title: 'Stock',
            dataIndex: 'stock',
            key: 'stock',
            sorter: {
                multiple: 1,
            }
        },
        {
            align: 'center',
            key: 'action',
            render: () => (
                <Space wrap>
                    <Button type="primary" icon={<EditFilled />}/>
                    <Button type="primary" icon={<DeleteFilled />} danger/>
                </Space>
            ),
        },
    ]

    return (
        <ExtendedTable 
            columns={columns}
            resource = {'products'}
        ></ExtendedTable>
    )
}

export default ProductsTable