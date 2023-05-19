import { Form, Table } from "antd";
import { useState } from "react";
import useService from "../../hooks/useService";
import AddRowBtn from "./Actions/AddRowBtn";
import { TableFormContext } from "../../contexts/TableFormContext";

const TableForm = ({ services, columns, Cell, AddRowComponent }) => {
  const [form] = Form.useForm();
  const [isForceRender, setIsForceRender] = useState(false);
  const forceRender = () => setIsForceRender((prev) => !prev);

  const { data } = useService({
    service: services.getData,
    deps: [isForceRender],
  });
  const [editingRow, setEditingRow] = useState("");

  return (
    <TableFormContext.Provider
      value={{
        forceRender,
        form,
        editingRow,
        setEditingRow,
        services,
      }}
    >
      <AddRowBtn AddRowComponent={AddRowComponent} />
      <Form form={form} component={false}>
        <Table
          components={{
            body: {
              cell: Cell,
            },
          }}
          bordered
          dataSource={data && [...data]}
          columns={columns(editingRow)}
          rowClassName="editable-row"
          pagination={{
            onChange: () => setEditingRow(""),
          }}
        />
      </Form>
    </TableFormContext.Provider>
  );
};
export default TableForm;
