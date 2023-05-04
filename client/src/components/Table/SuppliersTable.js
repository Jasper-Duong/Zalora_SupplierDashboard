import { Button, Space } from 'antd'
import { EditFilled, DeleteFilled } from '@ant-design/icons'
import ExtendedTable from './ExtendedTable'

const SuppliersTable = () => {    
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
            title: 'Email',
            dataIndex: 'email',
            key: 'email',
        },
        {
            title: 'Contact Number',
            dataIndex: 'contact_number',
            key: 'contact_number',
        },
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
            resource = {'suppliers'}
        ></ExtendedTable>
    )
}

export default SuppliersTable