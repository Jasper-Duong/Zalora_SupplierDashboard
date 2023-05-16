import { Button, Space, Table, Select, Input } from "antd";
import { DeleteFilled } from "@ant-design/icons";
import ExtendedTable from "./ExtendedTable";
import EditProductBtn from "../EditProduct/EditProductBtn";
import HomeHeader from "../../layout/HomeLayout/HomeHeader";
import { useState } from 'react'

import { getFilterOptions } from "../../services/filter";
const ProductsTable = () => {
  const [filterOptions, setFilterOptions] = useState([]);

  const handleFilterDropdownOpenChange = async (visible, attribute) => {
    if (visible) {
      const options = await getFilterOptions(`products/attribute/${attribute}`);
      setFilterOptions(options.data);
    }
  }

  const columns = [
    {
      title: "Name",
      dataIndex: "name",
      key: "name",
      sorter: {
        multiple: 2,
      },
    },
    {
      title: "Brand",
      dataIndex: "brand",
      key: "brand",
      filters: filterOptions,
      filterSearch: true,
      onFilterDropdownOpenChange: (visible) => visible && handleFilterDropdownOpenChange(visible, "brand"),
    },
    {
      title: "SKU",
      dataIndex: "sku",
      key: "sku",
    },
    {
      title: "Size",
      dataIndex: "size",
      key: "size",
      sorter: {
        multiple: 3,
      },
    },
    {
      title: "Color",
      dataIndex: "color",
      key: "color",
      filters: filterOptions,
      filterSearch: true,
      onFilterDropdownOpenChange: (visible) => visible && handleFilterDropdownOpenChange(visible, "color"),
    },
    {
      title: "Suppliers",
      dataIndex: "suppliers",
      key: "suppliers",
      filters: filterOptions,
      filterSearch: true,
      onFilterDropdownOpenChange: (visible) => visible && handleFilterDropdownOpenChange(visible, "color"),
      render: (suppliers) => {
        return (
          <>
            {suppliers.map((supplier) => (
              <p key={supplier}>{supplier.name}</p>
            ))}
          </>
        );
      },
    },
    Table.EXPAND_COLUMN,
    /*
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
      title: "Stock",
      dataIndex: "stock",
      key: "stock",
      sorter: {
        multiple: 1,
      },
    },
    {
      align: "center",
      key: "action",
      render: (_, record) => {
        return (
          <Space wrap>
            <EditProductBtn data={record} />
            <Button type="primary" icon={<DeleteFilled />} danger />
          </Space>
        );
      },
    },
  ];

  return (
    <>
      <HomeHeader title={"Product Table"}/>
      <ExtendedTable columns={columns} resource={"products"}></ExtendedTable>
    </>
  );
};

export default ProductsTable;
