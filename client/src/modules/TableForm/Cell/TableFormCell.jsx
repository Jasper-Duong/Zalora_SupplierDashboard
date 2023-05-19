import React from "react";
import NewRow from "../Rows/NewRow";
import EditableCell from "../Rows/EditableRow";

export default function TableFormCell({ adding, editing, ...restProps }) {
  return adding ? (
    <NewRow {...restProps} />
  ) : (
    <EditableCell editing={editing} {...restProps} />
  );
}
