/* eslint-disable react-hooks/exhaustive-deps */
import { useState, useEffect } from "react";
import { Table, Tag, Select } from "antd";
import qs from "qs";
import { request } from "../../config/axios";

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

const ExtendedTable = (props) => {
    const [data, setData] = useState(props.data);
    const [tableParams, setTableParams] = useState({
        pagination: {
            showQuickJumper: true,
            pageSizeOptions: ["10", "50", "100"],
            showSizeChanger: true,
            current: 1,
            pageSize: 10,
        },
        sorter: [],
        status: "true"
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
            status: params.status,
            sort: params.sorter.map(sort => sort.columnKey),
            //order: params.sorter.map(sort => sort.order === 'ascend' ? 'asc' : 'desc'),
            ...filters
        })
    }

    const handleTableChange = (pagination, filters, sorter) => {
        setTableParams({
            pagination,
            filters,
            sorter: Array.isArray(sorter) ? sorter : [sorter],
            status: tableParams.status
        });
    }

    const handleStatusChange = (value) => {
        setTableParams({
            ...tableParams,
            status: value
        });
      };


    useEffect(() => {
        async function fetchData() {
            try {
                const res = await request.get(`${props.resource}/?${decodeURIComponent(qs.stringify(getQueryParams(tableParams), { arrayFormat: 'brackets' }))}`)
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
    }, [JSON.stringify(tableParams), props.isForceRender])

    return (
        <>
        <Select
          size={"large"}
          defaultValue="true"
          style={{ width: 120 }}
          onChange={handleStatusChange}
          options={[
            {value: "true", label: "Active"},
            {value: "false", label: "Not Active"}
          ]}
        />
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
        </>
  );
};

export default ExtendedTable;
