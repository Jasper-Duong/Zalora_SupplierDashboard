import React, { useContext } from "react";
import { TableFormContext } from "../../../contexts/TableFormContext";
import SaveRowBtn from "../Actions/SaveRowBtn";
import ActionsBtns from "../Actions/ActionsBtns";

export default function TableFormActions({ record }) {
  const { editingKey } = useContext(TableFormContext);
  return editingKey === record.key ? (
    <SaveRowBtn />
  ) : (
    <ActionsBtns record={record} />
  );
}
