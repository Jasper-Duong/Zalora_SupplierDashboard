import React, { useContext } from "react";
import styled from "styled-components";
import EditRowBtn from "./EditRowBtn";
import DeleteRowBtn from "./DeleteRowBtn";
import { TableFormContext } from "../../../contexts/TableFormContext";
import { deleteStockApi } from "../../../services/stock";
import { message } from "antd";

export default function ActionsBtns({ record }) {
  const { form, forceRender, setEditingKey, editingKey } =
    useContext(TableFormContext);
  const handleEdit = () => {
    form.setFieldsValue({ ...record });
    setEditingKey(record.key);
  };
  const handleDelete = () => {
    try {
      console.log("Deleting", record);
      deleteStockApi(1, record.key);
      message.success("Deleted Stock");
      forceRender();
    } catch (error) {
      console.log("Failed to delete");
    }
  };
  return (
    <Styled>
      <EditRowBtn
        editingKey={editingKey}
        handleEdit={handleEdit}
        record={record}
      />
      <DeleteRowBtn handleDelete={handleDelete} key={record.key} />
    </Styled>
  );
}
const Styled = styled.span`
  display: flex;
  gap: 0.5rem;
`;
