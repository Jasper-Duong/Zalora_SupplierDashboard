import { Popconfirm, Typography, message } from "antd";
import React, { useContext } from "react";
import styled from "styled-components";
import { TableFormContext } from "../../../contexts/TableFormContext";
import { updateStockApi } from "../../../services/stock";

export default function SaveRowBtn({ key }) {
  const { setEditingKey, form, forceRender, editingKey } = useContext(TableFormContext);
  const handleCancel = () => {
    setEditingKey("");
  };
  const handleSave = async () => {
    try {
      const row = await form.validateFields();
      const submitData = { key:editingKey, ...row };
      console.log("Updating", submitData, row);
      updateStockApi(1, editingKey, submitData);
      message.success("Updated Stock");
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
