import { Modal } from "antd";
import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function CustomModal({ isOpen, Child, title }) {
  const [open, setOpen] = useState(isOpen);
  const closeModal = () => setOpen(false);
  const navigate = useNavigate();
  const onCancel = () => {
    closeModal();
    setTimeout(() => navigate("/"), 300);
  };
  return (
    <Modal
      destroyOnClose
      title={title}
      onCancel={onCancel}
      open={open}
      footer={null}
    >
      <Child closeModal={closeModal} />
    </Modal>
  );
}
