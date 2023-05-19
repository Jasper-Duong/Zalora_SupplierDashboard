import { Popconfirm, Typography, message } from "antd";
import React, { useContext } from "react";
import styled from "styled-components";
import { TableFormContext } from "../../../contexts/TableFormContext";

export default function SaveRowBtn({ key }) {
  const { setEditingRow, form, forceRender, editingRow, services:{updateItem} } = useContext(TableFormContext);
  const handleCancel = () => {
    setEditingRow("");
  };
  const handleSave = async () => {
    try {
      const row = await form.validateFields();
      const submitData = { ...row };
      console.log("Updating", submitData);
      await updateItem(editingRow, submitData);
      message.success("Updated Successfully");
      handleCancel();
      forceRender();
    } catch (error) {
      console.log("Validated Failed: ", error);
    }
  };

  return (
    <Styled>
      <Typography.Link className="save-btn" onClick={() => handleSave(key)}>
        Save
      </Typography.Link>
      <Popconfirm
        title="Sure to cancel?"
        className="cancel-btn"
        onConfirm={handleCancel}
      >
        Cancel
      </Popconfirm>
    </Styled>
  );
}

const Styled = styled.span`
  .save-btn {
    margin-right: 0.5rem;
  }
  .cancel-btn {
    cursor: pointer;
    &:hover {
      opacity: 0.5;
    }
  }
`;
