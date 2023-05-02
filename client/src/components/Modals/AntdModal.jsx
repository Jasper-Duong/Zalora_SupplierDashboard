import { Button, Modal } from "antd";
import styled from "styled-components";
import React, { useState } from "react";

const AntdModal = ({ openModalBtnText, title, BodyComponent }) => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const showModal = () => {
    setIsModalOpen(true);
  };
  const closeModal = () => {
    setIsModalOpen(false);
  };
  const handleOk = () => {
    setIsModalOpen(false);
  };
  const handleCancel = () => {
    setIsModalOpen(false);
  };
  return (
    <Styled>
      <Button type="primary" onClick={showModal}>
        {openModalBtnText || "Open Modal"}
      </Button>
      <Modal
        destroyOnClose
        title={title || "Basic Modal"}
        open={isModalOpen}
        onOk={handleOk}
        onCancel={handleCancel}
        footer={null}
      >
        <BodyComponent closeModal={closeModal} />
      </Modal>
    </Styled>
  );
};
export default AntdModal;

const Styled = styled.div``;
