import React, { useContext } from "react";
import styled from "styled-components";
import EditRowBtn from "./EditRowBtn";
import DeleteRowBtn from "./DeleteRowBtn";
import { TableFormContext } from "../../../contexts/TableFormContext";
import { message } from "antd";

export default function ActionsBtns({ record }) {
  const {
    form,
    forceRender,
    setEditingRow,
    editingRow,
    services: { deleteItem },
  } = useContext(TableFormContext);
  const handleEdit = () => {
    form.setFieldsValue({ ...record });
    setEditingRow(record.id);
  };
  const handleDelete = async () => {
    try {
      console.log("Deleting", record);
      await deleteItem(record.id);
      message.success("Deleted Successfully");
      forceRender();
    } catch (error) {
      console.log("Failed to delete");
    }
  };
  return (
    <Styled>
      <EditRowBtn editingRow={editingRow} handleEdit={handleEdit} />
      <DeleteRowBtn handleDelete={handleDelete} />
    </Styled>
  );
}
const Styled = styled.span`
  display: flex;
  gap: 0.5rem;
`;
