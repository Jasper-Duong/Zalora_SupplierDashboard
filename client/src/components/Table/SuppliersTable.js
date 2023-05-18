import { Button, Space, Table, Input } from "antd";
import { DeleteFilled } from "@ant-design/icons";
import ExtendedTable from "./ExtendedTable";
import AntdModal from "../Modals/AntdModal";
import EditSupplierBtn from "../EditSupplier/EditSupplierBtn";
import EditSupplier from "../EditSupplier/EditSupplier";
import HomeHeader from "../../layout/HomeLayout/HomeHeader";

const CustomFilterDropdown = ({ setSelectedKeys, selectedKeys, confirm, clearFilters, close }) => {
  return (
    <div style={{ padding: 8 }} onKeyDown={(e) => e.stopPropagation()}>
        <Input
          placeholder={`Search `}
          value={selectedKeys[0]}
          onChange={(e) => {}}
          onPressEnter={() => {}}
          style={{ marginBottom: 8, display: 'block' }}
        />
        <Space>
          <Button
            type="primary"
            onClick={() => {}}
            //icon={<SearchOutlined />}
            size="small"
            style={{ width: 90 }}
          >
            Search
          </Button>
          <Button
            onClick={() => {}}
            size="small"
            style={{ width: 90 }}
          >
            Reset
          </Button>
        </Space>
      </div>
  )
}
const SuppliersTable = () => {
  const columns = [
    {
      title: "Name",
      dataIndex: "name",
      key: "name",
      sorter: {
        multiple: 2,
      },
      filterDropdown: CustomFilterDropdown
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
    {
      title: "Addresses",
      dataIndex: "addresses",
      key: "addresses",
      render: (addresses) => {
        return (
          <p>{addresses.length}</p>
        )
      }
    },
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
