import { Form, InputNumber } from "antd";
import TableFormSelect from "../Input/TableFormSelect";
import { useContext } from "react";
import { TableFormContext } from "../../../contexts/TableFormContext";

export default function EditableCell({
  editing,
  dataIndex,
  title,
  inputType,
  record,
  index,
  children,
  ...restProps
}) {
  const { form } = useContext(TableFormContext);
  const inputNode =
    inputType === "number" ? <InputNumber /> : <TableFormSelect form={form} />;
  return (
    <td {...restProps}>
      {editing ? (
        <Form.Item
          name={dataIndex}
          style={{
            margin: 0,
          }}
          rules={[
            {
              required: true,
              message: `Please Input ${title}!`,
            },
          ]}
        >
          {inputNode}
        </Form.Item>
      ) : (
        children
      )}
    </td>
  );
}
