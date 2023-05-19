import React, { useContext } from "react";
import { TableFormContext } from "../../../contexts/TableFormContext";
import SaveRowBtn from "../Actions/SaveRowBtn";
import ActionsBtns from "../Actions/ActionsBtns";

export default function TableFormActions({ record }) {
  const { editingRow } = useContext(TableFormContext);
  return editingRow === record.id ? (
    <SaveRowBtn />
  ) : (
    <ActionsBtns record={record} />
  );
}
