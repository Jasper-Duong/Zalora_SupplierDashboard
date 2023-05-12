import { Button } from "antd";
import React, { useState } from "react";

export default function MultiForms({ forceRender, data, EditForm, AddForm }) {
  // console.log({data});
  const [isAdd, setIsAdd] = useState(false);
  const renderEditForms = () =>
    data.map((item, idx) => (
      <EditForm forceRender={forceRender} item={item} key={idx} />
    ));
  return (
    <>
      {renderEditForms()}
      {isAdd ? (
        <AddForm forceRender={forceRender} setIsAdd={setIsAdd} />
      ) : (
        <Button onClick={() => setIsAdd(true)}>Add new +</Button>
      )}
    </>
  );
}
