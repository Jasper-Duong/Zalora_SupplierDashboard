import { Button, Space, Table, Select, Input } from "antd";
import { DeleteFilled } from "@ant-design/icons";
import ExtendedTable from "./ExtendedTable";
import EditProductBtn from "../EditProduct/EditProductBtn";
import HomeHeader from "../../layout/HomeLayout/HomeHeader";
import { useState } from 'react'

import { getFilterOptions } from "../../services/filter";
const ProductsTable = () => {
  const [filterOptions, setFilterOptions] = useState({});

  const handleFilterDropdownOpenChange = async (visible, api, key) => {
    if (visible) {
      const options = await getFilterOptions(`${api}`);
      let filters = options.data
      filters = filters.reduce((accumulator, filter) => (
        [...accumulator, {
          "text": filter.name,
          "value": filter.name,
        }]
      ), [])
      setFilterOptions({...filterOptions, [key]: filters});
    }
  }

  const filterBox = ({ setSelectedKeys, confirm, selectedKeys, clearFilters }) => {
    const handleSearch = (selectedKeys, confirm) => {
      confirm();
    };
    
    const handleInputChange = (event) => {
      console.log(`Selected filter value: ${event.target.value}`);
      setSelectedKeys(event.target.value ? [event.target.value] : []);
    };
    
    return (
      <div>
        <Input
          placeholder={`Search ${columns[0].title}`}
          onPressEnter={() => handleSearch(selectedKeys, confirm)}
          onChange={handleInputChange}
        />
        <Button onClick={clearFilters}>Reset</Button>
      </div>)
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
      filters: filterOptions?.brand || [],
      filterSearch: true,
      //filterDropdown: (props) => {console.log("cuu")},
      onFilterDropdownOpenChange: (visible) => handleFilterDropdownOpenChange(visible, "products/attribute/brand", "brand"),
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
      filters: filterOptions?.color || [],
      filterSearch: true,
      onFilterDropdownOpenChange: (visible) => visible && handleFilterDropdownOpenChange(visible, "products/attribute/color", "color"),
    },
    {
      title: "Suppliers",
      dataIndex: "suppliers",
      key: "suppliers",
      filters: filterOptions?.suppliers || [],
      filterSearch: true,
      onFilterDropdownOpenChange: (visible) => visible && handleFilterDropdownOpenChange(visible, "suppliers/attribute/name", "suppliers"),
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
