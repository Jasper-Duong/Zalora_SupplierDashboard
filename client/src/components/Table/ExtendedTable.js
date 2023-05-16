import { useState, useEffect } from 'react'
import { Table, Tag } from 'antd'
import qs from 'qs';
import axios from '../../config/axios'

const ExpandedAddresses = (record) => (
    record.addresses.map(address => (
        <p key={address.id}>{address.address_info}</p>
    ))
)

const ExpandedSuppliers = (record) => (
    record.suppliers.map(supplier => (
        <Tag key={supplier.id}>{supplier.name} : {supplier.stock}</Tag>
    ))
)

const isRowExpandable = (record, resource) => (
    (resource === 'products' && record.suppliers.length > 0) || (resource === 'suppliers' && record.addresses.length > 0) ? true : false
)

var oldFilters = {}

const ExtendedTable = (props) => {
    const [data, setData] = useState(props.data)
    const [tableParams, setTableParams] = useState({
        pagination: {
            showQuickJumper: true,
            pageSizeOptions: ['10', '50', '100'],
            showSizeChanger: true,
            current: 1,
            pageSize: 10, 
        },
        sorter: []
    });

    const getQueryParams = (params) => {
        let filters = {}
        if (params.filters) {
            filters = Object.entries(params.filters).reduce((prev, [key, value]) => {
                if (value !== null) prev[key] = value
                return prev
            }, {})
        }
        return ({
            page: params.pagination?.current,
            limit: params.pagination?.pageSize,
            sort: params.sorter.map(sort => sort.columnKey),
            order: params.sorter.map(sort => sort.order === 'ascend' ? 'asc' : 'desc'),
            ...filters
        })
    }

    const handleTableChange = (pagination, filters, sorter) => {
        let newFilters = {}
        Object.keys(filters).forEach(key => {
            console.log(key, filters[key])
            newFilters[key] = filters[key] ? filters[key] : oldFilters[key]
        });
        oldFilters = {...newFilters}
        filters = {...newFilters}
        console.log(newFilters)
        setTableParams({
            pagination,
            filters,
            sorter: Array.isArray(sorter) ?  sorter : [sorter]
        })
    }


    useEffect(() => {
        async function fetchData() {
            try {
                const res = await axios.get(`${props.resource}/?${decodeURIComponent(qs.stringify(getQueryParams(tableParams), { arrayFormat: 'brackets' }))}`)
                setData(res.data[`${props.resource}`])
                setTableParams({
                    ...tableParams,
                    pagination: {
                        ...tableParams.pagination,
                        total: res.data.total, 
                    }
                })
            } catch (err) {
                console.log(err)
            }
        }
        fetchData()
    }, [JSON.stringify(tableParams)])

    return (
        <Table 
            dataSource={data} 
            columns={props.columns}
            pagination={tableParams.pagination}
            onChange={handleTableChange}
            rowKey = "id"
            expandable={{
                expandedRowRender: props.resource === 'products' ? ExpandedSuppliers : ExpandedAddresses,
                rowExpandable: (record) => isRowExpandable(record, props.resource)
            }}
        ></Table>
    )
}

export default ExtendedTable