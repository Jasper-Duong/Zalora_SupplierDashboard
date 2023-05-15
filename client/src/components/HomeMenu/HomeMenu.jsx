import { DropboxOutlined, UserOutlined } from "@ant-design/icons";
import { Menu } from "antd";
import { useNavigate } from "react-router-dom";

function getItem(label, key, icon, children) {
  return {
    key,
    icon,
    children,
    label,
  };
}
const items = [
  getItem("Products", "products", <DropboxOutlined />, [
    getItem("Products Table", "products/table"),
    getItem("Add Product", "products/add"),
  ]),
  getItem("Suppliers", "suppliers", <UserOutlined />, [
    getItem("Suppliers Table", "suppliers/table"),
    getItem("Add Supplier", "suppliers/add"),
  ]),
];

export default function HomeMenu() {
  const navigate = useNavigate();
  const handleSelect = ({ key }) => {
    navigate(`/${key}`);
  };
  return (
    <Menu
      theme="dark"
      defaultSelectedKeys={["products/table"]}
      mode="inline"
      items={items}
      onSelect={handleSelect}
    />
  );
}
