import { useState, useEffect } from 'react'
import { Table } from 'antd'
import axios from 'axios'
import qs from 'qs';

const ExtendedTable = (props) => {
    const [data, setData] = useState(props.data)
    const [tableParams, setTableParams] = useState({
        pagination: {
            showQuickJumper: true,
            pageSizeOptions: ['10', '50', '100'],
            showSizeChanger: true,
            current: 1,
            pageSize: 50, 
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
        setTableParams({
            pagination,
            filters,
            sorter: Array.isArray(sorter) ?  sorter : [sorter]
        })
    }

    useEffect(() => {
        async function fetchData() {
            try {
                console.log(tableParams.filters)
                const res = await axios.get(`https://644f18eaba9f39c6ab5d2217.mockapi.io/${props.resource}?${decodeURIComponent(qs.stringify(getQueryParams(tableParams), { arrayFormat: 'brackets' }))}`)
                //const res = await axios.get(`http://localhost:4000/products?${qs.stringify(getQueryParams(tableParams), { arrayFormat: 'comma' }).replace(/%2C/g, ',')}`)
                setData(res.data)
                setTableParams({
                    ...tableParams,
                    pagination: {
                        ...tableParams.pagination,
                        total: 200, 
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
        ></Table>
    )
}

export default ExtendedTable