import { Button, Space } from "antd";
import { EditFilled, DeleteFilled } from "@ant-design/icons";
import ExtendedTable from "./ExtendedTable";
import AntdModal from "../Modals/AntdModal";
import EditProduct from "../EditProduct/EditProduct";
import EditProductBtn from "../EditProduct/EditProductBtn";
import { useDispatch } from "react-redux";
import { setSelectedProduct } from "../../store/actions/products.action";

const ProductsTable = () => {
  const dispatch = useDispatch();
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
      filters: [
        { text: "Brand 1", value: "Brand 1" },
        { text: "Brand 2", value: "Brand 2" },
      ],
      filterSearch: true,
    },
    {
      title: "SKU",
      dataIndex: "SKU",
      key: "SKU",
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
        { text: "Đen,", value: "Đen," },
        { text: "Hồng", value: "Hồng" },
      ],
      filterSearch: true,
    },
    {
      title: "Suppliers",
      dataIndex: "suppliers",
      key: "suppliers",
      filters: [
        { text: "Supplier 1", value: "Supplier 1" },
        { text: "Supplier 2", value: "Supplier 2" },
      ],
      filterSearch: true,
      render: (suppliers) => {
        return (
          <>
            {suppliers.map((supplier) => (
              <p key={supplier}>{supplier.toUpperCase()}</p>
            ))}
          </>
        );
      },
    },
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
        // console.log(record);
        dispatch(setSelectedProduct(record));
        return (
          <Space wrap>
            <AntdModal
              ShowModalBtn={EditProductBtn}
              title="Edit Product"
              BodyComponent={EditProduct}
              data={record}
            />

            <Button type="primary" icon={<DeleteFilled />} danger />
          </Space>
        );
      },
    },
  ];

  return (
    <ExtendedTable columns={columns} resource={"products"}></ExtendedTable>
  );
};

export default ProductsTable;
