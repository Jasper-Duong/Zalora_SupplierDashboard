import { Form } from "antd";
import React from "react";

export default function NewRow({ dataIndex, inputNode, title, ...restProps }) {
  return (
    <td>
      <Form.Item
        name={dataIndex}
        rules={[{ required: true, message: `Please input ${title}` }]}
      >
        {inputNode}
      </Form.Item>
    </td>
  );
}
