import { Button, Popconfirm, Space, Table, message } from "antd";
import { DeleteFilled, EditOutlined } from "@ant-design/icons";
import ExtendedTable from "./ExtendedTable";
import HomeHeader from "../../layout/HomeLayout/HomeHeader";
import { Link } from "react-router-dom";
import { useState } from "react";
import { deleteSupplierApi } from "../../services/supplier";
import CustomFilterDropdown from "./CustomFilterDropdown";

const SuppliersTable = () => {
  const [isForceRender, setIsForceRender] = useState(false);
  const forceRender = () => setIsForceRender((prev) => !prev);
  const handleDeleteRecord = async (record) => {
    try {
      await deleteSupplierApi(record.id);
      message.success(`Deleted ${record.name}`);
      forceRender();
    } catch (error) {
      console.log(error);
    }
  };
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
      title: "Email",
      dataIndex: "email",
      key: "email",
      ...CustomFilterDropdown("email"),
    },
    {
      title: "Contact Number",
      dataIndex: "contact_number",
      key: "contact_number",
      ...CustomFilterDropdown("contact_number"),
    },
    {
      title: "Stock",
      dataIndex: "stock",
      key: "stock",
      /*sorter: {
        multiple: 1,
      },*/
    },
    Table.EXPAND_COLUMN,
    {
      title: "Addresses",
      dataIndex: "addresses",
      key: "addresses",
      render: (addresses) => {
        return <p>{addresses.length}</p>;
      },
    },
    {
      align: "center",
      key: "action",
      render: (_, record) => (
        <Space wrap>
          <Link to={`/suppliers/edit/${record.id}`}>
            <Button type="primary" icon={<EditOutlined />} />
          </Link>
          <Popconfirm
            title={"Sure to delete?"}
            onConfirm={() => handleDeleteRecord(record)}
          >
            <Button type="primary" icon={<DeleteFilled />} danger />
          </Popconfirm>
        </Space>
      ),
    },
  ];

  return (
    <>
      <HomeHeader title={"Supplier Table"} />
      <ExtendedTable
        isForceRender={isForceRender}
        columns={columns}
        resource={"suppliers"}
      ></ExtendedTable>
    </>
  );
};

export default SuppliersTable;
