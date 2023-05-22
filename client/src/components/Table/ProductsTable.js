import { Button, Space, Table, Popconfirm, message } from "antd";
import { DeleteFilled } from "@ant-design/icons";
import ExtendedTable from "./ExtendedTable";
import EditProductBtn from "../EditProduct/EditProductBtn";
import HomeHeader from "../../layout/HomeLayout/HomeHeader";
import { useState } from "react";

import { getFilterOptions } from "../../services/filter";
import { deleteProductApi } from "../../services/product";
import CustomFilterDropdown from "./CustomFilterDropdown";

const ProductsTable = () => {
  const [filterOptions, setFilterOptions] = useState([]);

  const [isForceRender, setIsForceRender] = useState(false);
  const forceRender = () => setIsForceRender((prev) => !prev);
  const handleDeleteRecord = async (record) => {
    try {
      await deleteProductApi(record.id);
      message.success(`Deleted ${record.name}`);
      forceRender();
    } catch (error) {
      console.log(error);
    }
  };

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

  const columns = [
    {
      title: "Name",
      dataIndex: "name",
      key: "name",
      /*sorter: {
        multiple: 1,
      },*/
      ...CustomFilterDropdown("name"),
    },
    {
      title: "Brand",
      dataIndex: "brand",
      key: "brand",
      ...CustomFilterDropdown("brand"),
    },
    {
      title: "SKU",
      dataIndex: "sku",
      key: "sku",
      ...CustomFilterDropdown("sku"),
    },
    {
      title: "Size",
      dataIndex: "size",
      key: "size",
      filters: filterOptions?.size || [],
      filterSearch: true,
      onFilterDropdownOpenChange: (visible) => handleFilterDropdownOpenChange(visible, "products/attribute/size", "size"),
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
      //filters: filterOptions?.suppliers || [],
      filterSearch: true,
      onFilterDropdownOpenChange: (visible) => visible && handleFilterDropdownOpenChange(visible, "suppliers/attribute/name", "suppliers"),
      render: (suppliers) => suppliers.length
    },
    Table.EXPAND_COLUMN,
    {
      title: "Stock",
      dataIndex: "stock",
      key: "stock",
      /*sorter: {
        multiple: 2,
      },*/
    },
    {
      align: "center",
      key: "action",
      render: (_, record) => {
        return (
          <Space wrap>
            <EditProductBtn data={record} />
            <Popconfirm
              title={"Sure to delete?"}
              onConfirm={() => handleDeleteRecord(record)}
            >
              <Button type="primary" icon={<DeleteFilled />} danger />
            </Popconfirm>
          </Space>
        );
      },
    },
  ];
  return (
    <>
      <HomeHeader title={"Product Table"} />
      <ExtendedTable
        isForceRender={isForceRender}
        columns={columns}
        resource={"products"}
      ></ExtendedTable>
    </>
  );
};

export default ProductsTable;
