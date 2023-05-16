import { Button, Space, Table } from "antd";
import { DeleteFilled } from "@ant-design/icons";
import ExtendedTable from "./ExtendedTable";
import AntdModal from "../Modals/AntdModal";
import EditSupplierBtn from "../EditSupplier/EditSupplierBtn";
import EditSupplier from "../EditSupplier/EditSupplier";
import HomeHeader from "../../layout/HomeLayout/HomeHeader";

const SuppliersTable = () => {
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
      title: "Email",
      dataIndex: "email",
      key: "email",
    },
    {
      title: "Contact Number",
      dataIndex: "contact_number",
      key: "contact_number",
    },
    {
      title: "Stock",
      dataIndex: "stock",
      key: "stock",
      sorter: {
        multiple: 1,
      },
    },
    Table.EXPAND_COLUMN,
    {/*
      title: "Addresses",
      dataIndex: "addresses",
      key: "addresses",
  */},
    {
      align: "center",
      key: "action",
      render: (_, record) => (
        <Space wrap>
          <AntdModal
            ShowModalBtn={EditSupplierBtn}
            BodyComponent={EditSupplier}
            data={record}
            title={"Edit Supplier"}
          />
          <Button type="primary" icon={<DeleteFilled />} danger />
        </Space>
      ),
    },
  ];

  return (
    <>
      <HomeHeader title={"Supplier Table"} />
      <ExtendedTable columns={columns} resource={"suppliers"}></ExtendedTable>
    </>
  );
};

export default SuppliersTable;
