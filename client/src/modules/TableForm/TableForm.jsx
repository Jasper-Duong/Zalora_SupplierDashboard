import { Form, Table } from "antd";
import { useState } from "react";
import useService from "../../hooks/useService";
import AddRowBtn from "./Actions/AddRowBtn";
import { TableFormContext } from "../../contexts/TableFormContext";
import TableFormColumns from "./Columns/TableFormColumns";
import EditableCell from "./Rows/EditableRow";

const TableForm = ({ service }) => {
  const [form] = Form.useForm();
  const [isForceRender, setIsForceRender] = useState(false);
  const forceRender = () => setIsForceRender((prev) => !prev);

  const { data } = useService({
    service,
    deps: [isForceRender],
  });
  const [editingKey, setEditingKey] = useState("");

  return (
    <TableFormContext.Provider
      value={{
        forceRender,
        form,
        editingKey,
        setEditingKey,
      }}
    >
      <AddRowBtn forceRender={forceRender} />
      <Form form={form} component={false}>
        <Table
          components={{
            body: {
              cell: EditableCell,
            },
          }}
          bordered
          dataSource={data && [...data]}
          columns={TableFormColumns(editingKey)}
          rowClassName="editable-row"
          pagination={{
            onChange: () => setEditingKey(""),
          }}
        />
      </Form>
    </TableFormContext.Provider>
  );
};
export default TableForm;
