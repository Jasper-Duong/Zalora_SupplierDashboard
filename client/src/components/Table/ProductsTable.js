import { Button, Space, Table, Select, Input } from "antd";
import { DeleteFilled } from "@ant-design/icons";
import ExtendedTable from "./ExtendedTable";
import EditProductBtn from "../EditProduct/EditProductBtn";
import HomeHeader from "../../layout/HomeLayout/HomeHeader";
import { useState } from 'react'

import { getFilterOptions } from "../../services/filter";
const ProductsTable = () => {
  const [filterOptions, setFilterOptions] = useState([]);
  const [filterValues, setFilterValues] = useState({});

  const handleFilterDropdownOpenChange = async (visible, api) => {
    if (visible) {
      const options = await getFilterOptions(`${api}`);
      let filters = options.data
      filters = filters.reduce((accumulator, filter) => (
        [...accumulator, {
          "text": filter.name,
          "value": filter.name,
        }]
      ), [])
      setFilterOptions(filters);
    } else {
      setFilterOptions([])
    }
  }

  const handleFilter = (value) => {
    console.log("help")
    /*setFilterValues(prevValues => ({
      ...prevValues,
      ["brand"]: [...(prevValues["brand"] || []), value],
    }));*/
    // return true or false to indicate whether the record should be included in the filtered result
  };

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
      onFilterDropdownOpenChange: (visible) => visible && handleFilterDropdownOpenChange(visible, "products/attribute/brand"),
      //filteredValue: filterValues.brand || null,
      //onFilter: handleFilter,
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
      filters: [
        {text: "white", value: "white"}
      ],
      filterSearch: true,
      onFilterDropdownOpenChange: (visible) => visible && handleFilterDropdownOpenChange(visible, "products/attribute/color"),
    },
    {
      title: "Suppliers",
      dataIndex: "suppliers",
      key: "suppliers",
      filters: filterOptions,
      filterSearch: true,
      onFilterDropdownOpenChange: (visible) => visible && handleFilterDropdownOpenChange(visible, "suppliers/attribute/name"),
      render: (suppliers) => suppliers.length
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
